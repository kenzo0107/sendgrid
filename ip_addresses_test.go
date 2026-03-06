package sendgrid

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetIPAddresses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`[
			{
				"ip":"192.168.1.1",
				"pools":["pool1"],
				"warmup":true,
				"start_date":1609459200,
				"subusers":["subuser1"],
				"rdns":"example.com",
				"assigned_at":1609459300
			}]`))
	})

	ctx := context.Background()
	expected, err := client.GetIPAddresses(ctx, &InputGetIPAddresses{
		IP:                 "192.168.1.1",
		ExcludeWhitelabels: true,
		Limit:              100,
		Offset:             10,
		Subuser:            "subuser1",
		SortByDirection:    "asc",
	})
	if err != nil {
		assert.Fail(t, "GetIPAddresses returned an error: %v", err)
		return
	}

	want := []*IPAddress{{
		IP:           "192.168.1.1",
		Pools:        []string{"pool1"},
		Warmup:       true,
		StartDate:    1609459200,
		Subusers:     []string{"subuser1"},
		Rdns:         "example.com",
		AssignedAt:   1609459300,
		Whitelabeled: false,
	}}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIPAddresses_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ips", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetIPAddresses(ctx, &InputGetIPAddresses{})
	assert.Error(t, err)
}

func TestGetIPAddresses_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetIPAddresses(ctx, &InputGetIPAddresses{Limit: -1})
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetAssignedIPAddresses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/assigned", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`[
			{
				"ip":"192.168.1.1",
				"pools":["pool1"],
				"warmup":true,
				"start_date":1609459200
			}]`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}

	})

	ctx := context.Background()
	expected, err := client.GetAssignedIPAddresses(ctx)
	if err != nil {
		assert.Fail(t, "GetAssignedIPAddresses returned an error: %v", err)
		return
	}

	want := []*AssignedIPAddress{{
		IP:        "192.168.1.1",
		Pools:     []string{"pool1"},
		Warmup:    true,
		StartDate: 1609459200,
	}}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetAssignedIPAddresses_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ips/assigned", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetAssignedIPAddresses(ctx)
	assert.Error(t, err)
}

func TestGetAssignedIPAddresses_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetAssignedIPAddresses(ctx)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetRemainingIPCount(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/remaining", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"results":[
				{
					"remaining":5,
					"period":"month",
					"price_per_ip":100
				}
			]
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	expected, err := client.GetRemainingIPCount(ctx)
	if err != nil {
		assert.Fail(t, "GetRemainingIPCount returned an error: %v", err)
		return
	}

	want := &OutputGetRemainingIPCount{
		Results: []*RemainingIPCount{{
			Remaining:  5,
			Period:     "month",
			PricePerIP: 100,
		}},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetRemainingIPCount_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ips/remaining", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetRemainingIPCount(ctx)
	assert.Error(t, err)
}

func TestGetRemainingIPCount_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetRemainingIPCount(ctx)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetIPAddress(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"ip":"192.168.1.1",
			"pools":["pool1"],
			"warmup":true,
			"start_date":1609459200,
			"subusers":["subuser1"],
			"rdns":"example.com",
			"whitelabeled":false
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	expected, err := client.GetIPAddress(ctx, "192.168.1.1")
	if err != nil {
		assert.Fail(t, "GetIPAddress returned an error: %v", err)
		return
	}

	want := &OutputGetIPAddress{
		IP:           "192.168.1.1",
		Pools:        []string{"pool1"},
		Warmup:       true,
		StartDate:    1609459200,
		Subusers:     []string{"subuser1"},
		Rdns:         "example.com",
		Whitelabeled: false,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIPAddress_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ips/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetIPAddress(ctx, "192.168.1.1")
	assert.Error(t, err)
}

func TestGetIPAddress_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetIPAddress(ctx, "192.168.1.1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}
