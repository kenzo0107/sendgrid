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

func TestGetPartnerSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/partner_settings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"title": "Partner title",
					"enabled": true,
					"name": "partner_name",
					"description": "A description of a partner."
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetPartnerSettings(context.TODO(), &InputGetPartnerSettings{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetPartnerSettings{
		Result: []*PartnerSetting{
			{
				Title:       "Partner title",
				Enabled:     true,
				Name:        "partner_name",
				Description: "A description of a partner.",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetPartnerSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/partner_settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetPartnerSettings(context.TODO(), &InputGetPartnerSettings{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetPartnerSettings_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetPartnerSettings(context.TODO(), &InputGetPartnerSettings{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
