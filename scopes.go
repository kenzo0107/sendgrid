package sendgrid

import (
	"context"
)

type OutputGetScopes struct {
	Scopes []string `json:"scopes,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/api-key-permissions/retrieve-a-list-of-scopes-for-which-this-user-has-access
func (c *Client) GetScopes(ctx context.Context) (*OutputGetScopes, error) {
	req, err := c.NewRequest("GET", "/scopes", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetScopes)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
