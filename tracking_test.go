package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestGetTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result":[
				{
					"title":"Click Tracking",
					"enabled":false,
					"name":"click",
					"description":"Overwrites every link to track every click in emails."
				},
				{
					"title":"Google Analytics",
					"enabled":false,
					"name":"google_analytics",
					"description":"Track your conversion rates and ROI with Google Analytics."
				},
				{
					"title":"Open Tracking",
					"enabled":true,
					"name":"open",
					"description":"Appends an invisible image to HTML emails to track emails that have been opened."
				},
				{
					"title":"Subscription Tracking",
					"enabled":false,
					"name":"subscription",
					"description":"Adds unsubscribe links to the bottom of the text and HTML emails.  Future emails won't be delivered to unsubscribed users."
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTrackingSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTrackingSettings{
		Result: []*ResultGetTrackingSettings{
			{
				Name:        "click",
				Title:       "Click Tracking",
				Description: "Overwrites every link to track every click in emails.",
				Enabled:     false,
			},
			{
				Name:        "google_analytics",
				Title:       "Google Analytics",
				Description: "Track your conversion rates and ROI with Google Analytics.",
				Enabled:     false,
			},
			{
				Name:        "open",
				Title:       "Open Tracking",
				Description: "Appends an invisible image to HTML emails to track emails that have been opened.",
				Enabled:     true,
			},
			{
				Name:        "subscription",
				Title:       "Subscription Tracking",
				Description: "Adds unsubscribe links to the bottom of the text and HTML emails.  Future emails won't be delivered to unsubscribed users.",
				Enabled:     false,
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTrackingSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetClickTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/click", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true,
			"enable_text":true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetClickTrackingSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetClickTrackingSettings{
		EnableText: true,
		Enabled:    true,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetClickTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/click", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetClickTrackingSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateClickTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/click", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true,
			"enable_text":true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateClickTrackingSettings(context.TODO(), &InputUpdateClickTrackingSettings{
		Enabled: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateClickTrackingSettings{
		EnableText: true,
		Enabled:    true,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateClickTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/click", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateClickTrackingSettings(context.TODO(), &InputUpdateClickTrackingSettings{
		Enabled: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetOpenTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/open", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetOpenTrackingSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetOpenTrackingSettings{
		Enabled: true,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetOpenTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/open", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetOpenTrackingSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateOpenTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/open", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateOpenTrackingSettings(context.TODO(), &InputUpdateOpenTrackingSettings{
		Enabled: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateOpenTrackingSettings{
		Enabled: true,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateOpenTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/open", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateOpenTrackingSettings(context.TODO(), &InputUpdateOpenTrackingSettings{
		Enabled: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetGoogleAnalyticsSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/google_analytics", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true,
			"utm_source":"sendgrid.com",
			"utm_medium":"email",
			"utm_term":"",
			"utm_content":"",
			"utm_campaign":"sendgrid-email"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetGoogleAnalyticsSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetGoogleAnalyticsSettings{
		Enabled:     true,
		UTMSource:   "sendgrid.com",
		UTMMedium:   "email",
		UTMTerm:     "",
		UTMContent:  "",
		UTMCampaign: "sendgrid-email",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetGoogleAnalyticsSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/google_analytics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetGoogleAnalyticsSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateGoogleAnalyticsSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/google_analytics", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":true,
			"utm_source":"sendgrid.com",
			"utm_medium":"email",
			"utm_term":"",
			"utm_content":"",
			"utm_campaign":"sendgrid-email"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateGoogleAnalyticsSettings(context.TODO(), &InputUpdateGoogleAnalyticsSettings{
		Enabled:     true,
		UTMSource:   "sendgrid.com",
		UTMMedium:   "email",
		UTMTerm:     "",
		UTMContent:  "",
		UTMCampaign: "sendgrid-email",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateGoogleAnalyticsSettings{
		Enabled:     true,
		UTMSource:   "sendgrid.com",
		UTMMedium:   "email",
		UTMTerm:     "",
		UTMContent:  "",
		UTMCampaign: "sendgrid-email",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateGoogleAnalyticsSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/google_analytics", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateGoogleAnalyticsSettings(context.TODO(), &InputUpdateGoogleAnalyticsSettings{
		Enabled:     true,
		UTMSource:   "sendgrid.com",
		UTMMedium:   "email",
		UTMTerm:     "",
		UTMContent:  "",
		UTMCampaign: "sendgrid-email",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSubscriptionTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/subscription", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":false,
			"html_content":"",
			"landing":"",
			"plain_content":"",
			"replace":"",
			"url":null
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubscriptionTrackingSettings(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSubscriptionTrackingSettings{
		Enabled:      false,
		HTMLContent:  "",
		Landing:      "",
		PlainContent: "",
		Replace:      "",
		URL:          "",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSubscriptionTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/subscription", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubscriptionTrackingSettings(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSubscriptionTrackingSettings(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/subscription", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"enabled":false,
			"html_content":"",
			"landing":"",
			"plain_content":"",
			"replace":"",
			"url":null
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateSubscriptionTrackingSettings(context.TODO(), &InputUpdateSubscriptionTrackingSettings{
		Enabled: false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateSubscriptionTrackingSettings{
		Enabled:      false,
		HTMLContent:  "",
		Landing:      "",
		PlainContent: "",
		Replace:      "",
		URL:          "",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateSubscriptionTrackingSettings_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tracking_settings/subscription", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateSubscriptionTrackingSettings(context.TODO(), &InputUpdateSubscriptionTrackingSettings{
		Enabled: false,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
