package sendgrid

import (
	"context"
)

type MailSetting struct {
	Title       string `json:"title,omitempty"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type InputGetMailSettings struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

type OutputGetMailSettings struct {
	Result []*MailSetting `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-all-mail-settings
func (c *Client) GetMailSettings(ctx context.Context, input *InputGetMailSettings) (*OutputGetMailSettings, error) {
	path, err := c.AddOptions("/mail_settings", input)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetAddressWhitelistMailSettings struct {
	Enabled bool     `json:"enabled"`
	List    []string `json:"list,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-address-whitelist-mail-settings
func (c *Client) GetAddressWhitelistMailSettings(ctx context.Context) (*OutputGetAddressWhitelistMailSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/address_whitelist", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAddressWhitelistMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateAddressWhitelistMailSettings struct {
	Enabled bool     `json:"enabled"`
	List    []string `json:"list,omitempty"`
}

type OutputUpdateAddressWhitelistMailSettings struct {
	Enabled bool     `json:"enabled"`
	List    []string `json:"list,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/update-address-whitelist-mail-settings
func (c *Client) UpdateAddressWhitelistMailSettings(ctx context.Context, input *InputUpdateAddressWhitelistMailSettings) (*OutputUpdateAddressWhitelistMailSettings, error) {
	req, err := c.NewRequest("PATCH", "/mail_settings/address_whitelist", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateAddressWhitelistMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetFooterMailSettings struct {
	Enabled      bool   `json:"enabled"`
	HTMLContent  string `json:"html_content,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-footer-mail-settings
func (c *Client) GetFooterMailSettings(ctx context.Context) (*OutputGetFooterMailSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/footer", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetFooterMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateFooterMailSettings struct {
	Enabled      bool   `json:"enabled"`
	HTMLContent  string `json:"html_content,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
}

type OutputUpdateFooterMailSettings struct {
	Enabled      bool   `json:"enabled"`
	HTMLContent  string `json:"html_content,omitempty"`
	PlainContent string `json:"plain_content,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/update-footer-mail-settings
func (c *Client) UpdateFooterMailSettings(ctx context.Context, input *InputUpdateFooterMailSettings) (*OutputUpdateFooterMailSettings, error) {
	req, err := c.NewRequest("PATCH", "/mail_settings/footer", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateFooterMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetForwardBounceMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-forward-bounce-mail-settings
func (c *Client) GetForwardBounceMailSettings(ctx context.Context) (*OutputGetForwardBounceMailSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/forward_bounce", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetForwardBounceMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateForwardBounceMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

type OutputUpdateForwardBounceMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/update-forward-bounce-mail-settings
func (c *Client) UpdateForwardBounceMailSettings(ctx context.Context, input *InputUpdateForwardBounceMailSettings) (*OutputUpdateForwardBounceMailSettings, error) {
	req, err := c.NewRequest("PATCH", "/mail_settings/forward_bounce", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateForwardBounceMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetForwardSpamMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-forward-spam-mail-settings
func (c *Client) GetForwardSpamMailSettings(ctx context.Context) (*OutputGetForwardSpamMailSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/forward_spam", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetForwardSpamMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateForwardSpamMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

type OutputUpdateForwardSpamMailSettings struct {
	Enabled bool   `json:"enabled"`
	Email   string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/update-forward-spam-mail-settings
func (c *Client) UpdateForwardSpamMailSettings(ctx context.Context, input *InputUpdateForwardSpamMailSettings) (*OutputUpdateForwardSpamMailSettings, error) {
	req, err := c.NewRequest("PATCH", "/mail_settings/forward_spam", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateForwardSpamMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetTemplateMailSettings struct {
	Enabled     bool   `json:"enabled"`
	HTMLContent string `json:"html_content,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/retrieve-legacy-template-mail-settings
func (c *Client) GetTemplateMailSettings(ctx context.Context) (*OutputGetTemplateMailSettings, error) {
	req, err := c.NewRequest("GET", "/mail_settings/template", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTemplateMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateTemplateMailSettings struct {
	Enabled     bool   `json:"enabled"`
	HTMLContent string `json:"html_content,omitempty"`
}

type OutputUpdateTemplateMailSettings struct {
	Enabled     bool   `json:"enabled"`
	HTMLContent string `json:"html_content,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/mail-settings/update-template-mail-settings
func (c *Client) UpdateTemplateMailSettings(ctx context.Context, input *InputUpdateTemplateMailSettings) (*OutputUpdateTemplateMailSettings, error) {
	req, err := c.NewRequest("PATCH", "/mail_settings/template", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateTemplateMailSettings)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
