package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/testuser", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"id": 12345,
			"username": "testuser",
			"email": "testuser@example.com",
			"disabled": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubuser(context.TODO(), "testuser")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &Subuser{
		ID:       12345,
		Username: "testuser",
		Email:    "testuser@example.com",
		Disabled: false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/testuser", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuser(context.TODO(), "testuser")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

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

// NewRequest Error Tests
func TestGetSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSubuser(context.TODO(), "testuser")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSubusers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputGetSubusers{
		Username: "test",
		Limit:    50,
	}
	_, err := client.GetSubusers(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSubuserReputations_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSubuserReputations(context.TODO(), "testuser")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputCreateSubuser{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Ips:      []string{"192.168.1.1"},
	}
	_, err := client.CreateSubuser(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateSubuserStatus_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateSubuserStatus{
		Disabled: false,
	}
	err := client.UpdateSubuserStatus(context.TODO(), "testuser", input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateSubuserIps_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ips := []string{"192.168.1.1", "192.168.1.2"}
	err := client.UpdateSubuserIps(context.TODO(), "testuser", ips)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteSubuser(context.TODO(), "testuser")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
