package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestCreateWebhookSecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := fmt.Fprint(w, `{
			"policy": {
				"id": "policy-1",
				"name": "test policy",
				"oauth": {
					"client_id": "client-id",
					"token_url": "https://example.com/token",
					"scopes": ["scope1"]
				},
				"signature": {
					"public_key": "public-key"
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateWebhookSecurityPolicy(context.TODO(), &InputCreateWebhookSecurityPolicy{
		Name: "test policy",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateWebhookSecurityPolicy{
		Policy: &WebhookSecurityPolicy{
			ID:   "policy-1",
			Name: "test policy",
			OAuth: &WebhookSecurityPolicyOAuth{
				ClientID: "client-id",
				TokenURL: "https://example.com/token",
				Scopes:   []string{"scope1"},
			},
			Signature: &WebhookSecurityPolicySignature{
				PublicKey: "public-key",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateWebhookSecurityPolicy_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateWebhookSecurityPolicy(context.TODO(), &InputCreateWebhookSecurityPolicy{Name: "test policy"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateWebhookSecurityPolicy_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CreateWebhookSecurityPolicy(context.TODO(), &InputCreateWebhookSecurityPolicy{Name: "test policy"})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetWebhookSecurityPolicies(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"policies": [
				{
					"id": "policy-1",
					"name": "test policy"
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetWebhookSecurityPolicies(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetWebhookSecurityPolicies{
		Policies: []*WebhookSecurityPolicy{
			{ID: "policy-1", Name: "test policy"},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetWebhookSecurityPolicies_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWebhookSecurityPolicies(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWebhookSecurityPolicies_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetWebhookSecurityPolicies(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetWebhookSecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"policy": {
				"id": "policy-1",
				"name": "test policy"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetWebhookSecurityPolicy(context.TODO(), "policy-1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetWebhookSecurityPolicy{
		Policy: &WebhookSecurityPolicy{ID: "policy-1", Name: "test policy"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetWebhookSecurityPolicy_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetWebhookSecurityPolicy(context.TODO(), "policy-1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetWebhookSecurityPolicy_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetWebhookSecurityPolicy(context.TODO(), "policy-1")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateWebhookSecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"policy": {
				"id": "policy-1",
				"name": "updated policy"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateWebhookSecurityPolicy(context.TODO(), "policy-1", &InputUpdateWebhookSecurityPolicy{Name: "updated policy"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateWebhookSecurityPolicy{
		Policy: &WebhookSecurityPolicy{ID: "policy-1", Name: "updated policy"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateWebhookSecurityPolicy_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateWebhookSecurityPolicy(context.TODO(), "policy-1", &InputUpdateWebhookSecurityPolicy{Name: "updated policy"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateWebhookSecurityPolicy_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateWebhookSecurityPolicy(context.TODO(), "policy-1", &InputUpdateWebhookSecurityPolicy{Name: "updated policy"})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteWebhookSecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if err := client.DeleteWebhookSecurityPolicy(context.TODO(), "policy-1"); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteWebhookSecurityPolicy_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/security/policies/policy-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.DeleteWebhookSecurityPolicy(context.TODO(), "policy-1"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteWebhookSecurityPolicy_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteWebhookSecurityPolicy(context.TODO(), "policy-1")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
