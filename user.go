package sendgrid

import (
	"context"
)

type OutputGetUserAccount struct {
	Type               string  `json:"type,omitempty"`
	Reputation         float64 `json:"reputation,omitempty"`
	IsResellerCustomer bool    `json:"is_reseller_customer,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/get-a-users-account-information
func (c *Client) GetUserAccount(ctx context.Context) (*OutputGetUserAccount, error) {
	req, err := c.NewRequest("GET", "/user/account", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetUserAccount)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetUserCredits struct {
	Remain         int64  `json:"remain,omitempty"`
	Total          int64  `json:"total,omitempty"`
	Overage        int64  `json:"overage,omitempty"`
	Used           int64  `json:"used,omitempty"`
	LastReset      string `json:"last_reset,omitempty"`
	NextReset      string `json:"next_reset,omitempty"`
	ResetFrequency string `json:"reset_frequency,omitempty"`
	IsHardLimit    bool   `json:"is_hard_limit,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/retrieve-your-credit-balance
func (c *Client) GetUserCredits(ctx context.Context) (*OutputGetUserCredits, error) {
	req, err := c.NewRequest("GET", "/user/credits", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetUserCredits)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetUserEmail struct {
	Email string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/retrieve-your-account-email-address
func (c *Client) GetUserEmail(ctx context.Context) (*OutputGetUserEmail, error) {
	req, err := c.NewRequest("GET", "/user/email", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetUserEmail)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateUserEmail struct {
	Email string `json:"email,omitempty"`
}

type OutputUpdateUserEmail struct {
	Email string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/update-your-account-email-address
func (c *Client) UpdateUserEmail(ctx context.Context, input *InputUpdateUserEmail) (*OutputUpdateUserEmail, error) {
	req, err := c.NewRequest("PUT", "/user/email", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateUserEmail)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateUserPassword struct {
	NewPassword string `json:"new_password,omitempty"`
	OldPassword string `json:"old_password,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/update-your-password
func (c *Client) UpdateUserPassword(ctx context.Context, input *InputUpdateUserPassword) error {
	req, err := c.NewRequest("PUT", "/user/password", input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type OutputGetUserProfile struct {
	Address          string `json:"address,omitempty"`
	Address2         string `json:"address2,omitempty"`
	City             string `json:"city,omitempty"`
	Company          string `json:"company,omitempty"`
	Country          string `json:"country,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Phone            string `json:"phone,omitempty"`
	State            string `json:"state,omitempty"`
	Website          string `json:"website,omitempty"`
	Zip              string `json:"zip,omitempty"`
	AuthyID          int64  `json:"authy_id,omitempty"`
	MultifactorPhone string `json:"multifactor_phone,omitempty"`
	Type             string `json:"type,omitempty"`
	UserID           int64  `json:"userid,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/get-a-users-profile
func (c *Client) GetUserProfile(ctx context.Context) (*OutputGetUserProfile, error) {
	req, err := c.NewRequest("GET", "/user/profile", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetUserProfile)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateUserProfile struct {
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Company   string `json:"company,omitempty"`
	Country   string `json:"country,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	State     string `json:"state,omitempty"`
	Website   string `json:"website,omitempty"`
	Zip       string `json:"zip,omitempty"`
}

type OutputUpdateUserProfile struct {
	Address   string `json:"address,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Company   string `json:"company,omitempty"`
	Country   string `json:"country,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	State     string `json:"state,omitempty"`
	Website   string `json:"website,omitempty"`
	Zip       string `json:"zip,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/update-a-users-profile
func (c *Client) UpdateUserProfile(ctx context.Context, input *InputUpdateUserProfile) (*OutputUpdateUserProfile, error) {
	req, err := c.NewRequest("PATCH", "/user/profile", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateUserProfile)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetUsername struct {
	Username string `json:"username,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/retrieve-your-username
func (c *Client) GetUsername(ctx context.Context) (*OutputGetUsername, error) {
	req, err := c.NewRequest("GET", "/user/username", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetUsername)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateUsername struct {
	Username string `json:"username,omitempty"`
}

type OutputUpdateUsername struct {
	Username string `json:"username,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/user-account/update-your-username
func (c *Client) UpdateUsername(ctx context.Context, input *InputUpdateUsername) (*OutputUpdateUsername, error) {
	req, err := c.NewRequest("PUT", "/user/username", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateUsername)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
