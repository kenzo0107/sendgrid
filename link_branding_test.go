package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/pkg/errors"
	"net/url"
	"strings"
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

func TestGetBrandedLink_AutoSSL(t *testing.T) {
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
			"default": false,
			"legacy": false,
			"dns": {
				"domain_cname": {
					"valid": false,
					"type": "cname",
					"host": "abc.examle.com",
					"data": "click.sendgrid.net"
				},
				"owner_cname": {
					"valid": false,
					"type": "cname",
					"host": "9876543.examle.com",
					"data": "sendgrid.net"
				},
				"acme_challenge_cname": {
					"valid": false,
					"type": "cname",
					"host": "_acme-challenge.abc.examle.com",
					"data": "abc.examle.com.0123456789abcdef.dcv.cloudflare.com"
				},
				"cf_custom_hostname_txt": {
					"valid": false,
					"type": "txt",
					"host": "_cf-custom-hostname.abc.examle.com",
					"data": "00000000-0000-0000-0000-000000000000"
				}
			},
			"last_validation_attempt_at": 1794228307,
			"region": "global",
			"auto_ssl": true,
			"auto_ssl_available": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetBrandedLink(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetBrandedLink{
		ID:                      1234567,
		UserID:                  9876543,
		Domain:                  "examle.com",
		Subdomain:               "abc",
		Username:                "dummy",
		Region:                  "global",
		AutoSSL:                 true,
		AutoSSLAvailable:        false,
		LastValidationAttemptAt: 1794228307,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Type: "cname",
				Host: "abc.examle.com",
				Data: "click.sendgrid.net",
			},
			OwnerCname: Record{
				Type: "cname",
				Host: "9876543.examle.com",
				Data: "sendgrid.net",
			},
			AcmeChallengeCname: Record{
				Type: "cname",
				Host: "_acme-challenge.abc.examle.com",
				Data: "abc.examle.com.0123456789abcdef.dcv.cloudflare.com",
			},
			CfCustomHostnameTxt: Record{
				Type: "txt",
				Host: "_cf-custom-hostname.abc.examle.com",
				Data: "00000000-0000-0000-0000-000000000000",
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

func TestCreateBrandedLink_AutoSSL(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if reqBody["auto_ssl"] != true {
			t.Errorf("Expected auto_ssl to be true, got %v", reqBody["auto_ssl"])
		}
		if reqBody["region"] != "global" {
			t.Errorf("Expected region to be global, got %v", reqBody["region"])
		}

		if _, err := fmt.Fprint(w, `{
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
					"data": "click.sendgrid.net"
				},
				"owner_cname": {
					"valid": false,
					"type": "cname",
					"host": "9876543.examle.com",
					"data": "sendgrid.net"
				},
				"acme_challenge_cname": {
					"valid": false,
					"type": "cname",
					"host": "_acme-challenge.abc.examle.com",
					"data": "abc.examle.com.0123456789abcdef.dcv.cloudflare.com"
				},
				"cf_custom_hostname_txt": {
					"valid": false,
					"type": "txt",
					"host": "_cf-custom-hostname.abc.examle.com",
					"data": "00000000-0000-0000-0000-000000000000"
				}
			},
			"region": "global",
			"auto_ssl": true,
			"auto_ssl_available": false
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.CreateBrandedLink(context.TODO(), &InputCreateBrandedLink{
		Domain:    "examle.com",
		Subdomain: "abc",
		Region:    "global",
		AutoSSL:   true,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputCreateBrandedLink{
		ID:               1234567,
		UserID:           9876543,
		Domain:           "examle.com",
		Subdomain:        "abc",
		Username:         "dummy",
		Region:           "global",
		AutoSSL:          true,
		AutoSSLAvailable: false,
		DNS: DNSBrandedLink{
			DomainCname: Record{
				Type: "cname",
				Host: "abc.examle.com",
				Data: "click.sendgrid.net",
			},
			OwnerCname: Record{
				Type: "cname",
				Host: "9876543.examle.com",
				Data: "sendgrid.net",
			},
			AcmeChallengeCname: Record{
				Type: "cname",
				Host: "_acme-challenge.abc.examle.com",
				Data: "abc.examle.com.0123456789abcdef.dcv.cloudflare.com",
			},
			CfCustomHostnameTxt: Record{
				Type: "txt",
				Host: "_cf-custom-hostname.abc.examle.com",
				Data: "00000000-0000-0000-0000-000000000000",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestCreateBrandedLink_OmitsAutoSSLAndRegionWhenUnset(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links", func(w http.ResponseWriter, r *http.Request) {
		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		for _, key := range []string{"auto_ssl", "region"} {
			if _, ok := reqBody[key]; ok {
				t.Errorf("Expected %s to be omitted from the request body, got %v", key, reqBody[key])
			}
		}

		if _, err := fmt.Fprint(w, testJsonBrandedLink); err != nil {
			t.Fatal(err)
		}
	})

	if _, err := client.CreateBrandedLink(context.TODO(), &InputCreateBrandedLink{
		Domain:    "examle.com",
		Subdomain: "abc",
	}); err != nil {
		t.Errorf("Unexpected error: %s", err)
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

func TestValidateBrandedLink_AutoSSL(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"valid": false,
			"validation_results": {
				"domain_cname": {
					"valid": false,
					"reason": "Expected CNAME record for \"abc.examle.com\" to match \"click.sendgrid.net\", but got \"\"."
				},
				"owner_cname": {
					"valid": false,
					"reason": "Expected CNAME record for \"9876543.examle.com\" to match \"sendgrid.net\", but got \"\"."
				},
				"acme_challenge_cname": {
					"valid": false,
					"reason": "Expected CNAME record for \"_acme-challenge.abc.examle.com\" to match \"abc.examle.com.0123456789abcdef.dcv.cloudflare.com\", but got \"\"."
				},
				"cf_custom_hostname_txt": {
					"valid": false,
					"reason": "Unable to lookup TXT record for \"_cf-custom-hostname.abc.examle.com\"."
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
				Reason: "Expected CNAME record for \"abc.examle.com\" to match \"click.sendgrid.net\", but got \"\".",
			},
			OwnerCname: ValidationResult{
				Reason: "Expected CNAME record for \"9876543.examle.com\" to match \"sendgrid.net\", but got \"\".",
			},
			AcmeChallengeCname: ValidationResult{
				Reason: "Expected CNAME record for \"_acme-challenge.abc.examle.com\" to match \"abc.examle.com.0123456789abcdef.dcv.cloudflare.com\", but got \"\".",
			},
			CfCustomHostnameTxt: ValidationResult{
				Reason: "Unable to lookup TXT record for \"_cf-custom-hostname.abc.examle.com\".",
			},
		},
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(ErrIncorrectResponse, errors.New(pretty.Compare(want, expected)))
	}
}

func TestValidateBrandedLink_AutoSSLValid(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/links/1234567/validate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"valid": true,
			"validation_results": {
				"domain_cname": {"valid": true, "reason": null},
				"owner_cname": {"valid": true, "reason": null},
				"acme_challenge_cname": {"valid": true, "reason": null},
				"cf_custom_hostname_txt": {"valid": true, "reason": null}
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
		Valid: true,
		ValidationResults: ValidationResultsBrandedLink{
			DomainCname:         ValidationResult{Valid: true},
			OwnerCname:          ValidationResult{Valid: true},
			AcmeChallengeCname:  ValidationResult{Valid: true},
			CfCustomHostnameTxt: ValidationResult{Valid: true},
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

// NewRequest Error Tests
func TestGetBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetBrandedLink(context.TODO(), 12345)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetDefaultBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetDefaultBrandedLink(context.TODO())
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetBrandedLinks_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetBrandedLinks(context.TODO(), &InputGetBrandedLinks{Limit: 50})
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestGetSubuserBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.GetSubuserBrandedLink(context.TODO(), "test-id")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestCreateBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	input := &InputCreateBrandedLink{
		Domain:    "example.com",
		Subdomain: "mail",
	}
	_, err := client.CreateBrandedLink(context.TODO(), input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestValidateBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	_, err := client.ValidateBrandedLink(context.TODO(), 12345)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestAssociateBrandedLinkWithSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	input := &InputAssociateBrandedLinkWithSubuser{
		Username: "testuser",
	}
	_, err := client.AssociateBrandedLinkWithSubuser(context.TODO(), 12345, input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDisassociateBrandedLinkWithSubuser_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	err := client.DisassociateBrandedLinkWithSubuser(context.TODO(), "testuser")
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestUpdateBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	input := &InputUpdateBrandedLink{
		Default: true,
	}
	_, err := client.UpdateBrandedLink(context.TODO(), 12345, input)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}

func TestDeleteBrandedLink_NewRequestError(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	originalBaseURL := client.baseURL
	invalidURL, _ := url.Parse("https://api.example.com/v3/")
	client.baseURL = invalidURL

	// Call method with appropriate parameters
	err := client.DeleteBrandedLink(context.TODO(), 12345)
	if err == nil {
		t.Error("Expected error for invalid baseURL")
	}
	if err != nil && !strings.Contains(err.Error(), "trailing slash") {
		t.Errorf("Expected error message to contain 'trailing slash', got %v", err.Error())
	}

	client.baseURL = originalBaseURL
}
