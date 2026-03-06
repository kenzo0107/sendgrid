package sendgrid

import (
	"context"
	"fmt"
	"net/url"
)

type InputStartIPWarmup struct {
	IP string `json:"ip"`
}

type IPWarmup struct {
	IP        string `json:"ip,omitempty"`
	StartDate int64  `json:"start_date,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/start-warming-up-an-ip-address
func (c *Client) StartIPWarmup(ctx context.Context, input *InputStartIPWarmup) ([]*IPWarmup, error) {
	req, err := c.NewRequest("POST", "/ips/warmup", input)
	if err != nil {
		return nil, err
	}

	var r []*IPWarmup
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/stop-warming-up-an-ip-address
func (c *Client) StopIPWarmup(ctx context.Context, ip string) error {
	path := fmt.Sprintf("/ips/warmup/%s", url.QueryEscape(ip))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/retrieve-ip-warmup-status
func (c *Client) GetIPWarmupStatus(ctx context.Context, ip string) ([]*IPWarmup, error) {
	path := fmt.Sprintf("/ips/warmup/%s", url.QueryEscape(ip))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var r []*IPWarmup
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/retrieve-all-ips-currently-in-warmup
func (c *Client) GetAllIPWarmup(ctx context.Context) ([]*IPWarmup, error) {
	req, err := c.NewRequest("GET", "/ips/warmup", nil)
	if err != nil {
		return nil, err
	}

	var r []*IPWarmup
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
