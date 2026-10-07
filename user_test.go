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

func TestGetUserAccount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/account", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{"type": "paid", "reputation": 100, "is_reseller_customer": true}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserAccount(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetUserAccount{Type: "paid", Reputation: 100, IsResellerCustomer: true}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserAccount_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/account", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetUserAccount(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserAccount_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetUserAccount(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetUserCredits(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/credits", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"remain": 200,
			"total": 200,
			"overage": 0,
			"used": 0,
			"last_reset": "2013-01-01",
			"next_reset": "2013-02-01",
			"reset_frequency": "monthly",
			"is_hard_limit": true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserCredits(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetUserCredits{
		Remain:         200,
		Total:          200,
		LastReset:      "2013-01-01",
		NextReset:      "2013-02-01",
		ResetFrequency: "monthly",
		IsHardLimit:    true,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserCredits_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/credits", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetUserCredits(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserCredits_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetUserCredits(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetUserEmail(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"email": "test@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserEmail(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetUserEmail{Email: "test@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserEmail_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/email", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetUserEmail(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserEmail_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetUserEmail(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateUserEmail(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/email", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if _, err := fmt.Fprint(w, `{"email": "example@example.com"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateUserEmail(context.TODO(), &InputUpdateUserEmail{Email: "example@example.com"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateUserEmail{Email: "example@example.com"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateUserEmail_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/email", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateUserEmail(context.TODO(), &InputUpdateUserEmail{Email: "example@example.com"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUserEmail_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateUserEmail(context.TODO(), &InputUpdateUserEmail{Email: "example@example.com"})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateUserPassword(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/password", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if _, err := fmt.Fprint(w, `{}`); err != nil {
			t.Fatal(err)
		}
	})

	err := client.UpdateUserPassword(context.TODO(), &InputUpdateUserPassword{
		NewPassword: "new_password",
		OldPassword: "old_password",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestUpdateUserPassword_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/password", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.UpdateUserPassword(context.TODO(), &InputUpdateUserPassword{
		NewPassword: "new_password",
		OldPassword: "old_password",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUserPassword_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.UpdateUserPassword(context.TODO(), &InputUpdateUserPassword{
		NewPassword: "new_password",
		OldPassword: "old_password",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetUserProfile(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"address": "814 West Chapman Avenue",
			"city": "Orange",
			"company": "SendGrid",
			"country": "US",
			"first_name": "Test",
			"last_name": "User",
			"phone": "555-555-5555",
			"state": "CA",
			"website": "http://www.sendgrid.com",
			"zip": "92868",
			"authy_id": 0,
			"multifactor_phone": "",
			"type": "free",
			"userid": 1234
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUserProfile(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetUserProfile{
		Address:   "814 West Chapman Avenue",
		City:      "Orange",
		Company:   "SendGrid",
		Country:   "US",
		FirstName: "Test",
		LastName:  "User",
		Phone:     "555-555-5555",
		State:     "CA",
		Website:   "http://www.sendgrid.com",
		Zip:       "92868",
		Type:      "free",
		UserID:    1234,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUserProfile_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetUserProfile(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUserProfile_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetUserProfile(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateUserProfile(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"address": "814 West Chapman Avenue",
			"city": "Orange",
			"company": "SendGrid",
			"country": "US",
			"first_name": "Example",
			"last_name": "User",
			"phone": "555-555-5555",
			"state": "CA",
			"website": "http://www.sendgrid.com",
			"zip": "92868"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateUserProfile(context.TODO(), &InputUpdateUserProfile{FirstName: "Example"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateUserProfile{
		Address:   "814 West Chapman Avenue",
		City:      "Orange",
		Company:   "SendGrid",
		Country:   "US",
		FirstName: "Example",
		LastName:  "User",
		Phone:     "555-555-5555",
		State:     "CA",
		Website:   "http://www.sendgrid.com",
		Zip:       "92868",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateUserProfile_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateUserProfile(context.TODO(), &InputUpdateUserProfile{FirstName: "Example"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUserProfile_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateUserProfile(context.TODO(), &InputUpdateUserProfile{FirstName: "Example"})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetUsername(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{"username": "test_username", "user_id": 1}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetUsername(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetUsername{Username: "test_username", UserID: 1}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetUsername_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/username", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetUsername(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetUsername_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetUsername(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateUsername(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/username", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if _, err := fmt.Fprint(w, `{"username": "test_username"}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateUsername(context.TODO(), &InputUpdateUsername{Username: "test_username"})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateUsername{Username: "test_username"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateUsername_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/username", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateUsername(context.TODO(), &InputUpdateUsername{Username: "test_username"})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateUsername_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateUsername(context.TODO(), &InputUpdateUsername{Username: "test_username"})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
