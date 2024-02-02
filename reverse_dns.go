package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type InputGetReverseDNSs struct {
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
	IP     string `json:"ip,omitempty"`
}

type OutputGetReverseDNS struct {
	ID                      int64   `json:"id,omitempty"`
	IP                      string  `json:"ip,omitempty"`
	RDNS                    string  `json:"rdns,omitempty"`
	Users                   []*User `json:"users,omitempty"`
	Subdomain               string  `json:"subdomain,omitempty"`
	Domain                  string  `json:"domain,omitempty"`
	Valid                   bool    `json:"valid,omitempty"`
	Legacy                  bool    `json:"legacy,omitempty"`
	LastValidationAttemptAt int64   `json:"last_validation_attempt_at,omitempty"`
	ARecord                 ARecord `json:"a_record,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
}

type ARecord struct {
	Valid bool   `json:"valid,omitempty"`
	Type  string `json:"type,omitempty"`
	Host  string `json:"host,omitempty"`
	Data  string `json:"data,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/reverse-dns/retrieve-all-reverse-dns-records
func (c *Client) GetReverseDNSs(ctx context.Context, input *InputGetReverseDNSs) ([]*OutputGetReverseDNS, error) {
	u, err := url.Parse("/whitelabel/ips")
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
	if input.IP != "" {
		q.Set("ip", input.IP)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*OutputGetReverseDNS{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/reverse-dns/retrieve-a-reverse-dns-record
func (c *Client) GetReverseDNS(ctx context.Context, id int64) (*OutputGetReverseDNS, error) {
	path := fmt.Sprintf("/whitelabel/ips/%v", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetReverseDNS)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateReverseDNS struct {
	IP        string `json:"ip,omitempty"`
	Subdomain string `json:"subdomain,omitempty"`
	Domain    string `json:"domain,omitempty"`
}

type OutputCreateReverseDNS struct {
	ID                      int64   `json:"id,omitempty"`
	IP                      string  `json:"ip,omitempty"`
	RDNS                    string  `json:"rdns,omitempty"`
	Users                   []*User `json:"users,omitempty"`
	Subdomain               string  `json:"subdomain,omitempty"`
	Domain                  string  `json:"domain,omitempty"`
	Valid                   bool    `json:"valid,omitempty"`
	Legacy                  bool    `json:"legacy,omitempty"`
	LastValidationAttemptAt int64   `json:"last_validation_attempt_at,omitempty"`
	ARecord                 ARecord `json:"a_record,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/reverse-dns/set-up-reverse-dns
func (c *Client) CreateReverseDNS(ctx context.Context, input *InputCreateReverseDNS) (*OutputCreateReverseDNS, error) {
	req, err := c.NewRequest("GET", "/whitelabel/ips", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateReverseDNS)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputValidateReverseDNS struct {
	ID                int64                       `json:"id,omitempty"`
	Valid             bool                        `json:"valid,omitempty"`
	ValidationResults ValidationResultsReverseDNS `json:"validation_results,omitempty"`
}

type ValidationResultsReverseDNS struct {
	ARecordValidationResults ARecordValidationResults `json:"a_record,omitempty"`
}

type ARecordValidationResults struct {
	Valid  bool   `json:"valid,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/reverse-dns/set-up-reverse-dns
func (c *Client) ValidateReverseDNS(ctx context.Context, id int64) (*OutputValidateReverseDNS, error) {
	path := fmt.Sprintf("/whitelabel/ips/%v/validate", id)
	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputValidateReverseDNS)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/reverse-dns/delete-a-reverse-dns-record
func (c *Client) DeleteReverseDNS(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/whitelabel/ips/%v", id)
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
