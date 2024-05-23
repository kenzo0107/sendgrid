package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type DomainAuthentication struct {
	ID                      int64                         `json:"id,omitempty"`
	UserID                  int64                         `json:"user_id,omitempty"`
	Subdomain               string                        `json:"subdomain,omitempty"`
	Domain                  string                        `json:"domain,omitempty"`
	Username                string                        `json:"username,omitempty"`
	IPs                     []string                      `json:"ips,omitempty"`
	CustomSpf               bool                          `json:"custom_spf,omitempty"`
	Default                 bool                          `json:"default,omitempty"`
	Legacy                  bool                          `json:"legacy,omitempty"`
	AutomaticSecurity       bool                          `json:"automatic_security,omitempty"`
	Valid                   bool                          `json:"valid,omitempty"`
	DNS                     DNS                           `json:"dns,omitempty"`
	Subusers                []SubuserSenderAuthentication `json:"subusers,omitempty"`
	LastValidationAttemptAt int64                         `json:"last_validation_attempt_at,omitempty"`
}

type DNS struct {
	MailCname Record `json:"mail_cname,omitempty"`
	Dkim1     Record `json:"dkim1,omitempty"`
	Dkim2     Record `json:"dkim2,omitempty"`
}

type Record struct {
	Valid bool   `json:"valid,omitempty"`
	Type  string `json:"type,omitempty"`
	Host  string `json:"host,omitempty"`
	Data  string `json:"data,omitempty"`
}

