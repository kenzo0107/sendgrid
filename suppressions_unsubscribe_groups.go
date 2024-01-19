package sendgrid

import (
	"context"
	"fmt"
	"strconv"
)

type SuppressionGroup struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	IsDefault       bool   `json:"is_default,omitempty"`
	Unsubscribes    int64  `json:"unsubscribes,omitempty"`
	LastEmailSentAt string `json:"last_email_sent_at,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/suppressions-unsubscribe-groups/get-information-on-a-single-suppression-group
func (c *Client) GetSuppressionGroup(ctx context.Context, id int64) (*SuppressionGroup, error) {
	path := fmt.Sprintf("/asm/groups/%s", strconv.FormatInt(id, 10))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(SuppressionGroup)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/suppressions-unsubscribe-groups/retrieve-all-suppression-groups-associated-with-the-user
func (c *Client) GetSuppressionGroups(ctx context.Context) ([]*SuppressionGroup, error) {
	req, err := c.NewRequest("GET", "/asm/groups", nil)
	if err != nil {
		return nil, err
	}

	r := []*SuppressionGroup{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateSuppressionGroup struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
}

type OutputCreateSuppressionGroup struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/suppressions-unsubscribe-groups/create-a-new-suppression-group
func (c *Client) CreateSuppressionGroup(ctx context.Context, input *InputCreateSuppressionGroup) (*OutputCreateSuppressionGroup, error) {
	req, err := c.NewRequest("POST", "/asm/groups", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSuppressionGroup)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputUpdateSuppressionGroup struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
}

type OutputUpdateSuppressionGroup struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	IsDefault       bool   `json:"is_default,omitempty"`
	LastEmailSentAt string `json:"last_email_sent_at,omitempty"`
	Unsubscribes    int64  `json:"unsubscribes,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/suppressions-unsubscribe-groups/update-a-suppression-group
func (c *Client) UpdateSuppressionGroup(ctx context.Context, id int64, input *InputUpdateSuppressionGroup) (*OutputUpdateSuppressionGroup, error) {
	path := fmt.Sprintf("/asm/groups/%s", strconv.FormatInt(id, 10))

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSuppressionGroup)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// see: https://docs.sendgrid.com/api-reference/suppressions-unsubscribe-groups/delete-a-suppression-group
func (c *Client) DeleteSuppressionGroup(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/asm/groups/%s", strconv.FormatInt(id, 10))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
