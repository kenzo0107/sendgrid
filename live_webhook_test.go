//go:build live

package sendgrid

import (
	"context"
	"net/http"
	"slices"
	"testing"
	"time"
)

// Webhook Security Policy は更新・削除の結果が GET に反映されるまで数秒かかるため、1 回の再取得では判定できません。
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

	policies, err := client.GetWebhookSecurityPolicies(ctx)
	if err != nil {
		t.Fatalf("GetWebhookSecurityPolicies: %v", err)
	}
	assertNoUnknownFields[OutputGetWebhookSecurityPolicies](t, rec)
	if !slices.ContainsFunc(policies.Policies, func(p *WebhookSecurityPolicy) bool { return p.ID == id }) {
		t.Errorf("一覧に作成した policy %s がありません", id)
	}

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
