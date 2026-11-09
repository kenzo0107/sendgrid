package sendgrid

import (
	"context"
	"fmt"
	"strings"
)

// Stat represents statistics data
type Stat struct {
	Date  string     `json:"date,omitempty"`
	Stats []StatItem `json:"stats,omitempty"`
}

// StatItem represents individual statistic metrics
type StatItem struct {
	Metrics StatMetrics `json:"metrics,omitempty"`
	Name    string      `json:"name,omitempty"`
	Type    string      `json:"type,omitempty"`
}

// StatMetrics represents metric data
type StatMetrics struct {
	Blocks           int `json:"blocks,omitempty"`
	BounceDrops      int `json:"bounce_drops,omitempty"`
	Bounces          int `json:"bounces,omitempty"`
	Clicks           int `json:"clicks,omitempty"`
	DeferredDrops    int `json:"deferred_drops,omitempty"`
	Delivered        int `json:"delivered,omitempty"`
	InvalidEmails    int `json:"invalid_emails,omitempty"`
	Opens            int `json:"opens,omitempty"`
	Processed        int `json:"processed,omitempty"`
	Requests         int `json:"requests,omitempty"`
	SpamReportDrops  int `json:"spam_report_drops,omitempty"`
	SpamReports      int `json:"spam_reports,omitempty"`
	UniqueClicks     int `json:"unique_clicks,omitempty"`
	UniqueOpens      int `json:"unique_opens,omitempty"`
	UnsubscribeDrops int `json:"unsubscribe_drops,omitempty"`
	Unsubscribes     int `json:"unsubscribes,omitempty"`
}

// GlobalStat represents global statistics
type GlobalStat struct {
	Date  string      `json:"date,omitempty"`
	Stats StatMetrics `json:"stats,omitempty"`
}

// CategoryStat represents category statistics
type CategoryStat struct {
	Date  string     `json:"date,omitempty"`
	Stats []StatItem `json:"stats,omitempty"`
}

// SubuserStat represents subuser statistics
type SubuserStat struct {
	Date  string     `json:"date,omitempty"`
	Stats []StatItem `json:"stats,omitempty"`
}

// StatsOptions represents query parameters for stats requests
type StatsOptions struct {
	StartDate   string   `url:"start_date,omitempty"`
	EndDate     string   `url:"end_date,omitempty"`
	Aggregation string   `url:"aggregated_by,omitempty"`
	Categories  []string `url:"-"`
	Subusers    []string `url:"-"`
	Tags        []string `url:"-"`
	Limit       int      `url:"limit,omitempty"`
	Offset      int      `url:"offset,omitempty"`
}

// GetGlobalStats retrieves global email statistics
// see: https://www.twilio.com/docs/sendgrid/api-reference/stats/retrieve-global-email-statistics
func (c *Client) GetGlobalStats(ctx context.Context, opts *StatsOptions) ([]GlobalStat, error) {
	path := "/stats"

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

	var stats []GlobalStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetCategoryStats retrieves category email statistics
// see: https://www.twilio.com/docs/sendgrid/api-reference/categories-statistics/retrieve-email-statistics-for-categories
func (c *Client) GetCategoryStats(ctx context.Context, categories []string, opts *StatsOptions) ([]CategoryStat, error) {
	path := fmt.Sprintf("/categories/stats?categories=%s", escapeQueryList(categories))

	if opts != nil {
		var err error
		tempPath, err := c.AddOptions("", opts)
		if err != nil {
			return nil, err
		}
		if tempPath != "" {
			if strings.Contains(path, "?") {
				path += "&" + strings.TrimPrefix(tempPath, "?")
			} else {
				path += tempPath
			}
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var stats []CategoryStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetCategorySums retrieves category sums
// see: https://www.twilio.com/docs/sendgrid/api-reference/categories-statistics/retrieve-sums-of-email-stats-for-each-category
func (c *Client) GetCategorySums(ctx context.Context, opts *StatsOptions) ([]CategoryStat, error) {
	path := "/categories/stats/sums"

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

	var stats []CategoryStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSubuserStats retrieves subuser email statistics
// see: https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-email-statistics-for-your-subusers
func (c *Client) GetSubuserStats(ctx context.Context, subusers []string, opts *StatsOptions) ([]SubuserStat, error) {
	path := fmt.Sprintf("/subusers/stats?subusers=%s", escapeQueryList(subusers))

	if opts != nil {
		var err error
		tempPath, err := c.AddOptions("", opts)
		if err != nil {
			return nil, err
		}
		if tempPath != "" {
			if strings.Contains(path, "?") {
				path += "&" + strings.TrimPrefix(tempPath, "?")
			} else {
				path += tempPath
			}
		}
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var stats []SubuserStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSubuserSums retrieves subuser sums
// see: https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-sums-of-email-stats-for-each-subuser
func (c *Client) GetSubuserSums(ctx context.Context, opts *StatsOptions) ([]SubuserStat, error) {
	path := "/subusers/stats/sums"

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

	var stats []SubuserStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetSubuserMonthlyStats retrieves monthly subuser statistics
// see: https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-monthly-stats-for-all-subusers
func (c *Client) GetSubuserMonthlyStats(ctx context.Context, opts *StatsOptions) ([]SubuserStat, error) {
	path := "/subusers/stats/monthly"

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

	var stats []SubuserStat
	if err := c.Do(ctx, req, &stats); err != nil {
		return nil, err
	}

	return stats, nil
}
