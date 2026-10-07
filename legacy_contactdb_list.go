package sendgrid

import (
	"context"
	"fmt"
)

type LegacyList struct {
	ID             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	RecipientCount int64  `json:"recipient_count,omitempty"`
}

type LegacyListRecipientCustomField struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type LegacyListRecipient struct {
	ID           string                            `json:"id,omitempty"`
	CreatedAt    int64                             `json:"created_at,omitempty"`
	CustomFields []*LegacyListRecipientCustomField `json:"custom_fields,omitempty"`
	Email        string                            `json:"email,omitempty"`
	FirstName    string                            `json:"first_name,omitempty"`
	LastName     string                            `json:"last_name,omitempty"`
	LastClicked  int64                             `json:"last_clicked,omitempty"`
	LastEmailed  int64                             `json:"last_emailed,omitempty"`
	LastOpened   int64                             `json:"last_opened,omitempty"`
	UpdatedAt    int64                             `json:"updated_at,omitempty"`
}

type InputCreateLegacyList struct {
	Name string `json:"name"`
}

type OutputCreateLegacyList LegacyList

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/create-a-list
func (c *Client) CreateLegacyList(ctx context.Context, input *InputCreateLegacyList) (*OutputCreateLegacyList, error) {
	req, err := c.NewRequest("POST", "/contactdb/lists", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateLegacyList)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacyLists struct {
	Lists []*LegacyList `json:"lists,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/retrieve-all-lists
func (c *Client) GetLegacyLists(ctx context.Context) (*OutputGetLegacyLists, error) {
	req, err := c.NewRequest("GET", "/contactdb/lists", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyLists)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputDeleteLegacyLists []int64

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/delete-multiple-lists
func (c *Client) DeleteLegacyLists(ctx context.Context, input *InputDeleteLegacyLists) error {
	req, err := c.NewRequest("DELETE", "/contactdb/lists", input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type OutputGetLegacyList LegacyList

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/retrieve-a-single-list
func (c *Client) GetLegacyList(ctx context.Context, id int64) (*OutputGetLegacyList, error) {
	path := fmt.Sprintf("/contactdb/lists/%d", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyList)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateLegacyList struct {
	Name string `json:"name"`
}

type OutputUpdateLegacyList LegacyList

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/update-a-list
func (c *Client) UpdateLegacyList(ctx context.Context, id int64, input *InputUpdateLegacyList) (*OutputUpdateLegacyList, error) {
	path := fmt.Sprintf("/contactdb/lists/%d", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateLegacyList)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputDeleteLegacyList struct {
	DeleteContacts bool `url:"delete_contacts,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/delete-a-list
func (c *Client) DeleteLegacyList(ctx context.Context, id int64, opts *InputDeleteLegacyList) error {
	path, err := c.AddOptions(fmt.Sprintf("/contactdb/lists/%d", id), opts)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type InputGetLegacyListRecipients struct {
	Page     int `url:"page,omitempty"`
	PageSize int `url:"page_size,omitempty"`
}

type OutputGetLegacyListRecipients struct {
	Recipients     []*LegacyListRecipient `json:"recipients,omitempty"`
	RecipientCount int64                  `json:"recipient_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/retrieve-all-recipients-on-a-list
func (c *Client) GetLegacyListRecipients(ctx context.Context, id int64, opts *InputGetLegacyListRecipients) (*OutputGetLegacyListRecipients, error) {
	path, err := c.AddOptions(fmt.Sprintf("/contactdb/lists/%d/recipients", id), opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyListRecipients)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputAddRecipientsToLegacyList []string

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/add-multiple-recipients-to-a-list
func (c *Client) AddRecipientsToLegacyList(ctx context.Context, id int64, input *InputAddRecipientsToLegacyList) error {
	path := fmt.Sprintf("/contactdb/lists/%d/recipients", id)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/add-a-single-recipient-to-a-list
func (c *Client) AddRecipientToLegacyList(ctx context.Context, listID int64, recipientID string) error {
	path := fmt.Sprintf("/contactdb/lists/%d/recipients/%s", listID, recipientID)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/lists/delete-a-single-recipient-from-a-single-list
func (c *Client) RemoveRecipientFromLegacyList(ctx context.Context, listID int64, recipientID string) error {
	path := fmt.Sprintf("/contactdb/lists/%d/recipients/%s", listID, recipientID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
