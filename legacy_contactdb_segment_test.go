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

func TestCreateLegacySegment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"name": "Last Name Miller",
			"list_id": 4,
			"conditions": [
				{
					"field": "last_name",
					"value": "Miller",
					"operator": "eq",
					"and_or": ""
				}
			],
			"recipient_count": 0
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateLegacySegment(context.TODO(), &InputCreateLegacySegment{
		Name:   "Last Name Miller",
		ListID: 4,
		Conditions: []*LegacySegmentCondition{
			{
				Field:    "last_name",
				Value:    "Miller",
				Operator: "eq",
				AndOr:    "",
			},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateLegacySegment{
		ID:     1,
		Name:   "Last Name Miller",
		ListID: 4,
		Conditions: []*LegacySegmentCondition{
			{
				Field:    "last_name",
				Value:    "Miller",
				Operator: "eq",
				AndOr:    "",
			},
		},
		RecipientCount: 0,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateLegacySegment_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateLegacySegment(context.TODO(), &InputCreateLegacySegment{
		Name: "Last Name Miller",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateLegacySegment_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CreateLegacySegment(context.TODO(), &InputCreateLegacySegment{
		Name: "Last Name Miller",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacySegments(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"segments": [
				{
					"id": 1234,
					"name": "Age segments < 25",
					"conditions": [
						{
							"field": "age",
							"value": "25",
							"operator": "lt"
						}
					],
					"recipient_count": 8
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacySegments(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacySegments{
		Segments: []*LegacySegment{
			{
				ID:   1234,
				Name: "Age segments < 25",
				Conditions: []*LegacySegmentCondition{
					{
						Field:    "age",
						Value:    "25",
						Operator: "lt",
					},
				},
				RecipientCount: 8,
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacySegments_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacySegments(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacySegments_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacySegments(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacySegment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1,
			"name": "Last Name Miller",
			"list_id": 4,
			"conditions": [
				{
					"field": "last_name",
					"value": "Miller",
					"operator": "eq",
					"and_or": ""
				}
			],
			"recipient_count": 1
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacySegment(context.TODO(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacySegment{
		ID:     1,
		Name:   "Last Name Miller",
		ListID: 4,
		Conditions: []*LegacySegmentCondition{
			{
				Field:    "last_name",
				Value:    "Miller",
				Operator: "eq",
				AndOr:    "",
			},
		},
		RecipientCount: 1,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacySegment_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacySegment(context.TODO(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacySegment_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacySegment(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateLegacySegment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/5", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 5,
			"name": "The Millers",
			"list_id": 5,
			"conditions": [
				{
					"field": "last_name",
					"value": "Miller",
					"operator": "eq",
					"and_or": ""
				}
			],
			"recipient_count": 1
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateLegacySegment(context.TODO(), 5, &InputUpdateLegacySegment{
		Name:   "The Millers",
		ListID: 5,
		Conditions: []*LegacySegmentCondition{
			{
				Field:    "last_name",
				Value:    "Miller",
				Operator: "eq",
				AndOr:    "",
			},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateLegacySegment{
		ID:     5,
		Name:   "The Millers",
		ListID: 5,
		Conditions: []*LegacySegmentCondition{
			{
				Field:    "last_name",
				Value:    "Miller",
				Operator: "eq",
				AndOr:    "",
			},
		},
		RecipientCount: 1,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateLegacySegment_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/5", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateLegacySegment(context.TODO(), 5, &InputUpdateLegacySegment{
		Name: "The Millers",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateLegacySegment_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateLegacySegment(context.TODO(), 5, &InputUpdateLegacySegment{
		Name: "The Millers",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteLegacySegment(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteLegacySegment(context.TODO(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteLegacySegment_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteLegacySegment(context.TODO(), 1)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacySegment_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteLegacySegment(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacySegmentRecipients(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"recipients": [
				{
					"created_at": 1422313607,
					"email": "jones@example.com",
					"first_name": null,
					"id": "YUBh",
					"last_clicked": null,
					"last_emailed": null,
					"last_name": "Jones",
					"last_opened": null,
					"updated_at": 1422313790,
					"custom_fields": [
						{
							"id": 23,
							"name": "pet",
							"value": "Indiana",
							"type": "text"
						}
					]
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacySegmentRecipients(context.TODO(), 1, &InputGetLegacySegmentRecipients{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacySegmentRecipients{
		Recipients: []*LegacySegmentRecipient{
			{
				ID:        "YUBh",
				Email:     "jones@example.com",
				LastName:  "Jones",
				CreatedAt: 1422313607,
				UpdatedAt: 1422313790,
				CustomFields: []map[string]interface{}{
					{
						"id":    float64(23),
						"name":  "pet",
						"value": "Indiana",
						"type":  "text",
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacySegmentRecipients_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/contactdb/segments/1/recipients", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacySegmentRecipients(context.TODO(), 1, nil)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacySegmentRecipients_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacySegmentRecipients(context.TODO(), 1, nil)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
