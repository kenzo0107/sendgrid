package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetSSOCertificate struct {
	ID                int64  `json:"id,omitempty"`
	PublicCertificate string `json:"public_certificate,omitempty"`
	NotBefore         int64  `json:"not_before,omitempty"`
	NotAfter          int64  `json:"not_after,omitempty"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/certificates/get-an-sso-certificate
func (c *Client) GetSSOCertificate(ctx context.Context, id int64) (*OutputGetSSOCertificate, error) {
	path := fmt.Sprintf("/sso/certificates/%v", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSSOCertificate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type SSOCertificate struct {
	ID                int64  `json:"id,omitempty"`
	PublicCertificate string `json:"public_certificate,omitempty"`
	NotBefore         int64  `json:"not_before,omitempty"`
	NotAfter          int64  `json:"not_after,omitempty"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/certificates/get-all-sso-certificates-by-integration
func (c *Client) GetSSOCertificates(ctx context.Context, integrationID string) ([]*SSOCertificate, error) {
	path := fmt.Sprintf("/sso/integrations/%s/certificates", integrationID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := []*SSOCertificate{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateSSOCertificate struct {
	PublicCertificate string `json:"public_certificate,omitempty"`
	Enabled           bool   `json:"enabled"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

type OutputCreateSSOCertificate struct {
	ID                int64  `json:"id,omitempty"`
	PublicCertificate string `json:"public_certificate,omitempty"`
	NotBefore         int64  `json:"not_before,omitempty"`
	NotAfter          int64  `json:"not_after,omitempty"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

func (c *Client) CreateSSOCertificate(ctx context.Context, input *InputCreateSSOCertificate) (*OutputCreateSSOCertificate, error) {
	req, err := c.NewRequest("POST", "/sso/certificates", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSSOCertificate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateSSOCertificate struct {
	PublicCertificate string `json:"public_certificate,omitempty"`
	Enabled           bool   `json:"enabled"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

type OutputUpdateSSOCertificate struct {
	ID                int64  `json:"id,omitempty"`
	PublicCertificate string `json:"public_certificate,omitempty"`
	NotBefore         int64  `json:"not_before,omitempty"`
	NotAfter          int64  `json:"not_after,omitempty"`
	IntegrationID     string `json:"integration_id,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/certificates/update-sso-certificate
func (c *Client) UpdateSSOCertificate(ctx context.Context, id int64, input *InputUpdateSSOCertificate) (*OutputUpdateSSOCertificate, error) {
	path := fmt.Sprintf("/sso/certificates/%v", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSSOCertificate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/certificates/delete-an-sso-certificate
func (c *Client) DeleteSSOCertificate(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/sso/certificates/%v", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
