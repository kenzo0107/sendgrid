package sendgrid

import (
	"context"
	"fmt"
)

// LegacySenderAddress represents the from or reply_to address of a Legacy
// Marketing Campaigns sender identity.
type LegacySenderAddress struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type LegacySenderVerified struct {
	Status bool   `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// LegacySender represents a Legacy Marketing Campaigns sender identity.
type LegacySender struct {
	ID        int64                 `json:"id,omitempty"`
	Nickname  string                `json:"nickname,omitempty"`
	From      *LegacySenderAddress  `json:"from,omitempty"`
	ReplyTo   *LegacySenderAddress  `json:"reply_to,omitempty"`
	Address   string                `json:"address,omitempty"`
	Address2  string                `json:"address_2,omitempty"`
	City      string                `json:"city,omitempty"`
	State     string                `json:"state,omitempty"`
	Zip       string                `json:"zip,omitempty"`
	Country   string                `json:"country,omitempty"`
	Verified  *LegacySenderVerified `json:"verified,omitempty"`
	UpdatedAt int64                 `json:"updated_at,omitempty"`
	CreatedAt int64                 `json:"created_at,omitempty"`
	Locked    bool                  `json:"locked,omitempty"`
}

type InputCreateLegacySender struct {
	Nickname string               `json:"nickname,omitempty"`
	From     *LegacySenderAddress `json:"from,omitempty"`
	ReplyTo  *LegacySenderAddress `json:"reply_to,omitempty"`
	Address  string               `json:"address,omitempty"`
	Address2 string               `json:"address_2,omitempty"`
	City     string               `json:"city,omitempty"`
	State    string               `json:"state,omitempty"`
	Zip      string               `json:"zip,omitempty"`
	Country  string               `json:"country,omitempty"`
}

type OutputCreateLegacySender LegacySender

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/create-a-sender-identity
func (c *Client) CreateLegacySender(ctx context.Context, input *InputCreateLegacySender) (*OutputCreateLegacySender, error) {
	req, err := c.NewRequest("POST", "/senders", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateLegacySender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/get-all-sender-identities
func (c *Client) GetLegacySenders(ctx context.Context) ([]*LegacySender, error) {
	req, err := c.NewRequest("GET", "/senders", nil)
	if err != nil {
		return nil, err
	}

	r := []*LegacySender{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacySender LegacySender

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/view-a-sender-identity
func (c *Client) GetLegacySender(ctx context.Context, id int64) (*OutputGetLegacySender, error) {
	path := fmt.Sprintf("/senders/%d", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacySender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateLegacySender struct {
	Nickname string               `json:"nickname,omitempty"`
	From     *LegacySenderAddress `json:"from,omitempty"`
	ReplyTo  *LegacySenderAddress `json:"reply_to,omitempty"`
	Address  string               `json:"address,omitempty"`
	Address2 string               `json:"address_2,omitempty"`
	City     string               `json:"city,omitempty"`
	State    string               `json:"state,omitempty"`
	Zip      string               `json:"zip,omitempty"`
	Country  string               `json:"country,omitempty"`
}

type OutputUpdateLegacySender LegacySender

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/update-a-sender-identity
func (c *Client) UpdateLegacySender(ctx context.Context, id int64, input *InputUpdateLegacySender) (*OutputUpdateLegacySender, error) {
	path := fmt.Sprintf("/senders/%d", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateLegacySender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/delete-a-sender-identity
func (c *Client) DeleteLegacySender(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/senders/%d", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/sender-identities-api/resend-sender-identity-verification
func (c *Client) ResendLegacySenderVerification(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/senders/%d/resend_verification", id)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
