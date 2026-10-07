package sendgrid

import (
	"context"
	"fmt"
)

// SendIPPoolRef represents an IP Pool an IP address is assigned to.
// It is used in the newer /send_ips API (IP Address Management), which is
// distinct from the older /ips and /ips/pools APIs.
type SendIPPoolRef struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// SendIP represents an IP address on your account as returned by the
// newer /send_ips API (IP Address Management).
type SendIP struct {
	IP               string           `json:"ip,omitempty"`
	Pools            []*SendIPPoolRef `json:"pools,omitempty"`
	IsAutoWarmup     bool             `json:"is_auto_warmup,omitempty"`
	IsParentAssigned bool             `json:"is_parent_assigned,omitempty"`
	UpdatedAt        int64            `json:"updated_at,omitempty"`
	IsEnabled        bool             `json:"is_enabled,omitempty"`
	IsLeased         bool             `json:"is_leased,omitempty"`
	AddedAt          int64            `json:"added_at,omitempty"`
	Region           string           `json:"region,omitempty"`
}

type InputGetSendIPs struct {
	IP               string `url:"ip,omitempty"`
	Limit            int    `url:"limit,omitempty"`
	AfterKey         int    `url:"after_key,omitempty"`
	BeforeKey        string `url:"before_key,omitempty"`
	IsLeased         bool   `url:"is_leased,omitempty"`
	IsEnabled        bool   `url:"is_enabled,omitempty"`
	IsParentAssigned bool   `url:"is_parent_assigned,omitempty"`
	Pool             string `url:"pool,omitempty"`
	StartAddedAt     int    `url:"start_added_at,omitempty"`
	EndAddedAt       int    `url:"end_added_at,omitempty"`
	Region           string `url:"region,omitempty"`
	IncludeRegion    bool   `url:"include_region,omitempty"`
}

type OutputGetSendIPsMetadataNextParams struct {
	AfterKey         string `json:"after_key,omitempty"`
	BeforeKey        string `json:"before_key,omitempty"`
	IP               string `json:"ip,omitempty"`
	IsLeased         bool   `json:"is_leased,omitempty"`
	IsEnabled        bool   `json:"is_enabled,omitempty"`
	IsParentAssigned bool   `json:"is_parent_assigned,omitempty"`
	Pool             string `json:"pool,omitempty"`
	StartAddedAt     string `json:"start_added_at,omitempty"`
	EndAddedAt       string `json:"end_added_at,omitempty"`
	Limit            string `json:"limit,omitempty"`
	Region           string `json:"region,omitempty"`
	IncludeRegion    string `json:"include_region,omitempty"`
}

type OutputGetSendIPsMetadata struct {
	NextParams *OutputGetSendIPsMetadataNextParams `json:"next_params,omitempty"`
}

