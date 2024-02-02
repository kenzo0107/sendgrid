package sendgrid

import (
	"context"
)

type OutputGetTrackingSettings struct {
	Result []*ResultGetTrackingSettings `json:"result,omitempty"`
}

type ResultGetTrackingSettings struct {
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/retrieve-tracking-settings
func (c *Client) GetTrackingSettings(ctx context.Context) (*OutputGetTrackingSettings, error) {
	req, err := c.NewRequest("GET", "/tracking_settings", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetClickTrackingSettings struct {
	EnableText bool `json:"enable_text,omitempty"`
	Enabled    bool `json:"enabled,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/retrieve-click-track-settings
func (c *Client) GetClickTrackingSettings(ctx context.Context) (*OutputGetClickTrackingSettings, error) {
	req, err := c.NewRequest("GET", "/tracking_settings/click", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetClickTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateClickTrackingSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

type OutputUpdateClickTrackingSettings struct {
	EnableText bool `json:"enable_text,omitempty"`
	Enabled    bool `json:"enabled,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/update-click-tracking-settings
func (c *Client) UpdateClickTrackingSettings(ctx context.Context, input *InputUpdateClickTrackingSettings) (*OutputUpdateClickTrackingSettings, error) {
	req, err := c.NewRequest("PATCH", "/tracking_settings/click", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateClickTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetOpenTrackingSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/get-open-tracking-settings
func (c *Client) GetOpenTrackingSettings(ctx context.Context) (*OutputGetOpenTrackingSettings, error) {
	req, err := c.NewRequest("GET", "/tracking_settings/open", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetOpenTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateOpenTrackingSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

type OutputUpdateOpenTrackingSettings struct {
	Enabled bool `json:"enabled,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/update-open-tracking-settings
func (c *Client) UpdateOpenTrackingSettings(ctx context.Context, input *InputUpdateOpenTrackingSettings) (*OutputUpdateOpenTrackingSettings, error) {
	req, err := c.NewRequest("PATCH", "/tracking_settings/open", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateOpenTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetGoogleAnalyticsSettings struct {
	Enabled     bool   `json:"enabled,omitempty"`
	UTMCampaign string `json:"utm_campaign,omitempty"`
	UTMContent  string `json:"utm_content,omitempty"`
	UTMMedium   string `json:"utm_medium,omitempty"`
	UTMSource   string `json:"utm_source,omitempty"`
	UTMTerm     string `json:"utm_term,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/retrieve-google-analytics-settings
func (c *Client) GetGoogleAnalyticsSettings(ctx context.Context) (*OutputGetGoogleAnalyticsSettings, error) {
	req, err := c.NewRequest("GET", "/tracking_settings/google_analytics", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetGoogleAnalyticsSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateGoogleAnalyticsSettings struct {
	Enabled     bool   `json:"enabled,omitempty"`
	UTMCampaign string `json:"utm_campaign,omitempty"`
	UTMContent  string `json:"utm_content,omitempty"`
	UTMMedium   string `json:"utm_medium,omitempty"`
	UTMSource   string `json:"utm_source,omitempty"`
	UTMTerm     string `json:"utm_term,omitempty"`
}

type OutputUpdateGoogleAnalyticsSettings struct {
	Enabled     bool   `json:"enabled,omitempty"`
	UTMCampaign string `json:"utm_campaign,omitempty"`
	UTMContent  string `json:"utm_content,omitempty"`
	UTMMedium   string `json:"utm_medium,omitempty"`
	UTMSource   string `json:"utm_source,omitempty"`
	UTMTerm     string `json:"utm_term,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/update-google-analytics-settings
func (c *Client) UpdateGoogleAnalyticsSettings(ctx context.Context, input *InputUpdateGoogleAnalyticsSettings) (*OutputUpdateGoogleAnalyticsSettings, error) {
	req, err := c.NewRequest("PATCH", "/tracking_settings/google_analytics", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateGoogleAnalyticsSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetSubscriptionTrackingSettings struct {
	Enabled      bool   `json:"enabled,omitempty"`
	HTMLContent  string `json:"html_content,omitempty"`
	Landing      string `json:"landing,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
	Replace      string `json:"replace,omitempty"`
	URL          string `json:"url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/retrieve-subscription-tracking-settings
func (c *Client) GetSubscriptionTrackingSettings(ctx context.Context) (*OutputGetSubscriptionTrackingSettings, error) {
	req, err := c.NewRequest("GET", "/tracking_settings/subscription", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSubscriptionTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateSubscriptionTrackingSettings struct {
	Enabled      bool   `json:"enabled,omitempty"`
	HTMLContent  string `json:"html_content,omitempty"`
	Landing      string `json:"landing,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
	Replace      string `json:"replace,omitempty"`
	URL          string `json:"url,omitempty"`
}

type OutputUpdateSubscriptionTrackingSettings struct {
	Enabled      bool   `json:"enabled,omitempty"`
	HTMLContent  string `json:"html_content,omitempty"`
	Landing      string `json:"landing,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
	Replace      string `json:"replace,omitempty"`
	URL          string `json:"url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-tracking/update-subscription-tracking-settings
func (c *Client) UpdateSubscriptionTrackingSettings(ctx context.Context, input *InputUpdateSubscriptionTrackingSettings) (*OutputUpdateSubscriptionTrackingSettings, error) {
	req, err := c.NewRequest("PATCH", "/tracking_settings/subscription", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSubscriptionTrackingSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
