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

func TestGetBounceSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"enabled": true,
			"soft_bounces": 100,
			"hard_bounces": 200
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetBounceSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetBounceSettings{
		Enabled:     true,
		SoftBounces: 100,
		HardBounces: 200,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetBounceSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetBounceSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateBounceSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"enabled": true,
			"soft_bounces": 150,
			"hard_bounces": 250
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputUpdateBounceSettings{
		Enabled:     true,
		SoftBounces: 150,
		HardBounces: 250,
	}

	expected, err := client.UpdateBounceSettings(context.TODO(), input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateBounceSettings{
		Enabled:     true,
		SoftBounces: 150,
		HardBounces: 250,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateBounceSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputUpdateBounceSettings{
		Enabled:     true,
		SoftBounces: 150,
		HardBounces: 250,
	}

	_, err := client.UpdateBounceSettings(context.TODO(), input)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

// NewRequest Error Tests
func TestGetBounceSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetBounceSettings(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateBounceSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateBounceSettings{
		Enabled:     true,
		SoftBounces: 150,
		HardBounces: 250,
	}

	_, err := client.UpdateBounceSettings(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
