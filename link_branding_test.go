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

const testJsonBrandedLink = `{
		"id": 1234567,
		"user_id": 9876543,
		"domain": "examle.com",
		"subdomain": "abc",
		"username": "dummy",
		"valid": false,
		"default": false,
		"legacy": false,
		"dns": {
		  "domain_cname": {
			"valid": false,
			"type": "cname",
			"host": "abc.examle.com",
			"data": "sendgrid.net"
		  },
		  "owner_cname": {
			"valid": false,
			"type": "cname",
			"host": "9876543.examle.com",
			"data": "sendgrid.net"
		  }
		}
	  }`

func TestGetBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetBrandedLink(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetBrandedLink{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   false,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetBrandedLink(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetDefaultBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/default", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetDefaultBrandedLink(context.TODO())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetDefaultBrandedLink{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   false,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetDefaultBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/default", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetDefaultBrandedLink(context.TODO())
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetBrandedLinks(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		testJson := fmt.Sprintf("[%s]", testJsonBrandedLink)
		if _, err := fmt.Fprint(w, testJson); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetBrandedLinks(context.TODO(), &InputGetBrandedLinks{
		Limit: 1,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*BrandedLink{
		{
			ID:        1234567,
			UserID:    9876543,
			Domain:    "examle.com",
			Subdomain: "abc",
			Username:  "dummy",
			Valid:     false,
			Default:   false,
			Legacy:    false,
			DNS: DNSBrandedLink{
				DomainCname: Record{
					Valid: false,
					Type:  "cname",
					Host:  "abc.examle.com",
					Data:  "sendgrid.net",
				},
				OwnerCname: Record{
					Valid: false,
					Type:  "cname",
					Host:  "9876543.examle.com",
					Data:  "sendgrid.net",
				},
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetBrandedLinks_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetBrandedLinks(context.TODO(), &InputGetBrandedLinks{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetSubuserBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Add("username", "subuser_name")
		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetSubuserBrandedLink(context.TODO(), "subuser_name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetSubuserBrandedLink{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   false,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetSubuserBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Add("username", "subuser_name")
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetSubuserBrandedLink(context.TODO(), "subuser_name")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestCreateBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateBrandedLink(context.TODO(), &InputCreateBrandedLink{
		Domain:    "examle.com",
		Subdomain: "abc",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateBrandedLink{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   false,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.CreateBrandedLink(context.TODO(), &InputCreateBrandedLink{
		Domain:    "examle.com",
		Subdomain: "abc",
		Default:   true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestValidateBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"valid": false,
			"validation_results": {
				"domain_cname": {
					"valid": false,
					"reason": "Expected CNAME record for \"abc.examle.com\" to match \"sendgrid.net\", but got \"abc.examle.com.\"."
				},
				"owner_cname": {
					"valid": false,
					"reason": "Expected CNAME record for \"9876543.examle.com\" to match \"sendgrid.net\", but got \"9876543.examle.com.\"."
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.ValidateBrandedLink(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputValidateBrandedLink{
		ID:    1234567,
		Valid: false,
		ValidationResults: ValidationResultsBrandedLink{
			DomainCname: ValidationResult{
				Valid:  false,
				Reason: "Expected CNAME record for \"abc.examle.com\" to match \"sendgrid.net\", but got \"abc.examle.com.\".",
			},
			OwnerCname: ValidationResult{
				Valid:  false,
				Reason: "Expected CNAME record for \"9876543.examle.com\" to match \"sendgrid.net\", but got \"9876543.examle.com.\".",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestValidateBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ValidateBrandedLink(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAssociateBrandedLinkWithSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/subuser", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AssociateBrandedLinkWithSubuser(context.TODO(), 1234567, &InputAssociateBrandedLinkWithSubuser{
		Username: "subuser_name",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAssociateBrandedLinkWithSubuser{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   false,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestAssociateBrandedLinkWithSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/subuser", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AssociateBrandedLinkWithSubuser(context.TODO(), 1234567, &InputAssociateBrandedLinkWithSubuser{
		Username: "subuser_name",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDisassociateBrandedLinkWithSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Add("username", "subuser_name")
		if _, err := fmt.Fprint(w, ""); err != nil {
			t.Fatal(err)
		}
	})

	err := client.DisassociateBrandedLinkWithSubuser(context.TODO(), "subuser_name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDisassociateBrandedLinkWithSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Add("username", "subuser_name")
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DisassociateBrandedLinkWithSubuser(context.TODO(), "subuser_name")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"user_id": 9876543,
			"domain": "examle.com",
			"subdomain": "abc",
			"username": "dummy",
			"valid": false,
			"default": true,
			"legacy": false,
			"dns": {
			  "domain_cname": {
				"valid": false,
				"type": "cname",
				"host": "abc.examle.com",
				"data": "sendgrid.net"
			  },
			  "owner_cname": {
				"valid": false,
				"type": "cname",
				"host": "9876543.examle.com",
				"data": "sendgrid.net"
			  }
			}
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateBrandedLink(context.TODO(), 1234567, &InputUpdateBrandedLink{
		Default: true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateBrandedLink{
		ID:        1234567,
		UserID:    9876543,
		Domain:    "examle.com",
		Subdomain: "abc",
		Username:  "dummy",
		Valid:     false,
		Default:   true,
		Legacy:    false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "abc.examle.com",
				Data:  "sendgrid.net",
			},
			OwnerCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "9876543.examle.com",
				Data:  "sendgrid.net",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateBrandedLink(context.TODO(), 1234567, &InputUpdateBrandedLink{
		Default: true,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteBrandedLink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, ""); err != nil {
			t.Fatal(err)
		}
	})

	err := client.DeleteBrandedLink(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteBrandedLink_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteBrandedLink(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
