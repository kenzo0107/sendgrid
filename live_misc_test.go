//go:build live

package sendgrid

import (
	"context"
	"net/http"
	"slices"
	"strings"
	"testing"
	"time"
)

// skipIfStatus は直近レスポンスが指定ステータスのいずれかでエラーになった場合に Skip します。
func skipIfStatus(t *testing.T, rec *recordingClient, err error, statuses ...int) {
	t.Helper()

	if err != nil && slices.Contains(statuses, rec.lastStatus) {
		t.Skipf("このアカウントでは利用できない API のため skip (HTTP %d): %s", rec.lastStatus, liveBody(rec))
	}
}

// assertNoUnknownFieldsQuiet は assertNoUnknownFields と同じ検証を、レスポンスボディを出力せずに行います。
// 個人情報を含むレスポンス (user 系) で使います。
func assertNoUnknownFieldsQuiet[T any](t *testing.T, rec *recordingClient) {
	t.Helper()

	if msg := checkResponseShape[T](rec.lastBody); msg != "" {
		t.Error(msg)
	}
}

func liveDate(daysAgo int) string {
	return time.Now().UTC().AddDate(0, 0, -daysAgo).Format("2006-01-02")
}

func TestLiveMiscPreBuiltDesign(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	list, err := client.GetPreBuiltDesigns(ctx, &InputGetPreBuiltDesigns{PageSize: 10, Summary: true})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetPreBuiltDesigns: %v", err)
	}
	if len(list.Result) == 0 {
		t.Skip("Pre-built design が 0 件のため skip")
	}
	id := list.Result[0].ID

	got, err := client.GetPreBuiltDesign(ctx, id)
	if err != nil {
		t.Fatalf("GetPreBuiltDesign: %v", err)
	}
	assertNoUnknownFields[OutputGetPreBuiltDesign](t, rec)
	if got.ID != id {
		t.Errorf("id = %q, want %q", got.ID, id)
	}

	name := liveName("prebuilt-design")
	dup, err := client.DuplicatePreBuiltDesign(ctx, id, &InputDuplicatePreBuiltDesign{Name: name})
	if err != nil {
		t.Fatalf("DuplicatePreBuiltDesign: %v", err)
	}
	if dup.ID == "" || dup.ID == id {
		t.Fatalf("複製結果が想定外です: id=%q", dup.ID)
	}
	deleteDesign := deleteOnce(t, "Design "+dup.ID, func() error {
		return client.DeleteDesign(context.Background(), dup.ID)
	})
	assertNoUnknownFields[OutputDuplicatePreBuiltDesign](t, rec)
	if dup.Name != name {
		t.Errorf("name = %q, want %q", dup.Name, name)
	}

	copied, err := client.GetDesign(ctx, dup.ID)
	if err != nil {
		t.Fatalf("GetDesign: %v", err)
	}
	if copied.Name != name {
		t.Errorf("複製後の再取得: name = %q, want %q", copied.Name, name)
	}

	if err := deleteDesign(); err != nil {
		t.Fatalf("DeleteDesign: %v", err)
	}
	_, err = client.GetDesign(ctx, dup.ID)
	assertDeleted(t, rec, err, "Design "+dup.ID)
}

func TestLiveMiscAuthenticatedDomains(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	domains, err := client.GetAuthenticatedDomains(ctx, &InputGetAuthenticatedDomains{Limit: 50})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetAuthenticatedDomains: %v", err)
	}
	assertNoUnknownFields[[]*DomainAuthentication](t, rec)
	if len(domains) == 0 {
		t.Skip("Authenticated Domain が 0 件のため DNS フィールドの検証は skip")
	}

	if _, err := client.GetAuthenticatedDomain(ctx, domains[0].ID); err != nil {
		t.Fatalf("GetAuthenticatedDomain: %v", err)
	}
	assertNoUnknownFields[OutputGetAuthenticatedDomain](t, rec)
}

func TestLiveMiscBounceClassifications(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	if _, err := client.GetBounceClassifications(ctx, &InputGetBounceClassifications{
		StartDate: liveDate(30),
		EndDate:   liveDate(0),
	}); err != nil {
		skipIfUnavailable(t, rec, err)
		t.Fatalf("GetBounceClassifications: %v", err)
	}
	assertNoUnknownFields[OutputGetBounceClassifications](t, rec)

	if _, err := client.GetBounceClassificationsByDomain(ctx, "Content", &InputGetBounceClassificationsByDomain{
		StartDate: liveDate(30),
		EndDate:   liveDate(0),
	}); err != nil {
		skipIfUnavailable(t, rec, err)
		t.Fatalf("GetBounceClassificationsByDomain: %v", err)
	}
	assertNoUnknownFields[OutputGetBounceClassificationsByDomain](t, rec)
}

