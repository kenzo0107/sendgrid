package sendgrid

import (
	"context"
	"fmt"
	"net/url"
)

// Bounce represents a bounced email
type Bounce struct {
	Created int64  `json:"created"`
	Email   string `json:"email"`
	Reason  string `json:"reason"`
	Status  string `json:"status"`
}

// Block represents a blocked email
type Block struct {
	Created int64  `json:"created"`
	Email   string `json:"email"`
	Reason  string `json:"reason"`
}

// SpamReport represents a spam report
type SpamReport struct {
	Created int64  `json:"created"`
	Email   string `json:"email"`
	IP      string `json:"ip"`
}

// InvalidEmail represents an invalid email
type InvalidEmail struct {
	Created int64  `json:"created"`
	Email   string `json:"email"`
	Reason  string `json:"reason"`
}

// OutputGetBounces represents the response for bounces list
type OutputGetBounces struct {
	Bounces []Bounce `json:"bounces,omitempty"`
}

// OutputGetBlocks represents the response for blocks list
type OutputGetBlocks struct {
	Blocks []Block `json:"blocks,omitempty"`
}

// OutputGetSpamReports represents the response for spam reports list
type OutputGetSpamReports struct {
	SpamReports []SpamReport `json:"spam_reports,omitempty"`
}

// OutputGetInvalidEmails represents the response for invalid emails list
type OutputGetInvalidEmails struct {
	InvalidEmails []InvalidEmail `json:"invalid_emails,omitempty"`
}

// InputDeleteSuppressions represents the request body for deleting suppressions
type InputDeleteSuppressions struct {
	Emails    []string `json:"emails,omitempty"`
	DeleteAll bool     `json:"delete_all,omitempty"`
}

