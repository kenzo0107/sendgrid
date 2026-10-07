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

func TestCreateLegacyCampaign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 986724,
			"title": "March Newsletter",
			"subject": "New Products for Spring!",
			"sender_id": 124451,
			"list_ids": [110, 124],
			"segment_ids": [110],
			"categories": ["spring line"],
			"suppression_group_id": 42,
			"custom_unsubscribe_url": "",
			"ip_pool": "marketing",
			"html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
			"plain_content": "Check out our spring line!",
			"status": "Draft"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateLegacyCampaign(context.TODO(), &InputCreateLegacyCampaign{
		Title:      "March Newsletter",
		Subject:    "New Products for Spring!",
		SenderID:   124451,
		ListIDs:    []int64{110, 124},
		SegmentIDs: []int64{110},
		Categories: []string{"spring line"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateLegacyCampaign{
		ID:                 986724,
		Title:              "March Newsletter",
		Subject:            "New Products for Spring!",
		SenderID:           124451,
		ListIDs:            []int64{110, 124},
		SegmentIDs:         []int64{110},
		Categories:         []string{"spring line"},
		SuppressionGroupID: 42,
		IPPool:             "marketing",
		HTMLContent:        "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
		PlainContent:       "Check out our spring line!",
		Status:             "Draft",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateLegacyCampaign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateLegacyCampaign(context.TODO(), &InputCreateLegacyCampaign{
		Title: "March Newsletter",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateLegacyCampaign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.CreateLegacyCampaign(context.TODO(), &InputCreateLegacyCampaign{
		Title: "March Newsletter",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacyCampaigns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"id": 986724,
					"title": "March Newsletter",
					"subject": "New Products for Spring!",
					"sender_id": 124451,
					"list_ids": [110, 124],
					"segment_ids": [110],
					"categories": ["spring line"],
					"suppression_group_id": 42,
					"ip_pool": "marketing",
					"html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
					"plain_content": "Check out our spring line!",
					"status": "Draft"
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyCampaigns(context.TODO(), &InputGetLegacyCampaigns{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyCampaigns{
		Result: []*LegacyCampaign{
			{
				ID:                 986724,
				Title:              "March Newsletter",
				Subject:            "New Products for Spring!",
				SenderID:           124451,
				ListIDs:            []int64{110, 124},
				SegmentIDs:         []int64{110},
				Categories:         []string{"spring line"},
				SuppressionGroupID: 42,
				IPPool:             "marketing",
				HTMLContent:        "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
				PlainContent:       "Check out our spring line!",
				Status:             "Draft",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyCampaigns_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyCampaigns(context.TODO(), &InputGetLegacyCampaigns{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyCampaigns_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyCampaigns(context.TODO(), &InputGetLegacyCampaigns{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacyCampaign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 986724,
			"title": "March Newsletter",
			"subject": "New Products for Spring!",
			"sender_id": 124451,
			"list_ids": [110],
			"segment_ids": [110],
			"categories": ["spring line"],
			"suppression_group_id": 42,
			"ip_pool": "marketing",
			"html_content": "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
			"plain_content": "Check out our spring line!",
			"status": "Draft"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyCampaign(context.TODO(), 986724)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyCampaign{
		ID:                 986724,
		Title:              "March Newsletter",
		Subject:            "New Products for Spring!",
		SenderID:           124451,
		ListIDs:            []int64{110},
		SegmentIDs:         []int64{110},
		Categories:         []string{"spring line"},
		SuppressionGroupID: 42,
		IPPool:             "marketing",
		HTMLContent:        "<html><head><title></title></head><body><p>Check out our spring line!</p></body></html>",
		PlainContent:       "Check out our spring line!",
		Status:             "Draft",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyCampaign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyCampaign(context.TODO(), 986724)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyCampaign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyCampaign(context.TODO(), 986724)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteLegacyCampaign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteLegacyCampaign(context.TODO(), 986724)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteLegacyCampaign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteLegacyCampaign(context.TODO(), 986724)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacyCampaign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteLegacyCampaign(context.TODO(), 986724)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateLegacyCampaign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 986724,
			"title": "May Newsletter",
			"subject": "New Products for Summer!",
			"sender_id": 124451,
			"list_ids": [110, 124],
			"segment_ids": [110],
			"categories": ["summer line"],
			"suppression_group_id": 42,
			"ip_pool": "marketing",
			"html_content": "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>",
			"plain_content": "Check out our summer line!",
			"status": "Draft"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateLegacyCampaign(context.TODO(), 986724, &InputUpdateLegacyCampaign{
		Title:        "May Newsletter",
		Subject:      "New Products for Summer!",
		Categories:   []string{"summer line"},
		HTMLContent:  "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>",
		PlainContent: "Check out our summer line!",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateLegacyCampaign{
		ID:                 986724,
		Title:              "May Newsletter",
		Subject:            "New Products for Summer!",
		SenderID:           124451,
		ListIDs:            []int64{110, 124},
		SegmentIDs:         []int64{110},
		Categories:         []string{"summer line"},
		SuppressionGroupID: 42,
		IPPool:             "marketing",
		HTMLContent:        "<html><head><title></title></head><body><p>Check out our summer line!</p></body></html>",
		PlainContent:       "Check out our summer line!",
		Status:             "Draft",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateLegacyCampaign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateLegacyCampaign(context.TODO(), 986724, &InputUpdateLegacyCampaign{
		Title: "May Newsletter",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateLegacyCampaign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateLegacyCampaign(context.TODO(), 986724, &InputUpdateLegacyCampaign{
		Title: "May Newsletter",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestScheduleLegacyCampaign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234,
			"send_at": 1489771528,
			"status": "Scheduled"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.ScheduleLegacyCampaign(context.TODO(), 986724, &InputScheduleLegacyCampaign{
		SendAt: 1489771528,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputScheduleLegacyCampaign{
		ID:     1234,
		SendAt: 1489771528,
		Status: "Scheduled",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestScheduleLegacyCampaign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ScheduleLegacyCampaign(context.TODO(), 986724, &InputScheduleLegacyCampaign{
		SendAt: 1489771528,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestScheduleLegacyCampaign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.ScheduleLegacyCampaign(context.TODO(), 986724, &InputScheduleLegacyCampaign{
		SendAt: 1489771528,
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateLegacyCampaignSchedule(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234,
			"send_at": 1489451436,
			"status": "Scheduled"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateLegacyCampaignSchedule(context.TODO(), 986724, &InputUpdateLegacyCampaignSchedule{
		SendAt: 1489451436,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateLegacyCampaignSchedule{
		ID:     1234,
		SendAt: 1489451436,
		Status: "Scheduled",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateLegacyCampaignSchedule_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateLegacyCampaignSchedule(context.TODO(), 986724, &InputUpdateLegacyCampaignSchedule{
		SendAt: 1489451436,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateLegacyCampaignSchedule_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.UpdateLegacyCampaignSchedule(context.TODO(), 986724, &InputUpdateLegacyCampaignSchedule{
		SendAt: 1489451436,
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetLegacyCampaignSchedule(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"send_at": 1490778528
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetLegacyCampaignSchedule(context.TODO(), 986724)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetLegacyCampaignSchedule{
		SendAt: 1490778528,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetLegacyCampaignSchedule_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetLegacyCampaignSchedule(context.TODO(), 986724)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetLegacyCampaignSchedule_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetLegacyCampaignSchedule(context.TODO(), 986724)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteLegacyCampaignSchedule(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteLegacyCampaignSchedule(context.TODO(), 986724)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteLegacyCampaignSchedule_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteLegacyCampaignSchedule(context.TODO(), 986724)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteLegacyCampaignSchedule_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteLegacyCampaignSchedule(context.TODO(), 986724)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestSendLegacyCampaignNow(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules/now", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234,
			"status": "Scheduled"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.SendLegacyCampaignNow(context.TODO(), 986724)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputSendLegacyCampaignNow{
		ID:     1234,
		Status: "Scheduled",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestSendLegacyCampaignNow_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules/now", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.SendLegacyCampaignNow(context.TODO(), 986724)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestSendLegacyCampaignNow_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.SendLegacyCampaignNow(context.TODO(), 986724)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestSendLegacyCampaignTest(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules/test", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"to": "your.email@example.com"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	err := client.SendLegacyCampaignTest(context.TODO(), 986724, &InputSendLegacyCampaignTest{
		To: "your.email@example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestSendLegacyCampaignTest_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/campaigns/986724/schedules/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.SendLegacyCampaignTest(context.TODO(), 986724, &InputSendLegacyCampaignTest{
		To: "your.email@example.com",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestSendLegacyCampaignTest_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.SendLegacyCampaignTest(context.TODO(), 986724, &InputSendLegacyCampaignTest{
		To: "your.email@example.com",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
