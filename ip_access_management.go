package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type AllowedIP struct {
	ID        int64  `json:"id,omitempty"`
	IP        string `json:"ip,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type OutputGetAllowedIP struct {
	Result []*AllowedIP `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-a-specific-allowed-ip
func (c *Client) GetAllowedIP(ctx context.Context, ruleID int64) (*OutputGetAllowedIP, error) {
	path := fmt.Sprintf("/access_settings/whitelist/%s", strconv.FormatInt(ruleID, 10))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAllowedIP)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetAllowedIPs struct {
	Result []*AllowedIP `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-a-list-of-currently-allowed-ips
func (c *Client) GetAllowedIPs(ctx context.Context) (*OutputGetAllowedIPs, error) {
	req, err := c.NewRequest("GET", "/access_settings/whitelist", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAllowedIPs)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type AccessActivity struct {
	Allowed    bool   `json:"allowed,omitempty"`
	AuthMethod string `json:"auth_method,omitempty"`
	FirstAt    int64  `json:"first_at,omitempty"`
	IP         string `json:"ip,omitempty"`
	LastAt     int64  `json:"last_at,omitempty"`
	Location   string `json:"location,omitempty"`
}

type InputGetAccessActivities struct {
	Limit int `json:"limit,omitempty"`
}

type OutputGetAccessActivities struct {
	Result []*AccessActivity `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-all-recent-access-attempts
func (c *Client) GetAccessActivities(ctx context.Context, input *InputGetAccessActivities) (*OutputGetAccessActivities, error) {
	u, _ := url.Parse("/access_settings/activity")

	if input != nil && input.Limit > 0 {
		q := u.Query()
		q.Set("limit", strconv.Itoa(input.Limit))
		u.RawQuery = q.Encode()
	}

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAccessActivities)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type AllowListIP struct {
	IP string `json:"ip,omitempty"`
}

type InputAddIPsToAllowList struct {
	IPs []AllowListIP `json:"ips,omitempty"`
}

type OutputAddIPsToAllowList struct {
	Result []*AllowedIP `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/add-one-or-more-ips-to-the-allow-list
func (c *Client) AddIPsToAllowList(ctx context.Context, input *InputAddIPsToAllowList) (*OutputAddIPsToAllowList, error) {
	req, err := c.NewRequest("POST", "/access_settings/whitelist", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddIPsToAllowList)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/remove-a-specific-ip-from-the-allowed-list
func (c *Client) RemoveIPFromAllowList(ctx context.Context, ruleID int64) error {
	path := fmt.Sprintf("/access_settings/whitelist/%d", ruleID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type InputRemoveIPsFromAllowList struct {
	IDs []int64 `json:"ids,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/remove-one-or-more-ips-from-the-allow-list
func (c *Client) RemoveIPsFromAllowList(ctx context.Context, input *InputRemoveIPsFromAllowList) error {
	req, err := c.NewRequest("DELETE", "/access_settings/whitelist", input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
