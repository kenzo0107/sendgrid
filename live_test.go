//go:build live

// 実際の SendGrid API を叩いて動作を確認する試験です。
// 試験専用アカウントの API キーで `mise run test-live` から実行してください。
//
// 必須の環境変数:
//   - SENDGRID_API_KEY:       試験専用アカウントの API キー
//   - SENDGRID_LIVE_USERNAME: 試験専用アカウントのユーザー名 (本番キーでの誤実行を防ぐためのガード)
//
// 任意の環境変数:
//   - SENDGRID_LIVE_PRIVATE_USERNAME: 個人のプライベートアカウントのユーザー名。
//     API キーのアカウントがこのユーザー名と一致しない場合は共有アカウントとみなし、
//     既存リソースや設定に影響し得る試験を skip し、レスポンスボディも出力しない
package sendgrid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

// liveResourcePrefix は試験で作成するリソース名の接頭辞です。削除漏れの特定に使います。
const liveResourcePrefix = "sdk-live-"

var (
	liveGuardOnce sync.Once
	liveGuardErr  error
	liveUsername  string
)

// recordingClient は直近のレスポンスボディを記録する httpClient です。
type recordingClient struct {
	client     *http.Client
	lastStatus int
	lastBody   []byte
}

func (r *recordingClient) Do(req *http.Request) (*http.Response, error) {
	// 通信エラー時に前のリクエストのステータスで skip や削除確認を誤判定しないよう、毎回リセットする
	r.lastStatus, r.lastBody = 0, nil
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	r.lastStatus = resp.StatusCode
	r.lastBody = body
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return resp, nil
}

// newLiveClient は実 API 向けの Client を返します。
// 前提が揃っていない場合は Skip し、試験専用アカウントでない場合は Fatal にします。
func newLiveClient(t *testing.T) (*Client, *recordingClient) {
	t.Helper()

	apiKey := os.Getenv("SENDGRID_API_KEY")
	username := os.Getenv("SENDGRID_LIVE_USERNAME")
	if apiKey == "" || username == "" {
		t.Skip("SENDGRID_API_KEY と SENDGRID_LIVE_USERNAME が未設定のため skip")
	}

	rec := &recordingClient{client: &http.Client{Timeout: 30 * time.Second}}
	client := New(apiKey, OptionHTTPClient(rec))

	liveGuardOnce.Do(func() {
		r, err := client.GetUsername(context.Background())
		if err != nil {
			liveGuardErr = fmt.Errorf("ユーザー名の取得に失敗: %w", err)
			return
		}
		if r.Username != username {
			liveGuardErr = fmt.Errorf("API キーのユーザー名 %q が SENDGRID_LIVE_USERNAME %q と一致しません", r.Username, username)
			return
		}
		liveUsername = r.Username
	})
	if liveGuardErr != nil {
		t.Fatalf("試験専用アカウントであることを確認できないため中止します: %v", liveGuardErr)
	}

	return client, rec
}

// liveName は試験用リソースの一意な名前を返します。
func liveName(suffix string) string {
	return fmt.Sprintf("%s%s-%d", liveResourcePrefix, suffix, time.Now().UnixNano())
}

// skipIfUnavailable は、プランやアカウント種別の制約で API が使えない (401/403) 場合に Skip します。
// API キー自体の有効性は newLiveClient のガードで確認済みのため、401/403 は機能が利用できないことを意味します。
func skipIfUnavailable(t *testing.T, rec *recordingClient, err error) {
	t.Helper()

	if err != nil && (rec.lastStatus == http.StatusUnauthorized || rec.lastStatus == http.StatusForbidden) {
		t.Skipf("このアカウントでは利用できない API のため skip (HTTP %d): %s", rec.lastStatus, liveBody(rec))
	}
}

// isPrivateAccount は環境変数のフラグではなく実際のユーザー名で判定します。
// mise は MISE_ENV 指定時も mise.local.toml を読み込むため、フラグだと別アカウントの実行に引き継がれてしまうためです。
func isPrivateAccount() bool {
	private := os.Getenv("SENDGRID_LIVE_PRIVATE_USERNAME")
	return private != "" && liveUsername == private
}

// liveBody は共有アカウントでは実際の宛先などを含み得るため、レスポンスボディを伏せて返します。
func liveBody(rec *recordingClient) string {
	if isPrivateAccount() {
		return string(rec.lastBody)
	}
	return "<共有アカウントのため非表示>"
}

// deleteOnce は試験の最後に明示的に削除するための関数を返します。
// 途中で失敗して呼ばれなかった場合に備え、未実行なら Cleanup で削除します。
func deleteOnce(t *testing.T, what string, del func() error) func() error {
	t.Helper()

	done := false
	t.Cleanup(func() {
		if done {
			return
		}
		if err := del(); err != nil {
			t.Errorf("%s の後片付けに失敗しました: %v", what, err)
		}
	})
	return func() error {
		done = true
		return del()
	}
}

