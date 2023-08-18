package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetSubusers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[{
			"id":12345678,
			"username":"dummy",
			"email":"dummy@example.com",
			"disabled": false
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubusers(context.TODO(), &InputGetSubusers{
		Username: "dummy",
		Limit:    1,
		Offset:   1,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Subuser{
		{
			ID:       12345678,
			Username: "dummy",
			Email:    "dummy@example.com",
			Disabled: false,
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSubusers_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubusers(context.TODO(), &InputGetSubusers{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSubuserReputations(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/reputations", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("usernames", "dummy")
		r.URL.RawQuery = q.Encode()
		if _, err := fmt.Fprint(w, `[{
			"reputation":100.0,
			"username":"dummy"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubuserReputations(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Reputation{
		{
			Reputation: 100.0,
			Username:   "dummy",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSubuserReputations_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/reputations", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("usernames", "dummy")
		r.URL.RawQuery = q.Encode()
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuserReputations(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"username":"dummy",
			"user_id":12345678,
			"email":"dummy3@example.com",
			"credit_allocation":{"type":"unlimited"}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateSubuser(context.TODO(), &InputCreateSubuser{
		Username: "dummy",
		Email:    "dummy3@example.com",
		Password: "dummy!123",
		Ips:      []string{"1.1.1.1"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateSubuser{
		UserID:   12345678,
		Username: "dummy",
		Email:    "dummy3@example.com",
		CreditAllocation: CreditAllocation{
			Type: "unlimited",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	_, err := client.CreateSubuser(context.TODO(), &InputCreateSubuser{
		Username: "dummy",
		Email:    "dummy3@example.com",
		Password: "dummy!123",
		Ips:      []string{"1.1.1.1"},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSubuserStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.UpdateSubuserStatus(context.TODO(), "dummy", &InputUpdateSubuserStatus{
		Disabled: false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestUpdateSubuserStatus_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	err := client.UpdateSubuserStatus(context.TODO(), "dummy", &InputUpdateSubuserStatus{
		Disabled: false,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSubuserIps(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.UpdateSubuserIps(context.TODO(), "dummy", []string{"1.1.1.1"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestUpdateSubuserIps_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	err := client.UpdateSubuserIps(context.TODO(), "dummy", []string{"1.1.1.1"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteSubuser(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})
	err := client.DeleteSubuser(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
