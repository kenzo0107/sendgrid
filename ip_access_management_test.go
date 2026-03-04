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

func TestGetAllowedIP(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist/1", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
			"id": 1,
			"ip": "192.168.1.1",
			"created_at": 1443651141,
			"updated_at": 1443651141
		}]
	}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetAllowedIP(context.TODO(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAllowedIP{
		Result: []*AllowedIP{
			{
				ID:        1,
				IP:        "192.168.1.1",
				CreatedAt: 1443651141,
				UpdatedAt: 1443651141,
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetAllowedIP_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAllowedIP(context.TODO(), 1)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetAllowedIP_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetAllowedIP(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetAllowedIPs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": 1,
				"ip": "192.168.1.1",
				"created_at": 1443651141,
				"updated_at": 1443651141
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetAllowedIPs(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAllowedIPs{
		Result: []*AllowedIP{
			{
				ID:        1,
				IP:        "192.168.1.1",
				CreatedAt: 1443651141,
				UpdatedAt: 1443651141,
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetAllowedIPs_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAllowedIPs(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetAllowedIPs_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetAllowedIPs(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestAddIPsToAllowList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": 1,
				"ip": "192.168.1.1",
				"created_at": 1443651141,
				"updated_at": 1443651141
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.AddIPsToAllowList(context.TODO(), &InputAddIPsToAllowList{
		IPs: []AllowListIP{
			{IP: "192.168.1.1"},
		},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAddIPsToAllowList{
		Result: []*AllowedIP{
			{
				ID:        1,
				IP:        "192.168.1.1",
				CreatedAt: 1443651141,
				UpdatedAt: 1443651141,
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestAddIPsToAllowList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddIPsToAllowList(context.TODO(), &InputAddIPsToAllowList{
		IPs: []AllowListIP{
			{IP: "192.168.1.1"},
		},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestAddIPsToAllowList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.AddIPsToAllowList(context.TODO(), &InputAddIPsToAllowList{
		IPs: []AllowListIP{
			{IP: "192.168.1.1"},
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

func TestRemoveIPsFromAllowList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.RemoveIPsFromAllowList(context.TODO(), &InputRemoveIPsFromAllowList{
		IDs: []int64{1, 2},
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestRemoveIPsFromAllowList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.RemoveIPsFromAllowList(context.TODO(), &InputRemoveIPsFromAllowList{
		IDs: []int64{1, 2},
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestRemoveIPsFromAllowList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.RemoveIPsFromAllowList(context.TODO(), &InputRemoveIPsFromAllowList{
		IDs: []int64{1, 2},
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetAccessActivity(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/activity", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"allowed": true,
				"auth_method": "basic",
				"first_at": 1443651141,
				"ip": "192.168.1.1",
				"last_at": 1443651141,
				"location": "US"
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.GetAccessActivities(context.TODO(), &InputGetAccessActivities{Limit: 20})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAccessActivities{
		Result: []*AccessActivity{
			{
				Allowed:    true,
				AuthMethod: "basic",
				FirstAt:    1443651141,
				IP:         "192.168.1.1",
				LastAt:     1443651141,
				Location:   "US",
			},
		},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestGetAccessActivity_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/activity", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAccessActivities(context.TODO(), &InputGetAccessActivities{Limit: 20})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetAccessActivities_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetAccessActivities(context.TODO(), &InputGetAccessActivities{Limit: 20})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestRemoveIPFromAllowList(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.RemoveIPFromAllowList(context.TODO(), 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestRemoveIPFromAllowList_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/access_settings/whitelist/1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.RemoveIPFromAllowList(context.TODO(), 1)
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestRemoveIPFromAllowList_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.RemoveIPFromAllowList(context.TODO(), 1)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