// assertDeleted は削除後の取得結果 err が「見つからない」ことを検証します。
// API によっては 404 ではなく 400 + "not found" を返すため (例: Suppression Group)、notFoundStatuses で追加のステータスを許容します。
func assertDeleted(t *testing.T, rec *recordingClient, err error, what string, notFoundStatuses ...int) {
	t.Helper()

	if err == nil {
		t.Errorf("%s が削除後も取得できます", what)
		return
	}
	if rec.lastStatus == http.StatusNotFound {
		return
	}
	if slices.Contains(notFoundStatuses, rec.lastStatus) && strings.Contains(err.Error(), "not found") {
		return
	}
	t.Errorf("%s の削除後の取得が想定外の結果でした (HTTP %d): %v", what, rec.lastStatus, err)
}

// assertNoUnknownFields は直近のレスポンスに、T に定義されていないフィールドが無いことを検証します。
func assertNoUnknownFields[T any](t *testing.T, rec *recordingClient) {
	t.Helper()

	if msg := checkResponseShape[T](rec.lastBody); msg != "" {
		t.Errorf("%s\nbody: %s", msg, liveBody(rec))
	}
}

// checkResponseShape は body を T にデコードできるか、T に未定義のフィールドが無いかを検証し、問題があればその内容を返します。
func checkResponseShape[T any](body []byte) string {
	var v T
	if err := json.Unmarshal(body, &v); err != nil {
		return fmt.Sprintf("レスポンスを %T にデコードできません: %v", v, err)
	}

	var raw any
	if err := json.Unmarshal(body, &raw); err != nil {
		return fmt.Sprintf("レスポンスが JSON ではありません: %v", err)
	}

	var paths []string
	collectUnknownFields(reflect.TypeOf(v), raw, "$", &paths)
	if len(paths) == 0 {
		return ""
	}
	sort.Strings(paths)
	paths = slices.Compact(paths)
	return fmt.Sprintf("レスポンスに %T で未定義のフィールドがあります: %s", v, strings.Join(paths, ", "))
}

// collectUnknownFields は JSON の値 raw を型 typ と照合し、typ に存在しないキーのパスを paths に追加します。
func collectUnknownFields(typ reflect.Type, raw any, path string, paths *[]string) {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	switch val := raw.(type) {
	case map[string]any:
		switch typ.Kind() {
		case reflect.Map:
			for k, child := range val {
				collectUnknownFields(typ.Elem(), child, path+"."+k, paths)
			}
		case reflect.Struct:
			fields := jsonFields(typ)
			for k, child := range val {
				ft, ok := fields[k]
				if !ok {
					*paths = append(*paths, path+"."+k)
					continue
				}
				collectUnknownFields(ft, child, path+"."+k, paths)
			}
		}
	case []any:
		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			for _, child := range val {
				collectUnknownFields(typ.Elem(), child, path+"[]", paths)
			}
		}
	}
}

// jsonFields は構造体の JSON キーとフィールド型の対応を返します。埋め込み構造体のフィールドも含みます。
func jsonFields(typ reflect.Type) map[string]reflect.Type {
	fields := map[string]reflect.Type{}
	for i := range typ.NumField() {
		f := typ.Field(i)
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name, _, _ := strings.Cut(tag, ",")
		if f.Anonymous && name == "" {
			ft := f.Type
			for ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				for k, v := range jsonFields(ft) {
					fields[k] = v
				}
				continue
			}
		}
		if !f.IsExported() {
			continue
		}
		if name == "" {
			name = f.Name
		}
		fields[name] = f.Type
	}
	return fields
}

func TestLiveGetScopes(t *testing.T) {
	client, rec := newLiveClient(t)

	r, err := client.GetScopes(context.Background())
	if err != nil {
		t.Fatalf("GetScopes: %v", err)
	}
	if len(r.Scopes) == 0 {
		t.Error("scopes が空です")
	}
	assertNoUnknownFields[OutputGetScopes](t, rec)
}

func TestLiveGetUserAccount(t *testing.T) {
	client, rec := newLiveClient(t)

	r, err := client.GetUserAccount(context.Background())
	if err != nil {
		t.Fatalf("GetUserAccount: %v", err)
	}
	if r.Type == "" {
		t.Error("type が空です")
	}
	assertNoUnknownFields[OutputGetUserAccount](t, rec)
}

func TestLiveTemplateLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	created, err := client.CreateTemplate(ctx, &InputCreateTemplate{
		Name:       liveName("template"),
		Generation: "dynamic",
	})
	if err != nil {
		t.Fatalf("CreateTemplate: %v", err)
	}
	if created.ID == "" {
		t.Fatal("作成した Template の ID が空です")
	}
	deleteTemplate := deleteOnce(t, "Template "+created.ID, func() error {
		return client.DeleteTemplate(context.Background(), created.ID)
	})
	assertNoUnknownFields[OutputCreateTemplate](t, rec)
	if created.ID == "" || created.Generation != "dynamic" {
		t.Fatalf("作成結果が想定外です: %+v", created)
	}

	got, err := client.GetTemplate(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetTemplate: %v", err)
	}
	assertNoUnknownFields[OutputGetTemplate](t, rec)
	if got.Name != created.Name {
		t.Errorf("name = %q, want %q", got.Name, created.Name)
	}

	list, err := client.GetTemplates(ctx, &InputGetTemplates{Generations: "dynamic", PageSize: 200})
	if err != nil {
		t.Fatalf("GetTemplates: %v", err)
	}
	assertNoUnknownFields[OutputGetTemplates](t, rec)
	if !slices.ContainsFunc(list.Templates, func(tpl Template) bool { return tpl.ID == created.ID }) {
		t.Errorf("一覧に作成したテンプレート %s がありません", created.ID)
	}

	newName := liveName("template-updated")
	updated, err := client.UpdateTemplate(ctx, created.ID, &InputUpdateTemplate{Name: newName})
	if err != nil {
		t.Fatalf("UpdateTemplate: %v", err)
	}
	assertNoUnknownFields[OutputUpdateTemplate](t, rec)
	if updated.Name != newName {
		t.Errorf("UpdateTemplate の戻り値: name = %q, want %q", updated.Name, newName)
	}
	got, err = client.GetTemplate(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetTemplate: %v", err)
	}
	if got.Name != newName {
		t.Errorf("更新後の再取得: name = %q, want %q", got.Name, newName)
	}

	if err := deleteTemplate(); err != nil {
		t.Fatalf("DeleteTemplate: %v", err)
	}
	_, err = client.GetTemplate(ctx, created.ID)
	assertDeleted(t, rec, err, "Template "+created.ID)
}

func TestLiveSuppressionGroupLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	created, err := client.CreateSuppressionGroup(ctx, &InputCreateSuppressionGroup{
		Name:        liveName("group"),
		Description: "created by kenzo0107/sendgrid live test",
	})
	if err != nil {
		t.Fatalf("CreateSuppressionGroup: %v", err)
	}
	if created.ID == 0 {
		t.Fatal("作成した Suppression Group の ID が空です")
	}
	deleteGroup := deleteOnce(t, fmt.Sprintf("Suppression Group %d", created.ID), func() error {
		return client.DeleteSuppressionGroup(context.Background(), created.ID)
	})
	assertNoUnknownFields[OutputCreateSuppressionGroup](t, rec)
	if created.ID == 0 || created.IsDefault {
		t.Fatalf("作成結果が想定外です: %+v", created)
	}

	got, err := client.GetSuppressionGroup(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetSuppressionGroup: %v", err)
	}
	assertNoUnknownFields[SuppressionGroup](t, rec)
	if got.Name != created.Name {
		t.Errorf("name = %q, want %q", got.Name, created.Name)
	}

	groups, err := client.GetSuppressionGroups(ctx)
	if err != nil {
		t.Fatalf("GetSuppressionGroups: %v", err)
	}
	assertNoUnknownFields[[]*SuppressionGroup](t, rec)
	if !slices.ContainsFunc(groups, func(g *SuppressionGroup) bool { return g.ID == created.ID }) {
		t.Errorf("一覧に作成したグループ %d がありません", created.ID)
	}

	newDesc := "updated by kenzo0107/sendgrid live test"
	updated, err := client.UpdateSuppressionGroup(ctx, created.ID, &InputUpdateSuppressionGroup{
		Name:        created.Name,
		Description: newDesc,
	})
	if err != nil {
		t.Fatalf("UpdateSuppressionGroup: %v", err)
	}
	assertNoUnknownFields[OutputUpdateSuppressionGroup](t, rec)
	if updated.Description != newDesc {
		t.Errorf("UpdateSuppressionGroup の戻り値: description = %q, want %q", updated.Description, newDesc)
	}
	got, err = client.GetSuppressionGroup(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetSuppressionGroup: %v", err)
	}
	if got.Description != newDesc {
		t.Errorf("更新後の再取得: description = %q, want %q", got.Description, newDesc)
	}

	if err := deleteGroup(); err != nil {
		t.Fatalf("DeleteSuppressionGroup: %v", err)
	}
	_, err = client.GetSuppressionGroup(ctx, created.ID)
	assertDeleted(t, rec, err, fmt.Sprintf("Suppression Group %d", created.ID), http.StatusBadRequest)
}
