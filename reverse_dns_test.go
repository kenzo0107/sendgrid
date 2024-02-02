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

func TestGetReverseDNSs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("limit", "1")
		q.Set("offset", "10")
		q.Set("ip", "150.150.150.150")
		r.URL.RawQuery = q.Encode()

		if _, err := fmt.Fprint(w, `[{
			"id": 1234567,
			"ip": "150.150.150.150",
			"rdns": "o1.email.example.com",
			"users": [
				{
					"user_id": 12345678,
					"username": "dummy"
				},
				{
					"user_id": 123456782,
					"username": "dummy2"
				}
			],
			"subdomain": "email",
			"domain": "example.com",
			"valid": true,
			"legacy":false,
			"last_validation_attempt_at":1575947780,
			"a_record": {
				"valid":true,
				"type":"a",
				"host":"o1.email.example.com",
				"data":"150.150.150.150"
			}
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetReverseDNSs(context.TODO(), &InputGetReverseDNSs{
		Limit:  1,
		Offset: 10,
		IP:     "150.150.150.150",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*OutputGetReverseDNS{
		{
			ID:   1234567,
			IP:   "150.150.150.150",
			RDNS: "o1.email.example.com",
			Users: []*User{
				{
					UserID:   12345678,
					Username: "dummy",
				},
				{
					UserID:   123456782,
					Username: "dummy2",
				},
			},
			Subdomain:               "email",
			Domain:                  "example.com",
			Valid:                   true,
			Legacy:                  false,
			LastValidationAttemptAt: 1575947780,
			ARecord: ARecord{
				Valid: true,
				Type:  "a",
				Host:  "o1.email.example.com",
				Data:  "150.150.150.150",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetReverseDNSs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetReverseDNSs(context.TODO(), &InputGetReverseDNSs{
		Limit:  1,
		Offset: 10,
		IP:     "150.150.150.150",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetReverseDNS(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"ip": "150.150.150.150",
			"rdns": "o1.email.example.com",
			"users": [
				{
					"user_id": 12345678,
					"username": "dummy"
				},
				{
					"user_id": 123456782,
					"username": "dummy2"
				}
			],
			"subdomain": "email",
			"domain": "example.com",
			"valid": true,
			"legacy":false,
			"last_validation_attempt_at":1575947780,
			"a_record": {
				"valid":true,
				"type":"a",
				"host":"o1.email.example.com",
				"data":"150.150.150.150"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetReverseDNS(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetReverseDNS{
		ID:   1234567,
		IP:   "150.150.150.150",
		RDNS: "o1.email.example.com",
		Users: []*User{
			{
				UserID:   12345678,
				Username: "dummy",
			},
			{
				UserID:   123456782,
				Username: "dummy2",
			},
		},
		Subdomain:               "email",
		Domain:                  "example.com",
		Valid:                   true,
		Legacy:                  false,
		LastValidationAttemptAt: 1575947780,
		ARecord: ARecord{
			Valid: true,
			Type:  "a",
			Host:  "o1.email.example.com",
			Data:  "150.150.150.150",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetReverseDNS_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetReverseDNS(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateReverseDNS(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"ip": "150.150.150.150",
			"rdns": "o1.email.example.com",
			"users": [
				{
					"user_id": 12345678,
					"username": "dummy"
				},
				{
					"user_id": 123456782,
					"username": "dummy2"
				}
			],
			"subdomain": "email",
			"domain": "example.com",
			"valid": true,
			"legacy":false,
			"last_validation_attempt_at":1575947780,
			"a_record": {
				"valid":true,
				"type":"a",
				"host":"o1.email.example.com",
				"data":"150.150.150.150"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateReverseDNS(context.TODO(), &InputCreateReverseDNS{
		IP:        "150.150.150.150",
		Subdomain: "email",
		Domain:    "example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateReverseDNS{
		ID:   1234567,
		IP:   "150.150.150.150",
		RDNS: "o1.email.example.com",
		Users: []*User{
			{
				UserID:   12345678,
				Username: "dummy",
			},
			{
				UserID:   123456782,
				Username: "dummy2",
			},
		},
		Subdomain:               "email",
		Domain:                  "example.com",
		Valid:                   true,
		Legacy:                  false,
		LastValidationAttemptAt: 1575947780,
		ARecord: ARecord{
			Valid: true,
			Type:  "a",
			Host:  "o1.email.example.com",
			Data:  "150.150.150.150",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateReverseDNS_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateReverseDNS(context.TODO(), &InputCreateReverseDNS{
		IP:        "150.150.150.150",
		Subdomain: "email",
		Domain:    "example.com",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestValidateReverseDNS(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id":1234567,
			"valid":true,
			"validation_results":{
				"a_record":{
					"valid":true,
					"reason":null
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.ValidateReverseDNS(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputValidateReverseDNS{
		ID:    1234567,
		Valid: true,
		ValidationResults: ValidationResultsReverseDNS{
			ARecordValidationResults: ARecordValidationResults{
				Valid:  true,
				Reason: "",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestValidateReverseDNS_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ValidateReverseDNS(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteReverseDNS(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteReverseDNS(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteReverseDNS_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/ips/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteReverseDNS(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
