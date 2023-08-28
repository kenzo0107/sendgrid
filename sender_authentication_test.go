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

func TestGetAuthenticatedDomains(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("limit", "1")
		q.Set("offset", "10")
		q.Set("exclude_subusers", "true")
		q.Set("username", "dummy")
		q.Set("domain", "example.com")
		r.URL.RawQuery = q.Encode()

		if _, err := fmt.Fprint(w, `[
			{
				"id": 1234567,
				"user_id": 9876543,
				"subdomain": "em1234",
				"domain": "example.com",
				"username": "dummy",
				"ips": [],
				"custom_spf": false,
				"default": false,
				"legacy": false,
				"automatic_security": true,
				"valid": false,
				"dns": {
					"mail_cname": {
						"valid": false,
						"type": "cname",
						"host": "em1234.example.com",
						"data": "u1234567.wl123.sendgrid.net"
					},
					"dkim1": {
						"valid": false,
						"type":  "cname",
						"host":  "s1._domainkey.example.com",
						"data":  "s1.domainkey.u1234567.wl123.sendgrid.net"
					},
					"dkim2": {
						"valid": false,
						"type":  "cname",
						"host":  "s2._domainkey.example.com",
						"data":  "s2.domainkey.u1234567.wl123.sendgrid.net"
					}
				},
				"last_validation_attempt_at": 1608242131
			}
		]`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetAuthenticatedDomains(context.TODO(), &InputGetAuthenticatedDomains{
		Limit:           1,
		Offset:          10,
		ExcludeSubusers: true,
		Username:        "dummy",
		Domain:          "example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := []*DomainAuthentication{
		{
			ID:                1234567,
			UserID:            9876543,
			Subdomain:         "em1234",
			Domain:            "example.com",
			Username:          "dummy",
			IPs:               []string{},
			CustomSpf:         false,
			Default:           false,
			Legacy:            false,
			AutomaticSecurity: true,
			Valid:             false,
			DNS: DNS{
				MailCname: Record{
					Valid: false,
					Type:  "cname",
					Host:  "em1234.example.com",
					Data:  "u1234567.wl123.sendgrid.net",
				},
				Dkim1: Record{
					Valid: false,
					Type:  "cname",
					Host:  "s1._domainkey.example.com",
					Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
				},
				Dkim2: Record{
					Valid: false,
					Type:  "cname",
					Host:  "s2._domainkey.example.com",
					Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
				},
			},
			LastValidationAttemptAt: 1608242131,
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetAuthenticatedDomains_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAuthenticatedDomains(context.TODO(), &InputGetAuthenticatedDomains{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetDefaultAuthentication(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/default", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("domain", "sendgrid.net")
		r.URL.RawQuery = q.Encode()

		if _, err := fmt.Fprint(w, `{
			"id": 0,
			"user_id": 0,
			"subdomain": "",
			"domain": "sendgrid.net",
			"username": "",
			"ips": null,
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": false,
			"valid": false,
			"dns": {}
		  }`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetDefaultAuthentication(context.TODO(), &InputGetDefaultAuthentication{
		Domain: "sendgrid.net",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetDefaultAuthentication{
		ID:                0,
		UserID:            0,
		Subdomain:         "",
		Domain:            "sendgrid.net",
		Username:          "",
		IPs:               nil,
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: false,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "",
				Host:  "",
				Data:  "",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "",
				Host:  "",
				Data:  "",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "",
				Host:  "",
				Data:  "",
			},
		},
		Subusers:                []SubuserSenderAuthentication(nil),
		LastValidationAttemptAt: 0,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetDefaultAuthentication_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/default", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetDefaultAuthentication(context.TODO(), &InputGetDefaultAuthentication{})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetAuthenticatedDomain(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/1234567", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": [],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns": {
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1": {
					"valid": false,
					"type":  "cname",
					"host":  "s1._domainkey.example.com",
					"data":  "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2": {
					"valid": false,
					"type":  "cname",
					"host":  "s2._domainkey.example.com",
					"data":  "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetAuthenticatedDomain(context.TODO(), 1234567)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAuthenticatedDomain{
		ID:                1234567,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetAuthenticatedDomain_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/1234567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAuthenticatedDomain(context.TODO(), 1234567)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAuthenticatedDomain(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 12345678,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": [],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns":{
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1":{
					"valid": false,
					"type": "cname",
					"host": "s1._domainkey.example.com",
					"data": "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2":{
					"valid": false,
					"type": "cname",
					"host": "s2._domainkey.example.com",
					"data": "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AuthenticateDomain(context.TODO(), &InputAuthenticateDomain{
		Domain: "example.com",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAuthenticateDomain{
		ID:                12345678,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestAuthenticatedDomain_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AuthenticateDomain(context.TODO(), &InputAuthenticateDomain{
		Domain: "example.com",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAddIPToAuthenticatedDomain(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/ips", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 12345678,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": ["127.0.0.1"],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns":{
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1":{
					"valid": false,
					"type": "cname",
					"host": "s1._domainkey.example.com",
					"data": "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2":{
					"valid": false,
					"type": "cname",
					"host": "s2._domainkey.example.com",
					"data": "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			},
			"last_validation_attempt_at": 1608242131
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AddIPToAuthenticatedDomain(context.TODO(), 12345678, &InputAddIPToAuthenticatedDomain{
		IP: "127.0.0.1",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAddIPToAuthenticatedDomain{
		ID:                12345678,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{"127.0.0.1"},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
		LastValidationAttemptAt: 1608242131,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestAddIPToAuthenticatedDomain_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/ips", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AddIPToAuthenticatedDomain(context.TODO(), 12345678, &InputAddIPToAuthenticatedDomain{
		IP: "127.0.0.1",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestRemoveIPFromAuthenticatedDomain(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.RemoveIPFromAuthenticatedDomain(context.TODO(), 12345678, "127.0.0.1")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestRemoveIPFromAuthenticatedDomain_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/ips/127.0.0.1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.RemoveIPFromAuthenticatedDomain(context.TODO(), 12345678, "127.0.0.1")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestValidateDomainAuthentication(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/validate", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 12345678,
			"valid": false,
			"validation_results": {
				"mail_cname": {
					"valid": false,
					"reason": "Expected CNAME for \"em1234.example.com\" to match \"u1234567.wl123.sendgrid.net\"."
				},
				"dkim1": {
					"valid": false,
					"reason": "Expected CNAME for \"s1._domainkey.example.com\" to match \"s1.domainkey.u1234567.wl123.sendgrid.net\"."
				},
				"dkim2": {
					"valid": false,
					"reason": "Expected CNAME for \"s2._domainkey.ajinomoto.com\" to match \"s2.domainkey.u4707187.wl188.sendgrid.net\"."
				},
				"spf": {
					"valid": false,
					"reason": ""
				}
			}
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.ValidateDomainAuthentication(context.TODO(), 12345678)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputValidateDomainAuthentication{
		ID:    12345678,
		Valid: false,
		ValidationResults: ValidationResults{
			MailCname: ValidationResult{
				Valid:  false,
				Reason: "Expected CNAME for \"em1234.example.com\" to match \"u1234567.wl123.sendgrid.net\".",
			},
			Dkim1: ValidationResult{
				Valid:  false,
				Reason: "Expected CNAME for \"s1._domainkey.example.com\" to match \"s1.domainkey.u1234567.wl123.sendgrid.net\".",
			},
			Dkim2: ValidationResult{
				Valid:  false,
				Reason: "Expected CNAME for \"s2._domainkey.ajinomoto.com\" to match \"s2.domainkey.u4707187.wl188.sendgrid.net\".",
			},
			SPF: ValidationResult{
				Valid:  false,
				Reason: "",
			},
		},
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestValidateDomainAuthentication_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678/validate", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.ValidateDomainAuthentication(context.TODO(), 12345678)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestUpdateDomainAuthentication(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": [],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns": {
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1": {
					"valid": false,
					"type":  "cname",
					"host":  "s1._domainkey.example.com",
					"data":  "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2": {
					"valid": false,
					"type":  "cname",
					"host":  "s2._domainkey.example.com",
					"data":  "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			},
			"last_validation_attempt_at": 1608242131
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.UpdateDomainAuthentication(context.TODO(), 12345678, &InputUpdateDomainAuthentication{
		Default:   false,
		CustomSpf: false,
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputUpdateDomainAuthentication{
		ID:                1234567,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
		LastValidationAttemptAt: 1608242131,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestUpdateDomainAuthentication_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.UpdateDomainAuthentication(context.TODO(), 12345678, &InputUpdateDomainAuthentication{
		Default:   false,
		CustomSpf: false,
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDeleteAuthenticatedDomain(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DeleteAuthenticatedDomain(context.TODO(), 12345678)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDeleteAuthenticatedDomain_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/12345678", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DeleteAuthenticatedDomain(context.TODO(), 12345678)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestGetAuthenticatedDomainAssociatedWithSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("username", "dummy")
		r.URL.RawQuery = q.Encode()

		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": [],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns": {
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1": {
					"valid": false,
					"type":  "cname",
					"host":  "s1._domainkey.example.com",
					"data":  "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2": {
					"valid": false,
					"type":  "cname",
					"host":  "s2._domainkey.example.com",
					"data":  "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			},
			"last_validation_attempt_at": 1608242131
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.GetAuthenticatedDomainAssociatedWithSubuser(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputGetAuthenticatedDomainAssociatedWithSubuser{
		ID:                1234567,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
		LastValidationAttemptAt: 1608242131,
	}

	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestGetAuthenticatedDomainAssociatedWithSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/subuser?username=%s", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.GetAuthenticatedDomainAssociatedWithSubuser(context.TODO(), "")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestAssociateAuthenticatedDomainWithSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/1234567/subuser", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, `{
			"id": 1234567,
			"user_id": 9876543,
			"subdomain": "em1234",
			"domain": "example.com",
			"username": "dummy",
			"ips": [],
			"custom_spf": false,
			"default": false,
			"legacy": false,
			"automatic_security": true,
			"valid": false,
			"dns": {
				"mail_cname": {
					"valid": false,
					"type": "cname",
					"host": "em1234.example.com",
					"data": "u1234567.wl123.sendgrid.net"
				},
				"dkim1": {
					"valid": false,
					"type":  "cname",
					"host":  "s1._domainkey.example.com",
					"data":  "s1.domainkey.u1234567.wl123.sendgrid.net"
				},
				"dkim2": {
					"valid": false,
					"type":  "cname",
					"host":  "s2._domainkey.example.com",
					"data":  "s2.domainkey.u1234567.wl123.sendgrid.net"
				}
			},
			"last_validation_attempt_at": 1608242131
		}`); err != nil {
			t.Fatal(err)
		}
	})

	expected, err := client.AssociateAuthenticatedDomainWithSubuser(context.TODO(), 1234567, &InputAssociateAuthenticatedDomainWithSubuser{
		Username: "dummy",
	})
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	want := &OutputAssociateAuthenticatedDomainWithSubuser{
		ID:                1234567,
		UserID:            9876543,
		Subdomain:         "em1234",
		Domain:            "example.com",
		Username:          "dummy",
		IPs:               []string{},
		CustomSpf:         false,
		Default:           false,
		Legacy:            false,
		AutomaticSecurity: true,
		Valid:             false,
		DNS: DNS{
			MailCname: Record{
				Valid: false,
				Type:  "cname",
				Host:  "em1234.example.com",
				Data:  "u1234567.wl123.sendgrid.net",
			},
			Dkim1: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s1._domainkey.example.com",
				Data:  "s1.domainkey.u1234567.wl123.sendgrid.net",
			},
			Dkim2: Record{
				Valid: false,
				Type:  "cname",
				Host:  "s2._domainkey.example.com",
				Data:  "s2.domainkey.u1234567.wl123.sendgrid.net",
			},
		},
		LastValidationAttemptAt: 1608242131,
	}
	if !reflect.DeepEqual(want, expected) {
		t.Fatal(errors.New(pretty.Compare(want, expected)))
	}
}

func TestAssociateAuthenticatedDomainWithSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/1234567/subuser", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := client.AssociateAuthenticatedDomainWithSubuser(context.TODO(), 1234567, &InputAssociateAuthenticatedDomainWithSubuser{
		Username: "dummy",
	})
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestDisassociateAuthenticatedDomainWithSubuser(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("username", "dummy")
		r.URL.RawQuery = q.Encode()

		w.WriteHeader(http.StatusNoContent)
	})

	err := client.DisassociateAuthenticatedDomainFromSubuser(context.TODO(), "dummy")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
}

func TestDisassociateAuthenticatedDomainWithSubuser_Failed(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/whitelabel/domains/subuser", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("username", "dummy")
		r.URL.RawQuery = q.Encode()

		w.WriteHeader(http.StatusInternalServerError)
	})

	err := client.DisassociateAuthenticatedDomainFromSubuser(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}
