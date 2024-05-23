package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type OutputGetSSOIntegration struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration,omitempty"`
	LastUpdated          int64  `json:"last_updated,omitempty"`
	SingleSignonURL      string `json:"single_signon_url,omitempty"`
	AudienceURL          string `json:"audience_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/single-sign-on-settings/get-an-sso-integration
func (c *Client) GetSSOIntegration(ctx context.Context, id string) (*OutputGetSSOIntegration, error) {
	path := fmt.Sprintf("/sso/integrations/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSSOIntegration)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetSSOIntegrations struct {
	Si bool `json:"si,omitempty"`
}

type SSOIntegration struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration,omitempty"`
	LastUpdated          int64  `json:"last_updated,omitempty"`
	SingleSignonURL      string `json:"single_signon_url,omitempty"`
	AudienceURL          string `json:"audience_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/single-sign-on-settings/get-all-sso-integrations
func (c *Client) GetSSOIntegrations(ctx context.Context, input *InputGetSSOIntegrations) ([]*SSOIntegration, error) {
	u, err := url.Parse("/sso/integrations")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if input.Si {
		q.Set("si", strconv.FormatBool(input.Si))
	}

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*SSOIntegration{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateSSOIntegration struct {
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration"`
}

type OutputCreateSSOIntegration struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration,omitempty"`
	LastUpdated          int64  `json:"last_updated,omitempty"`
	SingleSignonURL      string `json:"single_signon_url,omitempty"`
	AudienceURL          string `json:"audience_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/single-sign-on-settings/create-an-sso-integration
func (c *Client) CreateSSOIntegration(ctx context.Context, input *InputCreateSSOIntegration) (*OutputCreateSSOIntegration, error) {
	req, err := c.NewRequest("POST", "/sso/integrations", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSSOIntegration)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateSSOIntegration struct {
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration"`
}

type OutputUpdateSSOIntegration struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Enabled              bool   `json:"enabled,omitempty"`
	SigninURL            string `json:"signin_url,omitempty"`
	SignoutURL           string `json:"signout_url,omitempty"`
	EntityID             string `json:"entity_id,omitempty"`
	CompletedIntegration bool   `json:"completed_integration,omitempty"`
	LastUpdated          int64  `json:"last_updated,omitempty"`
	SingleSignonURL      string `json:"single_signon_url,omitempty"`
	AudienceURL          string `json:"audience_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/single-sign-on-settings/update-an-sso-integration
func (c *Client) UpdateSSOIntegration(ctx context.Context, id string, input *InputUpdateSSOIntegration) (*OutputUpdateSSOIntegration, error) {
	path := fmt.Sprintf("/sso/integrations/%s", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSSOIntegration)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/single-sign-on-settings/delete-an-sso-integration
func (c *Client) DeleteSSOIntegration(ctx context.Context, id string) error {
	path := fmt.Sprintf("/sso/integrations/%s", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
