package sendgrid

import (
	"context"
	"fmt"
)

type OutputGetTemplateVersion struct {
	ID                   string  `json:"id,omitempty"`
	TemplateID           string  `json:"template_id,omitempty"`
	Active               int     `json:"active,omitempty"`
	Name                 string  `json:"name,omitempty"`
	HTMLContent          string  `json:"html_content,omitempty"`
	PlainContent         string  `json:"plain_content,omitempty"`
	GeneratePlainContent bool    `json:"generate_plain_content,omitempty"`
	Subject              string  `json:"subject,omitempty"`
	Editor               string  `json:"editor,omitempty"`
	TestData             string  `json:"test_data,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	Warnings             Warning `json:"warnings,omitempty"`
	ThumbnailURL         string  `json:"thumbnail_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates-versions/retrieve-a-specific-transactional-template-version
func (c *Client) GetTemplateVersion(ctx context.Context, templateID, versionID string) (*OutputGetTemplateVersion, error) {
	path := fmt.Sprintf("/templates/%s/versions/%s", templateID, versionID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTemplateVersion)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateTemplateVersion struct {
	Active               int    `json:"active,omitempty"`
	Name                 string `json:"name,omitempty"`
	HTMLContent          string `json:"html_content,omitempty"`
	PlainContent         string `json:"plain_content,omitempty"`
	GeneratePlainContent bool   `json:"generate_plain_content,omitempty"`
	Subject              string `json:"subject,omitempty"`
	Editor               string `json:"editor,omitempty"`
	TestData             string `json:"test_data,omitempty"`
}

type OutputCreateTemplateVersion struct {
	ID                   string    `json:"id,omitempty"`
	TemplateID           string    `json:"template_id,omitempty"`
	Active               int       `json:"active,omitempty"`
	Name                 string    `json:"name,omitempty"`
	HTMLContent          string    `json:"html_content,omitempty"`
	PlainContent         string    `json:"plain_content,omitempty"`
	GeneratePlainContent bool      `json:"generate_plain_content,omitempty"`
	Subject              string    `json:"subject,omitempty"`
	Editor               string    `json:"editor,omitempty"`
	TestData             string    `json:"test_data,omitempty"`
	UpdatedAt            string    `json:"updated_at,omitempty"`
	Warnings             []Warning `json:"warnings,omitempty"`
	ThumbnailURL         string    `json:"thumbnail_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates-versions/create-a-new-transactional-template-version
func (c *Client) CreateTemplateVersion(ctx context.Context, templateID string, input *InputCreateTemplateVersion) (*OutputCreateTemplateVersion, error) {
	path := fmt.Sprintf("/templates/%s/versions", templateID)
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateTemplateVersion)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateTemplateVersion struct {
	Active               int    `json:"active,omitempty"`
	Name                 string `json:"name,omitempty"`
	HTMLContent          string `json:"html_content,omitempty"`
	PlainContent         string `json:"plain_content,omitempty"`
	GeneratePlainContent bool   `json:"generate_plain_content,omitempty"`
	Subject              string `json:"subject,omitempty"`
	Editor               string `json:"editor,omitempty"`
	TestData             string `json:"test_data,omitempty"`
}

type OutputUpdateTemplateVersion struct {
	ID                   string    `json:"id,omitempty"`
	TemplateID           string    `json:"template_id,omitempty"`
	Active               int       `json:"active,omitempty"`
	Name                 string    `json:"name,omitempty"`
	HTMLContent          string    `json:"html_content,omitempty"`
	PlainContent         string    `json:"plain_content,omitempty"`
	GeneratePlainContent bool      `json:"generate_plain_content,omitempty"`
	Subject              string    `json:"subject,omitempty"`
	Editor               string    `json:"editor,omitempty"`
	TestData             string    `json:"test_data,omitempty"`
	UpdatedAt            string    `json:"updated_at,omitempty"`
	Warnings             []Warning `json:"warnings,omitempty"`
	ThumbnailURL         string    `json:"thumbnail_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates-versions/edit-a-transactional-template-version
func (c *Client) UpdateTemplateVersion(ctx context.Context, templateID, versionID string, input *InputUpdateTemplateVersion) (*OutputUpdateTemplateVersion, error) {
	path := fmt.Sprintf("/templates/%s/versions/%s", templateID, versionID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateTemplateVersion)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputActivateTemplateVersion struct {
	ID                   string    `json:"id,omitempty"`
	TemplateID           string    `json:"template_id,omitempty"`
	Active               int       `json:"active,omitempty"`
	Name                 string    `json:"name,omitempty"`
	HTMLContent          string    `json:"html_content,omitempty"`
	PlainContent         string    `json:"plain_content,omitempty"`
	GeneratePlainContent bool      `json:"generate_plain_content,omitempty"`
	Subject              string    `json:"subject,omitempty"`
	Editor               string    `json:"editor,omitempty"`
	TestData             string    `json:"test_data,omitempty"`
	UpdatedAt            string    `json:"updated_at,omitempty"`
	Warnings             []Warning `json:"warnings,omitempty"`
	ThumbnailURL         string    `json:"thumbnail_url,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates-versions/activate-a-transactional-template-version
func (c *Client) ActivateTemplateVersion(ctx context.Context, templateID, versionID string) (*OutputActivateTemplateVersion, error) {
	path := fmt.Sprintf("/templates/%s/versions/%s/activate", templateID, versionID)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputActivateTemplateVersion)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates-versions/delete-a-transactional-template-version
func (c *Client) DeleteTemplateVersion(ctx context.Context, templateID, versionID string) error {
	path := fmt.Sprintf("/templates/%s/versions/%s", templateID, versionID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
