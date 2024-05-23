package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetEventWebhook struct {
	ID               string `json:"id,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	URL              string `json:"url,omitempty"`
	GroupResubscribe bool   `json:"group_resubscribe,omitempty"`
	Delivered        bool   `json:"delivered,omitempty"`
	GroupUnsubscribe bool   `json:"group_unsubscribe,omitempty"`
	SpamReport       bool   `json:"spam_report,omitempty"`
	Bounce           bool   `json:"bounce,omitempty"`
	Deferred         bool   `json:"deferred,omitempty"`
	Unsubscribe      bool   `json:"unsubscribe,omitempty"`
	Processed        bool   `json:"processed,omitempty"`
	Open             bool   `json:"open,omitempty"`
	Click            bool   `json:"click,omitempty"`
	Dropped          bool   `json:"dropped,omitempty"`
	FriendlyName     string `json:"friendly_name,omitempty"`
	OAuthClientID    string `json:"oauth_client_id,omitempty"`
	OAuthTokenURL    string `json:"oauth_token_url,omitempty"`
	PublicKey        string `json:"public_key,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/webhooks/get-an-event-webhook
func (c *Client) GetEventWebhook(ctx context.Context, id string) (*OutputGetEventWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetEventWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type EventWebhook struct {
	ID               string `json:"id,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	URL              string `json:"url,omitempty"`
	GroupResubscribe bool   `json:"group_resubscribe,omitempty"`
	Delivered        bool   `json:"delivered,omitempty"`
	GroupUnsubscribe bool   `json:"group_unsubscribe,omitempty"`
	SpamReport       bool   `json:"spam_report,omitempty"`
	Bounce           bool   `json:"bounce,omitempty"`
	Deferred         bool   `json:"deferred,omitempty"`
	Unsubscribe      bool   `json:"unsubscribe,omitempty"`
	Processed        bool   `json:"processed,omitempty"`
	Open             bool   `json:"open,omitempty"`
	Click            bool   `json:"click,omitempty"`
	Dropped          bool   `json:"dropped,omitempty"`
	FriendlyName     string `json:"friendly_name,omitempty"`
	OAuthClientID    string `json:"oauth_client_id,omitempty"`
	OAuthTokenURL    string `json:"oauth_token_url,omitempty"`
	PublicKey        string `json:"public_key,omitempty"`
}

type OutputGetEventWebhooks struct {
	MaxAllowed int             `json:"max_allowed,omitempty"`
	Webhooks   []*EventWebhook `json:"webhooks,omitempty"`
}

// https://docs.sendgrid.com/api-reference/webhooks/get-all-event-webhooks
func (c *Client) GetEventWebhooks(ctx context.Context) (*OutputGetEventWebhooks, error) {
	req, err := c.NewRequest("GET", "/user/webhooks/event/settings/all", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetEventWebhooks)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateEventWebhook struct {
	Enabled           bool   `json:"enabled"`
	URL               string `json:"url,omitempty"`
	GroupResubscribe  bool   `json:"group_resubscribe"`
	Delivered         bool   `json:"delivered"`
	GroupUnsubscribe  bool   `json:"group_unsubscribe"`
	SpamReport        bool   `json:"spam_report"`
	Bounce            bool   `json:"bounce"`
	Deferred          bool   `json:"deferred"`
	Unsubscribe       bool   `json:"unsubscribe"`
	Processed         bool   `json:"processed"`
	Open              bool   `json:"open"`
	Click             bool   `json:"click"`
	Dropped           bool   `json:"dropped"`
	FriendlyName      string `json:"friendly_name,omitempty"`
	OAuthClientID     string `json:"oauth_client_id,omitempty"`
	OAuthClientSecret string `json:"oauth_client_secret,omitempty"`
	OAuthTokenURL     string `json:"oauth_token_url,omitempty"`
}

type OutputCreateEventWebhook struct {
	ID                  string `json:"id,omitempty"`
	Enabled             bool   `json:"enabled,omitempty"`
	URL                 string `json:"url,omitempty"`
	AccountStatusChange bool   `json:"account_status_change,omitempty"`
	GroupResubscribe    bool   `json:"group_resubscribe,omitempty"`
	Delivered           bool   `json:"delivered,omitempty"`
	GroupUnsubscribe    bool   `json:"group_unsubscribe,omitempty"`
	SpamReport          bool   `json:"spam_report,omitempty"`
	Bounce              bool   `json:"bounce,omitempty"`
	Deferred            bool   `json:"deferred,omitempty"`
	Unsubscribe         bool   `json:"unsubscribe,omitempty"`
	Processed           bool   `json:"processed,omitempty"`
	Open                bool   `json:"open,omitempty"`
	Click               bool   `json:"click,omitempty"`
	Dropped             bool   `json:"dropped,omitempty"`
	FriendlyName        string `json:"friendly_name,omitempty"`
	CreatedDate         string `json:"created_date,omitempty"`
	UpdatedDate         string `json:"updated_date,omitempty"`
	OAuthClientID       string `json:"oauth_client_id,omitempty"`
	OAuthTokenURL       string `json:"oauth_token_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/webhooks/create-an-event-webhook
func (c *Client) CreateEventWebhook(ctx context.Context, input *InputCreateEventWebhook) (*OutputCreateEventWebhook, error) {
	req, err := c.NewRequest("POST", "/user/webhooks/event/settings", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateEventWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateEventWebhook struct {
	Enabled           bool   `json:"enabled"`
	URL               string `json:"url,omitempty"`
	GroupResubscribe  bool   `json:"group_resubscribe"`
	Delivered         bool   `json:"delivered"`
	GroupUnsubscribe  bool   `json:"group_unsubscribe"`
	SpamReport        bool   `json:"spam_report"`
	Bounce            bool   `json:"bounce"`
	Deferred          bool   `json:"deferred"`
	Unsubscribe       bool   `json:"unsubscribe"`
	Processed         bool   `json:"processed"`
	Open              bool   `json:"open"`
	Click             bool   `json:"click"`
	Dropped           bool   `json:"dropped"`
	FriendlyName      string `json:"friendly_name,omitempty"`
	OAuthClientID     string `json:"oauth_client_id,omitempty"`
	OAuthClientSecret string `json:"oauth_client_secret,omitempty"`
	OAuthTokenURL     string `json:"oauth_token_url,omitempty"`
}

type OutputUpdateEventWebhook struct {
	ID               string `json:"id,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	URL              string `json:"url,omitempty"`
	GroupResubscribe bool   `json:"group_resubscribe,omitempty"`
	Delivered        bool   `json:"delivered,omitempty"`
	GroupUnsubscribe bool   `json:"group_unsubscribe,omitempty"`
	SpamReport       bool   `json:"spam_report,omitempty"`
	Bounce           bool   `json:"bounce,omitempty"`
	Deferred         bool   `json:"deferred,omitempty"`
	Unsubscribe      bool   `json:"unsubscribe,omitempty"`
	Processed        bool   `json:"processed,omitempty"`
	Open             bool   `json:"open,omitempty"`
	Click            bool   `json:"click,omitempty"`
	Dropped          bool   `json:"dropped,omitempty"`
	FriendlyName     string `json:"friendly_name,omitempty"`
	CreatedDate      string `json:"created_date,omitempty"`
	UpdatedDate      string `json:"updated_date,omitempty"`
	OAuthClientID    string `json:"oauth_client_id,omitempty"`
	OAuthTokenURL    string `json:"oauth_token_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/webhooks/update-an-event-webhook
func (c *Client) UpdateEventWebhook(ctx context.Context, id string, input *InputUpdateEventWebhook) (*OutputUpdateEventWebhook, error) {
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)
	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateEventWebhook)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/webhooks/delete-an-event-webhook
func (c *Client) DeleteEventWebhook(ctx context.Context, id string) error {
	path := fmt.Sprintf("/user/webhooks/event/settings/%s", id)
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
