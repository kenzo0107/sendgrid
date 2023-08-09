package sendgrid

import (
	"context"
	"fmt"
)

type User struct {
	Username  string   `json:"username,omitempty"`
	Email     string   `json:"email,omitempty"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	Address   string   `json:"address,omitempty"`
	Address2  string   `json:"address2,omitempty"`
	City      string   `json:"city,omitempty"`
	State     string   `json:"state,omitempty"`
	Zip       string   `json:"zip,omitempty"`
	Country   string   `json:"country,omitempty"`
	Company   string   `json:"company,omitempty"`
	Website   string   `json:"website,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	IsAdmin   bool     `json:"is_admin,omitempty"`
	IsSSO     bool     `json:"is_sso,omitempty"`
	UserType  string   `json:"user_type,omitempty"`
	Scopes    []string `json:"scopes,omitempty"`

	IsReadOnly     bool   `json:"is_read_only,omitempty"`
	Token          string `json:"token,omitempty"`
	ExpirationDate int    `json:"expiration_date,omitempty"`
}

func (c *Client) GetTeammate(ctx context.Context, username string) (*User, error) {
	u := fmt.Sprintf("/teammates/%s", username)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (c *Client) GetUsernameByEmail(ctx context.Context, email string) (username string, err error) {
	users, err := c.GetTeammates(ctx)
	if err != nil {
		return "", err
	}

	for _, user := range users {
		if user.Email != email {
			continue
		}
		username = user.Username
	}
	return username, nil
}

func (c *Client) GetTeammates(ctx context.Context) ([]*User, error) {
	req, err := c.NewRequest("GET", "/teammates", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Users []*User `json:"result,omitempty"`
	}

	r := new(Response)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.Users, nil
}

func (c *Client) GetPendingTeammates(ctx context.Context) ([]*User, error) {
	req, err := c.NewRequest("GET", "/teammates/pending", nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Users []*User `json:"result,omitempty"`
	}

	r := new(Response)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.Users, nil
}

type InputInviteTeammate struct {
	Email   *string   `json:"email"`
	IsAdmin *bool     `json:"is_admin"`
	Scopes  []*string `json:"scopes"`
}

func (c *Client) InviteTeammate(ctx context.Context, input *InputInviteTeammate) (*User, error) {
	req, err := c.NewRequest("POST", "/teammates", input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
}

type InputUpdateTeammatePermissions struct {
	IsAdmin *bool     `json:"is_admin"`
	Scopes  []*string `json:"scopes"`
}

func (c *Client) UpdateTeammatePermissions(ctx context.Context, username string, input *InputUpdateTeammatePermissions) (*User, error) {
	u := fmt.Sprintf("/teammates/%s", username)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	user := new(User)
	if err := c.Do(ctx, req, &user); err != nil {
		return nil, err
	}
	return user, nil
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
