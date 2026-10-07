package sendgrid

import (
	"context"
)

type PartnerSetting struct {
	Title       string `json:"title,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type OutputGetPartnerSettings struct {
	Result []*PartnerSetting `json:"result,omitempty"`
}

type InputGetPartnerSettings struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/partner-settings/returns-a-list-of-all-partner-settings
func (c *Client) GetPartnerSettings(ctx context.Context, input *InputGetPartnerSettings) (*OutputGetPartnerSettings, error) {
	path, err := c.AddOptions("/partner_settings", input)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetPartnerSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
