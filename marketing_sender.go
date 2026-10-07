package sendgrid

import (
	"context"
	"fmt"
)

// MarketingSenderAddress represents the from or reply_to address of a
// Marketing Campaigns Sender.
type MarketingSenderAddress struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type MarketingSenderVerified struct {
	Status bool   `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// MarketingSender represents a Marketing Campaigns Sender.
type MarketingSender struct {
	ID        int64                    `json:"id,omitempty"`
	Nickname  string                   `json:"nickname,omitempty"`
	From      *MarketingSenderAddress  `json:"from,omitempty"`
	ReplyTo   *MarketingSenderAddress  `json:"reply_to,omitempty"`
	Address   string                   `json:"address,omitempty"`
	Address2  string                   `json:"address_2,omitempty"`
	City      string                   `json:"city,omitempty"`
	State     string                   `json:"state,omitempty"`
	Zip       string                   `json:"zip,omitempty"`
	Country   string                   `json:"country,omitempty"`
	Verified  *MarketingSenderVerified `json:"verified,omitempty"`
	Locked    bool                     `json:"locked,omitempty"`
	UpdatedAt int64                    `json:"updated_at,omitempty"`
	CreatedAt int64                    `json:"created_at,omitempty"`
}

type InputCreateMarketingSender struct {
	Nickname string                  `json:"nickname,omitempty"`
	From     *MarketingSenderAddress `json:"from,omitempty"`
	ReplyTo  *MarketingSenderAddress `json:"reply_to,omitempty"`
	Address  string                  `json:"address,omitempty"`
	Address2 string                  `json:"address_2,omitempty"`
	City     string                  `json:"city,omitempty"`
	State    string                  `json:"state,omitempty"`
	Zip      string                  `json:"zip,omitempty"`
	Country  string                  `json:"country,omitempty"`
}

type OutputCreateMarketingSender MarketingSender

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/create-a-sender
func (c *Client) CreateMarketingSender(ctx context.Context, input *InputCreateMarketingSender) (*OutputCreateMarketingSender, error) {
	req, err := c.NewRequest("POST", "/marketing/senders", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateMarketingSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/get-a-list-of-all-senders
func (c *Client) GetMarketingSenders(ctx context.Context) ([]*MarketingSender, error) {
	req, err := c.NewRequest("GET", "/marketing/senders", nil)
	if err != nil {
		return nil, err
	}

	r := []*MarketingSender{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetMarketingSender MarketingSender

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/get-a-specific-sender
func (c *Client) GetMarketingSender(ctx context.Context, id int64) (*OutputGetMarketingSender, error) {
	path := fmt.Sprintf("/marketing/senders/%d", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetMarketingSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateMarketingSender struct {
	Nickname string                  `json:"nickname,omitempty"`
	From     *MarketingSenderAddress `json:"from,omitempty"`
	ReplyTo  *MarketingSenderAddress `json:"reply_to,omitempty"`
	Address  string                  `json:"address,omitempty"`
	Address2 string                  `json:"address_2,omitempty"`
	City     string                  `json:"city,omitempty"`
	State    string                  `json:"state,omitempty"`
	Zip      string                  `json:"zip,omitempty"`
	Country  string                  `json:"country,omitempty"`
}

type OutputUpdateMarketingSender MarketingSender

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/update-a-sender
func (c *Client) UpdateMarketingSender(ctx context.Context, id int64, input *InputUpdateMarketingSender) (*OutputUpdateMarketingSender, error) {
	path := fmt.Sprintf("/marketing/senders/%d", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateMarketingSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/delete-a-sender
func (c *Client) DeleteMarketingSender(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/marketing/senders/%d", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/marketing-campaigns-senders/resend-a-sender-verification
func (c *Client) ResendMarketingSenderVerification(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/marketing/senders/%d/resend_verification", id)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
