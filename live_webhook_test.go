//go:build live

package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"testing"
	"time"
)

// Webhook Security Policy は作成・更新・削除の結果が GET や一覧に反映されるまで数秒かかるため、1 回の再取得では判定できません。
func eventually(t *testing.T, timeout time.Duration, cond func() bool) bool {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for {
		if cond() {
			return true
		}
		if time.Now().After(deadline) {
			return false
		}
		time.Sleep(2 * time.Second)
	}
}

func TestLiveWebhookSecurityPolicyLifecycle(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	created, err := client.CreateWebhookSecurityPolicy(ctx, &InputCreateWebhookSecurityPolicy{
		Name: liveName("policy"),
		OAuth: &WebhookSecurityPolicyOAuth{
			ClientID:     "sdk-live-client",
			ClientSecret: "sdk-live-secret",
			TokenURL:     "https://example.com/token",
			Scopes:       []string{"webhooks:read"},
		},
		Signature: &WebhookSecurityPolicySignature{Enabled: true},
	})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("CreateWebhookSecurityPolicy: %v", err)
	}
	if created.Policy == nil || created.Policy.ID == "" {
		t.Fatalf("作成結果が想定外です: %+v", created)
	}
	id := created.Policy.ID
	deletePolicy := deleteOnce(t, "Webhook Security Policy "+id, func() error {
		return client.DeleteWebhookSecurityPolicy(context.Background(), id)
	})
	assertNoUnknownFields[OutputCreateWebhookSecurityPolicy](t, rec)
	if created.Policy.OAuth == nil || created.Policy.OAuth.ClientSecret != "" {
		t.Errorf("oauth は client_secret を除いて返る想定です: %+v", created.Policy.OAuth)
	}
	if created.Policy.Signature == nil || created.Policy.Signature.PublicKey == "" {
		t.Errorf("signature の public_key が返っていません: %+v", created.Policy.Signature)
	}

	got, err := client.GetWebhookSecurityPolicy(ctx, id)
	if err != nil {
		t.Fatalf("GetWebhookSecurityPolicy: %v", err)
	}
	assertNoUnknownFields[OutputGetWebhookSecurityPolicy](t, rec)
	if got.Policy.Name != created.Policy.Name {
		t.Errorf("name = %q, want %q", got.Policy.Name, created.Policy.Name)
	}

	if !eventually(t, 30*time.Second, func() bool {
		policies, err := client.GetWebhookSecurityPolicies(ctx)
		return err == nil && slices.ContainsFunc(policies.Policies, func(p *WebhookSecurityPolicy) bool { return p.ID == id })
	}) {
		t.Errorf("一覧に作成した policy %s がありません", id)
	}
	assertNoUnknownFields[OutputGetWebhookSecurityPolicies](t, rec)

	newName := liveName("policy-updated")
	updated, err := client.UpdateWebhookSecurityPolicy(ctx, id, &InputUpdateWebhookSecurityPolicy{
		Name:      newName,
		Signature: &WebhookSecurityPolicySignature{Enabled: false},
	})
	if err != nil {
		t.Fatalf("UpdateWebhookSecurityPolicy: %v", err)
	}
	assertNoUnknownFields[OutputUpdateWebhookSecurityPolicy](t, rec)
	if updated.Policy.Name != newName || updated.Policy.Signature != nil || updated.Policy.OAuth == nil {
		t.Errorf("UpdateWebhookSecurityPolicy の戻り値が想定外です (name を変更し signature のみ外す想定): %+v", updated.Policy)
	}
	if !eventually(t, 30*time.Second, func() bool {
		got, err = client.GetWebhookSecurityPolicy(ctx, id)
		return err == nil && got.Policy.Name == newName && got.Policy.Signature == nil
	}) {
		t.Errorf("更新後の再取得が想定外です: %+v (err: %v)", got, err)
	}

	if err := deletePolicy(); err != nil {
		t.Fatalf("DeleteWebhookSecurityPolicy: %v", err)
	}
	eventually(t, 30*time.Second, func() bool {
		_, err = client.GetWebhookSecurityPolicy(ctx, id)
		return rec.lastStatus == http.StatusNotFound
	})
	assertDeleted(t, rec, err, "Webhook Security Policy "+id)
}

