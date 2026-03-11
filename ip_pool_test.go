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

func TestGetIPPools(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`[
			{"name":"pool1"},
			{"name":"pool2"}
		]`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	expected, err := client.GetIPPools(ctx)
	if err != nil {
		assert.Fail(t, "GetIPPools returned an error: %v", err)
		return
	}

	want := []*IPPool{
		{Name: "pool1"},
		{Name: "pool2"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestTestGetIPPools_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetIPPools(ctx)
	assert.Error(t, err)
}

func TestTestGetIPPools_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetIPPools(ctx)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestGetIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"pool_name":"pool1",
			"ips":[
				"192.168.0.1",
				"192.168.0.2"
			]
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	expected, err := client.GetIPPool(ctx, "pool1")
	if err != nil {
		assert.Fail(t, "GetIPPool returned an error: %v", err)
		return
	}

	want := &OutputGetIPPool{
		PoolName: "pool1",
		IPs:      []string{"192.168.0.1", "192.168.0.2"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	_, err := client.GetIPPool(ctx, "pool1")
	assert.Error(t, err)
}

func TestGetIPPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	_, err := client.GetIPPool(ctx, "pool1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestCreateIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"name":"pool1"
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	input := &InputCreateIPPool{Name: "pool1"}
	expected, err := client.CreateIPPool(ctx, input)
	if err != nil {
		assert.Fail(t, "CreateIPPool returned an error: %v", err)
		return
	}

	want := &OutputCreateIPPool{Name: "pool1"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	input := &InputCreateIPPool{Name: "pool1"}
	_, err := client.CreateIPPool(ctx, input)
	assert.Error(t, err)
}

func TestCreateIPPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	input := &InputCreateIPPool{Name: "pool1"}
	_, err := client.CreateIPPool(ctx, input)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestUpdateIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"name":"pool1"
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	input := &InputUpdateIPPool{Name: "pool1"}
	expected, err := client.UpdateIPPool(ctx, "pool1", input)
	if err != nil {
		assert.Fail(t, "UpdateIPPool returned an error: %v", err)
		return
	}

	want := &OutputUpdateIPPool{Name: "pool1"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	input := &InputUpdateIPPool{Name: "pool1"}
	_, err := client.UpdateIPPool(ctx, "pool1", input)
	assert.Error(t, err)
}

func TestUpdateIPPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	input := &InputUpdateIPPool{Name: "pool1"}
	_, err := client.UpdateIPPool(ctx, "pool1", input)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestDeleteIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.DeleteIPPool(ctx, "pool1")
	if err != nil {
		assert.Fail(t, "DeleteIPPool returned an error: %v", err)
	}
}

func TestDeleteIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	err := client.DeleteIPPool(ctx, "pool1")
	assert.Error(t, err)
}

func TestDeleteIPPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	err := client.DeleteIPPool(ctx, "pool1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestAddIPToPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1/ips", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{
			"ip":"192.168.0.1"
		}`)); err != nil {
			assert.Fail(t, "Failed to write response: %v", err)
		}
	})

	ctx := context.Background()
	input := &InputAddIPToPool{IP: "192.168.0.1"}
	expected, err := client.AddIPToPool(ctx, "pool1", input)
	if err != nil {
		assert.Fail(t, "AddIPToPool returned an error: %v", err)
		return
	}

	want := &OutputAddIPToPool{IP: "192.168.0.1"}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddIPToPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	input := &InputAddIPToPool{IP: "192.168.0.1"}
	_, err := client.AddIPToPool(ctx, "pool1", input)
	assert.Error(t, err)
}

func TestAddIPToPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	input := &InputAddIPToPool{IP: "192.168.0.1"}
	_, err := client.AddIPToPool(ctx, "pool1", input)
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}

func TestRemoveIPFromPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1/ips/192.168.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := client.RemoveIPFromPool(ctx, "pool1", "192.168.0.1")
	if err != nil {
		assert.Fail(t, "RemoveIPFromPool returned an error: %v", err)
	}
}

func TestRemoveIPFromPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ips/pools/pool1/ips/192.168.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "Internal server error"}`))
	})

	ctx := context.Background()
	err := client.RemoveIPFromPool(ctx, "pool1", "192.168.0.1")
	assert.Error(t, err)
}

func TestRemoveIPFromPool_NewRequestFailed(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	ctx := context.Background()
	err := client.RemoveIPFromPool(ctx, "pool1", "192.168.0.1")
	assert.Error(t, err)

	client.baseURL = originalBaseURL
}
