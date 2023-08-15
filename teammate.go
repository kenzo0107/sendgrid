package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetTeammate struct {
	Username  string   `json:"username,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
	UserType  string   `json:"user_type,omitempty"`
	IsAdmin   bool     `json:"is_admin,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	Website   string   `json:"website,omitempty"`
	Address   string   `json:"address,omitempty"`
	Address2  string   `json:"address2,omitempty"`
	City      string   `json:"city,omitempty"`
	State     string   `json:"state,omitempty"`
	Zip       string   `json:"zip,omitempty"`
	Country   string   `json:"country,omitempty"`
}

func (c *Client) GetTeammate(ctx context.Context, username string) (*OutputGetTeammate, error) {
	u := fmt.Sprintf("/teammates/%s", username)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTeammate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type Teammate struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	UserType  string `json:"user_type,omitempty"`
	IsAdmin   bool   `json:"is_admin,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Website   string `json:"website,omitempty"`
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
}

type OutputGetTeammates struct {
	Teammates []Teammate `json:"result,omitempty"`
}

func (c *Client) GetTeammates(ctx context.Context) (*OutputGetTeammates, error) {
	req, err := c.NewRequest("GET", "/teammates", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTeammates)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type PendingTeammate struct {
	Email          string   `json:"email,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
	IsAdmin        bool     `json:"is_admin,omitempty"`
	Token          string   `json:"token,omitempty"`
	ExpirationDate int      `json:"expiration_date,omitempty"`
}

type OutputGetPendingTeammates struct {
	PendingTeammates []PendingTeammate `json:"result,omitempty"`
}

func (c *Client) GetPendingTeammates(ctx context.Context) (*OutputGetPendingTeammates, error) {
	req, err := c.NewRequest("GET", "/teammates/pending", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetPendingTeammates)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputInviteTeammate struct {
	Email   string   `json:"email"`
	IsAdmin bool     `json:"is_admin"`
	Scopes  []string `json:"scopes"`
}

type OutputInviteTeammate struct {
	Token   string   `json:"token,omitempty"`
	Email   string   `json:"email"`
	IsAdmin bool     `json:"is_admin"`
	Scopes  []string `json:"scopes"`
}

func (c *Client) InviteTeammate(ctx context.Context, input *InputInviteTeammate) (*OutputInviteTeammate, error) {
	req, err := c.NewRequest("POST", "/teammates", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputInviteTeammate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateTeammatePermissions struct {
	IsAdmin bool     `json:"is_admin"`
	Scopes  []string `json:"scopes"`
}

type OutputUpdateTeammatePermissions struct {
	Username  string   `json:"username,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Email     string   `json:"email,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`
	UserType  string   `json:"user_type,omitempty"`
	IsAdmin   bool     `json:"is_admin,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	Website   string   `json:"website,omitempty"`
	Address   string   `json:"address,omitempty"`
	Address2  string   `json:"address2,omitempty"`
	City      string   `json:"city,omitempty"`
	State     string   `json:"state,omitempty"`
	Zip       string   `json:"zip,omitempty"`
	Country   string   `json:"country,omitempty"`
}

func (c *Client) UpdateTeammatePermissions(ctx context.Context, username string, input *InputUpdateTeammatePermissions) (*OutputUpdateTeammatePermissions, error) {
	u := fmt.Sprintf("/teammates/%s", username)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateTeammatePermissions)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DeleteTeammate(ctx context.Context, username string) error {
	u := fmt.Sprintf("/teammates/%s", username)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeletePendingTeammate(ctx context.Context, token string) error {
	u := fmt.Sprintf("/teammates/pending/%s", token)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
