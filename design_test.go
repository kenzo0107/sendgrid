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

func TestGetDesigns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "12345678-90ab-1234-56cd-efghijk78901",
				"updated_at": "2024-05-22T01:59:57Z",
				"created_at": "2024-05-22T01:59:57Z",
				"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
				"name": "example",
				"editor": "code",
				"subject": "Hello",
				"categories": ["welcome"],
				"generate_plain_content": true
			}],
			"_metadata": {
				"count": 1,
				"prev": "",
				"next": "",
				"self": "https://api.sendgrid.com/v3/designs?page_token=xxx"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetDesigns(context.TODO(), &InputGetDesigns{})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetDesigns{
		Result: []*Design{
			{
				ID:                   "12345678-90ab-1234-56cd-efghijk78901",
				UpdatedAt:            "2024-05-22T01:59:57Z",
				CreatedAt:            "2024-05-22T01:59:57Z",
				ThumbnailURL:         "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
				Name:                 "example",
				Editor:               "code",
				Subject:              "Hello",
				Categories:           []string{"welcome"},
				GeneratePlainContent: true,
			},
		},
		Metadata: _Metadata{
			Count: 1,
			Prev:  "",
			Next:  "",
			Self:  "https://api.sendgrid.com/v3/designs?page_token=xxx",
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetDesigns_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetDesigns(context.TODO(), &InputGetDesigns{})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetDesign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "12345678-90ab-1234-56cd-efghijk78901",
			"updated_at": "2024-05-22T01:59:57Z",
			"created_at": "2024-05-22T01:59:57Z",
			"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
			"name": "example",
			"editor": "code",
			"html_content": "<html><body><h1>Hello, World!</h1></body></html>",
			"plain_content": "",
			"generate_plain_content": false,
			"subject": "",
			"categories": []
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetDesign{
		ID:                   "12345678-90ab-1234-56cd-efghijk78901",
		UpdatedAt:            "2024-05-22T01:59:57Z",
		CreatedAt:            "2024-05-22T01:59:57Z",
		ThumbnailURL:         "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
		Name:                 "example",
		Editor:               "code",
		HTMLContent:          "<html><body><h1>Hello, World!</h1></body></html>",
		PlainContent:         "",
		GeneratePlainContent: false,
		Subject:              "",
		Categories:           []string{},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetDesign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestCreateDesign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "12345678-90ab-1234-56cd-efghijk78901",
			"updated_at": "2024-05-22T01:59:57Z",
			"created_at": "2024-05-22T01:59:57Z",
			"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
			"name": "example",
			"editor": "code",
			"html_content": "<html><body><h1>Hello, World!</h1></body></html>",
			"plain_content": "",
			"generate_plain_content": false,
			"subject": "",
			"categories": []
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.CreateDesign(context.TODO(), &InputCreateDesign{
		Name:        "example",
		Editor:      "code",
		HTMLContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateDesign{
		ID:                   "12345678-90ab-1234-56cd-efghijk78901",
		UpdatedAt:            "2024-05-22T01:59:57Z",
		CreatedAt:            "2024-05-22T01:59:57Z",
		ThumbnailURL:         "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
		Name:                 "example",
		Editor:               "code",
		HTMLContent:          "<html><body><h1>Hello, World!</h1></body></html>",
		PlainContent:         "",
		GeneratePlainContent: false,
		Subject:              "",
		Categories:           []string{},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestCreateDesign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateDesign(context.TODO(), &InputCreateDesign{
		Name:        "example",
		Editor:      "code",
		HTMLContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestUpdateDesign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "12345678-90ab-1234-56cd-efghijk78901",
			"updated_at": "2024-05-22T01:59:57Z",
			"created_at": "2024-05-22T01:59:57Z",
			"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
			"name": "example",
			"editor": "code",
			"html_content": "<html><body><h1>Hello, World!</h1></body></html>",
			"plain_content": "",
			"generate_plain_content": false,
			"subject": "",
			"categories": []
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.UpdateDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901", &InputUpdateDesign{
		Name:        "example",
		HTMLContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateDesign{
		ID:                   "12345678-90ab-1234-56cd-efghijk78901",
		UpdatedAt:            "2024-05-22T01:59:57Z",
		CreatedAt:            "2024-05-22T01:59:57Z",
		ThumbnailURL:         "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
		Name:                 "example",
		Editor:               "code",
		HTMLContent:          "<html><body><h1>Hello, World!</h1></body></html>",
		PlainContent:         "",
		GeneratePlainContent: false,
		Subject:              "",
		Categories:           []string{},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestUpdateDesign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901", &InputUpdateDesign{
		Name:        "example",
		HTMLContent: "<html><body><h1>Hello, World!</h1></body></html>",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDeleteDesign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteDesign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901")
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

// NewRequest Error Tests for Design methods
func TestGetDesigns_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetDesigns(context.TODO(), &InputGetDesigns{
		PageToken: "test",
		PageSize:  10,
		Summary:   true,
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetDesign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetDesign(context.TODO(), "test-design-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateDesign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputCreateDesign{
		Name:   "Test Design",
		Editor: "code",
	}
	_, err := client.CreateDesign(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateDesign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	input := &InputUpdateDesign{
		Name: "Updated Design",
	}
	_, err := client.UpdateDesign(context.TODO(), "test-design-id", input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteDesign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	err := client.DeleteDesign(context.TODO(), "test-design-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDuplicateDesign(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "98765432-10ab-4321-56cd-efghijk78901",
			"updated_at": "2024-05-22T01:59:57Z",
			"created_at": "2024-05-22T01:59:57Z",
			"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
			"name": "example copy",
			"editor": "code",
			"html_content": "<html><body><h1>Hello, World!</h1></body></html>",
			"plain_content": "",
			"generate_plain_content": false,
			"subject": "",
			"categories": []
		}`); err != nil {
			t.Fatal(err)
		}
	})

	r, err := client.DuplicateDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901", &InputDuplicateDesign{
		Name:   "example copy",
		Editor: "code",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputDuplicateDesign{
		ID:                   "98765432-10ab-4321-56cd-efghijk78901",
		UpdatedAt:            "2024-05-22T01:59:57Z",
		CreatedAt:            "2024-05-22T01:59:57Z",
		ThumbnailURL:         "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
		Name:                 "example copy",
		Editor:               "code",
		HTMLContent:          "<html><body><h1>Hello, World!</h1></body></html>",
		PlainContent:         "",
		GeneratePlainContent: false,
		Subject:              "",
		Categories:           []string{},
	}

	if !reflect.DeepEqual(want, r) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, r)))
	}
}

func TestDuplicateDesign_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/12345678-90ab-1234-56cd-efghijk78901", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DuplicateDesign(context.TODO(), "12345678-90ab-1234-56cd-efghijk78901", &InputDuplicateDesign{
		Name: "example copy",
	})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestDuplicateDesign_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.DuplicateDesign(context.TODO(), "test-design-id", &InputDuplicateDesign{
		Name: "test",
	})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetPreBuiltDesigns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/pre-builts", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id": "12345678-90ab-1234-56cd-efghijk78901",
				"updated_at": "2024-05-22T01:59:57Z",
				"created_at": "2024-05-22T01:59:57Z",
				"thumbnail_url": "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
				"name": "pre-built-example",
				"editor": "design"
			}],
			"_metadata": {
				"count": 1,
				"prev": "",
				"next": "",
				"self": "https://api.sendgrid.com/v3/designs/pre-builts?page_token=xxx"
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetPreBuiltDesigns(context.TODO(), &InputGetPreBuiltDesigns{
		PageToken: "test",
		PageSize:  10,
		Summary:   true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetPreBuiltDesigns{
		Result: []*Design{
			{
				ID:           "12345678-90ab-1234-56cd-efghijk78901",
				UpdatedAt:    "2024-05-22T01:59:57Z",
				CreatedAt:    "2024-05-22T01:59:57Z",
				ThumbnailURL: "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
				Name:         "pre-built-example",
				Editor:       "design",
			},
		},
		Metadata: _Metadata{
			Count: 1,
			Prev:  "",
			Next:  "",
			Self:  "https://api.sendgrid.com/v3/designs/pre-builts?page_token=xxx",
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetPreBuiltDesigns_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/designs/pre-builts", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetPreBuiltDesigns(context.TODO(), &InputGetPreBuiltDesigns{})
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

func TestGetPreBuiltDesigns_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	_, err := client.GetPreBuiltDesigns(context.TODO(), &InputGetPreBuiltDesigns{})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
