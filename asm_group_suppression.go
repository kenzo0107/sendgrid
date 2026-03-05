package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type InputAddSuppressionsToGroup struct {
	RecipientEmails []string `json:"recipient_emails,omitempty"`
}

type OutputAddSuppressionsToGroup struct {
	RecipientEmails []string `json:"recipient_emails,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/suppressions-suppressions/add-suppressions-to-a-suppression-group
func (c *Client) AddSuppressionsToGroup(ctx context.Context, groupID int64, input *InputAddSuppressionsToGroup) (*OutputAddSuppressionsToGroup, error) {
	path := fmt.Sprintf("/asm/groups/%s/suppressions", strconv.FormatInt(groupID, 10))

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddSuppressionsToGroup)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ASMSuppressionGroup struct {
	Description string `json:"description,omitempty"`
	ID          int64  `json:"id,omitempty"`
	IsDefault   bool   `json:"is_default,omitempty"`
	Name        string `json:"name,omitempty"`
	Suppressed  bool   `json:"suppressed,omitempty"`
}

type OutputGetSuppressionGroupsByEmail struct {
	Suppressions []*ASMSuppressionGroup `json:"suppressions,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppression-groups-for-an-email-address
func (c *Client) GetSuppressionGroupsByEmail(ctx context.Context, email string) (*OutputGetSuppressionGroupsByEmail, error) {
	path := fmt.Sprintf("/asm/suppressions/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSuppressionGroupsByEmail)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppressions-for-a-suppression-group
func (c *Client) GetSuppressionsForSuppressionGroup(ctx context.Context, groupID string) ([]string, error) {
	path := fmt.Sprintf("/asm/groups/%s/suppressions", groupID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := []string{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type InputSearchGroupSuppressions struct {
	RecipientEmails []string `json:"recipient_emails,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/search-for-suppressions-within-a-group
func (c *Client) SearchForSuppressionsWithinGroup(ctx context.Context, groupID string, input *InputSearchGroupSuppressions) ([]string, error) {
	path := fmt.Sprintf("/asm/groups/%s/suppressions/search", groupID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := []string{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type Suppression struct {
	Email     string `json:"email,omitempty"`
	GroupID   int64  `json:"group_id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppressions
func (c *Client) GetSuppressions(ctx context.Context) ([]*Suppression, error) {
	req, err := c.NewRequest("GET", "/asm/suppressions", nil)
	if err != nil {
		return nil, err
	}

	r := []*Suppression{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) DeleteSuppressionFromGroup(ctx context.Context, groupID, email string) error {
	path := fmt.Sprintf("/asm/groups/%s/suppressions/%s", groupID, email)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
