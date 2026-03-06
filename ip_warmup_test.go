package sendgrid

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartIPWarmup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`[
			{"ip":"192.168.0.1", "start_date":1700000000}
		]`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	expected, err := client.StartIPWarmup(ctx, &InputStartIPWarmup{IP: "192.168.0.1"})
	if err != nil {
		assert.Fail(t, "StartIPWarmup returned an error: %v", err)
	}

	want := []*IPWarmup{
		{IP: "192.168.0.1", StartDate: 1700000000},
	}

	if !assert.Equal(t, want, expected) {
		assert.Fail(t, "StartIPWarmup returned unexpected result: got %v, want %v", expected, want)
	}
}

func TestStartIPWarmup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, err := client.StartIPWarmup(ctx, &InputStartIPWarmup{IP: "192.168.0.1"})
	if err == nil {
		assert.Fail(t, "Expected error, got nil")
	}
}

func TestStartIPWarmup_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.StartIPWarmup(ctx, nil)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestStopIPWarmup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.StopIPWarmup(ctx, "192.168.0.1")
	if err != nil {
		assert.Fail(t, "StopIPWarmup returned an error: %v", err)
	}
}

func TestStopIPWarmup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})

	ctx := context.Background()
	err := client.StopIPWarmup(ctx, "192.168.0.1")
	assert.Error(t, err)
}

func TestStopIPWarmup_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	err := client.StopIPWarmup(ctx, "192.168.0.1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetIPWarmupStatus(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup/192.168.0.1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`[
			{"ip":"192.168.0.1", "start_date":1700000000}
		]`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	status, err := client.GetIPWarmupStatus(ctx, "192.168.0.1")
	if err != nil {
		assert.Fail(t, "GetIPWarmupStatus returned an error: %v", err)
	}

	want := []*IPWarmup{
		{
			IP:        "192.168.0.1",
			StartDate: 1700000000,
		},
	}

	if !assert.Equal(t, want, status) {
		assert.Fail(t, "GetIPWarmupStatus returned unexpected result: got %v, want %v", status, want)
	}
}

func TestGetIPWarmupStatus_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup/192.168.0.1", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, err := client.GetIPWarmupStatus(ctx, "192.168.0.1")
	assert.Error(t, err)
}

func TestGetIPWarmupStatus_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetIPWarmupStatus(ctx, "192.168.0.1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetAllIPWarmup(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`[
			{"ip":"192.168.0.1", "start_date":1700000000},
			{"ip":"192.168.0.2", "start_date":1700000001}
		]`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	status, err := client.GetAllIPWarmup(ctx)
	if err != nil {
		assert.Fail(t, "GetAllIPWarmup returned an error: %v", err)
	}

	want := []*IPWarmup{
		{
			IP:        "192.168.0.1",
			StartDate: 1700000000,
		},
		{
			IP:        "192.168.0.2",
			StartDate: 1700000001,
		},
	}

	if !assert.Equal(t, want, status) {
		assert.Fail(t, "GetAllIPWarmup returned unexpected result: got %v, want %v", status, want)
	}
}

func TestGetAllIPWarmup_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/warmup", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})

	ctx := context.Background()
	_, err := client.GetAllIPWarmup(ctx)
	assert.Error(t, err)
}

func TestGetAllIPWarmup_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetAllIPWarmup(ctx)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}
