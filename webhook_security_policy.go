package sendgrid

import (
	"context"
	"fmt"
)

type WebhookSecurityPolicyOAuth struct {
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
	TokenURL     string   `json:"token_url,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
}

type WebhookSecurityPolicySignature struct {
	Enabled   bool   `json:"enabled,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

type WebhookSecurityPolicy struct {
	ID        string                          `json:"id,omitempty"`
	Name      string                          `json:"name,omitempty"`
	OAuth     *WebhookSecurityPolicyOAuth     `json:"oauth,omitempty"`
	Signature *WebhookSecurityPolicySignature `json:"signature,omitempty"`
}

type InputCreateWebhookSecurityPolicy struct {
	Name      string                          `json:"name,omitempty"`
	OAuth     *WebhookSecurityPolicyOAuth     `json:"oauth,omitempty"`
	Signature *WebhookSecurityPolicySignature `json:"signature,omitempty"`
}

type OutputCreateWebhookSecurityPolicy struct {
	Policy *WebhookSecurityPolicy `json:"policy,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/webhook-security-policies/create-a-new-webhook-security-policy
func (c *Client) CreateWebhookSecurityPolicy(ctx context.Context, input *InputCreateWebhookSecurityPolicy) (*OutputCreateWebhookSecurityPolicy, error) {
	req, err := c.NewRequest("POST", "/user/webhooks/security/policies", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateWebhookSecurityPolicy)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetWebhookSecurityPolicies struct {
	Policies []*WebhookSecurityPolicy `json:"policies,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/webhook-security-policies/retrieve-all-webhook-security-policies-for-your-account
func (c *Client) GetWebhookSecurityPolicies(ctx context.Context) (*OutputGetWebhookSecurityPolicies, error) {
	req, err := c.NewRequest("GET", "/user/webhooks/security/policies", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetWebhookSecurityPolicies)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetWebhookSecurityPolicy struct {
	Policy *WebhookSecurityPolicy `json:"policy,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/webhook-security-policies/retrieve-a-specific-webhook-security-policy
func (c *Client) GetWebhookSecurityPolicy(ctx context.Context, id string) (*OutputGetWebhookSecurityPolicy, error) {
	path := fmt.Sprintf("/user/webhooks/security/policies/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetWebhookSecurityPolicy)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateWebhookSecurityPolicy struct {
	Name      string                          `json:"name,omitempty"`
	OAuth     *WebhookSecurityPolicyOAuth     `json:"oauth,omitempty"`
	Signature *WebhookSecurityPolicySignature `json:"signature,omitempty"`
}

type OutputUpdateWebhookSecurityPolicy struct {
	Policy *WebhookSecurityPolicy `json:"policy,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/webhook-security-policies/update-an-existing-webhook-security-policy
func (c *Client) UpdateWebhookSecurityPolicy(ctx context.Context, id string, input *InputUpdateWebhookSecurityPolicy) (*OutputUpdateWebhookSecurityPolicy, error) {
	path := fmt.Sprintf("/user/webhooks/security/policies/%s", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateWebhookSecurityPolicy)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/webhook-security-policies/delete-a-specific-webhook-security-policy
func (c *Client) DeleteWebhookSecurityPolicy(ctx context.Context, id string) error {
	path := fmt.Sprintf("/user/webhooks/security/policies/%s", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
