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

// CreateLegacyList

func TestCreateLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"name": "your list name",
			"recipient_count": 0
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateLegacyList(context.TODO(), &InputCreateLegacyList{
		Name: "your list name",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateLegacyList{
		ID:             1,
		Name:           "your list name",
		RecipientCount: 0,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateLegacyList(context.TODO(), &InputCreateLegacyList{
		Name: "your list name",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CreateLegacyList(context.TODO(), &InputCreateLegacyList{
		Name: "your list name",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// GetLegacyLists

func TestGetLegacyLists(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"lists": [
				{
					"id": 1,
					"name": "the jones",
					"recipient_count": 1
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyLists(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyLists{
		Lists: []*LegacyList{
			{
				ID:             1,
				Name:           "the jones",
				RecipientCount: 1,
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyLists_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyLists(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyLists_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyLists(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// DeleteLegacyLists

func TestDeleteLegacyLists(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	input := InputDeleteLegacyLists([]int64{1, 2, 3, 4})
	if err := client.DeleteLegacyLists(context.TODO(), &input); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteLegacyLists_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := InputDeleteLegacyLists([]int64{1, 2, 3, 4})
	if err := client.DeleteLegacyLists(context.TODO(), &input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacyLists_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := InputDeleteLegacyLists([]int64{1, 2, 3, 4})
	err := client.DeleteLegacyLists(context.TODO(), &input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// GetLegacyList

func TestGetLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"name": "listname",
			"recipient_count": 0
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyList(context.TODO(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyList{
		ID:             1,
		Name:           "listname",
		RecipientCount: 0,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyList(context.TODO(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyList(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// UpdateLegacyList

func TestUpdateLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"id": 1234,
			"name": "2016 iPhone Users",
			"recipient_count": 0
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateLegacyList(context.TODO(), 1234, &InputUpdateLegacyList{
		Name: "2016 iPhone Users",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateLegacyList{
		ID:             1234,
		Name:           "2016 iPhone Users",
		RecipientCount: 0,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1234", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateLegacyList(context.TODO(), 1234, &InputUpdateLegacyList{
		Name: "2016 iPhone Users",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateLegacyList(context.TODO(), 1234, &InputUpdateLegacyList{
		Name: "2016 iPhone Users",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// DeleteLegacyList

func TestDeleteLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusAccepted)
	})

	if err := client.DeleteLegacyList(context.TODO(), 1, &InputDeleteLegacyList{DeleteContacts: true}); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.DeleteLegacyList(context.TODO(), 1, &InputDeleteLegacyList{DeleteContacts: true}); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteLegacyList(context.TODO(), 1, &InputDeleteLegacyList{DeleteContacts: true})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// GetLegacyListRecipients

func TestGetLegacyListRecipients(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"recipients": [
				{
					"created_at": 1433348344,
					"custom_fields": [
						{
							"id": 6234,
							"name": "age",
							"type": "number",
							"value": null
						}
					],
					"email": "example@example.com",
					"first_name": "Example",
					"id": "ZGVWfyZWsuYmFpbmVzQHNlbmRmCmLkLmNv==",
					"last_clicked": 1438616117,
					"last_emailed": 1438613272,
					"last_name": "User",
					"last_opened": 1438616109,
					"updated_at": 1438616119
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyListRecipients(context.TODO(), 1, &InputGetLegacyListRecipients{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyListRecipients{
		Recipients: []*LegacyListRecipient{
			{
				CreatedAt: 1433348344,
				CustomFields: []*LegacyListRecipientCustomField{
					{
						ID:   6234,
						Name: "age",
						Type: "number",
					},
				},
				Email:       "example@example.com",
				FirstName:   "Example",
				ID:          "ZGVWfyZWsuYmFpbmVzQHNlbmRmCmLkLmNv==",
				LastClicked: 1438616117,
				LastEmailed: 1438613272,
				LastName:    "User",
				LastOpened:  1438616109,
				UpdatedAt:   1438616119,
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyListRecipients_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyListRecipients(context.TODO(), 1, &InputGetLegacyListRecipients{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyListRecipients_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyListRecipients(context.TODO(), 1, &InputGetLegacyListRecipients{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// AddRecipientsToLegacyList

func TestAddRecipientsToLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
	})

	input := InputAddRecipientsToLegacyList([]string{"ZW1haWwxQGV4YW1wbGUuY29t", "ZW1haWwyQHRlc3QubmV0"})
	if err := client.AddRecipientsToLegacyList(context.TODO(), 1, &input); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestAddRecipientsToLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	input := InputAddRecipientsToLegacyList([]string{"ZW1haWwxQGV4YW1wbGUuY29t"})
	if err := client.AddRecipientsToLegacyList(context.TODO(), 1, &input); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddRecipientsToLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := InputAddRecipientsToLegacyList([]string{"ZW1haWwxQGV4YW1wbGUuY29t"})
	err := client.AddRecipientsToLegacyList(context.TODO(), 1, &input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// AddRecipientToLegacyList

func TestAddRecipientToLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients/recipient123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
	})

	if err := client.AddRecipientToLegacyList(context.TODO(), 1, "recipient123"); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestAddRecipientToLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients/recipient123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.AddRecipientToLegacyList(context.TODO(), 1, "recipient123"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddRecipientToLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.AddRecipientToLegacyList(context.TODO(), 1, "recipient123")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

// RemoveRecipientFromLegacyList

func TestRemoveRecipientFromLegacyList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients/recipient123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if err := client.RemoveRecipientFromLegacyList(context.TODO(), 1, "recipient123"); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestRemoveRecipientFromLegacyList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/lists/1/recipients/recipient123", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.RemoveRecipientFromLegacyList(context.TODO(), 1, "recipient123"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestRemoveRecipientFromLegacyList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.RemoveRecipientFromLegacyList(context.TODO(), 1, "recipient123")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