type OutputGetSendIPs struct {
	Result   []*SendIP                 `json:"result,omitempty"`
	Metadata *OutputGetSendIPsMetadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-a-list-of-all-ip-addresses-on-your-account
func (c *Client) GetSendIPs(ctx context.Context, opts *InputGetSendIPs) (*OutputGetSendIPs, error) {
	path, err := c.AddOptions("/send_ips/ips", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIPs)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputAddSendIP struct {
	IsAutoWarmup     bool     `json:"is_auto_warmup"`
	IsParentAssigned bool     `json:"is_parent_assigned"`
	Subusers         []string `json:"subusers,omitempty"`
	Region           string   `json:"region,omitempty"`
	IncludeRegion    bool     `json:"include_region,omitempty"`
}

type OutputAddSendIP struct {
	IP               string   `json:"ip,omitempty"`
	IsAutoWarmup     bool     `json:"is_auto_warmup,omitempty"`
	IsParentAssigned bool     `json:"is_parent_assigned,omitempty"`
	Subusers         []string `json:"subusers,omitempty"`
	Region           string   `json:"region,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/add-a-twilio-sendgrid-ip-address
func (c *Client) AddSendIP(ctx context.Context, input *InputAddSendIP) (*OutputAddSendIP, error) {
	req, err := c.NewRequest("POST", "/send_ips/ips", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddSendIP)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetSendIP struct {
	IP               string           `json:"ip,omitempty"`
	IsParentAssigned bool             `json:"is_parent_assigned,omitempty"`
	IsAutoWarmup     bool             `json:"is_auto_warmup,omitempty"`
	Pools            []*SendIPPoolRef `json:"pools,omitempty"`
	AddedAt          int64            `json:"added_at,omitempty"`
	UpdatedAt        int64            `json:"updated_at,omitempty"`
	IsEnabled        bool             `json:"is_enabled,omitempty"`
	IsLeased         bool             `json:"is_leased,omitempty"`
	Region           string           `json:"region,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-details-for-an-ip-address
func (c *Client) GetSendIP(ctx context.Context, ip string) (*OutputGetSendIP, error) {
	path := fmt.Sprintf("/send_ips/ips/%s", ip)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIP)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateSendIP struct {
	IsAutoWarmup     bool `json:"is_auto_warmup,omitempty"`
	IsParentAssigned bool `json:"is_parent_assigned,omitempty"`
	IsEnabled        bool `json:"is_enabled,omitempty"`
}

type OutputUpdateSendIP struct {
	IP               string `json:"ip,omitempty"`
	IsAutoWarmup     bool   `json:"is_auto_warmup,omitempty"`
	IsParentAssigned bool   `json:"is_parent_assigned,omitempty"`
	IsEnabled        bool   `json:"is_enabled,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/update-details-for-an-ip-address
func (c *Client) UpdateSendIP(ctx context.Context, ip string, input *InputUpdateSendIP) (*OutputUpdateSendIP, error) {
	path := fmt.Sprintf("/send_ips/ips/%s", ip)

	req, err := c.NewRequest("PATCH", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSendIP)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetSendIPSubusers struct {
	AfterKey int `url:"after_key,omitempty"`
	Limit    int `url:"limit,omitempty"`
}

type OutputGetSendIPSubusersMetadataNextParams struct {
	AfterKey string `json:"after_key,omitempty"`
	Limit    string `json:"limit,omitempty"`
}

type OutputGetSendIPSubusersMetadata struct {
	NextParams *OutputGetSendIPSubusersMetadataNextParams `json:"next_params,omitempty"`
}

type OutputGetSendIPSubusers struct {
	Result   []string                         `json:"result,omitempty"`
	Metadata *OutputGetSendIPSubusersMetadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-a-list-of-subusers-assigned-to-an-ip
func (c *Client) GetSendIPSubusers(ctx context.Context, ip string, opts *InputGetSendIPSubusers) (*OutputGetSendIPSubusers, error) {
	path, err := c.AddOptions(fmt.Sprintf("/send_ips/ips/%s/subusers", ip), opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIPSubusers)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputBatchAddSendIPSubusers struct {
	Subusers []string `json:"subusers"`
}

type OutputBatchAddSendIPSubusers struct {
	IP       string   `json:"ip,omitempty"`
	Subusers []string `json:"subusers,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/assign-a-batch-of-subusers-to-an-ip
func (c *Client) BatchAddSendIPSubusers(ctx context.Context, ip string, input *InputBatchAddSendIPSubusers) (*OutputBatchAddSendIPSubusers, error) {
	path := fmt.Sprintf("/send_ips/ips/%s/subusers:batchAdd", ip)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputBatchAddSendIPSubusers)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputBatchDeleteSendIPSubusers struct {
	Subusers []string `json:"subusers"`
}

// OutputBatchDeleteSendIPSubusers is intentionally empty: the API responds
// with 204 No Content on success.
type OutputBatchDeleteSendIPSubusers struct{}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/delete-a-batch-of-subusers-from-an-ip
func (c *Client) BatchDeleteSendIPSubusers(ctx context.Context, ip string, input *InputBatchDeleteSendIPSubusers) (*OutputBatchDeleteSendIPSubusers, error) {
	path := fmt.Sprintf("/send_ips/ips/%s/subusers:batchDelete", ip)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputBatchDeleteSendIPSubusers)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type SendIPPool struct {
	Name         string   `json:"name,omitempty"`
	ID           string   `json:"id,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	IPsPreview   []string `json:"ips_preview,omitempty"`
	TotalIPCount int      `json:"total_ip_count,omitempty"`
}

type InputGetSendIPPools struct {
	Limit         int    `url:"limit,omitempty"`
	AfterKey      int    `url:"after_key,omitempty"`
	IP            string `url:"ip,omitempty"`
	Region        string `url:"region,omitempty"`
	IncludeRegion bool   `url:"include_region,omitempty"`
}

type OutputGetSendIPPoolsMetadataNextParams struct {
	AfterKey      string `json:"after_key,omitempty"`
	IP            string `json:"ip,omitempty"`
	Limit         string `json:"limit,omitempty"`
	Region        string `json:"region,omitempty"`
	IncludeRegion string `json:"include_region,omitempty"`
}

type OutputGetSendIPPoolsMetadata struct {
	NextParams *OutputGetSendIPPoolsMetadataNextParams `json:"next_params,omitempty"`
}

type OutputGetSendIPPools struct {
	Result   []*SendIPPool                 `json:"result,omitempty"`
	Metadata *OutputGetSendIPPoolsMetadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-all-ip-pools-that-have-associated-ips
func (c *Client) GetSendIPPools(ctx context.Context, opts *InputGetSendIPPools) (*OutputGetSendIPPools, error) {
	path, err := c.AddOptions("/send_ips/pools", opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIPPools)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateSendIPPool struct {
	Name string   `json:"name"`
	IPs  []string `json:"ips,omitempty"`
}

type OutputCreateSendIPPool struct {
	Name string   `json:"name,omitempty"`
	ID   string   `json:"id,omitempty"`
	IPs  []string `json:"ips,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/create-an-ip-pool-with-a-name-and-ip-assignments
func (c *Client) CreateSendIPPool(ctx context.Context, input *InputCreateSendIPPool) (*OutputCreateSendIPPool, error) {
	req, err := c.NewRequest("POST", "/send_ips/pools", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateSendIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type SendIPPoolIPCountByRegion struct {
	Region string `json:"region,omitempty"`
	Count  int    `json:"count,omitempty"`
}

type OutputGetSendIPPool struct {
	Name            string                       `json:"name,omitempty"`
	ID              string                       `json:"id,omitempty"`
	IPsPreview      []string                     `json:"ips_preview,omitempty"`
	TotalIPCount    int                          `json:"total_ip_count,omitempty"`
	IPCountByRegion []*SendIPPoolIPCountByRegion `json:"ip_count_by_region,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-details-for-an-ip-pool
func (c *Client) GetSendIPPool(ctx context.Context, poolID string) (*OutputGetSendIPPool, error) {
	path := fmt.Sprintf("/send_ips/pools/%s", poolID)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateSendIPPool struct {
	Name string `json:"name"`
}

type OutputUpdateSendIPPool struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/update-an-ip-pool-name
func (c *Client) UpdateSendIPPool(ctx context.Context, poolID string, input *InputUpdateSendIPPool) (*OutputUpdateSendIPPool, error) {
	path := fmt.Sprintf("/send_ips/pools/%s", poolID)

	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateSendIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/delete-ip-pool
func (c *Client) DeleteSendIPPool(ctx context.Context, poolID string) error {
	path := fmt.Sprintf("/send_ips/pools/%s", poolID)

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type InputGetSendIPPoolIPs struct {
	Limit         int  `url:"limit,omitempty"`
	AfterKey      int  `url:"after_key,omitempty"`
	IncludeRegion bool `url:"include_region,omitempty"`
}

type SendIPPoolIP struct {
	IP     string           `json:"ip,omitempty"`
	Region string           `json:"region,omitempty"`
	Pools  []*SendIPPoolRef `json:"pools,omitempty"`
}

type OutputGetSendIPPoolIPsMetadataNextParams struct {
	AfterKey      string `json:"after_key,omitempty"`
	Limit         string `json:"limit,omitempty"`
	IncludeRegion string `json:"include_region,omitempty"`
}

type OutputGetSendIPPoolIPsMetadata struct {
	NextParams *OutputGetSendIPPoolIPsMetadataNextParams `json:"next_params,omitempty"`
}

type OutputGetSendIPPoolIPs struct {
	Result   []*SendIPPoolIP                 `json:"result,omitempty"`
	Metadata *OutputGetSendIPPoolIPsMetadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/get-ips-assigned-to-an-ip-pool
func (c *Client) GetSendIPPoolIPs(ctx context.Context, poolID string, opts *InputGetSendIPPoolIPs) (*OutputGetSendIPPoolIPs, error) {
	path, err := c.AddOptions(fmt.Sprintf("/send_ips/pools/%s/ips", poolID), opts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSendIPPoolIPs)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputBatchAddSendIPPoolIPs struct {
	IPs []string `json:"ips"`
}

type OutputBatchAddSendIPPoolIPs struct {
	Name string   `json:"name,omitempty"`
	ID   string   `json:"id,omitempty"`
	IPs  []string `json:"ips,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/add-a-batch-of-ips-to-an-ip-pool
func (c *Client) BatchAddSendIPPoolIPs(ctx context.Context, poolID string, input *InputBatchAddSendIPPoolIPs) (*OutputBatchAddSendIPPoolIPs, error) {
	path := fmt.Sprintf("/send_ips/pools/%s/ips:batchAdd", poolID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputBatchAddSendIPPoolIPs)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputBatchDeleteSendIPPoolIPs struct {
	IPs []string `json:"ips"`
}

// OutputBatchDeleteSendIPPoolIPs is intentionally empty: the API responds
// with 204 No Content on success.
type OutputBatchDeleteSendIPPoolIPs struct{}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-address-management/delete-a-batch-of-ips-from-an-ip-pool
func (c *Client) BatchDeleteSendIPPoolIPs(ctx context.Context, poolID string, input *InputBatchDeleteSendIPPoolIPs) (*OutputBatchDeleteSendIPPoolIPs, error) {
	path := fmt.Sprintf("/send_ips/pools/%s/ips:batchDelete", poolID)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputBatchDeleteSendIPPoolIPs)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}
