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

func TestCreateMarketingSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders", func(w http.ResponseWriter, r *http.Request) {
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

	input := &InputCreateMarketingSender{
		Nickname: "My Sender ID",
		From: &MarketingSenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &MarketingSenderAddress{
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
	expected, err := client.CreateMarketingSender(context.Background(), input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateMarketingSender{
		ID:       1,
		Nickname: "My Sender ID",
		From: &MarketingSenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &MarketingSenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &MarketingSenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestCreateMarketingSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputCreateMarketingSender{
		Nickname: "My Sender ID",
		Address:  "123 Elm St.",
		City:     "Denver",
		Country:  "United States",
	}
	_, err := client.CreateMarketingSender(context.Background(), input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateMarketingSender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputCreateMarketingSender{
		Nickname: "My Sender ID",
		Address:  "123 Elm St.",
		City:     "Denver",
		Country:  "United States",
	}
	_, err := client.CreateMarketingSender(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetMarketingSenders(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders", func(w http.ResponseWriter, r *http.Request) {
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

	expected, err := client.GetMarketingSenders(context.Background())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*MarketingSender{
		{
			ID:       1,
			Nickname: "My Sender ID",
			From: &MarketingSenderAddress{
				Email: "from@example.com",
				Name:  "Example INC",
			},
			ReplyTo: &MarketingSenderAddress{
				Email: "replyto@example.com",
				Name:  "Example INC",
			},
			Address:   "123 Elm St.",
			Address2:  "Apt. 456",
			City:      "Denver",
			State:     "Colorado",
			Zip:       "80202",
			Country:   "United States",
			Verified:  &MarketingSenderVerified{Status: true},
			UpdatedAt: 1449872165,
			CreatedAt: 1449872165,
			Locked:    false,
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetMarketingSenders_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetMarketingSenders(context.Background())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetMarketingSenders_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetMarketingSenders(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetMarketingSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
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

	expected, err := client.GetMarketingSender(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetMarketingSender{
		ID:       1,
		Nickname: "My Sender ID",
		From: &MarketingSenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &MarketingSenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &MarketingSenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestGetMarketingSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetMarketingSender(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetMarketingSender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetMarketingSender(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateMarketingSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
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

	input := &InputUpdateMarketingSender{
		Nickname: "Updated Sender ID",
	}
	expected, err := client.UpdateMarketingSender(context.Background(), 1, input)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateMarketingSender{
		ID:       1,
		Nickname: "Updated Sender ID",
		From: &MarketingSenderAddress{
			Email: "from@example.com",
			Name:  "Example INC",
		},
		ReplyTo: &MarketingSenderAddress{
			Email: "replyto@example.com",
			Name:  "Example INC",
		},
		Address:   "123 Elm St.",
		Address2:  "Apt. 456",
		City:      "Denver",
		State:     "Colorado",
		Zip:       "80202",
		Country:   "United States",
		Verified:  &MarketingSenderVerified{Status: true},
		UpdatedAt: 1449872165,
		CreatedAt: 1449872165,
		Locked:    false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse)
	}
}

func TestUpdateMarketingSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := &InputUpdateMarketingSender{
		Nickname: "Updated Sender ID",
	}
	_, err := client.UpdateMarketingSender(context.Background(), 1, input)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateMarketingSender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateMarketingSender{
		Nickname: "Updated Sender ID",
	}
	_, err := client.UpdateMarketingSender(context.TODO(), 1, input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteMarketingSender(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteMarketingSender(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteMarketingSender_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteMarketingSender(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteMarketingSender_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteMarketingSender(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestResendMarketingSenderVerification(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1/resend_verification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.ResendMarketingSenderVerification(context.Background(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestResendMarketingSenderVerification_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/senders/1/resend_verification", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.ResendMarketingSenderVerification(context.Background(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestResendMarketingSenderVerification_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.ResendMarketingSenderVerification(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
