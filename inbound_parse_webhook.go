package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetInboundParseWebhooks struct {
	Result []*InboundParseWebhook `json:"result,omitempty"`
}

type InboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/retrieve-all-parse-settings
func (c *Client) GetInboundParseWebhooks(ctx context.Context) ([]*InboundParseWebhook, error) {
	req, err := c.NewRequest("GET", "/user/webhooks/parse/settings", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetInboundParseWebhooks)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.Result, nil
}

type OutputGetInboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/retrieve-a-specific-parse-setting
func (c *Client) GetInboundParseWebhook(ctx context.Context, hostname string) (*OutputGetInboundParseWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", hostname)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetInboundParseWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateInboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

type OutputCreateInboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/create-a-parse-setting
func (c *Client) CreateInboundParseWebhook(ctx context.Context, input *InputCreateInboundParseWebhook) (*OutputCreateInboundParseWebhook, error) {
	req, err := c.NewRequest("POST", "/user/webhooks/parse/settings", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateInboundParseWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateInboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

type OutputUpdateInboundParseWebhook struct {
	URL       string `json:"url,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	SpamCheck bool   `json:"spam_check,omitempty"`
	SendRaw   bool   `json:"send_raw,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/update-a-parse-setting
func (c *Client) UpdateInboundParseWebhook(ctx context.Context, hostname string, input *InputUpdateInboundParseWebhook) (*OutputUpdateInboundParseWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", hostname)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateInboundParseWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/delete-a-parse-setting
func (c *Client) DeleteInboundParseWebhook(ctx context.Context, hostname string) error {
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
