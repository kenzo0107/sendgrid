package sendgrid

import (
	"context"
	"fmt"
)

// LegacySegmentCondition represents a single condition used to build a Legacy Marketing Campaigns segment.
type LegacySegmentCondition struct {
	Field    string `json:"field,omitempty"`
	Value    string `json:"value,omitempty"`
	Operator string `json:"operator,omitempty"`
	AndOr    string `json:"and_or,omitempty"`
}

// LegacySegment represents a Legacy Marketing Campaigns contactdb segment.
type LegacySegment struct {
	ID             int64                     `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	ListID         int64                     `json:"list_id,omitempty"`
	Conditions     []*LegacySegmentCondition `json:"conditions,omitempty"`
	RecipientCount int64                     `json:"recipient_count,omitempty"`
}

type InputCreateLegacySegment struct {
	Name       string                    `json:"name,omitempty"`
	ListID     int64                     `json:"list_id,omitempty"`
	Conditions []*LegacySegmentCondition `json:"conditions,omitempty"`
}

type OutputCreateLegacySegment struct {
	ID             int64                     `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	ListID         int64                     `json:"list_id,omitempty"`
	Conditions     []*LegacySegmentCondition `json:"conditions,omitempty"`
	RecipientCount int64                     `json:"recipient_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/create-a-segment
func (c *Client) CreateLegacySegment(ctx context.Context, input *InputCreateLegacySegment) (*OutputCreateLegacySegment, error) {
	req, err := c.NewRequest("POST", "/contactdb/segments", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateLegacySegment)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacySegments struct {
	Segments []*LegacySegment `json:"segments,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/retrieve-all-segments
func (c *Client) GetLegacySegments(ctx context.Context) (*OutputGetLegacySegments, error) {
	req, err := c.NewRequest("GET", "/contactdb/segments", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacySegments)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacySegment struct {
	ID             int64                     `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	ListID         int64                     `json:"list_id,omitempty"`
	Conditions     []*LegacySegmentCondition `json:"conditions,omitempty"`
	RecipientCount int64                     `json:"recipient_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/retrieve-a-segment
func (c *Client) GetLegacySegment(ctx context.Context, id int64) (*OutputGetLegacySegment, error) {
	path := fmt.Sprintf("/contactdb/segments/%d", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacySegment)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateLegacySegment struct {
	Name       string                    `json:"name,omitempty"`
	ListID     int64                     `json:"list_id,omitempty"`
	Conditions []*LegacySegmentCondition `json:"conditions,omitempty"`
}

type OutputUpdateLegacySegment struct {
	ID             int64                     `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	ListID         int64                     `json:"list_id,omitempty"`
	Conditions     []*LegacySegmentCondition `json:"conditions,omitempty"`
	RecipientCount int64                     `json:"recipient_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/update-a-segment
func (c *Client) UpdateLegacySegment(ctx context.Context, id int64, input *InputUpdateLegacySegment) (*OutputUpdateLegacySegment, error) {
	path := fmt.Sprintf("/contactdb/segments/%d", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateLegacySegment)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/delete-a-segment
func (c *Client) DeleteLegacySegment(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/contactdb/segments/%d", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

// LegacySegmentRecipient represents a recipient returned by the Legacy Marketing Campaigns
// "Retrieve recipients on a segment" endpoint.
type LegacySegmentRecipient struct {
	ID           string                   `json:"id,omitempty"`
	Email        string                   `json:"email,omitempty"`
	FirstName    string                   `json:"first_name,omitempty"`
	LastName     string                   `json:"last_name,omitempty"`
	CreatedAt    int64                    `json:"created_at,omitempty"`
	UpdatedAt    int64                    `json:"updated_at,omitempty"`
	LastClicked  int64                    `json:"last_clicked,omitempty"`
	LastEmailed  int64                    `json:"last_emailed,omitempty"`
	LastOpened   int64                    `json:"last_opened,omitempty"`
	CustomFields []map[string]interface{} `json:"custom_fields,omitempty"`
}

type InputGetLegacySegmentRecipients struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"page_size,omitempty"`
}

type OutputGetLegacySegmentRecipients struct {
	Recipients     []*LegacySegmentRecipient `json:"recipients,omitempty"`
	RecipientCount int64                     `json:"recipient_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts-api-segments/retrieve-recipients-on-a-segment
func (c *Client) GetLegacySegmentRecipients(ctx context.Context, id int64, opts *InputGetLegacySegmentRecipients) (*OutputGetLegacySegmentRecipients, error) {
	path, err := c.AddOptions(fmt.Sprintf("/contactdb/segments/%d/recipients", id), opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacySegmentRecipients)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
