package sendgrid

import (
	"context"
	"fmt"
)

type APIKey struct {
	ApiKeyId string `json:"api_key_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

type OutputGetAPIKeys struct {
	APIKeys []APIKey `json:"result,omitempty"`
}

func (c *Client) GetAPIKeys(ctx context.Context) (*OutputGetAPIKeys, error) {
	req, err := c.NewRequest("GET", "/api_keys", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAPIKeys)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetAPIKey struct {
	ApiKeyId string   `json:"api_key_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Scopes   []string `json:"scopes,omitempty"`
}

func (c *Client) GetAPIKey(ctx context.Context, apiKeyId string) (*OutputGetAPIKey, error) {
	path := fmt.Sprintf("/api_keys/%s", apiKeyId)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAPIKey)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateAPIKey struct {
	Name   string   `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

type OutputCreateAPIKey struct {
	ApiKey   string   `json:"api_key,omitempty"`
	ApiKeyId string   `json:"api_key_id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Scopes   []string `json:"scopes,omitempty"`
}

func (c *Client) CreateAPIKey(ctx context.Context, input *InputCreateAPIKey) (*OutputCreateAPIKey, error) {
	req, err := c.NewRequest("POST", "/api_keys", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateAPIKey)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateAPIKeyName struct {
	Name string `json:"name"`
}

type OutputUpdateAPIKeyName struct {
	ApiKeyId string `json:"api_key_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

func (c *Client) UpdateAPIKeyName(ctx context.Context, apiKeyId string, input *InputUpdateAPIKeyName) (*OutputUpdateAPIKeyName, error) {
	u := fmt.Sprintf("/api_keys/%s", apiKeyId)

	req, err := c.NewRequest("PATCH", u, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateAPIKeyName)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateAPIKeyNameAndScopes struct {
	Name   string   `json:"name"`
	Scopes []string `json:"scopes,omitempty"`
}

type OutputUpdateAPIKeyNameAndScopes struct {
	ApiKeyId string   `json:"api_key_id,omitempty"`
	Name     string   `json:"name"`
	Scopes   []string `json:"scopes,omitempty"`
}

func (c *Client) UpdateAPIKeyNameAndScopes(ctx context.Context, apiKeyId string, input *InputUpdateAPIKeyNameAndScopes) (*OutputUpdateAPIKeyNameAndScopes, error) {
	u := fmt.Sprintf("/api_keys/%s", apiKeyId)

	req, err := c.NewRequest("PUT", u, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateAPIKeyNameAndScopes)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteAPIKey(ctx context.Context, apiKeyId string) error {
	u := fmt.Sprintf("/api_keys/%s", apiKeyId)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
