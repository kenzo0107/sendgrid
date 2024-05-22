package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
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
				"editor": "code"
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

	expected, err := client.GetDesigns(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetDesigns{
		Result: []*Design{
			{
				ID:           "12345678-90ab-1234-56cd-efghijk78901",
				UpdatedAt:    "2024-05-22T01:59:57Z",
				CreatedAt:    "2024-05-22T01:59:57Z",
				ThumbnailURL: "//us-east-2-production-thumbnail-bucket.s3.amazonaws.com/xxx.png",
				Name:         "example",
				Editor:       "code",
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

	_, err := client.GetDesigns(context.TODO())
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