func TestLiveMiscEngagementQualityScores(t *testing.T) {
	client, rec := newLiveClient(t)

	_, err := client.GetEngagementQualityScores(context.Background(), &InputGetEngagementQualityScores{
		From: liveDate(7),
		To:   liveDate(1),
	})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetEngagementQualityScores: %v", err)
	}
	assertNoUnknownFields[OutputGetEngagementQualityScores](t, rec)
}

func TestLiveMiscSubuserEngagementQualityScores(t *testing.T) {
	client, rec := newLiveClient(t)

	r, err := client.GetSubuserEngagementQualityScores(context.Background(), &InputGetSubuserEngagementQualityScores{
		Limit: 10,
		Date:  liveDate(2),
	})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetSubuserEngagementQualityScores: %v", err)
	}
	assertNoUnknownFields[OutputGetSubuserEngagementQualityScores](t, rec)
	if len(r.Result) == 0 {
		t.Log("Subuser のスコアが 0 件のため score の各フィールドは未検証")
	}
}

func TestLiveMiscSendIPs(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	ips, err := client.GetSendIPs(ctx, &InputGetSendIPs{Limit: 10})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetSendIPs: %v", err)
	}
	assertNoUnknownFields[OutputGetSendIPs](t, rec)

	if len(ips.Result) > 0 {
		ip := ips.Result[0].IP
		if _, err := client.GetSendIP(ctx, ip); err != nil {
			t.Fatalf("GetSendIP: %v", err)
		}
		assertNoUnknownFields[OutputGetSendIP](t, rec)

		if _, err := client.GetSendIPSubusers(ctx, ip, &InputGetSendIPSubusers{Limit: 10}); err != nil {
			t.Fatalf("GetSendIPSubusers: %v", err)
		}
		assertNoUnknownFields[OutputGetSendIPSubusers](t, rec)
	} else {
		t.Log("IP が 0 件のため GetSendIP / GetSendIPSubusers は未実行")
	}

	pools, err := client.GetSendIPPools(ctx, &InputGetSendIPPools{Limit: 10})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetSendIPPools: %v", err)
	}
	assertNoUnknownFields[OutputGetSendIPPools](t, rec)

	if len(pools.Result) == 0 {
		t.Log("IP Pool が 0 件のため GetSendIPPool / GetSendIPPoolIPs は未実行")
		return
	}
	poolID := pools.Result[0].ID
	if _, err := client.GetSendIPPool(ctx, poolID); err != nil {
		t.Fatalf("GetSendIPPool: %v", err)
	}
	assertNoUnknownFields[OutputGetSendIPPool](t, rec)

	if _, err := client.GetSendIPPoolIPs(ctx, poolID, &InputGetSendIPPoolIPs{Limit: 10}); err != nil {
		t.Fatalf("GetSendIPPoolIPs: %v", err)
	}
	assertNoUnknownFields[OutputGetSendIPPoolIPs](t, rec)
}

func TestLiveMiscPartnerSettings(t *testing.T) {
	client, rec := newLiveClient(t)

	if _, err := client.GetPartnerSettings(context.Background(), &InputGetPartnerSettings{Limit: 10}); err != nil {
		skipIfStatus(t, rec, err, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound)
		t.Fatalf("GetPartnerSettings: %v", err)
	}
	assertNoUnknownFields[OutputGetPartnerSettings](t, rec)
}

// user 系のレスポンスは個人情報を含むため、値やボディを出力しません。
func TestLiveMiscUser(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	fail := func(method string, err error) {
		t.Helper()
		// エラーメッセージにレスポンスボディが含まれうるため、ステータスのみ出力する
		_ = err
		t.Fatalf("%s が失敗しました (HTTP %d)", method, rec.lastStatus)
	}

	if _, err := client.GetUserCredits(ctx); err != nil {
		skipIfUnavailable(t, rec, err)
		fail("GetUserCredits", err)
	}
	assertNoUnknownFieldsQuiet[OutputGetUserCredits](t, rec)

	email, err := client.GetUserEmail(ctx)
	if err != nil {
		skipIfUnavailable(t, rec, err)
		fail("GetUserEmail", err)
	}
	assertNoUnknownFieldsQuiet[OutputGetUserEmail](t, rec)
	if !strings.Contains(email.Email, "@") {
		t.Error("email の形式が想定外です")
	}

	if _, err := client.GetUserProfile(ctx); err != nil {
		skipIfUnavailable(t, rec, err)
		fail("GetUserProfile", err)
	}
	assertNoUnknownFieldsQuiet[OutputGetUserProfile](t, rec)

	username, err := client.GetUsername(ctx)
	if err != nil {
		skipIfUnavailable(t, rec, err)
		fail("GetUsername", err)
	}
	assertNoUnknownFieldsQuiet[OutputGetUsername](t, rec)
	if username.Username == "" || username.UserID == 0 {
		t.Error("username / user_id が空です")
	}
}
