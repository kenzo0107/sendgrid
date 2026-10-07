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

func withInvalidBaseURL(client *Client, fn func()) {
	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL
	fn()
	client.baseURL = originalBaseURL
}

func assertNewRequestError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected error for invalid baseURL")
		return
	}
	if !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}
}

// GetSendIPs

func TestGetSendIPs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"ip": "127.0.0.1",
					"pools": [{"name": "marketing_pool", "id": "12345"}],
					"is_auto_warmup": true,
					"is_parent_assigned": true,
					"is_enabled": true,
					"is_leased": true,
					"added_at": 1664390835,
					"region": "us"
				}
			],
			"_metadata": {
				"next_params": {"after_key": null}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIPs(context.TODO(), &InputGetSendIPs{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIPs{
		Result: []*SendIP{
			{
				IP:               "127.0.0.1",
				Pools:            []*SendIPPoolRef{{Name: "marketing_pool", ID: "12345"}},
				IsAutoWarmup:     true,
				IsParentAssigned: true,
				IsEnabled:        true,
				IsLeased:         true,
				AddedAt:          1664390835,
				Region:           "us",
			},
		},
		Metadata: &OutputGetSendIPsMetadata{
			NextParams: &OutputGetSendIPsMetadataNextParams{},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIPs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIPs(context.TODO(), &InputGetSendIPs{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIPs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIPs(context.TODO(), &InputGetSendIPs{})
		assertNewRequestError(t, err)
	})
}

// AddSendIP

func TestAddSendIP(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		if _, err := fmt.Fprint(w, `{
			"ip": "127.0.0.1",
			"is_auto_warmup": true,
			"is_parent_assigned": true,
			"subusers": ["12345"]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AddSendIP(context.TODO(), &InputAddSendIP{
		IsAutoWarmup:     true,
		IsParentAssigned: true,
		Subusers:         []string{"12345"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAddSendIP{
		IP:               "127.0.0.1",
		IsAutoWarmup:     true,
		IsParentAssigned: true,
		Subusers:         []string{"12345"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddSendIP_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddSendIP(context.TODO(), &InputAddSendIP{
		IsAutoWarmup:     true,
		IsParentAssigned: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddSendIP_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.AddSendIP(context.TODO(), &InputAddSendIP{
			IsAutoWarmup:     true,
			IsParentAssigned: true,
		})
		assertNewRequestError(t, err)
	})
}

// GetSendIP

func TestGetSendIP(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"ip": "127.0.0.1",
			"is_parent_assigned": true,
			"is_auto_warmup": true,
			"pools": [{"id": "12345", "name": "transactional_pool"}],
			"added_at": 1664390835,
			"is_enabled": true,
			"is_leased": true,
			"region": "us"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIP(context.TODO(), "127.0.0.1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIP{
		IP:               "127.0.0.1",
		IsParentAssigned: true,
		IsAutoWarmup:     true,
		Pools:            []*SendIPPoolRef{{ID: "12345", Name: "transactional_pool"}},
		AddedAt:          1664390835,
		IsEnabled:        true,
		IsLeased:         true,
		Region:           "us",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIP_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIP(context.TODO(), "127.0.0.1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIP_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIP(context.TODO(), "127.0.0.1")
		assertNewRequestError(t, err)
	})
}

// UpdateSendIP

func TestUpdateSendIP(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		if _, err := fmt.Fprint(w, `{
			"ip": "127.0.0.1",
			"is_auto_warmup": true,
			"is_parent_assigned": true,
			"is_enabled": true
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateSendIP(context.TODO(), "127.0.0.1", &InputUpdateSendIP{
		IsAutoWarmup:     true,
		IsParentAssigned: true,
		IsEnabled:        true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateSendIP{
		IP:               "127.0.0.1",
		IsAutoWarmup:     true,
		IsParentAssigned: true,
		IsEnabled:        true,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateSendIP_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateSendIP(context.TODO(), "127.0.0.1", &InputUpdateSendIP{
		IsAutoWarmup: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSendIP_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.UpdateSendIP(context.TODO(), "127.0.0.1", &InputUpdateSendIP{
			IsAutoWarmup: true,
		})
		assertNewRequestError(t, err)
	})
}

// GetSendIPSubusers

func TestGetSendIPSubusers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"result": ["12345", "67890"],
			"_metadata": {"next_params": {"after_key": null}}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIPSubusers(context.TODO(), "127.0.0.1", &InputGetSendIPSubusers{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIPSubusers{
		Result: []string{"12345", "67890"},
		Metadata: &OutputGetSendIPSubusersMetadata{
			NextParams: &OutputGetSendIPSubusersMetadataNextParams{},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIPSubusers_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIPSubusers(context.TODO(), "127.0.0.1", &InputGetSendIPSubusers{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIPSubusers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIPSubusers(context.TODO(), "127.0.0.1", &InputGetSendIPSubusers{})
		assertNewRequestError(t, err)
	})
}

// BatchAddSendIPSubusers

func TestBatchAddSendIPSubusers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers:batchAdd", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := fmt.Fprint(w, `{
			"ip": "127.0.0.1",
			"subusers": ["12345", "67890"]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.BatchAddSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchAddSendIPSubusers{
		Subusers: []string{"12345", "67890"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputBatchAddSendIPSubusers{
		IP:       "127.0.0.1",
		Subusers: []string{"12345", "67890"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestBatchAddSendIPSubusers_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers:batchAdd", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.BatchAddSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchAddSendIPSubusers{
		Subusers: []string{"12345"},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestBatchAddSendIPSubusers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.BatchAddSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchAddSendIPSubusers{
			Subusers: []string{"12345"},
		})
		assertNewRequestError(t, err)
	})
}

// BatchDeleteSendIPSubusers

func TestBatchDeleteSendIPSubusers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers:batchDelete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	expected, err := client.BatchDeleteSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchDeleteSendIPSubusers{
		Subusers: []string{"12345"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputBatchDeleteSendIPSubusers{}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestBatchDeleteSendIPSubusers_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/ips/127.0.0.1/subusers:batchDelete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.BatchDeleteSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchDeleteSendIPSubusers{
		Subusers: []string{"12345"},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestBatchDeleteSendIPSubusers_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.BatchDeleteSendIPSubusers(context.TODO(), "127.0.0.1", &InputBatchDeleteSendIPSubusers{
			Subusers: []string{"12345"},
		})
		assertNewRequestError(t, err)
	})
}

// GetSendIPPools

func TestGetSendIPPools(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"name": "transactional_pool",
					"id": "12345",
					"ips_preview": ["127.0.0.1", "127.0.0.2"],
					"total_ip_count": 2
				}
			],
			"_metadata": {"next_params": {"after_key": null}}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIPPools(context.TODO(), &InputGetSendIPPools{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIPPools{
		Result: []*SendIPPool{
			{
				Name:         "transactional_pool",
				ID:           "12345",
				IPsPreview:   []string{"127.0.0.1", "127.0.0.2"},
				TotalIPCount: 2,
			},
		},
		Metadata: &OutputGetSendIPPoolsMetadata{
			NextParams: &OutputGetSendIPPoolsMetadataNextParams{},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIPPools_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIPPools(context.TODO(), &InputGetSendIPPools{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIPPools_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIPPools(context.TODO(), &InputGetSendIPPools{})
		assertNewRequestError(t, err)
	})
}

// CreateSendIPPool

func TestCreateSendIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		if _, err := fmt.Fprint(w, `{
			"name": "transactional_pool",
			"id": "12345",
			"ips": ["127.0.0.1", "127.0.0.2"]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateSendIPPool(context.TODO(), &InputCreateSendIPPool{
		Name: "transactional_pool",
		IPs:  []string{"127.0.0.1", "127.0.0.2"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateSendIPPool{
		Name: "transactional_pool",
		ID:   "12345",
		IPs:  []string{"127.0.0.1", "127.0.0.2"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateSendIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateSendIPPool(context.TODO(), &InputCreateSendIPPool{
		Name: "transactional_pool",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateSendIPPool_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.CreateSendIPPool(context.TODO(), &InputCreateSendIPPool{
			Name: "transactional_pool",
		})
		assertNewRequestError(t, err)
	})
}

// GetSendIPPool

func TestGetSendIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"name": "marketing_pool",
			"id": "12345",
			"ips_preview": ["127.0.0.1", "127.0.0.2"],
			"total_ip_count": 2
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIPPool(context.TODO(), "12345")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIPPool{
		Name:         "marketing_pool",
		ID:           "12345",
		IPsPreview:   []string{"127.0.0.1", "127.0.0.2"},
		TotalIPCount: 2,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIPPool(context.TODO(), "12345")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIPPool_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIPPool(context.TODO(), "12345")
		assertNewRequestError(t, err)
	})
}

// UpdateSendIPPool

func TestUpdateSendIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if _, err := fmt.Fprint(w, `{
			"name": "marketing_pool",
			"id": "12345"
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateSendIPPool(context.TODO(), "12345", &InputUpdateSendIPPool{
		Name: "marketing_pool",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateSendIPPool{
		Name: "marketing_pool",
		ID:   "12345",
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateSendIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateSendIPPool(context.TODO(), "12345", &InputUpdateSendIPPool{
		Name: "marketing_pool",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateSendIPPool_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.UpdateSendIPPool(context.TODO(), "12345", &InputUpdateSendIPPool{
			Name: "marketing_pool",
		})
		assertNewRequestError(t, err)
	})
}

// DeleteSendIPPool

func TestDeleteSendIPPool(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if err := client.DeleteSendIPPool(context.TODO(), "12345"); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestDeleteSendIPPool_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if err := client.DeleteSendIPPool(context.TODO(), "12345"); err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteSendIPPool_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		err := client.DeleteSendIPPool(context.TODO(), "12345")
		assertNewRequestError(t, err)
	})
}

// GetSendIPPoolIPs

func TestGetSendIPPoolIPs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if _, err := fmt.Fprint(w, `{
			"result": [
				{
					"ip": "127.0.0.1",
					"pools": [{"id": "12345", "name": "transactional_pool"}]
				}
			],
			"_metadata": {"next_params": {"after_key": null}}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSendIPPoolIPs(context.TODO(), "12345", &InputGetSendIPPoolIPs{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSendIPPoolIPs{
		Result: []*SendIPPoolIP{
			{
				IP:    "127.0.0.1",
				Pools: []*SendIPPoolRef{{ID: "12345", Name: "transactional_pool"}},
			},
		},
		Metadata: &OutputGetSendIPPoolIPsMetadata{
			NextParams: &OutputGetSendIPPoolIPsMetadataNextParams{},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSendIPPoolIPs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSendIPPoolIPs(context.TODO(), "12345", &InputGetSendIPPoolIPs{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSendIPPoolIPs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.GetSendIPPoolIPs(context.TODO(), "12345", &InputGetSendIPPoolIPs{})
		assertNewRequestError(t, err)
	})
}

// BatchAddSendIPPoolIPs

func TestBatchAddSendIPPoolIPs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips:batchAdd", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		if _, err := fmt.Fprint(w, `{
			"name": "marketing_pool",
			"id": "12345",
			"ips": ["127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4"]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.BatchAddSendIPPoolIPs(context.TODO(), "12345", &InputBatchAddSendIPPoolIPs{
		IPs: []string{"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputBatchAddSendIPPoolIPs{
		Name: "marketing_pool",
		ID:   "12345",
		IPs:  []string{"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4"},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestBatchAddSendIPPoolIPs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips:batchAdd", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.BatchAddSendIPPoolIPs(context.TODO(), "12345", &InputBatchAddSendIPPoolIPs{
		IPs: []string{"127.0.0.1"},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestBatchAddSendIPPoolIPs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.BatchAddSendIPPoolIPs(context.TODO(), "12345", &InputBatchAddSendIPPoolIPs{
			IPs: []string{"127.0.0.1"},
		})
		assertNewRequestError(t, err)
	})
}

// BatchDeleteSendIPPoolIPs

func TestBatchDeleteSendIPPoolIPs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips:batchDelete", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusNoContent)
	})

	expected, err := client.BatchDeleteSendIPPoolIPs(context.TODO(), "12345", &InputBatchDeleteSendIPPoolIPs{
		IPs: []string{"127.0.0.1"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputBatchDeleteSendIPPoolIPs{}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestBatchDeleteSendIPPoolIPs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/send_ips/pools/12345/ips:batchDelete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.BatchDeleteSendIPPoolIPs(context.TODO(), "12345", &InputBatchDeleteSendIPPoolIPs{
		IPs: []string{"127.0.0.1"},
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestBatchDeleteSendIPPoolIPs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	withInvalidBaseURL(client, func() {
		_, err := client.BatchDeleteSendIPPoolIPs(context.TODO(), "12345", &InputBatchDeleteSendIPPoolIPs{
			IPs: []string{"127.0.0.1"},
		})
		assertNewRequestError(t, err)
	})
}
