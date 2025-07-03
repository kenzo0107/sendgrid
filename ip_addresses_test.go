package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPAddressStruct(t *testing.T) {
	ip := IPAddress{
		IP:         "192.168.1.1",
		Pools:      []string{"pool1", "pool2"},
		Warmup:     true,
		StartDate:  1609459200,
		Subusers:   []string{"subuser1"},
		Rdns:       "example.com",
		AssignedAt: 1609459300,
	}

	assert.Equal(t, "192.168.1.1", ip.IP)
	assert.Len(t, ip.Pools, 2)
	assert.Equal(t, "pool1", ip.Pools[0])
	assert.Equal(t, "pool2", ip.Pools[1])
	assert.True(t, ip.Warmup)
	assert.Equal(t, int64(1609459200), ip.StartDate)
	assert.Len(t, ip.Subusers, 1)
	assert.Equal(t, "subuser1", ip.Subusers[0])
	assert.Equal(t, "example.com", ip.Rdns)
	assert.Equal(t, int64(1609459300), ip.AssignedAt)
}

func TestIPPoolStruct(t *testing.T) {
	pool := IPPool{
		Name: "test-pool",
	}

	assert.Equal(t, "test-pool", pool.Name)
}

func TestIPWarmupStatusStruct(t *testing.T) {
	status := IPWarmupStatus{
		IP:     "192.168.1.1",
		Warmup: true,
	}

	assert.Equal(t, "192.168.1.1", status.IP)
	assert.True(t, status.Warmup)
}

func TestInputAddIPToPoolStruct(t *testing.T) {
	input := InputAddIPToPool{
		IP: "192.168.1.1",
	}

	assert.Equal(t, "192.168.1.1", input.IP)
}

func TestInputAssignIPToSubuserStruct(t *testing.T) {
	input := InputAssignIPToSubuser{
		IPs: []string{"192.168.1.1", "192.168.1.2"},
	}

	assert.Len(t, input.IPs, 2)
	assert.Equal(t, "192.168.1.1", input.IPs[0])
	assert.Equal(t, "192.168.1.2", input.IPs[1])
}

func TestOutputAssignedIPsStruct(t *testing.T) {
	output := OutputAssignedIPs{
		IPs: []string{"192.168.1.1", "192.168.1.2"},
	}

	assert.Len(t, output.IPs, 2)
	assert.Equal(t, "192.168.1.1", output.IPs[0])
	assert.Equal(t, "192.168.1.2", output.IPs[1])
}

func TestGetIPAddresses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"ip":"192.168.1.1","pools":["pool1"],"warmup":true}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	ips, err := client.GetIPAddresses(ctx)
	assert.NoError(t, err)
	assert.Len(t, ips, 1)
	assert.Equal(t, "192.168.1.1", ips[0].IP)
	assert.Equal(t, []string{"pool1"}, ips[0].Pools)
	assert.True(t, ips[0].Warmup)
}

func TestGetIPAddresses_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetIPAddresses(ctx)
	assert.Error(t, err)
}

func TestGetAssignedIPAddresses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/assigned", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"ip":"192.168.1.2","pools":["pool2"],"warmup":false}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	ips, err := client.GetAssignedIPAddresses(ctx)
	assert.NoError(t, err)
	assert.Len(t, ips, 1)
	assert.Equal(t, "192.168.1.2", ips[0].IP)
	assert.Equal(t, []string{"pool2"}, ips[0].Pools)
	assert.False(t, ips[0].Warmup)
}

func TestGetAssignedIPAddresses_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetAssignedIPAddresses(ctx)
	assert.Error(t, err)
}

func TestGetRemainingIPCount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/remaining", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"remaining":5}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	result, err := client.GetRemainingIPCount(ctx)
	assert.NoError(t, err)
	assert.Equal(t, float64(5), result["remaining"])
}

func TestGetRemainingIPCount_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetRemainingIPCount(ctx)
	assert.Error(t, err)
}

func TestGetIPAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/192.168.1.1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ip":"192.168.1.1","pools":["pool1"],"warmup":true}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	ip, err := client.GetIPAddress(ctx, "192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.1.1", ip.IP)
	assert.Equal(t, []string{"pool1"}, ip.Pools)
	assert.True(t, ip.Warmup)
}

func TestGetIPAddress_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetIPAddress(ctx, "192.168.1.1")
	assert.Error(t, err)
}