func TestLiveInboundParseWebhookSecurityPolicy(t *testing.T) {
	client, rec := newLiveClient(t)
	ctx := context.Background()

	if !isPrivateAccount() {
		t.Skip("共有アカウントの Inbound Parse 設定を変更しないため skip")
	}

	domains, err := client.GetAuthenticatedDomains(ctx, &InputGetAuthenticatedDomains{Limit: 50})
	skipIfUnavailable(t, rec, err)
	if err != nil {
		t.Fatalf("GetAuthenticatedDomains: %v", err)
	}
	idx := slices.IndexFunc(domains, func(d *DomainAuthentication) bool { return d.Valid })
	if idx < 0 {
		t.Skip("Parse 設定の作成には検証済みの Authenticated Domain が必要なため skip")
	}
	hostname := fmt.Sprintf("%sparse-%d.%s", liveResourcePrefix, time.Now().UnixNano(), domains[idx].Domain)

	newPolicy := func(suffix string) string {
		t.Helper()
		created, err := client.CreateWebhookSecurityPolicy(ctx, &InputCreateWebhookSecurityPolicy{
			Name:      liveName(suffix),
			Signature: &WebhookSecurityPolicySignature{Enabled: true},
		})
		skipIfUnavailable(t, rec, err)
		if err != nil {
			t.Fatalf("CreateWebhookSecurityPolicy: %v", err)
		}
		id := created.Policy.ID
		t.Cleanup(func() {
			if err := client.DeleteWebhookSecurityPolicyWithOptions(context.Background(), id, &InputDeleteWebhookSecurityPolicy{Force: true}); err != nil && rec.lastStatus != http.StatusNotFound {
				t.Errorf("Webhook Security Policy %s の後片付けに失敗しました: %v", id, err)
			}
		})
		return id
	}
	// 途中で 401/403 になってテスト全体が skip にならないよう、検証を始める前に両方作っておく
	policyID := newPolicy("parse-policy")
	forcedID := newPolicy("parse-policy-forced")

	created, err := client.CreateInboundParseWebhook(ctx, &InputCreateInboundParseWebhook{
		Hostname:       hostname,
		URL:            "https://example.com/parse",
		SecurityPolicy: policyID,
	})
	if err != nil {
		t.Fatalf("CreateInboundParseWebhook: %v", err)
	}
	deleteParse := deleteOnce(t, "Inbound Parse "+hostname, func() error {
		return client.DeleteInboundParseWebhook(context.Background(), hostname)
	})
	assertNoUnknownFields[OutputCreateInboundParseWebhook](t, rec)
	if created.SecurityPolicy != policyID {
		t.Errorf("作成時の security_policy = %q, want %q", created.SecurityPolicy, policyID)
	}

	got, err := client.GetInboundParseWebhook(ctx, hostname)
	if err != nil {
		t.Fatalf("GetInboundParseWebhook: %v", err)
	}
	assertNoUnknownFields[OutputGetInboundParseWebhook](t, rec)
	if got.SecurityPolicy != policyID {
		t.Errorf("security_policy = %q, want %q", got.SecurityPolicy, policyID)
	}

	updated, err := client.UpdateInboundParseWebhook(ctx, hostname, &InputUpdateInboundParseWebhook{URL: "https://example.com/parse2"})
	if err != nil {
		t.Fatalf("UpdateInboundParseWebhook: %v", err)
	}
	assertNoUnknownFields[OutputUpdateInboundParseWebhook](t, rec)
	if updated.SecurityPolicy != policyID {
		t.Errorf("SecurityPolicy を省いた更新で security_policy = %q になりました, want %q", updated.SecurityPolicy, policyID)
	}

	// 参照が反映される前だと force なしでも削除が通ってしまい、やり直せないため、先に反映を待つ
	if !eventually(t, 30*time.Second, func() bool {
		got, err = client.GetInboundParseWebhook(ctx, hostname)
		return err == nil && got.SecurityPolicy == policyID
	}) {
		t.Fatalf("Parse 設定にポリシーの参照が反映されません: %+v (err: %v)", got, err)
	}
	if err := client.DeleteWebhookSecurityPolicy(ctx, policyID); err == nil || rec.lastStatus != http.StatusConflict {
		t.Fatalf("Parse 設定に付いたポリシーの force なし削除は 409 になる想定です (HTTP %d): %v", rec.lastStatus, err)
	}

	updated, err = client.UpdateInboundParseWebhook(ctx, hostname, &InputUpdateInboundParseWebhook{
		URL:            "https://example.com/parse2",
		SecurityPolicy: String(""),
	})
	if err != nil {
		t.Fatalf("UpdateInboundParseWebhook: %v", err)
	}
	if updated.SecurityPolicy != "" {
		t.Errorf(`security_policy = "" で外れる想定です: %q`, updated.SecurityPolicy)
	}

	if _, err := client.UpdateInboundParseWebhook(ctx, hostname, &InputUpdateInboundParseWebhook{
		URL:            "https://example.com/parse2",
		SecurityPolicy: String(forcedID),
	}); err != nil {
		t.Fatalf("UpdateInboundParseWebhook: %v", err)
	}
	if err := client.DeleteWebhookSecurityPolicyWithOptions(ctx, forcedID, &InputDeleteWebhookSecurityPolicy{Force: true}); err != nil {
		t.Fatalf("DeleteWebhookSecurityPolicyWithOptions(force): %v", err)
	}
	if !eventually(t, 30*time.Second, func() bool {
		got, err = client.GetInboundParseWebhook(ctx, hostname)
		return err == nil && got.SecurityPolicy == ""
	}) {
		t.Errorf("force 削除後も Parse 設定に参照が残っています: %+v (err: %v)", got, err)
	}

	if err := deleteParse(); err != nil {
		t.Fatalf("DeleteInboundParseWebhook: %v", err)
	}
}
