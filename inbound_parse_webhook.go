package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetInboundParsetWebhooks struct {
	Result []*InboundParsetWebhook `json:"result,omitempty"`
}

type InboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/retrieve-all-parse-settings
func (c *Client) GetInboundParsetWebhooks(ctx context.Context) ([]*InboundParsetWebhook, error) {
	req, err := c.NewRequest("GET", "/user/webhooks/parse/settings", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetInboundParsetWebhooks)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.Result, nil
}

type OutputGetInboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/retrieve-a-specific-parse-setting
func (c *Client) GetInboundParsetWebhook(ctx context.Context, hostname string) (*OutputGetInboundParsetWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", hostname)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetInboundParsetWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateInboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

type OutputCreateInboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/create-a-parse-setting
func (c *Client) CreateInboundParsetWebhook(ctx context.Context, input *InputCreateInboundParsetWebhook) (*OutputCreateInboundParsetWebhook, error) {
	req, err := c.NewRequest("POST", "/user/webhooks/parse/settings", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateInboundParsetWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateInboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

type OutputUpdateInboundParsetWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/update-a-parse-setting
func (c *Client) UpdateInboundParsetWebhook(ctx context.Context, hostname string, input *InputUpdateInboundParsetWebhook) (*OutputUpdateInboundParsetWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", hostname)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateInboundParsetWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/delete-a-parse-setting
func (c *Client) DeleteInboundParsetWebhook(ctx context.Context, hostname string) error {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", hostname)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
