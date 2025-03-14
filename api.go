package cagc

import (
	"context"
	"net/http"
)

// GetVersion gets the version of the Coolify API
func (c *Client) GetVersion(ctx context.Context) (string, error) {
	var version string
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/version", nil, &version)
	return version, err
}

// EnableAPI enables the Coolify API (requires root permissions)
func (c *Client) EnableAPI(ctx context.Context) (*MessageResponse, error) {
	var response MessageResponse
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/enable", nil, &response)
	return &response, err
}

// DisableAPI disables the Coolify API (requires root permissions)
func (c *Client) DisableAPI(ctx context.Context) (*MessageResponse, error) {
	var response MessageResponse
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/disable", nil, &response)
	return &response, err
}
