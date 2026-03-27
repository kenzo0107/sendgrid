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
			"disabled": false,
			"region": "global"
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
		Region:   "global",
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
		testMethod(t, r, "GET")

		// クエリパラメータの検証
		query := r.URL.Query()
		if query.Get("username") != "dummy" {
			t.Errorf("Expected username=dummy, got %s", query.Get("username"))
		}
		if query.Get("limit") != "1" {
			t.Errorf("Expected limit=1, got %s", query.Get("limit"))
		}
		if query.Get("offset") != "1" {
			t.Errorf("Expected offset=1, got %s", query.Get("offset"))
		}

		if _, err := fmt.Fprint(w, `[{
			"id":12345678,
			"username":"dummy",
			"email":"dummy@example.com",
			"disabled": false,
			"region": "global"
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
			Region:   "global",
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

func TestGetSubusers_WithRegionFilter(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		// クエリパラメータの検証
		query := r.URL.Query()
		if query.Get("region") != "eu" {
			t.Errorf("Expected region=eu, got %s", query.Get("region"))
		}

		if _, err := fmt.Fprint(w, `[{
			"id":11111111,
			"username":"eu-user",
			"email":"euuser@example.com",
			"disabled": false,
			"region": "eu"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubusers(context.TODO(), &InputGetSubusers{
		Region: "eu",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Subuser{
		{
			ID:       11111111,
			Username: "eu-user",
			Email:    "euuser@example.com",
			Disabled: false,
			Region:   "eu",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSubusers_WithIncludeRegion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		// クエリパラメータの検証
		query := r.URL.Query()
		if query.Get("include_region") != "true" {
			t.Errorf("Expected include_region=true, got %s", query.Get("include_region"))
		}

		if _, err := fmt.Fprint(w, `[{
			"id":22222222,
			"username":"global-user",
			"email":"globaluser@example.com",
			"disabled": false,
			"region": "global"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubusers(context.TODO(), &InputGetSubusers{
		IncludeRegion: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Subuser{
		{
			ID:       22222222,
			Username: "global-user",
			Email:    "globaluser@example.com",
			Disabled: false,
			Region:   "global",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetSubusers_WithAllParameters(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		// 全てのクエリパラメータの検証
		query := r.URL.Query()
		if query.Get("username") != "test-user" {
			t.Errorf("Expected username=test-user, got %s", query.Get("username"))
		}
		if query.Get("limit") != "10" {
			t.Errorf("Expected limit=10, got %s", query.Get("limit"))
		}
		if query.Get("offset") != "5" {
			t.Errorf("Expected offset=5, got %s", query.Get("offset"))
		}
		if query.Get("region") != "us" {
			t.Errorf("Expected region=us, got %s", query.Get("region"))
		}
		if query.Get("include_region") != "true" {
			t.Errorf("Expected include_region=true, got %s", query.Get("include_region"))
		}

		if _, err := fmt.Fprint(w, `[{
			"id":33333333,
			"username":"test-user",
			"email":"testuser@example.com",
			"disabled": false,
			"region": "us"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubusers(context.TODO(), &InputGetSubusers{
		Username:      "test-user",
		Limit:         10,
		Offset:        5,
		Region:        "us",
		IncludeRegion: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*Subuser{
		{
			ID:       33333333,
			Username: "test-user",
			Email:    "testuser@example.com",
			Disabled: false,
			Region:   "us",
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
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
			"credit_allocation":{"type":"unlimited"},
			"region":"global"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateSubuser(context.TODO(), &InputCreateSubuser{
		Username:      "dummy",
		Email:         "dummy3@example.com",
		Password:      "dummy!123",
		Ips:           []string{"1.1.1.1"},
		Region:        "global",
		IncludeRegion: true,
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
		Region: "global",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateSubuser_EU_Region(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"username":"eu-user",
			"user_id":98765432,
			"email":"euuser@example.com",
			"credit_allocation":{"type":"unlimited"},
			"region":"eu"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateSubuser(context.TODO(), &InputCreateSubuser{
		Username:      "eu-user",
		Email:         "euuser@example.com",
		Password:      "dummy!123",
		Ips:           []string{"1.1.1.1"},
		Region:        "eu",
		IncludeRegion: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateSubuser{
		UserID:   98765432,
		Username: "eu-user",
		Email:    "euuser@example.com",
		CreditAllocation: CreditAllocation{
			Type: "unlimited",
		},
		Region: "eu",
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
		Username:      "dummy",
		Email:         "dummy3@example.com",
		Password:      "dummy!123",
		Ips:           []string{"1.1.1.1"},
		Region:        "global",
		IncludeRegion: true,
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
		Username:      "testuser",
		Email:         "test@example.com",
		Password:      "password123",
		Ips:           []string{"192.168.1.1"},
		Region:        "global",
		IncludeRegion: true,
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

func TestGetCreditsForSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"type": "recurring",
			"reset_frequency": "daily",
			"remain": 500,
			"total": 1000,
			"used": 500
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetCreditsForSubuser(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetCreditsForSubuser{
		Type:           "recurring",
		ResetFrequency: "daily",
		Remain:         500,
		Total:          1000,
		Used:           500,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetCreditsForSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetCreditsForSubuser(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetCreditsForSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetCreditsForSubuser(context.TODO(), "dummy")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateCreditsForSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if _, err := fmt.Fprint(w, `{
			"type": "recurring",
			"reset_frequency": "daily",
			"remain": 500,
			"total": 1000,
			"used": 500
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputUpdateCreditsForSubuser{
		Type: "recurring",
	}

	expected, err := client.UpdateCreditsForSubuser(context.TODO(), "dummy", input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateCreditsForSubuser{
		Type:           "recurring",
		ResetFrequency: "daily",
		Remain:         500,
		Total:          1000,
		Used:           500,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateCreditsForSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputUpdateCreditsForSubuser{
		Type: "recurring",
	}
	_, err := client.UpdateCreditsForSubuser(context.TODO(), "dummy", input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateCreditsForSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateCreditsForSubuser{
		Type: "recurring",
	}
	_, err := client.UpdateCreditsForSubuser(context.TODO(), "dummy", input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateRemainingCreditsForSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits/remaining", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"type": "recurring",
			"reset_frequency": "daily",
			"remain": 600,
			"total": 1000,
			"used": 400
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateRemainingCreditsForSubuser(context.TODO(), "dummy", &InputUpdateRemainingCreditsForSubuser{
		AllocationUpdate: 100,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	want := &OutputUpdateRemainingCreditsForSubuser{
		Type:           "recurring",
		ResetFrequency: "daily",
		Remain:         600,
		Total:          1000,
		Used:           400,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateRemainingCreditsForSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/subusers/dummy/credits/remaining", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateRemainingCreditsForSubuser(context.TODO(), "dummy", &InputUpdateRemainingCreditsForSubuser{
		AllocationUpdate: 100,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateRemainingCreditsForSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateRemainingCreditsForSubuser(context.TODO(), "dummy", &InputUpdateRemainingCreditsForSubuser{
		AllocationUpdate: 100,
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