func TestGetIPPools(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/pools", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"name":"pool1"},{"name":"pool2"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	pools, err := client.GetIPPools(ctx)
	assert.NoError(t, err)
	assert.Len(t, pools, 2)
	assert.Equal(t, "pool1", pools[0].Name)
	assert.Equal(t, "pool2", pools[1].Name)
}

func TestGetIPPools_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetIPPools(ctx)
	assert.Error(t, err)
}

func TestCreateIPPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/ips/pools", r.URL.Path)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"name":"test-pool"}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	pool, err := client.CreateIPPool(ctx, "test-pool")
	assert.NoError(t, err)
	assert.Equal(t, "test-pool", pool.Name)
}

func TestCreateIPPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.CreateIPPool(ctx, "test-pool")
	assert.Error(t, err)
}

func TestGetIPPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/pools/test-pool", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"name":"test-pool"}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	pool, err := client.GetIPPool(ctx, "test-pool")
	assert.NoError(t, err)
	assert.Equal(t, "test-pool", pool.Name)
}

func TestGetIPPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetIPPool(ctx, "test-pool")
	assert.Error(t, err)
}

func TestUpdateIPPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/ips/pools/old-pool", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"name":"new-pool"}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	pool, err := client.UpdateIPPool(ctx, "old-pool", "new-pool")
	assert.NoError(t, err)
	assert.Equal(t, "new-pool", pool.Name)
}

func TestUpdateIPPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.UpdateIPPool(ctx, "old-pool", "new-pool")
	assert.Error(t, err)
}

func TestDeleteIPPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/ips/pools/test-pool", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteIPPool(ctx, "test-pool")
	assert.NoError(t, err)
}

func TestDeleteIPPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteIPPool(ctx, "test-pool")
	assert.Error(t, err)
}

func TestAddIPToPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/ips/pools/test-pool/ips", r.URL.Path)
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.AddIPToPool(ctx, "test-pool", "192.168.1.1")
	assert.NoError(t, err)
}

func TestAddIPToPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.AddIPToPool(ctx, "test-pool", "192.168.1.1")
	assert.Error(t, err)
}

func TestRemoveIPFromPool(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/ips/pools/test-pool/ips/192.168.1.1", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.RemoveIPFromPool(ctx, "test-pool", "192.168.1.1")
	assert.NoError(t, err)
}

func TestRemoveIPFromPool_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.RemoveIPFromPool(ctx, "test-pool", "192.168.1.1")
	assert.Error(t, err)
}

func TestStartIPWarmup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/ips/warmup/192.168.1.1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ip":"192.168.1.1","warmup":true}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	status, err := client.StartIPWarmup(ctx, "192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.1.1", status.IP)
	assert.True(t, status.Warmup)
}

func TestStartIPWarmup_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.StartIPWarmup(ctx, "192.168.1.1")
	assert.Error(t, err)
}

func TestStopIPWarmup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/ips/warmup/192.168.1.1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ip":"192.168.1.1","warmup":false}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	status, err := client.StopIPWarmup(ctx, "192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.1.1", status.IP)
	assert.False(t, status.Warmup)
}

func TestStopIPWarmup_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.StopIPWarmup(ctx, "192.168.1.1")
	assert.Error(t, err)
}

func TestGetIPWarmupStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/warmup/192.168.1.1", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ip":"192.168.1.1","warmup":true}`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	status, err := client.GetIPWarmupStatus(ctx, "192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.1.1", status.IP)
	assert.True(t, status.Warmup)
}

func TestGetIPWarmupStatus_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetIPWarmupStatus(ctx, "192.168.1.1")
	assert.Error(t, err)
}

func TestGetAllIPWarmupStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ips/warmup", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"ip":"192.168.1.1","warmup":true},{"ip":"192.168.1.2","warmup":false}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	statuses, err := client.GetAllIPWarmupStatus(ctx)
	assert.NoError(t, err)
	assert.Len(t, statuses, 2)
	assert.Equal(t, "192.168.1.1", statuses[0].IP)
	assert.True(t, statuses[0].Warmup)
	assert.Equal(t, "192.168.1.2", statuses[1].IP)
	assert.False(t, statuses[1].Warmup)
}

func TestGetAllIPWarmupStatus_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetAllIPWarmupStatus(ctx)
	assert.Error(t, err)
}
