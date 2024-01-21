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

func TestGetTemplateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "d-12345abcde",
			"template_id": "abc1234-12ab-34cd-56ef-78901abcde",
			"active": 1,
			"name": "dummy",
			"generate_plain_content": true,
			"updated_at": "2024-01-20 08:46:07",
			"editor": "code"
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTemplateVersion{
		ID:                   "d-12345abcde",
		TemplateID:           "abc1234-12ab-34cd-56ef-78901abcde",
		Active:               1,
		Name:                 "dummy",
		GeneratePlainContent: true,
		UpdatedAt:            "2024-01-20 08:46:07",
		Editor:               "code",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTemplateVersion_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateTemplateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "d-12345abcde",
			"template_id": "abc1234-12ab-34cd-56ef-78901abcde",
			"active": 0,
			"name": "dummy",
			"generate_plain_content": true,
			"editor": "code",
			"updated_at": "2024-01-20 08:46:07"
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateTemplateVersion(context.TODO(), "d-12345abcde", &InputCreateTemplateVersion{
		Active:               0,
		Name:                 "dummy",
		GeneratePlainContent: true,
		Editor:               "code",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateTemplateVersion{
		ID:                   "d-12345abcde",
		TemplateID:           "abc1234-12ab-34cd-56ef-78901abcde",
		Active:               0,
		Name:                 "dummy",
		HTMLContent:          "",
		PlainContent:         "",
		GeneratePlainContent: true,
		Subject:              "",
		Editor:               "code",
		TestData:             "",
		UpdatedAt:            "2024-01-20 08:46:07",
		Warnings:             []Warning(nil),
		ThumbnailURL:         "",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateTemplateVersion_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateTemplateVersion(context.TODO(), "d-12345abcde", &InputCreateTemplateVersion{
		Active:               0,
		Name:                 "dummy",
		GeneratePlainContent: true,
		Editor:               "code",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateTemplateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "abc1234-12ab-34cd-56ef-78901abcde",
			"template_id": "d-12345abcde",
			"active": 0,
			"name": "dummy2",
			"generate_plain_content": true,
			"updated_at": "2024-01-20 15:46:46",
			"editor": "code"
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde", &InputUpdateTemplateVersion{
		Active:               0,
		Name:                 "dummy2",
		GeneratePlainContent: true,
		Editor:               "code",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateTemplateVersion{
		ID:                   "abc1234-12ab-34cd-56ef-78901abcde",
		TemplateID:           "d-12345abcde",
		Active:               0,
		Name:                 "dummy2",
		HTMLContent:          "",
		PlainContent:         "",
		GeneratePlainContent: true,
		Subject:              "",
		Editor:               "code",
		TestData:             "",
		UpdatedAt:            "2024-01-20 15:46:46",
		Warnings:             []Warning(nil),
		ThumbnailURL:         "",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTemplateVersion_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde", &InputUpdateTemplateVersion{
		Active:               0,
		Name:                 "dummy2",
		GeneratePlainContent: true,
		Editor:               "code",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestActivateTemplateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde/activate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": "abc1234-12ab-34cd-56ef-78901abcde",
			"template_id": "d-12345abcde",
			"active": 1,
			"name": "dummy2",
			"generate_plain_content": true,
			"updated_at": "2024-01-20 15:46:46",
			"editor": "code"
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.ActivateTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputActivateTemplateVersion{
		ID:                   "abc1234-12ab-34cd-56ef-78901abcde",
		TemplateID:           "d-12345abcde",
		Active:               1,
		Name:                 "dummy2",
		HTMLContent:          "",
		PlainContent:         "",
		GeneratePlainContent: true,
		Subject:              "",
		Editor:               "code",
		TestData:             "",
		UpdatedAt:            "2024-01-20 15:46:46",
		Warnings:             []Warning(nil),
		ThumbnailURL:         "",
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestActivateTemplateVersion_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ActivateTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteTemplateVersion(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteTemplateVersion_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/abc1234-12ab-34cd-56ef-78901abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteTemplateVersion(context.TODO(), "d-12345abcde", "abc1234-12ab-34cd-56ef-78901abcde")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
