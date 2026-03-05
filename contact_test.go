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

func TestAddOrUpdateContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"job_id": "test-job-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.AddOrUpdateContacts(context.TODO(), &InputAddOrUpdateContacts{
		ListIDs: []string{"list-1"},
		Contacts: []*ContactRequest{
			{
				Email:     "test@example.com",
				FirstName: "Test",
				LastName:  "User",
			},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAddOrUpdateContacts{
		JobID: "test-job-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestAddOrUpdateContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddOrUpdateContacts(context.TODO(), &InputAddOrUpdateContacts{
		Contacts: []*ContactRequest{
			{Email: "test@example.com"},
		},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestAddOrUpdateContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.AddOrUpdateContacts(context.TODO(), &InputAddOrUpdateContacts{
		Contacts: []*ContactRequest{
			{Email: "test@example.com"},
		},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestImportContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/imports", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"job_id": "test-job-id",
			"upload_uri": "https://example.com/upload",
			"upload_headers": [{"header": "Content-Type", "value": "text/csv"}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.ImportContacts(context.TODO(), &InputImportContacts{
		FileType:      "csv",
		FieldMappings: []string{"e1", "e2"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputImportContacts{
		JobID:     "test-job-id",
		UploadURI: "https://example.com/upload",
		UploadHeaders: []*ImportContactsUploadHeader{
			{Header: "Content-Type", Value: "text/csv"},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestImportContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/imports", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ImportContacts(context.TODO(), &InputImportContacts{
		FileType:      "csv",
		FieldMappings: []string{"e1"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestImportContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.ImportContacts(context.TODO(), &InputImportContacts{
		FileType:      "csv",
		FieldMappings: []string{"e1"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetImportContactsStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/imports/test-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "test-id",
			"status": "completed",
			"job_type": "upsert",
			"results": {
				"requested_count": 100,
				"created_count": 80,
				"updated_count": 20,
				"deleted_count": 0,
				"errored_count": 0
			},
			"started_at": "2024-01-01T00:00:00Z",
			"finished_at": "2024-01-01T00:01:00Z"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetImportContactsStatus(context.TODO(), "test-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetImportContactsStatus{
		ID:      "test-id",
		Status:  "completed",
		JobType: "upsert",
		Results: &ImportContactsResults{
			RequestedCount: 100,
			CreatedCount:   80,
			UpdatedCount:   20,
		},
		StartedAt:  "2024-01-01T00:00:00Z",
		FinishedAt: "2024-01-01T00:01:00Z",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetImportContactsStatus_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/imports/test-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetImportContactsStatus(context.TODO(), "test-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetImportContactsStatus_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetImportContactsStatus(context.TODO(), "test-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetContact(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/test-contact-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"contact": {
				"id": "test-contact-id",
				"email": "test@example.com",
				"first_name": "Test",
				"last_name": "User"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetContact(context.TODO(), "test-contact-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetContact{
		Contact: &Contact{
			ID:        "test-contact-id",
			Email:     "test@example.com",
			FirstName: "Test",
			LastName:  "User",
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetContact_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/test-contact-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetContact(context.TODO(), "test-contact-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetContact_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetContact(context.TODO(), "test-contact-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetBatchedContactsByIDs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/batch", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "contact-1",
				"email": "test1@example.com"
			}, {
				"id": "contact-2",
				"email": "test2@example.com"
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetBatchedContactsByIDs(context.TODO(), &InputGetBatchedContactsByIDs{
		IDs: []string{"contact-1", "contact-2"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetBatchedContactsByIDs{
		Result: []*Contact{
			{ID: "contact-1", Email: "test1@example.com"},
			{ID: "contact-2", Email: "test2@example.com"},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetBatchedContactsByIDs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/batch", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetBatchedContactsByIDs(context.TODO(), &InputGetBatchedContactsByIDs{
		IDs: []string{"contact-1"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetBatchedContactsByIDs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetBatchedContactsByIDs(context.TODO(), &InputGetBatchedContactsByIDs{
		IDs: []string{"contact-1"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetContactsByEmails(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search/emails", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": {
				"test@example.com": {
					"contact": {
						"id": "contact-1",
						"email": "test@example.com",
						"first_name": "Test"
					}
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetContactsByEmails(context.TODO(), &InputGetContactsByEmails{
		Emails: []string{"test@example.com"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetContactsByEmails{
		Result: map[string]*ContactByEmailResult{
			"test@example.com": {
				Contact: &Contact{
					ID:        "contact-1",
					Email:     "test@example.com",
					FirstName: "Test",
				},
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetContactsByEmails_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search/emails", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetContactsByEmails(context.TODO(), &InputGetContactsByEmails{
		Emails: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetContactsByEmails_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetContactsByEmails(context.TODO(), &InputGetContactsByEmails{
		Emails: []string{"test@example.com"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestSearchContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "contact-1",
				"email": "test@example.com"
			}],
			"contact_count": 1
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.SearchContacts(context.TODO(), &InputSearchContacts{
		Query: "email LIKE '%example.com'",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputSearchContacts{
		Result: []*Contact{
			{ID: "contact-1", Email: "test@example.com"},
		},
		ContactCount: 1,
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestSearchContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.SearchContacts(context.TODO(), &InputSearchContacts{
		Query: "email LIKE '%example.com'",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestSearchContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.SearchContacts(context.TODO(), &InputSearchContacts{
		Query: "email LIKE '%example.com'",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSampleContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "contact-1",
				"email": "test@example.com"
			}],
			"contact_count": 1
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetSampleContacts(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSampleContacts{
		Result: []*Contact{
			{ID: "contact-1", Email: "test@example.com"},
		},
		ContactCount: 1,
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetSampleContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSampleContacts(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetSampleContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSampleContacts(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetTotalContactCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/count", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"contact_count": 1000,
			"billable_count": 500,
			"billable_breakdown": {
				"total": 500,
				"breakdown": {"marketing": 300, "transactional": 200}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetTotalContactCount(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTotalContactCount{
		ContactCount:  1000,
		BillableCount: 500,
		BillableBreakdown: &ContactBillableBreakdown{
			Total: 500,
			Breakdown: map[string]int64{
				"marketing":     300,
				"transactional": 200,
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetTotalContactCount_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/count", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTotalContactCount(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetTotalContactCount_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetTotalContactCount(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestExportContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "export-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.ExportContacts(context.TODO(), &InputExportContacts{
		ListIDs:  []string{"list-1"},
		FileType: "csv",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputExportContacts{
		ID: "export-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestExportContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ExportContacts(context.TODO(), &InputExportContacts{
		FileType: "csv",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestExportContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.ExportContacts(context.TODO(), &InputExportContacts{
		FileType: "csv",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetExportContactsStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports/export-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "export-id",
			"status": "ready",
			"created_at": "2024-01-01T00:00:00Z",
			"updated_at": "2024-01-01T00:01:00Z",
			"completed_at": "2024-01-01T00:01:00Z",
			"expires_at": "2024-01-02T00:00:00Z",
			"urls": ["https://example.com/download"],
			"contact_count": 100
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetExportContactsStatus(context.TODO(), "export-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetExportContactsStatus{
		ID:           "export-id",
		Status:       "ready",
		CreatedAt:    "2024-01-01T00:00:00Z",
		UpdatedAt:    "2024-01-01T00:01:00Z",
		CompletedAt:  "2024-01-01T00:01:00Z",
		ExpiresAt:    "2024-01-02T00:00:00Z",
		URLs:         []string{"https://example.com/download"},
		ContactCount: 100,
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetExportContactsStatus_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports/export-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetExportContactsStatus(context.TODO(), "export-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetExportContactsStatus_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetExportContactsStatus(context.TODO(), "export-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetAllExistingExports(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "export-1",
				"status": "ready",
				"created_at": "2024-01-01T00:00:00Z"
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetAllExistingExports(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAllExistingExports{
		Result: []*ContactExportJob{
			{
				ID:        "export-1",
				Status:    "ready",
				CreatedAt: "2024-01-01T00:00:00Z",
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetAllExistingExports_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/exports", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAllExistingExports(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetAllExistingExports_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetAllExistingExports(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteContacts(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"job_id": "delete-job-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.DeleteContacts(context.TODO(), &InputDeleteContacts{
		IDs: []string{"contact-1", "contact-2"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputDeleteContacts{
		JobID: "delete-job-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestDeleteContacts_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteContacts(context.TODO(), &InputDeleteContacts{
		IDs: []string{"contact-1"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteContacts_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.DeleteContacts(context.TODO(), &InputDeleteContacts{
		IDs: []string{"contact-1"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteContactIdentifier(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/test-contact-id/identifiers", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"job_id": "delete-identifier-job-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.DeleteContactIdentifier(context.TODO(), "test-contact-id", &InputDeleteContactIdentifier{
		IdentifierType:  "email",
		IdentifierValue: "test@example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputDeleteContactIdentifier{
		JobID: "delete-identifier-job-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestDeleteContactIdentifier_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/test-contact-id/identifiers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DeleteContactIdentifier(context.TODO(), "test-contact-id", &InputDeleteContactIdentifier{
		IdentifierType:  "email",
		IdentifierValue: "test@example.com",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteContactIdentifier_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.DeleteContactIdentifier(context.TODO(), "test-contact-id", &InputDeleteContactIdentifier{
		IdentifierType:  "email",
		IdentifierValue: "test@example.com",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetContactsByIdentifiers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search/identifiers/email", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": {
				"test@example.com": {
					"contact": {
						"id": "contact-1",
						"email": "test@example.com"
					}
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetContactsByIdentifiers(context.TODO(), "email", &InputGetContactsByIdentifiers{
		Identifiers: []string{"test@example.com"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetContactsByIdentifiers{
		Result: map[string]*ContactByIdentifierResult{
			"test@example.com": {
				Contact: &Contact{
					ID:    "contact-1",
					Email: "test@example.com",
				},
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetContactsByIdentifiers_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/marketing/contacts/search/identifiers/email", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetContactsByIdentifiers(context.TODO(), "email", &InputGetContactsByIdentifiers{
		Identifiers: []string{"test@example.com"},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetContactsByIdentifiers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetContactsByIdentifiers(context.TODO(), "email", &InputGetContactsByIdentifiers{
		Identifiers: []string{"test@example.com"},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
