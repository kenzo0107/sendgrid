package sendgrid

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type OutputGetInboundParseWebhooks struct {
	Result []*InboundParseWebhook `json:"result,omitempty"`
}

type InboundParseWebhook struct {
	URL            string `json:"url,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	SpamCheck      bool   `json:"spam_check,omitempty"`
	SendRaw        bool   `json:"send_raw,omitempty"`
	SecurityPolicy string `json:"security_policy,omitempty"`
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
	URL            string `json:"url,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	SpamCheck      bool   `json:"spam_check,omitempty"`
	SendRaw        bool   `json:"send_raw,omitempty"`
	SecurityPolicy string `json:"security_policy,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/retrieve-a-specific-parse-setting
func (c *Client) GetInboundParseWebhook(ctx context.Context, hostname string) (*OutputGetInboundParseWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", url.PathEscape(hostname))

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
	URL            string `json:"url,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	SpamCheck      bool   `json:"spam_check"`
	SendRaw        bool   `json:"send_raw"`
	SecurityPolicy string `json:"security_policy,omitempty"`
}

type OutputCreateInboundParseWebhook struct {
	URL            string `json:"url,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	SpamCheck      bool   `json:"spam_check,omitempty"`
	SendRaw        bool   `json:"send_raw,omitempty"`
	SecurityPolicy string `json:"security_policy,omitempty"`
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
	SpamCheck bool   `json:"spam_check"`
	SendRaw   bool   `json:"send_raw"`
	// A plain string would send "" on every update that does not set it, so nil is needed to leave the policy unchanged.
	SecurityPolicy *string `json:"security_policy,omitempty"`
}

// A separate detach field is not used because it could conflict with
// SecurityPolicy; "" can stand for null since an empty policy ID is never valid.
func (i InputUpdateInboundParseWebhook) MarshalJSON() ([]byte, error) {
	type alias InputUpdateInboundParseWebhook
	if i.SecurityPolicy == nil || *i.SecurityPolicy != "" {
		return json.Marshal(alias(i))
	}
	return json.Marshal(struct {
		alias
		SecurityPolicy *string `json:"security_policy"`
	}{alias: alias(i)})
}

type OutputUpdateInboundParseWebhook struct {
	URL            string `json:"url,omitempty"`
	Hostname       string `json:"hostname,omitempty"`
	SpamCheck      bool   `json:"spam_check,omitempty"`
	SendRaw        bool   `json:"send_raw,omitempty"`
	SecurityPolicy string `json:"security_policy,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/settings-inbound-parse/update-a-parse-setting
func (c *Client) UpdateInboundParseWebhook(ctx context.Context, hostname string, input *InputUpdateInboundParseWebhook) (*OutputUpdateInboundParseWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", url.PathEscape(hostname))

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
	path := fmt.Sprintf("/user/webhooks/parse/settings/%s", url.PathEscape(hostname))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
