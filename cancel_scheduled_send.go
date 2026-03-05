package sendgrid

import (
	"context"
	"fmt"
)

type OutputCreateBatchID struct {
	BatchID string `json:"batch_id,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/cancel-scheduled-sends/create-a-batch-id
func (c *Client) CreateBatchID(ctx context.Context) (*OutputCreateBatchID, error) {
	req, err := c.NewRequest("POST", "/mail/batch", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateBatchID)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputValidateBatchID struct {
	BatchID string `json:"batch_id,omitempty"`
}

func (c *Client) ValidateBatchID(ctx context.Context, batchID string) (*OutputValidateBatchID, error) {
	path := fmt.Sprintf("/mail/batch/%s", batchID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputValidateBatchID)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetScheduledSend struct {
	BatchID string `json:"batch_id,omitempty"`
	Status  string `json:"status,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/retrieve-scheduled-send
func (c *Client) GetScheduledSend(ctx context.Context, batchID string) (*OutputGetScheduledSend, error) {
	path := fmt.Sprintf("/user/scheduled_sends/%s", batchID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetScheduledSend)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ScheduledSend struct {
	BatchID string `json:"batch_id,omitempty"`
	Status  string `json:"status,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/retrieve-all-scheduled-sends
func (c *Client) GetScheduledSends(ctx context.Context) ([]*ScheduledSend, error) {
	req, err := c.NewRequest("GET", "/user/scheduled_sends", nil)
	if err != nil {
		return nil, err
	}

	r := []*ScheduledSend{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateScheduledSend struct {
	Status string `json:"status,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/cancel-scheduled-sends/update-a-scheduled-send
func (c *Client) UpdateScheduledSend(ctx context.Context, batchID string, input *InputUpdateScheduledSend) error {
	path := fmt.Sprintf("/user/scheduled_sends/%s", batchID)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type OutputCancelOrPauseScheduledSend struct {
	BatchID string `json:"batch_id,omitempty"`
	Status  string `json:"status,omitempty"`
}

// see: https://docs.sendgrid.com/api-reference/cancel-scheduled-sends/delete-a-cancellation-or-pause-of-a-scheduled-send
func (c *Client) CancelOrPauseScheduledSend(ctx context.Context, batchID string) (*OutputCancelOrPauseScheduledSend, error) {
	path := fmt.Sprintf("/user/scheduled_sends/%s", batchID)

	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputCancelOrPauseScheduledSend)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/cancel-or-pause-a-scheduled-send
func (c *Client) DeleteScheduledSend(ctx context.Context, batchID string) error {
	path := fmt.Sprintf("/user/scheduled_sends/%s", batchID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
