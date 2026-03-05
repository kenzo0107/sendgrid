package sendgrid

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestAddSuppressionsToGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/123/suppressions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := w.Write([]byte(`{
			"recipient_emails": [
				"test@example.com"
			]
		}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	input := &InputAddSuppressionsToGroup{
		RecipientEmails: []string{"test@example.com"},
	}
	expected, _ := client.AddSuppressionsToGroup(context.Background(), 123, input)

	want := &OutputAddSuppressionsToGroup{
		RecipientEmails: []string{"test@example.com"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddSuppressionsToGroup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/123/suppressions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddSuppressionsToGroup(context.Background(), 123, &InputAddSuppressionsToGroup{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestAddSuppressionsToGroup_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.AddSuppressionsToGroup(context.Background(), 123, &InputAddSuppressionsToGroup{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}

func TestGetSuppressionGroupsByEmail(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := w.Write([]byte(`{
			"suppressions": [
				{
					"description": "Test Description",
					"id": 123,
					"is_default": false,
					"name": "Test Group",
					"suppressed": true
				}
			]
		}`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	expected, err := client.GetSuppressionGroupsByEmail(context.Background(), "test@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := &OutputGetSuppressionGroupsByEmail{
		Suppressions: []*ASMSuppressionGroup{
			{
				Description: "Test Description",
				ID:          123,
				IsDefault:   false,
				Name:        "Test Group",
				Suppressed:  true,
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSuppressionGroupsByEmail_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSuppressionGroupsByEmail(context.Background(), "test@example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetSuppressionGroupsByEmail_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSuppressionGroupsByEmail(context.Background(), "test@example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}

func TestGetSuppressionsForSuppressionGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := w.Write([]byte(`[
			"test@example.com"
		]`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	expected, err := client.GetSuppressionsForSuppressionGroup(context.Background(), "group_id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"test@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSuppressionsForSuppressionGroup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSuppressionsForSuppressionGroup(context.Background(), "group_id")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetSuppressionsForSuppressionGroup_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSuppressionsForSuppressionGroup(context.Background(), "group_id")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}

func TestSearchForSuppressionsWithinGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := w.Write([]byte(`[
			"test@example.com"
		]`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	expected, err := client.SearchForSuppressionsWithinGroup(context.Background(), "group_id", &InputSearchGroupSuppressions{
		RecipientEmails: []string{"test@example.com"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"test@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestSearchForSuppressionsWithinGroup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.SearchForSuppressionsWithinGroup(context.Background(), "group_id", &InputSearchGroupSuppressions{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestSearchForSuppressionsWithinGroup_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.SearchForSuppressionsWithinGroup(context.Background(), "group_id", &InputSearchGroupSuppressions{
		RecipientEmails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}

func TestGetSuppressions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := w.Write([]byte(`[
			{
				"email": "test@example.com",
				"group_id": 123
			}
		]`)); err != nil {
			t.Fatalf("failed to write response: %v", err)
		}
	})

	expected, err := client.GetSuppressions(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []*Suppression{
		{
			Email:   "test@example.com",
			GroupID: 123,
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSuppressions_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/suppressions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSuppressions(context.Background())
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetSuppressions_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSuppressions(context.Background())
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}

func TestDeleteSuppressionFromGroup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions/email@example.com", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteSuppressionFromGroup(context.Background(), "group_id", "email@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteSuppressionFromGroup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/asm/groups/group_id/suppressions/email@example.com", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteSuppressionFromGroup(context.Background(), "group_id", "email@example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestDeleteSuppressionFromGroup_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteSuppressionFromGroup(context.Background(), "group_id", "email@example.com")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	client.baseURL = originalBaseURL
}
