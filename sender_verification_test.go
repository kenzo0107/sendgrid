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

func TestGetVerifiedSenders(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("limit", "1")
		q.Set("id", "12345678")
		q.Set("lastSeenID", "1000")
		r.URL.RawQuery = q.Encode()

		if _, err := fmt.Fprint(w, `{"results":[
			{
				"id": 12345678,
				"nickname": "dummy",
				"from_email": "dummy@example.com",
				"from_name": "dummy",
				"reply_to": "dummy@example.com",
				"reply_to_name": "dummy",
				"address": "dummy",
				"address2": "",
				"state": "",
				"city": "dummy",
				"zip": "",
				"country": "Japan",
				"verified": false,
				"locked": false
			}
		]}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetVerifiedSenders(context.TODO(), &InputGetVerifiedSenders{
		Limit:      1,
		ID:         12345678,
		LastSeenID: 1000,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	want := []*VerifiedSender{
		{
			ID:          12345678,
			Nickname:    "dummy",
			FromEmail:   "dummy@example.com",
			FromName:    "dummy",
			ReplyTo:     "dummy@example.com",
			ReplyToName: "dummy",
			Address:     "dummy",
			Address2:    "",
			State:       "",
			City:        "dummy",
			Zip:         "",
			Country:     "Japan",
			Verified:    false,
			Locked:      false,
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetVerifiedSenders_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetVerifiedSenders(context.TODO(), &InputGetVerifiedSenders{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateVerifiedSenderRequest(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 12345678,
			"nickname": "dummy",
			"from_email": "dummy@example.com",
			"from_name": "dummy",
			"reply_to": "dummy@example.com",
			"reply_to_name": "dummy",
			"address": "dummy",
			"address2": "",
			"state": "",
			"city": "dummy",
			"zip": "",
			"country": "Japan",
			"verified": false,
			"locked": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateVerifiedSenderRequest(context.TODO(), &InputCreateVerifiedSenderRequest{
		Nickname:    "dummy",
		FromEmail:   "dummy@example.com",
		FromName:    "dummy",
		ReplyTo:     "dummy@example.com",
		ReplyToName: "dummy",
		Address:     "dummy",
		Address2:    "",
		State:       "",
		City:        "dummy",
		Zip:         "",
		Country:     "Japan",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateVerifiedSenderRequest{
		ID:          12345678,
		Nickname:    "dummy",
		FromEmail:   "dummy@example.com",
		FromName:    "dummy",
		ReplyTo:     "dummy@example.com",
		ReplyToName: "dummy",
		Address:     "dummy",
		Address2:    "",
		State:       "",
		City:        "dummy",
		Zip:         "",
		Country:     "Japan",
		Verified:    false,
		Locked:      false,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateVerifiedSenderRequest_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateVerifiedSenderRequest(context.TODO(), &InputCreateVerifiedSenderRequest{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestResendVerifiedSenderRequest(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/resend/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	err := client.ResendVerifiedSenderRequest(context.TODO(), 12345678)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestResendVerifiedSenderRequest_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/resend/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.ResendVerifiedSenderRequest(context.TODO(), 12345678)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestVerifySenderRequest(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/verify/abcdefghijklmn", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.VerifySenderRequest(context.TODO(), "abcdefghijklmn")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestVerifySenderRequest_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/verify/abcdefghijklmn", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.VerifySenderRequest(context.TODO(), "abcdefghijklmn")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateVerifiedSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/12345678", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 12345678,
			"nickname": "dummy",
			"from_email": "dummy@example.com",
			"from_name": "dummy",
			"reply_to": "dummy@example.com",
			"reply_to_name": "dummy",
			"address": "dummy",
			"address2": "",
			"state": "",
			"city": "dummy",
			"zip": "",
			"country": "Japan",
			"verified": false,
			"locked": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateVerifiedSender(context.TODO(), 12345678, &InputUpdateVerifiedSender{
		Nickname:    "dummy",
		FromEmail:   "dummy@example.com",
		FromName:    "dummy",
		ReplyTo:     "dummy@example.com",
		ReplyToName: "dummy",
		Address:     "dummy",
		Address2:    "",
		State:       "",
		City:        "dummy",
		Zip:         "",
		Country:     "Japan",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateVerifiedSender{
		ID:          12345678,
		Nickname:    "dummy",
		FromEmail:   "dummy@example.com",
		FromName:    "dummy",
		ReplyTo:     "dummy@example.com",
		ReplyToName: "dummy",
		Address:     "dummy",
		Address2:    "",
		State:       "",
		City:        "dummy",
		Zip:         "",
		Country:     "Japan",
		Verified:    false,
		Locked:      false,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateVerifiedSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateVerifiedSender(context.TODO(), 12345678, &InputUpdateVerifiedSender{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteVerifiedSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteVerifiedSender(context.TODO(), 12345678)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteVerifiedSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteVerifiedSender(context.TODO(), 12345678)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCompletedStepsVerifiedSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/steps_completed", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"results": {
				"sender_verified": false,
				"domain_verified": false
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CompletedStepsVerifiedSender(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &CompletedStepsVerifiedSender{
		SenderVerified: false,
		DomainVerified: false,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestCompletedStepsVerifiedSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/steps_completed", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CompletedStepsVerifiedSender(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSenderVerificationDomainWarnList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/domains", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"results": {
				"sender_verified": false,
				"domain_verified": false
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSenderVerificationDomainWarnList(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &CompletedStepsVerifiedSender{
		SenderVerified: false,
		DomainVerified: false,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSenderVerificationDomainWarnList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/verified_senders/domains", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSenderVerificationDomainWarnList(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
