package sendgrid

import (
	"context"
	"fmt"
	"net/url"
)

// IPPool represents an IP pool
type IPPool struct {
	Name string `json:"name,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/retrieve-all-ip-pools
func (c *Client) GetIPPools(ctx context.Context) ([]*IPPool, error) {
	path := "/ips/pools"

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var pools []*IPPool
	if err := c.Do(ctx, req, &pools); err != nil {
		return nil, err
	}

	return pools, nil
}

type OutputGetIPPool struct {
	PoolName string   `json:"pool_name,omitempty"`
	IPs      []string `json:"ips,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/retrieve-all-the-ips-in-a-specified-pool
func (c *Client) GetIPPool(ctx context.Context, name string) (*OutputGetIPPool, error) {
	path := fmt.Sprintf("/ips/pools/%s", url.QueryEscape(name))

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputCreateIPPool struct {
	Name string `json:"name,omitempty"`
}

type OutputCreateIPPool struct {
	Name string `json:"name,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/create-an-ip-pool
func (c *Client) CreateIPPool(ctx context.Context, input *InputCreateIPPool) (*OutputCreateIPPool, error) {
	req, err := c.NewRequest("POST", "/ips/pools", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputCreateIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputUpdateIPPool struct {
	Name string `json:"name,omitempty"`
}

type OutputUpdateIPPool struct {
	Name string `json:"name,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/rename-an-ip-pool
func (c *Client) UpdateIPPool(ctx context.Context, name string, input *InputUpdateIPPool) (*OutputUpdateIPPool, error) {
	path := fmt.Sprintf("/ips/pools/%s", url.QueryEscape(name))

	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputUpdateIPPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/delete-an-ip-pool
func (c *Client) DeleteIPPool(ctx context.Context, name string) error {
	path := fmt.Sprintf("/ips/pools/%s", url.QueryEscape(name))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

type InputAddIPToPool struct {
	IP string `json:"ip"`
}

type OutputAddIPToPool struct {
	IP        string   `json:"ip,omitempty"`
	Pools     []string `json:"pools,omitempty"`
	Warmup    bool     `json:"warmup,omitempty"`
	StartDate int64    `json:"start_date,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/add-an-ip-address-to-a-pool
func (c *Client) AddIPToPool(ctx context.Context, poolName string, input *InputAddIPToPool) (*OutputAddIPToPool, error) {
	path := fmt.Sprintf("/ips/pools/%s/ips", url.QueryEscape(poolName))

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddIPToPool)
	if err := c.Do(ctx, req, r); err != nil {
		return nil, err
	}

	return r, nil
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/remove-an-ip-address-from-a-pool
func (c *Client) RemoveIPFromPool(ctx context.Context, poolName, ip string) error {
	path := fmt.Sprintf("/ips/pools/%s/ips/%s", url.QueryEscape(poolName), url.QueryEscape(ip))

	req, err := c.NewRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
