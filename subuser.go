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
}

type InputGetSubusers struct {
	Username string
	Limit    int
	Offset   int
}

func (c *Client) GetSubusers(ctx context.Context, input *InputGetSubusers) ([]*Subuser, error) {
	u, err := url.Parse("/subusers")
	if err != nil {
		return nil, err
	}

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
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Ips      []string `json:"ips"`
}

type OutputCreateSubuser struct {
	UserID             int64            `json:"user_id"`
	Username           string           `json:"username"`
	Email              string           `json:"email"`
	SignupSessionToken string           `json:"signup_session_token"`
	AuthorizationToken string           `json:"authorization_token"`
	Password           string           `json:"password"`
	CreditAllocation   CreditAllocation `json:"credit_allocation"`
}

type CreditAllocation struct {
	Type string `json:"type"`
}

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
