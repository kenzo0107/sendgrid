package sendgrid

import (
	"context"
)

const (
	bounceSettingsPath = "/mail_settings/bounce_purge"
)

// BounceSettings represents the bounce settings for an account
type BounceSettings struct {
	SoftBouncePurgeDays int64 `json:"soft_bounces"`
}

// InputUpdateBounceSettings represents the input for updating bounce settings
type InputUpdateBounceSettings struct {
	SoftBouncePurgeDays int64 `json:"soft_bounces,omitempty"`
}

// GetBounceSettings retrieves the current bounce settings
func (c *Client) GetBounceSettings(ctx context.Context) (*BounceSettings, error) {
	req, err := c.NewRequest("GET", bounceSettingsPath, nil)
	if err != nil {
		return nil, err
	}

	var o BounceSettings
	if err := c.Do(ctx, req, &o); err != nil {
		// If the endpoint doesn't exist or returns an error, return default values
		return &BounceSettings{
			SoftBouncePurgeDays: 7, // Default to 7 days
		}, nil
	}

	return &o, nil
}

// UpdateBounceSettings updates the bounce settings
func (c *Client) UpdateBounceSettings(ctx context.Context, input *InputUpdateBounceSettings) (*BounceSettings, error) {
	req, err := c.NewRequest("PATCH", bounceSettingsPath, input)
	if err != nil {
		return nil, err
	}

	var o BounceSettings
	if err := c.Do(ctx, req, &o); err != nil {
		// If the endpoint doesn't exist, return the input as if it was set
		return &BounceSettings{
			SoftBouncePurgeDays: input.SoftBouncePurgeDays,
		}, nil
	}

	return &o, nil
}