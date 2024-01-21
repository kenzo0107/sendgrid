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

func TestGetTemplate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id":"d-12345abcde",
			"name":"dummy",
			"generation":"dynamic",
			"updated_at":"2020-10-14 02:24:21",
			"versions": [{
				"id": "6692fdc5-803f-45fe-8f07-a2c330f6f28b",
				"template_id": "d-12345abcde",
				"name": "dummy",
				"subject": "dummy",
				"updated_at": "2020-10-20 05:11:38",
				"generate_plain_content": true,
				"html_content": "<html>\n<head>\n  <title>dummy</title>\n</head>\n<body>\ndummy\n</body>\n</html>\n",
				"plain_content": "",
				"editor": "code",
				"thumbnail_url": "//thumbnail-bucket.s3.amazonaws.com/dummy.png"
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTemplate(context.TODO(), "d-12345abcde")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTemplate{
		ID:         "d-12345abcde",
		Name:       "dummy",
		Generation: "dynamic",
		UpdatedAt:  "2020-10-14 02:24:21",
		Versions: []Version{
			{
				ID:                   "6692fdc5-803f-45fe-8f07-a2c330f6f28b",
				TemplateID:           "d-12345abcde",
				Name:                 "dummy",
				Subject:              "dummy",
				UpdatedAt:            "2020-10-20 05:11:38",
				GeneratePlainContent: true,
				HTMLContent:          "<html>\n<head>\n  <title>dummy</title>\n</head>\n<body>\ndummy\n</body>\n</html>\n",
				PlainContent:         "",
				Editor:               "code",
				ThumbnailURL:         "//thumbnail-bucket.s3.amazonaws.com/dummy.png",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTemplate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTemplate(context.TODO(), "d-12345abcde")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetTemplates(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"result": [{
				"id":"d-12345abcde",
				"name":"dummy",
				"generation":"dynamic",
				"updated_at":"2020-10-14 02:24:21",
				"versions": [{
					"id": "6692fdc5-803f-45fe-8f07-a2c330f6f28b",
					"template_id": "d-12345abcde",
					"name": "dummy",
					"subject": "dummy",
					"updated_at": "2020-10-20 05:11:38",
					"generate_plain_content": true,
					"html_content": "<html>\n<head>\n  <title>dummy</title>\n</head>\n<body>\ndummy\n</body>\n</html>\n",
					"plain_content": "",
					"editor": "code",
					"thumbnail_url": "//thumbnail-bucket.s3.amazonaws.com/dummy.png"
				}]
			}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetTemplates(context.TODO(), &InputGetTemplates{
		Generations: "dynamic",
		PageSize:    1,
		PageToken:   "dummy",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetTemplates{
		Templates: []Template{
			{
				ID:         "d-12345abcde",
				Name:       "dummy",
				Generation: "dynamic",
				UpdatedAt:  "2020-10-14 02:24:21",
				Versions: []Version{
					{
						ID:                   "6692fdc5-803f-45fe-8f07-a2c330f6f28b",
						TemplateID:           "d-12345abcde",
						Name:                 "dummy",
						Subject:              "dummy",
						UpdatedAt:            "2020-10-20 05:11:38",
						GeneratePlainContent: true,
						HTMLContent:          "<html>\n<head>\n  <title>dummy</title>\n</head>\n<body>\ndummy\n</body>\n</html>\n",
						PlainContent:         "",
						Editor:               "code",
						ThumbnailURL:         "//thumbnail-bucket.s3.amazonaws.com/dummy.png",
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetTemplates_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetTemplates(context.TODO(), &InputGetTemplates{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateTemplate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id":"d-12345abcde",
			"name":"dummy",
			"generation":"dynamic",
			"updated_at":"2020-10-14 02:24:21",
			"versions": [{}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateTemplate(context.TODO(), &InputCreateTemplate{
		Name:       "dummy",
		Generation: "dynamic",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateTemplate{
		ID:         "d-12345abcde",
		Name:       "dummy",
		Generation: "dynamic",
		UpdatedAt:  "2020-10-14 02:24:21",
		Versions: []Version{
			{
				ID:                   "",
				TemplateID:           "",
				Name:                 "",
				Subject:              "",
				UpdatedAt:            "",
				GeneratePlainContent: false,
				HTMLContent:          "",
				PlainContent:         "",
				Editor:               "",
				ThumbnailURL:         "",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateTemplate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateTemplate(context.TODO(), &InputCreateTemplate{
		Name:       "dummy",
		Generation: "dynamic",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDuplicateTemplate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id":"d-67890fghij",
			"name":"dummy2",
			"generation":"dynamic",
			"updated_at":"2020-10-15 02:24:21",
			"versions": [{}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.DuplicateTemplate(context.TODO(), "d-12345abcde", &InputDuplicateTemplate{
		Name: "dummy2",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputDuplicateTemplate{
		ID:         "d-67890fghij",
		Name:       "dummy2",
		Generation: "dynamic",
		UpdatedAt:  "2020-10-15 02:24:21",
		Versions: []Version{
			{
				ID:                   "",
				TemplateID:           "",
				Name:                 "",
				Subject:              "",
				UpdatedAt:            "",
				GeneratePlainContent: false,
				HTMLContent:          "",
				PlainContent:         "",
				Editor:               "",
				ThumbnailURL:         "",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestDuplicateTemplate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.DuplicateTemplate(context.TODO(), "d-12345abcde", &InputDuplicateTemplate{
		Name: "dummy",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateTemplate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id":"d-12345abcde",
			"name":"dummy2",
			"generation":"dynamic",
			"updated_at":"2020-10-14 02:24:21",
			"versions": [{}]
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateTemplate(context.TODO(), "d-12345abcde", &InputUpdateTemplate{
		Name: "dummy2",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateTemplate{
		ID:         "d-12345abcde",
		Name:       "dummy2",
		Generation: "dynamic",
		UpdatedAt:  "2020-10-14 02:24:21",
		Versions: []Version{
			{
				ID:                   "",
				TemplateID:           "",
				Name:                 "",
				Subject:              "",
				UpdatedAt:            "",
				GeneratePlainContent: false,
				HTMLContent:          "",
				PlainContent:         "",
				Editor:               "",
				ThumbnailURL:         "",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateTemplate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateTemplate(context.TODO(), "d-12345abcde", &InputUpdateTemplate{
		Name: "dummy",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteTemplate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteTemplate(context.TODO(), "d-12345abcde")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}
}

func TestDeleteTemplate_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/d-12345abcde", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteTemplate(context.TODO(), "d-12345abcde")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
