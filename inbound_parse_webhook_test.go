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

func TestGetInboundParsetWebhooks(t *testing.T) {
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

	expected, err := client.GetInboundParsetWebhooks(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*InboundParsetWebhook{
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

func TestGetInboundParsetWebhooks_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetInboundParsetWebhooks(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetInboundParsetWebhook(t *testing.T) {
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

	expected, err := client.GetInboundParsetWebhook(context.TODO(), "bar.foo")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetInboundParsetWebhook{
		URL:       "https://example.com",
		Hostname:  "bar.foo",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetInboundParsetWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/bar.foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetInboundParsetWebhook(context.TODO(), "bar.foo")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCreateInboundParsetWebhook(t *testing.T) {
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

	expected, err := client.CreateInboundParsetWebhook(context.TODO(), &InputCreateInboundParsetWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateInboundParsetWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateInboundParsetWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateInboundParsetWebhook(context.TODO(), &InputCreateInboundParsetWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateInboundParsetWebhook(t *testing.T) {
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

	expected, err := client.UpdateInboundParsetWebhook(context.TODO(), "foo.bar", &InputUpdateInboundParsetWebhook{
		URL:       "https://example.com",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateInboundParsetWebhook{
		URL:       "https://example.com",
		Hostname:  "foo.bar",
		SpamCheck: false,
		SendRaw:   false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateInboundParsetWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateInboundParsetWebhook(context.TODO(), "foo.bar", &InputUpdateInboundParsetWebhook{
		URL:       "https://example.com",
		SpamCheck: false,
		SendRaw:   false,
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteInboundParsetWebhook(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteInboundParsetWebhook(context.TODO(), "foo.bar")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteInboundParsetWebhook_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/webhooks/parse/settings/foo.bar", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteInboundParsetWebhook(context.TODO(), "foo.bar")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}
