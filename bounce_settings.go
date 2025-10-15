package sendgrid

import (
	"context"
	"log"
)

type OutputGetBounceSettings struct {
	Enabled     bool  `json:"enabled,omitempty"`
	SoftBounces int64 `json:"soft_bounces,omitempty"`
	HardBounces int64 `json:"hard_bounces,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/bounces/remove-bounces
func (c *Client) GetBounceSettings(ctx context.Context) (*OutputGetBounceSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/bounce_purge", nil)
	if err != nil {
		return nil, err
	}

	var o OutputGetBounceSettings
	if err := c.Do(ctx, req, &o); err != nil {
		return nil, err
	}

	return &o, nil
}

type InputUpdateBounceSettings struct {
	Enabled     bool  `json:"enabled"`
	SoftBounces int64 `json:"soft_bounces,omitempty"`
	HardBounces int64 `json:"hard_bounces,omitempty"`
}

type OutputUpdateBounceSettings struct {
	Enabled     bool  `json:"enabled,omitempty"`
	SoftBounces int64 `json:"soft_bounces,omitempty"`
	HardBounces int64 `json:"hard_bounces,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/settings-mail/update-bounce-purge-mail-settings
func (c *Client) UpdateBounceSettings(ctx context.Context, input *InputUpdateBounceSettings) (*OutputUpdateBounceSettings, error) {
	log.Printf("input: %#v", input)
	req, err := c.NewRequest("PATCH", "/mail_settings/bounce_purge", input)
	if err != nil {
		return nil, err
	}

	var o OutputUpdateBounceSettings
	if err := c.Do(ctx, req, &o); err != nil {
		return nil, err
	}
	return &o, nil
}
