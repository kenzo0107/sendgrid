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

func TestGetSSOIntegration(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("si"); got != "true" {
			t.Errorf("si query = %q, want %q", got, "true")
		}
		if _, err := fmt.Fprint(w, `{
			"id": "abcdef",
			"name": "dummy",
			"enabled": true,
			"signin_url": "https://example.com/signin",
			"signout_url": "https://example.com/signout",
			"entity_id": "https://example.com/entity",
			"completed_integration": true,
			"last_updated": 1586137600,
			"single_signon_url": "https://example.com/sso",
			"audience_url": "https://example.com/audience"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSSOIntegration(context.TODO(), "abcdef")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSSOIntegration{
		ID:                   "abcdef",
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
		LastUpdated:          1586137600,
		SingleSignonURL:      "https://example.com/sso",
		AudienceURL:          "https://example.com/audience",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSSOIntegration_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSSOIntegration(context.TODO(), "abcdef")
	if err == nil {
		t.Error("Expected an error but got nil")
		return
	}
}

func TestGetSSOIntegrations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("si"); got != "true" {
			t.Errorf("si query = %q, want %q", got, "true")
		}
		if _, err := fmt.Fprint(w, `[
			{
				"id": "abcdef",
				"name": "dummy",
				"enabled": true,
				"signin_url": "https://example.com/signin",
				"signout_url": "https://example.com/signout",
				"entity_id": "https://example.com/entity",
				"completed_integration": true,
				"last_updated": 1586137600,
				"single_signon_url": "https://example.com/sso",
				"audience_url": "https://example.com/audience"
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSSOIntegrations(context.TODO(), &InputGetSSOIntegrations{Si: true})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*SSOIntegration{
		{
			ID:                   "abcdef",
			Name:                 "dummy",
			Enabled:              true,
			SigninURL:            "https://example.com/signin",
			SignoutURL:           "https://example.com/signout",
			EntityID:             "https://example.com/entity",
			CompletedIntegration: true,
			LastUpdated:          1586137600,
			SingleSignonURL:      "https://example.com/sso",
			AudienceURL:          "https://example.com/audience",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSSOIntegrations_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSSOIntegrations(context.TODO(), &InputGetSSOIntegrations{Si: true})
	if err == nil {
		t.Error("Expected an error but got nil")
		return
	}
}

func TestCreateSSOIntegration(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "abcdef",
			"name": "dummy",
			"enabled": true,
			"signin_url": "https://example.com/signin",
			"signout_url": "https://example.com/signout",
			"entity_id": "https://example.com/entity",
			"completed_integration": true,
			"last_updated": 1586137600,
			"single_signon_url": "https://example.com/sso",
			"audience_url": "https://example.com/audience"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateSSOIntegration(context.TODO(), &InputCreateSSOIntegration{
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateSSOIntegration{
		ID:                   "abcdef",
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
		LastUpdated:          1586137600,
		SingleSignonURL:      "https://example.com/sso",
		AudienceURL:          "https://example.com/audience",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateSSOIntegration_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateSSOIntegration(context.TODO(), &InputCreateSSOIntegration{
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
	})
	if err == nil {
		t.Error("Expected an error but got nil")
		return
	}
}

func TestUpdateSSOIntegration(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("si"); got != "true" {
			t.Errorf("si query = %q, want %q", got, "true")
		}
		if _, err := fmt.Fprint(w, `{
			"id": "abcdef",
			"name": "dummy",
			"enabled": true,
			"signin_url": "https://example.com/signin",
			"signout_url": "https://example.com/signout",
			"entity_id": "https://example.com/entity",
			"completed_integration": true,
			"last_updated": 1586137600,
			"single_signon_url": "https://example.com/sso",
			"audience_url": "https://example.com/audience"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateSSOIntegration(context.TODO(), "abcdef", &InputUpdateSSOIntegration{
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateSSOIntegration{
		ID:                   "abcdef",
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
		LastUpdated:          1586137600,
		SingleSignonURL:      "https://example.com/sso",
		AudienceURL:          "https://example.com/audience",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateSSOIntegration_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateSSOIntegration(context.TODO(), "abcdef", &InputUpdateSSOIntegration{
		Name:                 "dummy",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "https://example.com/entity",
		CompletedIntegration: true,
	})
	if err == nil {
		t.Error("Expected an error but got nil")
		return
	}
}

func TestDeleteSSOIntegration(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteSSOIntegration(context.TODO(), "abcdef")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteSSOIntegration_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sso/integrations/abcdef", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteSSOIntegration(context.TODO(), "abcdef")
	if err == nil {
		t.Error("Expected an error but got nil")
		return
	}
}

// NewRequest Error Tests
func TestGetSSOIntegration_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSSOIntegration(context.TODO(), "test-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSSOIntegrations_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputGetSSOIntegrations{
		Si: true,
	}
	_, err := client.GetSSOIntegrations(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateSSOIntegration_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputCreateSSOIntegration{
		Name:                 "test-integration",
		Enabled:              true,
		SigninURL:            "https://example.com/signin",
		SignoutURL:           "https://example.com/signout",
		EntityID:             "test-entity",
		CompletedIntegration: false,
	}
	_, err := client.CreateSSOIntegration(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateSSOIntegration_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateSSOIntegration{
		Name:       "updated-integration",
		Enabled:    true,
		SigninURL:  "https://example.com/signin-updated",
		SignoutURL: "https://example.com/signout-updated",
		EntityID:   "updated-entity",
	}
	_, err := client.UpdateSSOIntegration(context.TODO(), "test-id", input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteSSOIntegration_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteSSOIntegration(context.TODO(), "test-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
