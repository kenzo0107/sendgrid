package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type OutputGetTemplate struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Generation string    `json:"generation,omitempty"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	Versions   []Version `json:"versions,omitempty"`
	Warning    Warning   `json:"warning,omitempty"`
}

type Version struct {
	ID                   string `json:"id,omitempty"`
	TemplateID           string `json:"template_id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Subject              string `json:"subject,omitempty"`
	UpdatedAt            string `json:"updated_at,omitempty"`
	GeneratePlainContent bool   `json:"generate_plain_content,omitempty"`
	HTMLContent          string `json:"html_content,omitempty"`
	PlainContent         string `json:"plain_content,omitempty"`
	Editor               string `json:"editor,omitempty"`
	ThumbnailURL         string `json:"thumbnail_url,omitempty"`
}

type Warning struct {
	Message string `json:"message,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/retrieve-a-single-transactional-template
func (c *Client) GetTemplate(ctx context.Context, id string) (*OutputGetTemplate, error) {
	path := fmt.Sprintf("/templates/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTemplate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetTemplates struct {
	Generations string
	PageSize    int
	PageToken   string
}

type OutputGetTemplates struct {
	Templates []Template `json:"result,omitempty"`
	Metadata  Metadata   `json:"_metadata,omitempty"`
}

type Template struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Generation string    `json:"generation,omitempty"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	Versions   []Version `json:"versions,omitempty"`
}

type Metadata struct {
	Prev  string `json:"prev,omitempty"`
	Self  string `json:"self,omitempty"`
	Next  string `json:"next,omitempty"`
	Count int    `json:"count,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/retrieve-paged-transactional-templates
func (c *Client) GetTemplates(ctx context.Context, input *InputGetTemplates) (*OutputGetTemplates, error) {
	u, err := url.Parse("/templates")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if input.Generations != "" {
		q.Set("generations", input.Generations)
	}
	if input.PageSize > 0 {
		q.Set("page_size", strconv.Itoa(input.PageSize))
	}
	if input.PageToken != "" {
		q.Set("page_token", input.PageToken)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTemplates)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateTemplate struct {
	Name       string `json:"name,omitempty"`
	Generation string `json:"generation,omitempty"`
}

type OutputCreateTemplate struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Generation string    `json:"generation,omitempty"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	Versions   []Version `json:"versions,omitempty"`
	Warning    Warning   `json:"warning,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/create-a-transactional-template
func (c *Client) CreateTemplate(ctx context.Context, input *InputCreateTemplate) (*OutputCreateTemplate, error) {
	req, err := c.NewRequest("POST", "/templates", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateTemplate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputDuplicateTemplate struct {
	Name string `json:"name,omitempty"`
}

type OutputDuplicateTemplate struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Generation string    `json:"generation,omitempty"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	Versions   []Version `json:"versions,omitempty"`
	Warning    Warning   `json:"warning,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/duplicate-a-transactional-template
func (c *Client) DuplicateTemplate(ctx context.Context, id string, input *InputDuplicateTemplate) (*OutputDuplicateTemplate, error) {
	path := fmt.Sprintf("/templates/%s", id)
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputDuplicateTemplate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateTemplate struct {
	Name string `json:"name,omitempty"`
}

type OutputUpdateTemplate struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Generation string    `json:"generation,omitempty"`
	UpdatedAt  string    `json:"updated_at,omitempty"`
	Versions   []Version `json:"versions,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/edit-a-transactional-template
func (c *Client) UpdateTemplate(ctx context.Context, id string, input *InputUpdateTemplate) (*OutputUpdateTemplate, error) {
	path := fmt.Sprintf("/templates/%s", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateTemplate)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/transactional-templates/delete-a-template
func (c *Client) DeleteTemplate(ctx context.Context, id string) error {
	path := fmt.Sprintf("/templates/%s", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
