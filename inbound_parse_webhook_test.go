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
