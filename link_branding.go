package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type OutputGetBrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

type DNSBrandedLink struct {
	DomainCname Record `json:"domain_cname,omitempty"`
	OwnerCname  Record `json:"owner_cname,omitempty"`
}

func (c *Client) GetBrandedLink(ctx context.Context, id int64) (*OutputGetBrandedLink, error) {
	path := fmt.Sprintf("/whitelabel/links/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type OutputGetDefaultBrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) GetDefaultBrandedLink(ctx context.Context) (*OutputGetDefaultBrandedLink, error) {
	req, err := c.NewRequest("GET", "/whitelabel/links/default", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetDefaultBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputGetBrandedLinks struct {
	Limit int
}

type BrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) GetBrandedLinks(ctx context.Context, input *InputGetBrandedLinks) ([]*BrandedLink, error) {
	u, err := url.Parse("/whitelabel/links")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if input.Limit > 0 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*BrandedLink{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type OutputGetSubuserBrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) GetSubuserBrandedLink(ctx context.Context, subuser string) (*OutputGetSubuserBrandedLink, error) {
	path := fmt.Sprintf("/whitelabel/links/subuser?username=%s", subuser)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSubuserBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputCreateBrandedLink struct {
	Domain    string `json:"domain,omitempty"`
	Subdomain string `json:"subdomain,omitempty"`
	Default   bool   `json:"default"`
}

type OutputCreateBrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) CreateBrandedLink(ctx context.Context, input *InputCreateBrandedLink) (*OutputCreateBrandedLink, error) {
	req, err := c.NewRequest("POST", "/whitelabel/links", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type OutputValidateBrandedLink struct {
	ID                int64                        `json:"id,omitempty"`
	Valid             bool                         `json:"valid,omitempty"`
	ValidationResults ValidationResultsBrandedLink `json:"validation_results,omitempty"`
}

type ValidationResultsBrandedLink struct {
	DomainCname ValidationResult `json:"domain_cname,omitempty"`
	OwnerCname  ValidationResult `json:"owner_cname,omitempty"`
}

func (c *Client) ValidateBrandedLink(ctx context.Context, id int64) (*OutputValidateBrandedLink, error) {
	path := fmt.Sprintf("/whitelabel/links/%s/validate", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputValidateBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputAssociateBrandedLinkWithSubuser struct {
	Username string `json:"username,omitempty"`
}

type OutputAssociateBrandedLinkWithSubuser struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) AssociateBrandedLinkWithSubuser(ctx context.Context, id int64, input *InputAssociateBrandedLinkWithSubuser) (*OutputAssociateBrandedLinkWithSubuser, error) {
	path := fmt.Sprintf("/whitelabel/links/%s/subuser", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAssociateBrandedLinkWithSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DisassociateBrandedLinkWithSubuser(ctx context.Context, username string) error {
	path := fmt.Sprintf("/whitelabel/links/subuser?username=%s", username)
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type InputUpdateBrandedLink struct {
	Default bool `json:"default"`
}

type OutputUpdateBrandedLink struct {
	ID        int64          `json:"id,omitempty"`
	Domain    string         `json:"domain,omitempty"`
	Subdomain string         `json:"subdomain,omitempty"`
	Username  string         `json:"username,omitempty"`
	UserID    int64          `json:"user_id,omitempty"`
	Default   bool           `json:"default,omitempty"`
	Valid     bool           `json:"valid,omitempty"`
	Legacy    bool           `json:"legacy,omitempty"`
	DNS       DNSBrandedLink `json:"dns,omitempty"`
}

func (c *Client) UpdateBrandedLink(ctx context.Context, id int64, input *InputUpdateBrandedLink) (*OutputUpdateBrandedLink, error) {
	path := fmt.Sprintf("/whitelabel/links/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateBrandedLink)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DeleteBrandedLink(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/whitelabel/links/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
