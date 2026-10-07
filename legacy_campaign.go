package sendgrid

import (
	"context"
	"fmt"
	"net/url"
)

// LegacyCampaign represents a Legacy Marketing Campaigns campaign resource.
type LegacyCampaign struct {
	ID                   int64    `json:"id,omitempty"`
	Title                string   `json:"title,omitempty"`
	Subject              string   `json:"subject,omitempty"`
	SenderID             int64    `json:"sender_id,omitempty"`
	ListIDs              []int64  `json:"list_ids,omitempty"`
	SegmentIDs           []int64  `json:"segment_ids,omitempty"`
	Categories           []string `json:"categories,omitempty"`
	SuppressionGroupID   int64    `json:"suppression_group_id,omitempty"`
	CustomUnsubscribeURL string   `json:"custom_unsubscribe_url,omitempty"`
	IPPool               string   `json:"ip_pool,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	Editor               string   `json:"editor,omitempty"`
	Status               string   `json:"status,omitempty"`
}

type InputCreateLegacyCampaign struct {
	Title                string   `json:"title"`
	Subject              string   `json:"subject,omitempty"`
	SenderID             int64    `json:"sender_id,omitempty"`
	ListIDs              []int64  `json:"list_ids,omitempty"`
	SegmentIDs           []int64  `json:"segment_ids,omitempty"`
	Categories           []string `json:"categories,omitempty"`
	SuppressionGroupID   int64    `json:"suppression_group_id,omitempty"`
	CustomUnsubscribeURL string   `json:"custom_unsubscribe_url,omitempty"`
	IPPool               string   `json:"ip_pool,omitempty"`
	HTMLContent          string   `json:"html_content,omitempty"`
	PlainContent         string   `json:"plain_content,omitempty"`
	Editor               string   `json:"editor,omitempty"`
}

type OutputCreateLegacyCampaign LegacyCampaign

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/create-a-campaign
func (c *Client) CreateLegacyCampaign(ctx context.Context, input *InputCreateLegacyCampaign) (*OutputCreateLegacyCampaign, error) {
	req, err := c.NewRequest("POST", "/campaigns", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateLegacyCampaign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetLegacyCampaigns struct {
	Limit  int64 `json:"limit,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

type OutputGetLegacyCampaigns struct {
	Result []*LegacyCampaign `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/retrieve-all-campaigns
func (c *Client) GetLegacyCampaigns(ctx context.Context, opts *InputGetLegacyCampaigns) (*OutputGetLegacyCampaigns, error) {
	u, _ := url.Parse("/campaigns")

	if opts != nil {
		q := u.Query()
		if opts.Limit > 0 {
			q.Set("limit", fmt.Sprintf("%d", opts.Limit))
		}
		if opts.Offset > 0 {
			q.Set("offset", fmt.Sprintf("%d", opts.Offset))
		}
		u.RawQuery = q.Encode()
	}

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyCampaigns)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacyCampaign LegacyCampaign

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/retrieve-a-single-campaign
func (c *Client) GetLegacyCampaign(ctx context.Context, id int64) (*OutputGetLegacyCampaign, error) {
	path := fmt.Sprintf("/campaigns/%d", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyCampaign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/delete-a-campaign
func (c *Client) DeleteLegacyCampaign(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/campaigns/%d", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type InputUpdateLegacyCampaign struct {
	Title        string   `json:"title,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	Categories   []string `json:"categories,omitempty"`
	HTMLContent  string   `json:"html_content,omitempty"`
	PlainContent string   `json:"plain_content,omitempty"`
}

type OutputUpdateLegacyCampaign LegacyCampaign

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/update-a-campaign
func (c *Client) UpdateLegacyCampaign(ctx context.Context, id int64, input *InputUpdateLegacyCampaign) (*OutputUpdateLegacyCampaign, error) {
	path := fmt.Sprintf("/campaigns/%d", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateLegacyCampaign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputScheduleLegacyCampaign struct {
	SendAt int64 `json:"send_at"`
}

type OutputScheduleLegacyCampaign struct {
	ID     int64  `json:"id,omitempty"`
	SendAt int64  `json:"send_at,omitempty"`
	Status string `json:"status,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/schedule-a-campaign
func (c *Client) ScheduleLegacyCampaign(ctx context.Context, id int64, input *InputScheduleLegacyCampaign) (*OutputScheduleLegacyCampaign, error) {
	path := fmt.Sprintf("/campaigns/%d/schedules", id)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputScheduleLegacyCampaign)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateLegacyCampaignSchedule struct {
	SendAt int64 `json:"send_at"`
}

type OutputUpdateLegacyCampaignSchedule struct {
	ID     int64  `json:"id,omitempty"`
	SendAt int64  `json:"send_at,omitempty"`
	Status string `json:"status,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/update-a-scheduled-campaign
func (c *Client) UpdateLegacyCampaignSchedule(ctx context.Context, id int64, input *InputUpdateLegacyCampaignSchedule) (*OutputUpdateLegacyCampaignSchedule, error) {
	path := fmt.Sprintf("/campaigns/%d/schedules", id)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateLegacyCampaignSchedule)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetLegacyCampaignSchedule struct {
	SendAt int64 `json:"send_at,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/view-scheduled-time-of-a-campaign
func (c *Client) GetLegacyCampaignSchedule(ctx context.Context, id int64) (*OutputGetLegacyCampaignSchedule, error) {
	path := fmt.Sprintf("/campaigns/%d/schedules", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetLegacyCampaignSchedule)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/unschedule-a-scheduled-campaign
func (c *Client) DeleteLegacyCampaignSchedule(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/campaigns/%d/schedules", id)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type OutputSendLegacyCampaignNow struct {
	ID     int64  `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/send-a-campaign
func (c *Client) SendLegacyCampaignNow(ctx context.Context, id int64) (*OutputSendLegacyCampaignNow, error) {
	path := fmt.Sprintf("/campaigns/%d/schedules/now", id)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputSendLegacyCampaignNow)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputSendLegacyCampaignTest struct {
	To string `json:"to"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/campaigns-api/send-a-test-campaign
func (c *Client) SendLegacyCampaignTest(ctx context.Context, id int64, input *InputSendLegacyCampaignTest) error {
	path := fmt.Sprintf("/campaigns/%d/schedules/test", id)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
