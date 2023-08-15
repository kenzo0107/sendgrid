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

func TestGetTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/username", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"username": "dummy",
			"email": "dummy@example.com",
			"first_name": "Kenzo",
			"last_name": "Tanaka",
			"address": "",
			"address2": "",
			"city": "",
			"state": "",
			"zip": "",
			"country": "",
			"website": "",
			"phone": "",
			"is_admin": false,
			"user_type": "teammate",
			"scopes": []
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeammate(context.TODO(), "username")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTeammate{
		Username:  "dummy",
		Email:     "dummy@example.com",
		FirstName: "Kenzo",
		LastName:  "Tanaka",
		Address:   "",
		Address2:  "",
		City:      "",
		State:     "",
		Zip:       "",
		Country:   "",
		Website:   "",
		Phone:     "",
		IsAdmin:   false,
		UserType:  "teammate",
		Scopes:    []string{},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/username", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTeammate(context.TODO(), "username")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetTeammates(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"result":[{
			"username": "dummy",
			"email": "dummy@example.com",
			"first_name": "Kenzo",
			"last_name": "Tanaka",
			"address": "",
			"address2": "",
			"city": "",
			"state": "",
			"zip": "",
			"country": "",
			"website": "",
			"phone": "",
			"is_admin": false,
			"user_type": "teammate",
			"scopes": []
		  }]}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeammates(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTeammates{
		Teammates: []Teammate{
			{
				Username:  "dummy",
				Email:     "dummy@example.com",
				FirstName: "Kenzo",
				LastName:  "Tanaka",
				Address:   "",
				Address2:  "",
				City:      "",
				State:     "",
				Zip:       "",
				Country:   "",
				Website:   "",
				Phone:     "",
				IsAdmin:   false,
				UserType:  "teammate",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetTeammates_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTeammates(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetPendingTeammates(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"result":[{
			"email": "dummy@example.com",
			"scopes": [],
			"is_admin": false,
			"token": "abcdefghi",
			"expiration_date": 1691502820
		  }]}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetPendingTeammates(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetPendingTeammates{
		PendingTeammates: []PendingTeammate{
			{
				Email:          "dummy@example.com",
				Scopes:         []string{},
				IsAdmin:        false,
				Token:          "abcdefghi",
				ExpirationDate: 1691502820,
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetPendingTeammates_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetPendingTeammates(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestInviteTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"email": "dummy@example.com",
			"scopes":[
				"user.profile.read",
				"user.profile.update"
			],
			"is_admin": false,
			"token": "abcdefghi"
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.InviteTeammate(context.TODO(), &InputInviteTeammate{
		Email:   "dummy@example.com",
		IsAdmin: false,
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	want := &OutputInviteTeammate{
		Token:   "abcdefghi",
		Email:   "dummy@example.com",
		IsAdmin: false,
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestInviteTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.InviteTeammate(context.TODO(), &InputInviteTeammate{
		Email:   "dummy@example.com",
		IsAdmin: false,
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateTeammatePermissions(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"username": "dummy",
			"first_name": "Kenzo",
			"last_name": "Tanaka",
			"email": "dummy@example.com",
			"scopes": [
				"user.profile.read",
				"user.profile.update"
			],
			"address":  "",
			"address2": "",
			"city":     "",
			"state":    "",
			"zip":      "",
			"country":  "",
			"website":  "",
			"phone":    "",
			"is_admin":  false,
			"user_type": "teammate"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTeammatePermissions(context.TODO(), "dummy", &InputUpdateTeammatePermissions{
		IsAdmin: false,
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	want := &OutputUpdateTeammatePermissions{
		Username:  "dummy",
		FirstName: "Kenzo",
		LastName:  "Tanaka",
		Email:     "dummy@example.com",
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
		Address:  "",
		Address2: "",
		City:     "",
		State:    "",
		Zip:      "",
		Country:  "",
		Website:  "",
		Phone:    "",
		IsAdmin:  false,
		UserType: "teammate",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTeammatePermissions_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateTeammatePermissions(context.TODO(), "dummy", &InputUpdateTeammatePermissions{
		IsAdmin: false,
		Scopes: []string{
			"user.profile.read",
		},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteTeammate(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestDeleteTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteTeammate(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeletePendingTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeletePendingTeammate(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestDeletePendingTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeletePendingTeammate(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