type SubuserSenderAuthentication struct {
	UserID   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type InputGetAuthenticatedDomains struct {
	Limit           int
	Offset          int
	ExcludeSubusers bool
	Username        string
	Domain          string
}

func (c *Client) GetAuthenticatedDomains(ctx context.Context, input *InputGetAuthenticatedDomains) ([]*DomainAuthentication, error) {
	u, err := url.Parse("/whitelabel/domains")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if input.Limit > 0 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}
	if input.Offset > 0 {
		q.Set("offset", strconv.Itoa(input.Offset))
	}
	if input.ExcludeSubusers {
		q.Set("exclude_subusers", "true")
	}
	if input.Username != "" {
		q.Set("username", input.Username)
	}
	if input.Domain != "" {
		q.Set("domain", input.Domain)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*DomainAuthentication{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetDefaultAuthentication struct {
	Domain string
}

type OutputGetDefaultAuthentication struct {
	ID                      int64                         `json:"id,omitempty"`
	UserID                  int64                         `json:"user_id,omitempty"`
	Subdomain               string                        `json:"subdomain,omitempty"`
	Domain                  string                        `json:"domain,omitempty"`
	Username                string                        `json:"username,omitempty"`
	IPs                     []string                      `json:"ips,omitempty"`
	CustomSpf               bool                          `json:"custom_spf,omitempty"`
	Default                 bool                          `json:"default,omitempty"`
	Legacy                  bool                          `json:"legacy,omitempty"`
	AutomaticSecurity       bool                          `json:"automatic_security,omitempty"`
	Valid                   bool                          `json:"valid,omitempty"`
	DNS                     DNS                           `json:"dns,omitempty"`
	Subusers                []SubuserSenderAuthentication `json:"subusers,omitempty"`
	LastValidationAttemptAt int64                         `json:"last_validation_attempt_at,omitempty"`
}

func (c *Client) GetDefaultAuthentication(ctx context.Context, input *InputGetDefaultAuthentication) (*OutputGetDefaultAuthentication, error) {
	u, err := url.Parse("/whitelabel/domains/default")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if input.Domain != "" {
		q.Set("domain", input.Domain)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetDefaultAuthentication)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetAuthenticatedDomain struct {
	ID                int64    `json:"id,omitempty"`
	UserID            int64    `json:"user_id,omitempty"`
	Subdomain         string   `json:"subdomain,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	Username          string   `json:"username,omitempty"`
	IPs               []string `json:"ips,omitempty"`
	CustomSpf         bool     `json:"custom_spf,omitempty"`
	Default           bool     `json:"default,omitempty"`
	Legacy            bool     `json:"legacy,omitempty"`
	AutomaticSecurity bool     `json:"automatic_security,omitempty"`
	Valid             bool     `json:"valid,omitempty"`
	DNS               DNS      `json:"dns,omitempty"`
}

func (c *Client) GetAuthenticatedDomain(ctx context.Context, domainId int64) (*OutputGetAuthenticatedDomain, error) {
	path := fmt.Sprintf("/whitelabel/domains/%s", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAuthenticatedDomain)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputAuthenticateDomain struct {
	Domain             string   `json:"domain,omitempty"`
	Subdomain          string   `json:"subdomain,omitempty"`
	Username           string   `json:"username,omitempty"`
	IPs                []string `json:"ips,omitempty"`
	CustomSpf          bool     `json:"custom_spf,omitempty"`
	Default            bool     `json:"default,omitempty"`
	AutomaticSecurity  bool     `json:"automatic_security,omitempty"`
	CustomDkimSelector string   `json:"custom_dkim_selector,omitempty"`
}

type OutputAuthenticateDomain struct {
	ID                int64    `json:"id,omitempty"`
	UserID            int64    `json:"user_id,omitempty"`
	Subdomain         string   `json:"subdomain,omitempty"`
	Domain            string   `json:"domain,omitempty"`
	Username          string   `json:"username,omitempty"`
	IPs               []string `json:"ips,omitempty"`
	CustomSpf         bool     `json:"custom_spf,omitempty"`
	Default           bool     `json:"default,omitempty"`
	Legacy            bool     `json:"legacy,omitempty"`
	AutomaticSecurity bool     `json:"automatic_security,omitempty"`
	Valid             bool     `json:"valid,omitempty"`
	DNS               DNS      `json:"dns,omitempty"`
}

func (c *Client) AuthenticateDomain(ctx context.Context, input *InputAuthenticateDomain) (*OutputAuthenticateDomain, error) {
	req, err := c.NewRequest("POST", "/whitelabel/domains", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAuthenticateDomain)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputAddIPToAuthenticatedDomain struct {
	IP string `json:"ip,omitempty"`
}

type OutputAddIPToAuthenticatedDomain struct {
	ID                      int64    `json:"id,omitempty"`
	UserID                  int64    `json:"user_id,omitempty"`
	Subdomain               string   `json:"subdomain,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	Username                string   `json:"username,omitempty"`
	IPs                     []string `json:"ips,omitempty"`
	CustomSpf               bool     `json:"custom_spf,omitempty"`
	Default                 bool     `json:"default,omitempty"`
	Legacy                  bool     `json:"legacy,omitempty"`
	AutomaticSecurity       bool     `json:"automatic_security,omitempty"`
	Valid                   bool     `json:"valid,omitempty"`
	DNS                     DNS      `json:"dns,omitempty"`
	LastValidationAttemptAt int64    `json:"last_validation_attempt_at,omitempty"`
}

// NOTE: The 'dns' key in the API response for adding an IP to the authenticated domain is different from what is documented.
// see: https://docs.sendgrid.com/api-reference/domain-authentication/add-an-ip-to-an-authenticated-domain#responses
func (c *Client) AddIPToAuthenticatedDomain(ctx context.Context, domainId int64, input *InputAddIPToAuthenticatedDomain) (*OutputAddIPToAuthenticatedDomain, error) {
	path := fmt.Sprintf("/whitelabel/domains/%s/ips", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddIPToAuthenticatedDomain)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// NOTE: The 'dns' key in the API response for removing an IP to the authenticated domain is different from what is documented.
// see: https://docs.sendgrid.com/api-reference/domain-authentication/remove-an-ip-from-an-authenticated-domain#responses
func (c *Client) RemoveIPFromAuthenticatedDomain(ctx context.Context, domainId int64, ip string) error {
	path := fmt.Sprintf("/whitelabel/domains/%s/ips/%s", strconv.FormatInt(domainId, 10), ip)
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type OutputValidateDomainAuthentication struct {
	ID                int64             `json:"id,omitempty"`
	Valid             bool              `json:"valid,omitempty"`
	ValidationResults ValidationResults `json:"validation_results,omitempty"`
}

type ValidationResults struct {
	MailCname ValidationResult `json:"mail_cname,omitempty"`
	Dkim1     ValidationResult `json:"dkim1,omitempty"`
	Dkim2     ValidationResult `json:"dkim2,omitempty"`
	SPF       ValidationResult `json:"spf,omitempty"`
}

type ValidationResult struct {
	Valid  bool   `json:"valid,omitempty"`
	Reason string `json:"reason,omitempty"`
}

func (c *Client) ValidateDomainAuthentication(ctx context.Context, domainId int64) (*OutputValidateDomainAuthentication, error) {
	path := fmt.Sprintf("/whitelabel/domains/%s/validate", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputValidateDomainAuthentication)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateDomainAuthentication struct {
	Default   bool `json:"default"`
	CustomSpf bool `json:"custom_spf"`
}

type OutputUpdateDomainAuthentication struct {
	ID                      int64                         `json:"id,omitempty"`
	UserID                  int64                         `json:"user_id,omitempty"`
	Subdomain               string                        `json:"subdomain,omitempty"`
	Domain                  string                        `json:"domain,omitempty"`
	Username                string                        `json:"username,omitempty"`
	IPs                     []string                      `json:"ips,omitempty"`
	CustomSpf               bool                          `json:"custom_spf,omitempty"`
	Default                 bool                          `json:"default,omitempty"`
	Legacy                  bool                          `json:"legacy,omitempty"`
	AutomaticSecurity       bool                          `json:"automatic_security,omitempty"`
	Valid                   bool                          `json:"valid,omitempty"`
	DNS                     DNS                           `json:"dns,omitempty"`
	Subusers                []SubuserSenderAuthentication `json:"subusers,omitempty"`
	LastValidationAttemptAt int64                         `json:"last_validation_attempt_at,omitempty"`
}

func (c *Client) UpdateDomainAuthentication(ctx context.Context, domainId int64, input *InputUpdateDomainAuthentication) (*OutputUpdateDomainAuthentication, error) {
	path := fmt.Sprintf("/whitelabel/domains/%s", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateDomainAuthentication)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DeleteAuthenticatedDomain(ctx context.Context, domainId int64) error {
	path := fmt.Sprintf("/whitelabel/domains/%s", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type OutputGetAuthenticatedDomainAssociatedWithSubuser struct {
	ID                      int64    `json:"id,omitempty"`
	UserID                  int64    `json:"user_id,omitempty"`
	Subdomain               string   `json:"subdomain,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	Username                string   `json:"username,omitempty"`
	IPs                     []string `json:"ips,omitempty"`
	CustomSpf               bool     `json:"custom_spf,omitempty"`
	Default                 bool     `json:"default,omitempty"`
	Legacy                  bool     `json:"legacy,omitempty"`
	AutomaticSecurity       bool     `json:"automatic_security,omitempty"`
	Valid                   bool     `json:"valid,omitempty"`
	DNS                     DNS      `json:"dns,omitempty"`
	LastValidationAttemptAt int64    `json:"last_validation_attempt_at,omitempty"`
}

func (c *Client) GetAuthenticatedDomainAssociatedWithSubuser(ctx context.Context, subuserName string) (*OutputGetAuthenticatedDomainAssociatedWithSubuser, error) {
	path := fmt.Sprintf("/whitelabel/domains/subuser?username=%s", subuserName)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAuthenticatedDomainAssociatedWithSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputAssociateAuthenticatedDomainWithSubuser struct {
	Username string `json:"username,omitempty"`
}

type OutputAssociateAuthenticatedDomainWithSubuser struct {
	ID                      int64    `json:"id,omitempty"`
	UserID                  int64    `json:"user_id,omitempty"`
	Subdomain               string   `json:"subdomain,omitempty"`
	Domain                  string   `json:"domain,omitempty"`
	Username                string   `json:"username,omitempty"`
	IPs                     []string `json:"ips,omitempty"`
	CustomSpf               bool     `json:"custom_spf,omitempty"`
	Default                 bool     `json:"default,omitempty"`
	Legacy                  bool     `json:"legacy,omitempty"`
	AutomaticSecurity       bool     `json:"automatic_security,omitempty"`
	Valid                   bool     `json:"valid,omitempty"`
	DNS                     DNS      `json:"dns,omitempty"`
	LastValidationAttemptAt int64    `json:"last_validation_attempt_at,omitempty"`
}

func (c *Client) AssociateAuthenticatedDomainWithSubuser(ctx context.Context, domainId int64, input *InputAssociateAuthenticatedDomainWithSubuser) (*OutputAssociateAuthenticatedDomainWithSubuser, error) {
	path := fmt.Sprintf("/whitelabel/domains/%s/subuser", strconv.FormatInt(domainId, 10))
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAssociateAuthenticatedDomainWithSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DisassociateAuthenticatedDomainFromSubuser(ctx context.Context, subuserName string) error {
	path := fmt.Sprintf("/whitelabel/domains/subuser?username=%s", subuserName)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
