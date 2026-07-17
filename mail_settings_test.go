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

func TestGetMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"result": [
				{"title": "Address Whitelist", "enabled": true, "name": "address_whitelist", "description": "desc"}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetMailSettings(context.TODO(), &InputGetMailSettings{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetMailSettings{
		Result: []*MailSetting{
			{Title: "Address Whitelist", Enabled: true, Name: "address_whitelist", Description: "desc"},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetMailSettings(context.TODO(), &InputGetMailSettings{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetMailSettings(context.TODO(), &InputGetMailSettings{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetAddressWhitelistMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/address_whitelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"enabled": true, "list": ["email1@example.com"]}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetAddressWhitelistMailSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAddressWhitelistMailSettings{Enabled: true, List: []string{"email1@example.com"}}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetAddressWhitelistMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/address_whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAddressWhitelistMailSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetAddressWhitelistMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetAddressWhitelistMailSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateAddressWhitelistMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/address_whitelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{"enabled": true, "list": ["email1@example.com"]}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateAddressWhitelistMailSettings(context.TODO(), &InputUpdateAddressWhitelistMailSettings{
		Enabled: true,
		List:    []string{"email1@example.com"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateAddressWhitelistMailSettings{Enabled: true, List: []string{"email1@example.com"}}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateAddressWhitelistMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/address_whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateAddressWhitelistMailSettings(context.TODO(), &InputUpdateAddressWhitelistMailSettings{Enabled: true})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateAddressWhitelistMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateAddressWhitelistMailSettings(context.TODO(), &InputUpdateAddressWhitelistMailSettings{Enabled: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetFooterMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/footer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"enabled": true, "html_content": "<p>html</p>", "plain_content": "plain"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetFooterMailSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetFooterMailSettings{Enabled: true, HTMLContent: "<p>html</p>", PlainContent: "plain"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetFooterMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/footer", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetFooterMailSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetFooterMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetFooterMailSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateFooterMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/footer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{"enabled": true, "html_content": "<p>html</p>", "plain_content": "plain"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateFooterMailSettings(context.TODO(), &InputUpdateFooterMailSettings{
		Enabled:      true,
		HTMLContent:  "<p>html</p>",
		PlainContent: "plain",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateFooterMailSettings{Enabled: true, HTMLContent: "<p>html</p>", PlainContent: "plain"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateFooterMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/footer", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateFooterMailSettings(context.TODO(), &InputUpdateFooterMailSettings{Enabled: true})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateFooterMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateFooterMailSettings(context.TODO(), &InputUpdateFooterMailSettings{Enabled: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetForwardBounceMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_bounce", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"enabled": true, "email": "email@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetForwardBounceMailSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetForwardBounceMailSettings{Enabled: true, Email: "email@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetForwardBounceMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_bounce", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetForwardBounceMailSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetForwardBounceMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetForwardBounceMailSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateForwardBounceMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_bounce", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{"enabled": true, "email": "email@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateForwardBounceMailSettings(context.TODO(), &InputUpdateForwardBounceMailSettings{
		Enabled: true,
		Email:   "email@example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateForwardBounceMailSettings{Enabled: true, Email: "email@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateForwardBounceMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_bounce", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateForwardBounceMailSettings(context.TODO(), &InputUpdateForwardBounceMailSettings{Enabled: true})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateForwardBounceMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateForwardBounceMailSettings(context.TODO(), &InputUpdateForwardBounceMailSettings{Enabled: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetForwardSpamMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_spam", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"enabled": true, "email": "email@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetForwardSpamMailSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetForwardSpamMailSettings{Enabled: true, Email: "email@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetForwardSpamMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_spam", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetForwardSpamMailSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetForwardSpamMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetForwardSpamMailSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateForwardSpamMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_spam", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{"enabled": true, "email": "email@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateForwardSpamMailSettings(context.TODO(), &InputUpdateForwardSpamMailSettings{
		Enabled: true,
		Email:   "email@example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateForwardSpamMailSettings{Enabled: true, Email: "email@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateForwardSpamMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/forward_spam", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateForwardSpamMailSettings(context.TODO(), &InputUpdateForwardSpamMailSettings{Enabled: true})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateForwardSpamMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateForwardSpamMailSettings(context.TODO(), &InputUpdateForwardSpamMailSettings{Enabled: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetTemplateMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/template", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"enabled": false, "html_content": "<p><% body %>Example</p>\n"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTemplateMailSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTemplateMailSettings{HTMLContent: "<p><% body %>Example</p>\n"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTemplateMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/template", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTemplateMailSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetTemplateMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetTemplateMailSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateTemplateMailSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/template", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{"enabled": true, "html_content": "<p>html</p>"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTemplateMailSettings(context.TODO(), &InputUpdateTemplateMailSettings{
		Enabled:     true,
		HTMLContent: "<p>html</p>",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateTemplateMailSettings{Enabled: true, HTMLContent: "<p>html</p>"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTemplateMailSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/template", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateTemplateMailSettings(context.TODO(), &InputUpdateTemplateMailSettings{Enabled: true})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateTemplateMailSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateTemplateMailSettings(context.TODO(), &InputUpdateTemplateMailSettings{Enabled: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
