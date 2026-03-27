package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type Subuser struct {
	ID       int64  `json:"id,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Region   string `json:"region,omitempty"`
}

// document link not found
func (c *Client) GetSubuser(ctx context.Context, username string) (*Subuser, error) {
	path := fmt.Sprintf("/subusers/%s", username)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(Subuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetSubusers struct {
	Username      string
	Limit         int
	Offset        int
	Region        string
	IncludeRegion bool
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/list-all-subusers
func (c *Client) GetSubusers(ctx context.Context, input *InputGetSubusers) ([]*Subuser, error) {
	u, _ := url.Parse("/subusers")

	q := u.Query()
	if input.Username != "" {
		q.Set("username", input.Username)
	}
	if input.Limit > 0 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}
	if input.Offset > 0 {
		q.Set("offset", strconv.Itoa(input.Offset))
	}
	if input.Region != "" {
		q.Set("region", input.Region)
	}
	if input.IncludeRegion {
		q.Set("include_region", "true")
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*Subuser{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type Reputation struct {
	Reputation float64 `json:"reputation,omitempty"`
	Username   string  `json:"username,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/retrieve-subuser-reputations
func (c *Client) GetSubuserReputations(ctx context.Context, usernames string) ([]*Reputation, error) {
	path := fmt.Sprintf("/subusers/reputations?usernames=%s", usernames)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := []*Reputation{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputCreateSubuser struct {
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Password      string   `json:"password"`
	Ips           []string `json:"ips"`
	Region        string   `json:"region"`
	IncludeRegion bool     `json:"include_region"`
}

type OutputCreateSubuser struct {
	UserID             int64            `json:"user_id"`
	Username           string           `json:"username"`
	Email              string           `json:"email"`
	SignupSessionToken string           `json:"signup_session_token"`
	AuthorizationToken string           `json:"authorization_token"`
	CreditAllocation   CreditAllocation `json:"credit_allocation"`
	Region             string           `json:"region"`
}

type CreditAllocation struct {
	Type string `json:"type"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/create-subuser
func (c *Client) CreateSubuser(ctx context.Context, input *InputCreateSubuser) (*OutputCreateSubuser, error) {
	req, err := c.NewRequest("POST", "/subusers", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateSubuserStatus struct {
	Disabled bool `json:"disabled"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/enabledisable-a-subuser
func (c *Client) UpdateSubuserStatus(ctx context.Context, username string, input *InputUpdateSubuserStatus) error {
	path := fmt.Sprintf("/subusers/%s", username)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/update-ips-assigned-to-a-subuser
func (c *Client) UpdateSubuserIps(ctx context.Context, username string, ips []string) error {
	path := fmt.Sprintf("/subusers/%s/ips", username)

	req, err := c.NewRequest("PUT", path, ips)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/delete-a-subuser
func (c *Client) DeleteSubuser(ctx context.Context, username string) error {
	path := fmt.Sprintf("/subusers/%s", username)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type OutputGetCreditsForSubuser struct {
	Type           string `json:"type,omitempty"`
	ResetFrequency string `json:"reset_frequency,omitempty"`
	Remain         int    `json:"remain,omitempty"`
	Total          int    `json:"total,omitempty"`
	Used           int    `json:"used,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/get-credits-for-subuser
func (c *Client) GetCreditsForSubuser(ctx context.Context, username string) (*OutputGetCreditsForSubuser, error) {
	path := fmt.Sprintf("/subusers/%s/credits", username)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetCreditsForSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateCreditsForSubuser struct {
	Type           string `json:"type,omitempty"`
	ResetFrequency string `json:"reset_frequency,omitempty"`
	Total          int    `json:"total,omitempty"`
}

type OutputUpdateCreditsForSubuser struct {
	Type           string `json:"type,omitempty"`
	ResetFrequency string `json:"reset_frequency,omitempty"`
	Remain         int    `json:"remain,omitempty"`
	Total          int    `json:"total,omitempty"`
	Used           int    `json:"used,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/update-credits-for-subuser
func (c *Client) UpdateCreditsForSubuser(ctx context.Context, username string, input *InputUpdateCreditsForSubuser) (*OutputUpdateCreditsForSubuser, error) {
	path := fmt.Sprintf("/subusers/%s/credits", username)

	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateCreditsForSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateRemainingCreditsForSubuser struct {
	AllocationUpdate int `json:"allocation_update,omitempty"`
}

type OutputUpdateRemainingCreditsForSubuser struct {
	Type           string `json:"type,omitempty"`
	ResetFrequency string `json:"reset_frequency,omitempty"`
	Remain         int    `json:"remain,omitempty"`
	Total          int    `json:"total,omitempty"`
	Used           int    `json:"used,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/update-remaining-credits-for-subuser
func (c *Client) UpdateRemainingCreditsForSubuser(ctx context.Context, username string, input *InputUpdateRemainingCreditsForSubuser) (*OutputUpdateRemainingCreditsForSubuser, error) {
	path := fmt.Sprintf("/subusers/%s/credits/remaining", username)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateRemainingCreditsForSubuser)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}
