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

func TestAddRecipientAddressesToGlobalSuppressions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"recipient_emails": ["test@example.com", "test2@example.com"]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.AddRecipientAddressesToGlobalSuppressions(context.TODO(), &InputAddRecipientAddressesToGlobalSuppressions{
		RecipientEmails: []string{"test@example.com", "test2@example.com"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAddRecipientAddressesToGlobalSuppressions{
		RecipientEmails: []string{"test@example.com", "test2@example.com"},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestAddRecipientAddressesToGlobalSuppressions_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddRecipientAddressesToGlobalSuppressions(context.TODO(), &InputAddRecipientAddressesToGlobalSuppressions{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestAddRecipientAddressesToGlobalSuppressions_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.AddRecipientAddressesToGlobalSuppressions(context.TODO(), &InputAddRecipientAddressesToGlobalSuppressions{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetGlobalSuppression(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"recipient_email": "test@example.com"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetGlobalSuppression(context.TODO(), "test@example.com")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetGlobalSuppression{
		RecipientEmail: "test@example.com",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetGlobalSuppression_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetGlobalSuppression(context.TODO(), "test@example.com")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetGlobalSuppression_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetGlobalSuppression(context.TODO(), "test@example.com")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetGlobalSuppressions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/suppression/unsubscribes", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[{
			"created": 1772349124,
			"email": "test@example.com"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetGlobalSuppressions(context.TODO(), &InputGetGlobalSuppressions{
		StartTime: 1772349124,
		EndTime:   1772435524,
		Offset:    10,
		Limit:     5,
		Email:     "test@example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*GlobalUnsubscribe{
		{
			Created: 1772349124,
			Email:   "test@example.com",
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetGlobalSuppressions_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/suppression/unsubscribes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetGlobalSuppressions(context.TODO(), &InputGetGlobalSuppressions{})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetGlobalSuppressions_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetGlobalSuppressions(context.TODO(), &InputGetGlobalSuppressions{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteGlobalSuppression(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteGlobalSuppression(context.TODO(), "test@example.com")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteGlobalSuppression_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/global/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteGlobalSuppression(context.TODO(), "test@example.com")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteGlobalSuppression_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteGlobalSuppression(context.TODO(), "test@example.com")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
