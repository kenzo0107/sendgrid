package sendgrid

import (
	"context"
	"fmt"
)

type Design struct {
	ID           string `json:"id,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	Name         string `json:"name,omitempty"`
	Editor       string `json:"editor,omitempty"`
}

type _Metadata struct {
	Prev  string `json:"prev,omitempty"`
	Self  string `json:"self,omitempty"`
	Next  string `json:"next,omitempty"`
	Count int64  `json:"count,omitempty"`
}

type OutputGetDesigns struct {
	Result   []*Design `json:"result,omitempty"`
	Metadata _Metadata `json:"_metadata,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/designs-api/list-designs
func (c *Client) GetDesigns(ctx context.Context) (*OutputGetDesigns, error) {
	req, err := c.NewRequest("GET", "/designs", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetDesigns)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetDesign struct {
	ID                   string   `json:"id,omitempty"`
	UpdatedAt            string   `json:"updated_at,omitempty"`
	CreatedAt            string   `json:"created_at,omitempty"`
	ThumbnailURL         string   `json:"thumbnail_url,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Editor               string   `json:"editor,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	GeneratePlainContent bool     `json:"generate_plain_content,omitempty"`
	Subject              string   `json:"subject,omitempty"`
	Categories           []string `json:"categories,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/designs-api/get-design
func (c *Client) GetDesign(ctx context.Context, id string) (*OutputGetDesign, error) {
	path := fmt.Sprintf("/designs/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetDesign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateDesign struct {
	Name                 string   `json:"name,omitempty"`
	Editor               string   `json:"editor,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	GeneratePlainContent bool     `json:"generate_plain_content"`
	Subject              string   `json:"subject,omitempty"`
	Categories           []string `json:"categories,omitempty"`
}

type OutputCreateDesign struct {
	ID                   string   `json:"id,omitempty"`
	UpdatedAt            string   `json:"updated_at,omitempty"`
	CreatedAt            string   `json:"created_at,omitempty"`
	ThumbnailURL         string   `json:"thumbnail_url,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Editor               string   `json:"editor,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	GeneratePlainContent bool     `json:"generate_plain_content,omitempty"`
	Subject              string   `json:"subject,omitempty"`
	Categories           []string `json:"categories,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/designs-api/create-design
func (c *Client) CreateDesign(ctx context.Context, input *InputCreateDesign) (*OutputCreateDesign, error) {
	req, err := c.NewRequest("POST", "/designs", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateDesign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateDesign struct {
	Name                 string   `json:"name,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	GeneratePlainContent bool     `json:"generate_plain_content"`
	Subject              string   `json:"subject,omitempty"`
	Categories           []string `json:"categories,omitempty"`
}

type OutputUpdateDesign struct {
	ID                   string   `json:"id,omitempty"`
	UpdatedAt            string   `json:"updated_at,omitempty"`
	CreatedAt            string   `json:"created_at,omitempty"`
	ThumbnailURL         string   `json:"thumbnail_url,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Editor               string   `json:"editor,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	GeneratePlainContent bool     `json:"generate_plain_content,omitempty"`
	Subject              string   `json:"subject,omitempty"`
	Categories           []string `json:"categories,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/designs-api/update-design
func (c *Client) UpdateDesign(ctx context.Context, id string, input *InputUpdateDesign) (*OutputUpdateDesign, error) {
	path := fmt.Sprintf("/designs/%s", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateDesign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/designs-api/delete-design
func (c *Client) DeleteDesign(ctx context.Context, id string) error {
	path := fmt.Sprintf("/designs/%s", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
