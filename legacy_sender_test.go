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

func TestCreateLegacySender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"nickname": "My Sender ID",
			"from": {
				"email": "from@example.com",
				"name": "Example INC"
			},
			"reply_to": {
				"email": "replyto@example.com",
				"name": "Example INC"
			},
			"address": "123 Elm St.",
			"address_2": "Apt. 456",
			"city": "Denver",
			"state": "Colorado",
			"zip": "80202",
			"country": "United States",
			"verified": {"status": true, "reason": null},
			"updated_at": 1449872165,
			"created_at": 1449872165,
			"locked": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputCreateLegacySender{
		Nickname: "My Sender ID",
		From: &LegacySenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &LegacySenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:  "123 Elm St.",
		Address2: "Apt. 456",
		City:     "Denver",
		State:    "Colorado",
		Zip:      "80202",
		Country:  "United States",
	}
	expected, err := client.CreateLegacySender(context.Background(), input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateLegacySender{
		ID:       1,
		Nickname: "My Sender ID",
		From: &LegacySenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &LegacySenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &LegacySenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateLegacySender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputCreateLegacySender{
		Nickname: "My Sender ID",
		Address:  "123 Elm St.",
		City:     "Denver",
		Country:  "United States",
	}
	_, err := client.CreateLegacySender(context.Background(), input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateLegacySender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputCreateLegacySender{
		Nickname: "My Sender ID",
		Address:  "123 Elm St.",
		City:     "Denver",
		Country:  "United States",
	}
	_, err := client.CreateLegacySender(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacySenders(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[
				{
					"id": 1,
					"nickname": "My Sender ID",
					"from": {
						"email": "from@example.com",
						"name": "Example INC"
					},
					"reply_to": {
						"email": "replyto@example.com",
						"name": "Example INC"
					},
					"address": "123 Elm St.",
					"address_2": "Apt. 456",
					"city": "Denver",
					"state": "Colorado",
					"zip": "80202",
					"country": "United States",
					"verified": {"status": true, "reason": null},
					"updated_at": 1449872165,
					"created_at": 1449872165,
					"locked": false
				}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacySenders(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*LegacySender{
		{
			ID:       1,
			Nickname: "My Sender ID",
			From: &LegacySenderAddress{
				Email: "from@example.com",
				Name:  "Example INC",
			},
			ReplyTo: &LegacySenderAddress{
				Email: "replyto@example.com",
				Name:  "Example INC",
			},
			Address:   "123 Elm St.",
			Address2:  "Apt. 456",
			City:      "Denver",
			State:     "Colorado",
			Zip:       "80202",
			Country:   "United States",
			Verified:  &LegacySenderVerified{Status: true},
			UpdatedAt: 1449872165,
			CreatedAt: 1449872165,
			Locked:    false,
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetLegacySenders_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacySenders(context.Background())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacySenders_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacySenders(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacySender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"nickname": "My Sender ID",
			"from": {
				"email": "from@example.com",
				"name": "Example INC"
			},
			"reply_to": {
				"email": "replyto@example.com",
				"name": "Example INC"
			},
			"address": "123 Elm St.",
			"address_2": "Apt. 456",
			"city": "Denver",
			"state": "Colorado",
			"zip": "80202",
			"country": "United States",
			"verified": {"status": true, "reason": null},
			"updated_at": 1449872165,
			"created_at": 1449872165,
			"locked": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacySender(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacySender{
		ID:       1,
		Nickname: "My Sender ID",
		From: &LegacySenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &LegacySenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &LegacySenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetLegacySender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacySender(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacySender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacySender(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateLegacySender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"nickname": "Updated Sender ID",
			"from": {
				"email": "from@example.com",
				"name": "Example INC"
			},
			"reply_to": {
				"email": "replyto@example.com",
				"name": "Example INC"
			},
			"address": "123 Elm St.",
			"address_2": "Apt. 456",
			"city": "Denver",
			"state": "Colorado",
			"zip": "80202",
			"country": "United States",
			"verified": {"status": true, "reason": null},
			"updated_at": 1449872165,
			"created_at": 1449872165,
			"locked": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	input := &InputUpdateLegacySender{
		Nickname: "Updated Sender ID",
	}
	expected, err := client.UpdateLegacySender(context.Background(), 1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateLegacySender{
		ID:       1,
		Nickname: "Updated Sender ID",
		From: &LegacySenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &LegacySenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &LegacySenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateLegacySender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputUpdateLegacySender{
		Nickname: "Updated Sender ID",
	}
	_, err := client.UpdateLegacySender(context.Background(), 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateLegacySender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateLegacySender{
		Nickname: "Updated Sender ID",
	}
	_, err := client.UpdateLegacySender(context.TODO(), 1, input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteLegacySender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteLegacySender(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteLegacySender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteLegacySender(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacySender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteLegacySender(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestResendLegacySenderVerification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1/resend_verification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.ResendLegacySenderVerification(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestResendLegacySenderVerification_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/senders/1/resend_verification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.ResendLegacySenderVerification(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestResendLegacySenderVerification_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.ResendLegacySenderVerification(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
