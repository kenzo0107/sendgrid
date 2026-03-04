package sendgrid

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
)

func TestCreateBatchID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail/batch", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"batch_id": "test-batch-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.CreateBatchID(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateBatchID{
		BatchID: "test-batch-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestCreateBatchID_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail/batch", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateBatchID(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCreateBatchID_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CreateBatchID(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestValidateBatchID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mail/batch/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"batch_id": "test-batch-id"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.ValidateBatchID(context.TODO(), "test-batch-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputValidateBatchID{
		BatchID: "test-batch-id",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestValidateBatchID_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ValidateBatchID(context.TODO(), "test-batch-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestValidateBatchID_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.ValidateBatchID(context.TODO(), "test-batch-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetScheduledSends(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `[{
			"batch_id": "test-batch-id",
			"status": "cancel"
		}]`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetScheduledSends(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*ScheduledSend{
		{
			BatchID: "test-batch-id",
			Status:  "cancel",
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetScheduledSends_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetScheduledSends(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetScheduledSends_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetScheduledSends(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetScheduledSend(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"batch_id": "test-batch-id",
			"status": "cancel"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetScheduledSend(context.TODO(), "test-batch-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetScheduledSend{
		BatchID: "test-batch-id",
		Status:  "cancel",
	}

	log.Printf("r: %#v", r)

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetScheduledSend_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetScheduledSend_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateScheduledSend(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.UpdateScheduledSend(context.TODO(), "test-batch-id", &InputUpdateScheduledSend{
		Status: "pause",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestUpdateScheduledSend_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.UpdateScheduledSend(context.TODO(), "test-batch-id", &InputUpdateScheduledSend{
		Status: "pause",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateScheduledSend_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.UpdateScheduledSend(context.TODO(), "test-batch-id", &InputUpdateScheduledSend{
		Status: "pause",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCancelOrPauseScheduledSend(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"batch_id": "test-batch-id",
			"status": "cancel"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.CancelOrPauseScheduledSend(context.TODO(), "test-batch-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCancelOrPauseScheduledSend{
		BatchID: "test-batch-id",
		Status:  "cancel",
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestCancelOrPauseScheduledSend_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CancelOrPauseScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCancelOrPauseScheduledSend_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CancelOrPauseScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteScheduledSend(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteScheduledSend(context.TODO(), "test-batch-id")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteScheduledSend_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/user/scheduled_sends/test-batch-id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteScheduledSend_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteScheduledSend(context.TODO(), "test-batch-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
