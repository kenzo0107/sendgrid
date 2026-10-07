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

func TestGetEngagementQualityScores(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/engagementquality/scores", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"score": 2,
					"username": "jdoe123",
					"date": "2006-01-02",
					"user_id": "180",
					"metrics": {
						"engagement_recency": 4,
						"open_rate": 1,
						"bounce_classification": 2,
						"bounce_rate": 4,
						"spam_rate": 4
					}
				}
			]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetEngagementQualityScores(context.TODO(), &InputGetEngagementQualityScores{
		From: "2006-01-02",
		To:   "2006-01-02",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetEngagementQualityScores{
		Result: []*EngagementQualityScore{
			{
				Score:    2,
				Username: "jdoe123",
				Date:     "2006-01-02",
				UserID:   "180",
				Metrics: &EngagementQualityMetrics{
					EngagementRecency:    4,
					OpenRate:             1,
					BounceClassification: 2,
					BounceRate:           4,
					SpamRate:             4,
				},
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetEngagementQualityScores_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/engagementquality/scores", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetEngagementQualityScores(context.TODO(), &InputGetEngagementQualityScores{
		From: "2006-01-02",
		To:   "2006-01-02",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetEngagementQualityScores_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetEngagementQualityScores(context.TODO(), &InputGetEngagementQualityScores{
		From: "2006-01-02",
		To:   "2006-01-02",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSubuserEngagementQualityScores(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/engagementquality/subusers/scores", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"score": 2,
					"username": "jdoe123",
					"date": "2021-12-31",
					"user_id": "180",
					"metrics": {
						"engagement_recency": 4,
						"open_rate": 1,
						"bounce_classification": 2,
						"bounce_rate": 4,
						"spam_rate": 4
					}
				}
			],
			"_metadata": {
				"next_params": {
					"after_key": "180"
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubuserEngagementQualityScores(context.TODO(), &InputGetSubuserEngagementQualityScores{
		Date: "2021-12-31",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSubuserEngagementQualityScores{
		Result: []*EngagementQualityScore{
			{
				Score:    2,
				Username: "jdoe123",
				Date:     "2021-12-31",
				UserID:   "180",
				Metrics: &EngagementQualityMetrics{
					EngagementRecency:    4,
					OpenRate:             1,
					BounceClassification: 2,
					BounceRate:           4,
					SpamRate:             4,
				},
			},
		},
		Metadata: &EngagementQualityMetadata{
			NextParams: &EngagementQualityNextParams{
				AfterKey: "180",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSubuserEngagementQualityScores_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/engagementquality/subusers/scores", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuserEngagementQualityScores(context.TODO(), &InputGetSubuserEngagementQualityScores{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSubuserEngagementQualityScores_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetSubuserEngagementQualityScores(context.TODO(), &InputGetSubuserEngagementQualityScores{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