// SuppressionListOptions represents query parameters for suppression list requests
type SuppressionListOptions struct {
	StartTime int64  `url:"start_time,omitempty"`
	EndTime   int64  `url:"end_time,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Offset    int    `url:"offset,omitempty"`
	Email     string `url:"email,omitempty"`
}

// GetBounces retrieves all bounces
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/retrieve-all-bounces
func (c *Client) GetBounces(ctx context.Context, opts *SuppressionListOptions) ([]Bounce, error) {
	path := "/suppression/bounces"

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var bounces []Bounce
	if err := c.Do(ctx, req, &bounces); err != nil {
		return nil, err
	}

	return bounces, nil
}

// GetBounce retrieves a specific bounce
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/retrieve-a-bounce
func (c *Client) GetBounce(ctx context.Context, email string) (*Bounce, error) {
	path := fmt.Sprintf("/suppression/bounces/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var bounces []Bounce
	if err := c.Do(ctx, req, &bounces); err != nil {
		return nil, err
	}

	if len(bounces) == 0 {
		return nil, fmt.Errorf("bounce not found for email: %s", email)
	}

	return &bounces[0], nil
}

// DeleteBounces deletes bounces
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/delete-bounces
func (c *Client) DeleteBounces(ctx context.Context, input *InputDeleteSuppressions) error {
	path := "/suppression/bounces"

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteBounce deletes a specific bounce
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/delete-a-bounce
func (c *Client) DeleteBounce(ctx context.Context, email string) error {
	path := fmt.Sprintf("/suppression/bounces/%s", url.QueryEscape(email))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// GetBlocks retrieves all blocks
// see: https://www.twilio.com/docs/sendgrid/api-reference/blocks/retrieve-all-blocks
func (c *Client) GetBlocks(ctx context.Context, opts *SuppressionListOptions) ([]Block, error) {
	path := "/suppression/blocks"

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var blocks []Block
	if err := c.Do(ctx, req, &blocks); err != nil {
		return nil, err
	}

	return blocks, nil
}

// GetBlock retrieves a specific block
// see: https://www.twilio.com/docs/sendgrid/api-reference/blocks/retrieve-a-specific-block
func (c *Client) GetBlock(ctx context.Context, email string) (*Block, error) {
	path := fmt.Sprintf("/suppression/blocks/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var blocks []Block
	if err := c.Do(ctx, req, &blocks); err != nil {
		return nil, err
	}

	if len(blocks) == 0 {
		return nil, fmt.Errorf("block not found for email: %s", email)
	}

	return &blocks[0], nil
}

// DeleteBlocks deletes blocks
// see: https://www.twilio.com/docs/sendgrid/api-reference/blocks/delete-blocks
func (c *Client) DeleteBlocks(ctx context.Context, input *InputDeleteSuppressions) error {
	path := "/suppression/blocks"

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteBlock deletes a specific block
// see: https://www.twilio.com/docs/sendgrid/api-reference/blocks/delete-a-specific-block
func (c *Client) DeleteBlock(ctx context.Context, email string) error {
	path := fmt.Sprintf("/suppression/blocks/%s", url.QueryEscape(email))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// GetSpamReports retrieves all spam reports
// see: https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/retrieve-all-spam-reports
func (c *Client) GetSpamReports(ctx context.Context, opts *SuppressionListOptions) ([]SpamReport, error) {
	path := "/suppression/spam_reports"

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var spamReports []SpamReport
	if err := c.Do(ctx, req, &spamReports); err != nil {
		return nil, err
	}

	return spamReports, nil
}

// GetSpamReport retrieves a specific spam report
// see: https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/retrieve-a-specific-spam-report
func (c *Client) GetSpamReport(ctx context.Context, email string) (*SpamReport, error) {
	path := fmt.Sprintf("/suppression/spam_reports/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var spamReports []SpamReport
	if err := c.Do(ctx, req, &spamReports); err != nil {
		return nil, err
	}

	if len(spamReports) == 0 {
		return nil, fmt.Errorf("spam report not found for email: %s", email)
	}

	return &spamReports[0], nil
}

// DeleteSpamReports deletes spam reports
// see: https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/delete-spam-reports
func (c *Client) DeleteSpamReports(ctx context.Context, input *InputDeleteSuppressions) error {
	path := "/suppression/spam_reports"

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteSpamReport deletes a specific spam report
// see: https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/delete-a-specific-spam-report
func (c *Client) DeleteSpamReport(ctx context.Context, email string) error {
	path := fmt.Sprintf("/suppression/spam_reports/%s", url.QueryEscape(email))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// GetInvalidEmails retrieves all invalid emails
// see: https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/retrieve-all-invalid-emails
func (c *Client) GetInvalidEmails(ctx context.Context, opts *SuppressionListOptions) ([]InvalidEmail, error) {
	path := "/suppression/invalid_emails"

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var invalidEmails []InvalidEmail
	if err := c.Do(ctx, req, &invalidEmails); err != nil {
		return nil, err
	}

	return invalidEmails, nil
}

// GetInvalidEmail retrieves a specific invalid email
// see: https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/retrieve-an-invalid-email
func (c *Client) GetInvalidEmail(ctx context.Context, email string) (*InvalidEmail, error) {
	path := fmt.Sprintf("/suppression/invalid_emails/%s", url.QueryEscape(email))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var invalidEmails []InvalidEmail
	if err := c.Do(ctx, req, &invalidEmails); err != nil {
		return nil, err
	}

	if len(invalidEmails) == 0 {
		return nil, fmt.Errorf("invalid email not found: %s", email)
	}

	return &invalidEmails[0], nil
}

// DeleteInvalidEmails deletes invalid emails
// see: https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/delete-invalid-emails
func (c *Client) DeleteInvalidEmails(ctx context.Context, input *InputDeleteSuppressions) error {
	path := "/suppression/invalid_emails"

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteInvalidEmail deletes a specific invalid email
// see: https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/delete-a-specific-invalid-email
func (c *Client) DeleteInvalidEmail(ctx context.Context, email string) error {
	path := fmt.Sprintf("/suppression/invalid_emails/%s", url.QueryEscape(email))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// BounceClassificationStat represents the bounce count for a single classification on a given date
type BounceClassificationStat struct {
	Classification string `json:"classification"`
	Count          int    `json:"count"`
}

// BounceClassificationsByDate represents bounce classification stats for a single date
type BounceClassificationsByDate struct {
	Date  string                     `json:"date"`
	Stats []BounceClassificationStat `json:"stats"`
}

// InputGetBounceClassifications represents query parameters for retrieving bounce classification totals
type InputGetBounceClassifications struct {
	StartDate string `url:"start_date,omitempty"`
	EndDate   string `url:"end_date,omitempty"`
}

// OutputGetBounceClassifications represents the response for bounce classification totals
type OutputGetBounceClassifications struct {
	Result []BounceClassificationsByDate `json:"result,omitempty"`
}

// BounceClassificationDomainStat represents the bounce count for a single domain on a given date
type BounceClassificationDomainStat struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

// BounceClassificationDomainStatsByDate represents bounce classification domain stats for a single date
type BounceClassificationDomainStatsByDate struct {
	Date  string                           `json:"date"`
	Stats []BounceClassificationDomainStat `json:"stats"`
}

// InputGetBounceClassificationsByDomain represents query parameters for retrieving bounce classification over time by domain stats
type InputGetBounceClassificationsByDomain struct {
	StartDate string `url:"start_date,omitempty"`
	EndDate   string `url:"end_date,omitempty"`
}

// OutputGetBounceClassificationsByDomain represents the response for bounce classification over time by domain stats
type OutputGetBounceClassificationsByDomain struct {
	Result []BounceClassificationDomainStatsByDate `json:"result,omitempty"`
}

// GetBounceClassifications retrieves bounce classification totals
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/retrieve-bounce-classification-totals
func (c *Client) GetBounceClassifications(ctx context.Context, opts *InputGetBounceClassifications) (*OutputGetBounceClassifications, error) {
	path := "/suppression/bounces/classifications"

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var output OutputGetBounceClassifications
	if err := c.Do(ctx, req, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

// GetBounceClassificationsByDomain retrieves bounce classification over time by domain stats
// see: https://www.twilio.com/docs/sendgrid/api-reference/bounces/retrieve-bounce-classification-over-time-by-domain-stats
func (c *Client) GetBounceClassificationsByDomain(ctx context.Context, classification string, opts *InputGetBounceClassificationsByDomain) (*OutputGetBounceClassificationsByDomain, error) {
	path := fmt.Sprintf("/suppression/bounces/classifications/%s", url.QueryEscape(classification))

	if opts != nil {
		var err error
		path, err = c.AddOptions(path, opts)
		if err != nil {
			return nil, err
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var output OutputGetBounceClassificationsByDomain
	if err := c.Do(ctx, req, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
