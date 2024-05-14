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

func TestGetEventWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "172af0f9-f165-4172-8a8c-25c16e8e8e25",
			"enabled": true,
			"url": "http://www.example.com",
			"group_resubscribe": true,
			"delivered": true,
			"group_unsubscribe": true,
			"spam_report": true,
			"bounce": true,
			"deferred": true,
			"unsubscribe": true,
			"processed": true,
			"open": true,
			"click": true,
			"dropped": true,
			"friendly_name": "example_name"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetEventWebhook{
		ID:               "172af0f9-f165-4172-8a8c-25c16e8e8e25",
		Enabled:          true,
		URL:              "http://www.example.com",
		GroupResubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "example_name",
		OAuthClientID:    "",
		OAuthTokenURL:    "",
		PublicKey:        "",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetEventWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetEventWebhooks(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/all", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"max_allowed": 10,
			"webhooks": [{
				"id": "172af0f9-f165-4172-8a8c-25c16e8e8e25",
				"enabled": true,
				"url": "http://www.example.com",
				"group_resubscribe": true,
				"delivered": true,
				"group_unsubscribe": true,
				"spam_report": true,
				"bounce": true,
				"deferred": true,
				"unsubscribe": true,
				"processed": true,
				"open": true,
				"click": true,
				"dropped": true,
				"friendly_name": "example_name"
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetEventWebhooks(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetEventWebhooks{
		MaxAllowed: 10,
		Webhooks: []*EventWebhook{
			{
				ID:               "172af0f9-f165-4172-8a8c-25c16e8e8e25",
				Enabled:          true,
				URL:              "http://www.example.com",
				GroupResubscribe: true,
				Delivered:        true,
				GroupUnsubscribe: true,
				SpamReport:       true,
				Bounce:           true,
				Deferred:         true,
				Unsubscribe:      true,
				Processed:        true,
				Open:             true,
				Click:            true,
				Dropped:          true,
				FriendlyName:     "example_name",
				OAuthClientID:    "",
				OAuthTokenURL:    "",
				PublicKey:        "",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetEventWebhooks_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/all", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetEventWebhooks(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCreateEventWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "172af0f9-f165-4172-8a8c-25c16e8e8e25",
			"url": "http://www.example.com",
			"enabled": true,
			"group_resubscribe": true,
			"delivered": true,
			"group_unsubscribe": true,
			"spam_report": true,
			"bounce": true,
			"deferred": true,
			"unsubscribe": true,
			"processed": true,
			"open": true,
			"click": true,
			"dropped": true,
			"friendly_name": "example_name"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputCreateEventWebhook{
		Enabled:          true,
		URL:              "http://www.example.com",
		GroupResubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "example_name",
	}
	expected, err := client.CreateEventWebhook(context.TODO(), input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateEventWebhook{
		ID:               "172af0f9-f165-4172-8a8c-25c16e8e8e25",
		URL:              "http://www.example.com",
		Enabled:          true,
		GroupResubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "example_name",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateEventWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputCreateEventWebhook{}
	_, err := client.CreateEventWebhook(context.TODO(), input)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateEventWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "172af0f9-f165-4172-8a8c-25c16e8e8e25",
			"url": "http://www.example.com",
			"enabled": true,
			"group_resubscribe": true,
			"delivered": true,
			"group_unsubscribe": true,
			"spam_report": true,
			"bounce": true,
			"deferred": true,
			"unsubscribe": true,
			"processed": true,
			"open": true,
			"click": true,
			"dropped": true,
			"friendly_name": "example_name"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputUpdateEventWebhook{
		Enabled:          true,
		URL:              "http://www.example.com",
		GroupResubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "example_name",
	}
	expected, err := client.UpdateEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateEventWebhook{
		ID:               "172af0f9-f165-4172-8a8c-25c16e8e8e25",
		URL:              "http://www.example.com",
		Enabled:          true,
		GroupResubscribe: true,
		Delivered:        true,
		GroupUnsubscribe: true,
		SpamReport:       true,
		Bounce:           true,
		Deferred:         true,
		Unsubscribe:      true,
		Processed:        true,
		Open:             true,
		Click:            true,
		Dropped:          true,
		FriendlyName:     "example_name",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateEventWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputUpdateEventWebhook{}
	_, err := client.UpdateEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25", input)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteEventWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteEventWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/event/settings/172af0f9-f165-4172-8a8c-25c16e8e8e25", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteEventWebhook(context.TODO(), "172af0f9-f165-4172-8a8c-25c16e8e8e25")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}
