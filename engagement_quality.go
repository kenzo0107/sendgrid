package sendgrid

import (
	"context"
)

type EngagementQualityMetrics struct {
	EngagementRecency    float64 `json:"engagement_recency,omitempty"`
	OpenRate             float64 `json:"open_rate,omitempty"`
	BounceClassification float64 `json:"bounce_classification,omitempty"`
	BounceRate           float64 `json:"bounce_rate,omitempty"`
	SpamRate             float64 `json:"spam_rate,omitempty"`
}

type EngagementQualityScore struct {
	UserID   string                    `json:"user_id,omitempty"`
	Username string                    `json:"username,omitempty"`
	Date     string                    `json:"date,omitempty"`
	Score    float64                   `json:"score,omitempty"`
	Metrics  *EngagementQualityMetrics `json:"metrics,omitempty"`
}

type InputGetEngagementQualityScores struct {
	From string `url:"from"`
	To   string `url:"to"`
}

type OutputGetEngagementQualityScores struct {
	Result []*EngagementQualityScore `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/sendgrid-engagement-quality-api/get-engagement-quality-scores
func (c *Client) GetEngagementQualityScores(ctx context.Context, input *InputGetEngagementQualityScores) (*OutputGetEngagementQualityScores, error) {
	path, err := c.AddOptions("/engagementquality/scores", input)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetEngagementQualityScores)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type EngagementQualityMetadata struct {
	NextParams *EngagementQualityNextParams `json:"next_params,omitempty"`
}

type EngagementQualityNextParams struct {
	AfterKey string `json:"after_key,omitempty"`
}

type InputGetSubuserEngagementQualityScores struct {
	Limit    int    `url:"limit,omitempty"`
	Date     string `url:"date,omitempty"`
	AfterKey string `url:"after_key,omitempty"`
}

type OutputGetSubuserEngagementQualityScores struct {
	Result   []*EngagementQualityScore  `json:"result,omitempty"`
	Metadata *EngagementQualityMetadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/sendgrid-engagement-quality-api/get-subusers-engagement-quality-scores
func (c *Client) GetSubuserEngagementQualityScores(ctx context.Context, input *InputGetSubuserEngagementQualityScores) (*OutputGetSubuserEngagementQualityScores, error) {
	path, err := c.AddOptions("/engagementquality/subusers/scores", input)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSubuserEngagementQualityScores)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
