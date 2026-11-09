package sendgrid

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func TestGetInboundParseWebhooks(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"hostname": "bar.foo",
					"url": "https://example.com",
					"spam_check": false,
					"send_raw": false
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetInboundParseWebhooks(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*InboundParseWebhook{
		{
			URL:       "https://example.com",
			Hostname:  "bar.foo",
			SpamCheck: false,
			SendRaw:   false,
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetInboundParseWebhooks_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetInboundParseWebhooks(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetInboundParseWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/bar.foo", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
					"hostname": "bar.foo",
					"url": "https://example.com",
					"spam_check": false,
					"send_raw": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetInboundParseWebhook(context.TODO(), "bar.foo")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetInboundParseWebhook{
		URL:       "https://example.com",
		Hostname:  "bar.foo",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetInboundParseWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/bar.foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetInboundParseWebhook(context.TODO(), "bar.foo")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCreateInboundParseWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"url": "https://example.com",
			"hostname": "foo.bar",
			"spam_check": false,
			"send_raw": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateInboundParseWebhook(context.TODO(), &InputCreateInboundParseWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateInboundParseWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateInboundParseWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateInboundParseWebhook(context.TODO(), &InputCreateInboundParseWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateInboundParseWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"url": "https://example.com",
			"hostname": "foo.bar",
			"spam_check": false,
			"send_raw": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateInboundParseWebhook(context.TODO(), "foo.bar", &InputUpdateInboundParseWebhook{
		URL:       "https://example.com",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateInboundParseWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateInboundParseWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateInboundParseWebhook(context.TODO(), "foo.bar", &InputUpdateInboundParseWebhook{
		URL:       "https://example.com",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteInboundParseWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteInboundParseWebhook(context.TODO(), "foo.bar")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteInboundParseWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteInboundParseWebhook(context.TODO(), "foo.bar")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

// NewRequest Error Tests
func TestGetInboundParseWebhooks_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetInboundParseWebhooks(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetInboundParseWebhook_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetInboundParseWebhook(context.TODO(), "mail.example.com")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateInboundParseWebhook_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	input := &InputCreateInboundParseWebhook{
		Hostname: "mail.example.com",
		URL:      "https://example.com/parse",
	}
	_, err := client.CreateInboundParseWebhook(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateInboundParseWebhook_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	input := &InputUpdateInboundParseWebhook{
		URL:       "https://example.com/updated-parse",
		SpamCheck: true,
	}
	_, err := client.UpdateInboundParseWebhook(context.TODO(), "mail.example.com", input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteInboundParseWebhook_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	err := client.DeleteInboundParseWebhook(context.TODO(), "mail.example.com")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateInboundParseWebhook_SecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		if want := `{"url":"https://example.com","hostname":"foo.bar","spam_check":false,"send_raw":false,"security_policy":"policy-1"}`; strings.TrimSpace(string(body)) != want {
			t.Errorf("request body = %s, want %s", body, want)
		}
		if _, err := fmt.Fprint(w, `{
			"url": "https://example.com",
			"hostname": "foo.bar",
			"spam_check": false,
			"send_raw": false,
			"security_policy": "policy-1"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateInboundParseWebhook(context.TODO(), &InputCreateInboundParseWebhook{
		URL:            "https://example.com",
		Hostname:       "foo.bar",
		SecurityPolicy: "policy-1",
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	want := &OutputCreateInboundParseWebhook{
		URL:            "https://example.com",
		Hostname:       "foo.bar",
		SecurityPolicy: "policy-1",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetInboundParseWebhook_SecurityPolicy(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"hostname": "foo.bar", "url": "https://example.com", "security_policy": "policy-1"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetInboundParseWebhook(context.TODO(), "foo.bar")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if expected.SecurityPolicy != "policy-1" {
		t.Errorf("SecurityPolicy = %q, want %q", expected.SecurityPolicy, "policy-1")
	}
}

func TestUpdateInboundParseWebhook_SecurityPolicy(t *testing.T) {
	tests := []struct {
		name           string
		securityPolicy *string
		wantBody       string
	}{
		{
			name:           "nil leaves the policy unchanged",
			securityPolicy: nil,
			wantBody:       `{"url":"https://example.com","spam_check":false,"send_raw":false}`,
		},
		{
			name:           "attaches the policy",
			securityPolicy: String("policy-1"),
			wantBody:       `{"url":"https://example.com","spam_check":false,"send_raw":false,"security_policy":"policy-1"}`,
		},
		{
			name:           "sends an empty string",
			securityPolicy: String(""),
			wantBody:       `{"url":"https://example.com","spam_check":false,"send_raw":false,"security_policy":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "PATCH")
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				if got := strings.TrimSpace(string(body)); got != tt.wantBody {
					t.Errorf("request body = %s, want %s", got, tt.wantBody)
				}
				if _, err := fmt.Fprint(w, `{"url": "https://example.com", "hostname": "foo.bar", "security_policy": "policy-1"}`); err != nil {
					t.Fatal(err)
				}
			})

			expected, err := client.UpdateInboundParseWebhook(context.TODO(), "foo.bar", &InputUpdateInboundParseWebhook{
				URL:            "https://example.com",
				SecurityPolicy: tt.securityPolicy,
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if expected.SecurityPolicy != "policy-1" {
				t.Errorf("SecurityPolicy = %q, want %q", expected.SecurityPolicy, "policy-1")
			}
		})
	}
}
