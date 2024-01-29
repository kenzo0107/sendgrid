package sendgrid

import (
	"context"
)

type OutputGetEnforceTLS struct {
	RequireTLS       bool    `json:"require_tls,omitempty"`
	RequireValidCert bool    `json:"require_valid_cert,omitempty"`
	Version          float64 `json:"version,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-enforced-tls/retrieve-current-enforced-tls-settings
func (c *Client) GetEnforceTLS(ctx context.Context) (*OutputGetEnforceTLS, error) {
	req, err := c.NewRequest("GET", "/user/settings/enforced_tls", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetEnforceTLS)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateEnforceTLS struct {
	RequireTLS       bool    `json:"require_tls,omitempty"`
	RequireValidCert bool    `json:"require_valid_cert,omitempty"`
	Version          float64 `json:"version,omitempty"`
}

type OutputUpdateEnforceTLS struct {
	RequireTLS       bool    `json:"require_tls,omitempty"`
	RequireValidCert bool    `json:"require_valid_cert,omitempty"`
	Version          float64 `json:"version,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-enforced-tls/update-enforced-tls-settings
func (c *Client) UpdateEnforceTLS(ctx context.Context, input *InputUpdateEnforceTLS) (*OutputUpdateEnforceTLS, error) {
	req, err := c.NewRequest("PATCH", "/user/settings/enforced_tls", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateEnforceTLS)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}
