package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type VerifiedSender struct {
	ID          int64  `json:"id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	FromEmail   string `json:"from_email,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	Locked      bool   `json:"locked,omitempty"`
}

type InputGetVerifiedSenders struct {
	Limit      int
	LastSeenID int
	ID         int64
}

type OutputGetVerifiedSenders struct {
	VerifiedSenders []*VerifiedSender `json:"results,omitempty"`
}

func (c *Client) GetVerifiedSenders(ctx context.Context, input *InputGetVerifiedSenders) ([]*VerifiedSender, error) {
	u, err := url.Parse("/verified_senders")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if input.Limit > 0 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}
	if input.LastSeenID > 0 {
		q.Set("lastSeenID", strconv.Itoa(input.LastSeenID))
	}
	if input.ID > 0 {
		q.Set("id", strconv.FormatInt(input.ID, 10))
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetVerifiedSenders)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.VerifiedSenders, nil
}

type InputCreateVerifiedSenderRequest struct {
	Nickname    string `json:"nickname,omitempty"`
	FromEmail   string `json:"from_email,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
}

type OutputCreateVerifiedSenderRequest struct {
	ID          int64  `json:"id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	FromEmail   string `json:"from_email,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	Locked      bool   `json:"locked,omitempty"`
}

func (c *Client) CreateVerifiedSenderRequest(ctx context.Context, input *InputCreateVerifiedSenderRequest) (*OutputCreateVerifiedSenderRequest, error) {
	req, err := c.NewRequest("POST", "/verified_senders", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateVerifiedSenderRequest)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) ResendVerifiedSenderRequest(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/verified_senders/resend/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("POST", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) VerifySenderRequest(ctx context.Context, token string) error {
	path := fmt.Sprintf("/verified_senders/verify/%s", token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type InputUpdateVerifiedSender struct {
	Nickname    string `json:"nickname,omitempty"`
	FromEmail   string `json:"from_email,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
}

type OutputUpdateVerifiedSender struct {
	ID          int64  `json:"id,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
	FromEmail   string `json:"from_email,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
	ReplyToName string `json:"reply_to_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Address2    string `json:"address2,omitempty"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	Verified    bool   `json:"verified,omitempty"`
	Locked      bool   `json:"locked,omitempty"`
}

func (c *Client) UpdateVerifiedSender(ctx context.Context, id int64, input *InputUpdateVerifiedSender) (*OutputUpdateVerifiedSender, error) {
	path := fmt.Sprintf("/verified_senders/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("PATCH", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateVerifiedSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) DeleteVerifiedSender(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/verified_senders/%s", strconv.FormatInt(id, 10))
	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}
	return nil
}

type OutputCompletedStepsVerifiedSender struct {
	CompletedStepsVerifiedSender *CompletedStepsVerifiedSender `json:"results,omitempty"`
}

type CompletedStepsVerifiedSender struct {
	SenderVerified bool `json:"sender_verified,omitempty"`
	DomainVerified bool `json:"domain_verified,omitempty"`
}

func (c *Client) CompletedStepsVerifiedSender(ctx context.Context) (*CompletedStepsVerifiedSender, error) {
	req, err := c.NewRequest("GET", "/verified_senders/steps_completed", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputCompletedStepsVerifiedSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.CompletedStepsVerifiedSender, nil
}

// see: https://docs.sendgrid.com/api-reference/sender-verification/domain-warn-list
// This endpoint returns a list of domains known to implement DMARC and categorizes them by failure type — hard failure or soft failure.
// Domains listed as hard failures will not deliver mail when used as a Sender Identity due to the domain's DMARC policy settings.
func (c *Client) GetSenderVerificationDomainWarnList(ctx context.Context) (*CompletedStepsVerifiedSender, error) {
	req, err := c.NewRequest("GET", "/verified_senders/domains", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputCompletedStepsVerifiedSender)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r.CompletedStepsVerifiedSender, nil
}
