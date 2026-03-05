package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type InputAddRecipientAddressesToGlobalSuppressions struct {
	RecipientEmails []string `json:"recipient_emails,omitempty"`
}

type OutputAddRecipientAddressesToGlobalSuppressions struct {
	RecipientEmails []string `json:"recipient_emails,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/add-recipient-addresses-to-the-global-suppression-group
func (c *Client) AddRecipientAddressesToGlobalSuppressions(ctx context.Context, input *InputAddRecipientAddressesToGlobalSuppressions) (*OutputAddRecipientAddressesToGlobalSuppressions, error) {
	req, err := c.NewRequest("POST", "/asm/suppressions/global", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddRecipientAddressesToGlobalSuppressions)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetGlobalSuppression struct {
	RecipientEmail string `json:"recipient_email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/retrieve-a-global-suppression
func (c *Client) GetGlobalSuppression(ctx context.Context, email string) (*OutputGetGlobalSuppression, error) {
	path := fmt.Sprintf("/asm/suppressions/global/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetGlobalSuppression)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetGlobalSuppressions struct {
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GlobalUnsubscribe struct {
	Created int64  `json:"created,omitempty"`
	Email   string `json:"email,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/retrieve-all-global-suppressions
func (c *Client) GetGlobalSuppressions(ctx context.Context, input *InputGetGlobalSuppressions) ([]*GlobalUnsubscribe, error) {
	u, _ := url.Parse("/suppression/unsubscribes")

	q := u.Query()
	if input.StartTime > 0 {
		q.Set("start_time", strconv.FormatInt(input.StartTime, 10))
	}
	if input.EndTime > 0 {
		q.Set("end_time", strconv.FormatInt(input.EndTime, 10))
	}
	if input.Limit > 0 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}
	if input.Offset > 0 {
		q.Set("offset", strconv.Itoa(input.Offset))
	}
	if input.Email != "" {
		q.Set("email", input.Email)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := []*GlobalUnsubscribe{}
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/delete-a-global-suppression
func (c *Client) DeleteGlobalSuppression(ctx context.Context, email string) error {
	path := fmt.Sprintf("/asm/suppressions/global/%s", url.QueryEscape(email))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}
