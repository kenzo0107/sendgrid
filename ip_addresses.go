package sendgrid

import (
	"context"
	"fmt"
	"net/url"
)

type IPAddress struct {
	IP           string   `json:"ip,omitempty"`
	Pools        []string `json:"pools,omitempty"`
	Warmup       bool     `json:"warmup,omitempty"`
	StartDate    int64    `json:"start_date,omitempty"`
	Subusers     []string `json:"subusers,omitempty"`
	Rdns         string   `json:"rdns,omitempty"`
	AssignedAt   int64    `json:"assigned_at,omitempty"`
	Whitelabeled bool     `json:"whitelabeled,omitempty"`
}

type InputGetIPAddresses struct {
	IP                 string `url:"ip,omitempty"`
	ExcludeWhitelabels bool   `url:"exclude_whitelabels,omitempty"`
	Limit              int    `url:"limit,omitempty"`
	Offset             int    `url:"offset,omitempty"`
	Subuser            string `url:"subuser,omitempty"`
	SortByDirection    string `url:"sort_by_direction,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address/retrieve-all-ip-addresses
func (c *Client) GetIPAddresses(ctx context.Context, input *InputGetIPAddresses) ([]*IPAddress, error) {
	u, _ := url.Parse("/ips")

	q := u.Query()
	if input.IP != "" {
		q.Set("ip", input.IP)
	}
	if input.ExcludeWhitelabels {
		q.Set("exclude_whitelabels", fmt.Sprintf("%t", input.ExcludeWhitelabels))
	}
	if input.Limit != 0 {
		q.Set("limit", fmt.Sprintf("%d", input.Limit))
	}
	if input.Offset != 0 {
		q.Set("offset", fmt.Sprintf("%d", input.Offset))
	}
	if input.Subuser != "" {
		q.Set("subuser", input.Subuser)
	}
	if input.SortByDirection != "" {
		q.Set("sort_by_direction", input.SortByDirection)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var ips []*IPAddress
	if err := c.Do(ctx, req, &ips); err != nil {
		return nil, err
	}

	return ips, nil
}

type AssignedIPAddress struct {
	IP        string   `json:"ip,omitempty"`
	Pools     []string `json:"pools,omitempty"`
	Warmup    bool     `json:"warmup,omitempty"`
	StartDate int64    `json:"start_date,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address/retrieve-all-assigned-ips
func (c *Client) GetAssignedIPAddresses(ctx context.Context) ([]*AssignedIPAddress, error) {
	path := "/ips/assigned"

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var ips []*AssignedIPAddress
	if err := c.Do(ctx, req, &ips); err != nil {
		return nil, err
	}

	return ips, nil
}

type RemainingIPCount struct {
	Remaining  int    `json:"remaining,omitempty"`
	Period     string `json:"period,omitempty"`
	PricePerIP int    `json:"price_per_ip,omitempty"`
}

type OutputGetRemainingIPCount struct {
	Results []*RemainingIPCount `json:"results,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address/get-remaining-ips-count
func (c *Client) GetRemainingIPCount(ctx context.Context) (*OutputGetRemainingIPCount, error) {
	req, err := c.NewRequest("GET", "/ips/remaining", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetRemainingIPCount)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetIPAddress struct {
	IP           string   `json:"ip,omitempty"`
	Subusers     []string `json:"subusers,omitempty"`
	Rdns         string   `json:"rdns,omitempty"`
	Pools        []string `json:"pools,omitempty"`
	Warmup       bool     `json:"warmup,omitempty"`
	StartDate    int64    `json:"start_date,omitempty"`
	Whitelabeled bool     `json:"whitelabeled,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address/retrieve-all-ip-pools-an-ip-address-belongs-to
func (c *Client) GetIPAddress(ctx context.Context, ip string) (*OutputGetIPAddress, error) {
	path := fmt.Sprintf("/ips/%s", url.QueryEscape(ip))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetIPAddress)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
