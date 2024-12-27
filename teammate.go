package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// consolidating normal teammate and SSO teammate fields
type Member struct {
	Teammate

	Company string   `json:"company,omitempty"`
	IsSSO   bool     `json:"is_sso,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
}

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

type InputGetTeammateSubuserAccess struct {
	AfterSubuserID int64  `json:"after_subuser_id,omitempty"`
	Limit          int64  `json:"limit,omitempty"`
	Username       string `json:"username,omitempty"`
}
type OutputGetTeammateSubuserAccess struct {
	HasRestrictedSubuserAccess bool                             `json:"has_restricted_subuser_access,omitempty"`
	SubuserAccess              []SubuserAccess                  `json:"subuser_access,omitempty"`
	Metadata                   MetadataGetTeammateSubuserAccess `json:"_metadata,omitempty"`
}

type SubuserAccess struct {
	ID             int64    `json:"id,omitempty"`
	Username       string   `json:"username,omitempty"`
	Email          string   `json:"email,omitempty"`
	Disabled       bool     `json:"disabled,omitempty"`
	PermissionType string   `json:"permission_type,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
}

type MetadataGetTeammateSubuserAccess struct {
	NextParams NextParams `json:"next_params,omitempty"`
}

type NextParams struct {
	Limit          int64  `json:"limit"`
	AfterSubuserID int64  `json:"after_subuser_id,omitempty"`
	Username       string `json:"username,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/teammates/get-teammate-subuser-access
func (c *Client) GetTeammateSubuserAccess(ctx context.Context, teammateName string, input *InputGetTeammateSubuserAccess) (*OutputGetTeammateSubuserAccess, error) {
	p := fmt.Sprintf("/teammates/%s/subuser_access", teammateName)
	u, err := url.Parse(p)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if input.AfterSubuserID > 0 {
		q.Set("after_subuser_id", strconv.FormatInt(input.AfterSubuserID, 10))
	}
	if input.Limit > 0 {
		q.Set("limit", strconv.FormatInt(input.Limit, 10))
	}
	if input.Username != "" {
		q.Set("username", input.Username)
	}

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTeammateSubuserAccess)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateSSOTeammate struct {
	Email                      string               `json:"email"`
	FirstName                  string               `json:"first_name"`
	LastName                   string               `json:"last_name"`
	IsAdmin                    bool                 `json:"is_admin"`
	IsSSO                      bool                 `json:"is_sso"`
	Persona                    string               `json:"persona,omitempty"`
	Scopes                     []string             `json:"scopes,omitempty"`
	HasRestrictedSubuserAccess bool                 `json:"has_restricted_subuser_access,omitempty"`
	SubuserAccess              []InputSubuserAccess `json:"subuser_access,omitempty"`
}

type InputSubuserAccess struct {
	ID             int64    `json:"id,omitempty"`
	PermissionType string   `json:"permission_type,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
}

type OutputCreateSSOTeammate struct {
	FirstName                  string                `json:"first_name,omitempty"`
	LastName                   string                `json:"last_name,omitempty"`
	Email                      string                `json:"email,omitempty"`
	IsAdmin                    bool                  `json:"is_admin,omitempty"`
	IsSSO                      bool                  `json:"is_sso,omitempty"`
	Scopes                     []string              `json:"scopes,omitempty"`
	HasRestrictedSubuserAccess bool                  `json:"has_restricted_subuser_access,omitempty"`
	SubuserAccess              []OutputSubuserAccess `json:"subuser_access,omitempty"`
}

type OutputSubuserAccess struct {
	ID             int64    `json:"id,omitempty"`
	Username       int64    `json:"username,omitempty"`
	Email          string   `json:"email,omitempty"`
	Disabled       bool     `json:"disabled,omitempty"`
	PermissionType string   `json:"permission_type,omitempty"`
	Scopes         []string `json:"scopes,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-teammates/create-sso-teammate
func (c *Client) CreateSSOTeammate(ctx context.Context, input *InputCreateSSOTeammate) (*OutputCreateSSOTeammate, error) {
	req, err := c.NewRequest("POST", "/sso/teammates", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSSOTeammate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateSSOTeammate struct {
	FirstName                  string               `json:"first_name"`
	LastName                   string               `json:"last_name"`
	IsAdmin                    bool                 `json:"is_admin"`
	Persona                    string               `json:"persona,omitempty"`
	Scopes                     []string             `json:"scopes,omitempty"`
	HasRestrictedSubuserAccess bool                 `json:"has_restricted_subuser_access,omitempty"`
	SubuserAccess              []InputSubuserAccess `json:"subuser_access,omitempty"`
}

type OutputUpdateSSOTeammate struct {
	Address                    string                `json:"address,omitempty"`
	Address2                   string                `json:"address2,omitempty"`
	City                       string                `json:"city,omitempty"`
	Company                    string                `json:"company,omitempty"`
	Country                    string                `json:"country,omitempty"`
	Username                   string                `json:"username,omitempty"`
	Phone                      string                `json:"phone,omitempty"`
	State                      string                `json:"state,omitempty"`
	UserType                   string                `json:"user_type,omitempty"`
	Website                    string                `json:"website,omitempty"`
	Zip                        string                `json:"zip,omitempty"`
	FirstName                  string                `json:"first_name,omitempty"`
	LastName                   string                `json:"last_name,omitempty"`
	Email                      string                `json:"email,omitempty"`
	IsAdmin                    bool                  `json:"is_admin,omitempty"`
	IsSSO                      bool                  `json:"is_sso,omitempty"`
	Scopes                     []string              `json:"scopes,omitempty"`
	HasRestrictedSubuserAccess bool                  `json:"has_restricted_subuser_access,omitempty"`
	SubuserAccess              []OutputSubuserAccess `json:"subuser_access,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-teammates/edit-an-sso-teammate
func (c *Client) UpdateSSOTeammate(ctx context.Context, username string, input *InputUpdateSSOTeammate) (*OutputUpdateSSOTeammate, error) {
	u := fmt.Sprintf("/sso/teammates/%s", username)
	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSSOTeammate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}
