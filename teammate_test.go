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

var testJSONTeammate = `{
	"username": "kenzo.tanaka",
	"email": "kenzo.tanaka@example.com",
	"first_name": "Kenzo",
	"last_name": "Tanaka",
	"address": "",
	"address2": "",
	"city": "",
	"state": "",
	"zip": "",
	"country": "",
	"company": "",
	"website": "",
	"phone": "",
	"is_admin": false,
	"is_sso": false,
	"user_type": "teammate",
	"scopes": [],
	"is_read_only": false
  }`

var testJSONTeammates = fmt.Sprintf(`{"result":[%s]}`, testJSONTeammate)

func getTestUser() *User {
	return &User{
		Username:   "kenzo.tanaka",
		Email:      "kenzo.tanaka@example.com",
		FirstName:  "Kenzo",
		LastName:   "Tanaka",
		Address:    "",
		Address2:   "",
		City:       "",
		State:      "",
		Zip:        "",
		Country:    "",
		Company:    "",
		Website:    "",
		Phone:      "",
		IsAdmin:    false,
		IsSSO:      false,
		UserType:   "teammate",
		Scopes:     []string{},
		IsReadOnly: false,
	}
}

var testJSONPendingTeammates = `{"result":[{
	"email": "kenzo.tanaka@example.com",
	"scopes": [],
	"is_admin": false,
	"token": "abcdefghi",
	"expiration_date": 1691502820
  }]}`

func getTestPendingTeammates() []*User {
	return []*User{
		{
			Email:          "kenzo.tanaka@example.com",
			Scopes:         []string{},
			IsAdmin:        false,
			Token:          "abcdefghi",
			ExpirationDate: 1691502820,
		},
	}
}

var testJSONInviteTeammate = `{
	"email": "kenzo.tanaka@example.com",
	"scopes":[
		"user.profile.read",
		"user.profile.update"
	],
	"is_admin": false,
	"token": "abcdefghi"
  }`

func getTestInviteTeammate() *User {
	return &User{
		Email: "kenzo.tanaka@example.com",
		Scopes: []string{
			"user.profile.read",
			"user.profile.update",
		},
		IsAdmin: false,
		Token:   "abcdefghi",
	}
}

func TestGetTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/username", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeammate); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeammate(context.TODO(), "username")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestUser()
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
		if _, err := fmt.Fprint(w, testJSONTeammates); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTeammates(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*User{getTestUser()}
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
		if _, err := fmt.Fprint(w, testJSONPendingTeammates); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetPendingTeammates(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := getTestPendingTeammates()
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
		if _, err := fmt.Fprint(w, testJSONInviteTeammate); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.InviteTeammate(context.TODO(), &InputInviteTeammate{
		Email:   "kenzo.tanaka@example.com",
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

	want := getTestInviteTeammate()

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
		Email:   "kenzo.tanaka@example.com",
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

	mux.HandleFunc("/teammates/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeammate); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTeammatePermissions(context.TODO(), "kenzo.tanaka", &InputUpdateTeammatePermissions{
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

	want := getTestUser()
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTeammatePermissions_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateTeammatePermissions(context.TODO(), "kenzo.tanaka", &InputUpdateTeammatePermissions{
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

	mux.HandleFunc("/teammates/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeammate); err != nil {
			t.Fatal(err)
		}
	})

	err := client.DeleteTeammate(context.TODO(), "kenzo.tanaka")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestDeleteTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteTeammate(context.TODO(), "kenzo.tanaka")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeletePendingTeammate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJSONTeammate); err != nil {
			t.Fatal(err)
		}
	})

	err := client.DeletePendingTeammate(context.TODO(), "kenzo.tanaka")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestDeletePendingTeammate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/pending/kenzo.tanaka", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeletePendingTeammate(context.TODO(), "kenzo.tanaka")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
