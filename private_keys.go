package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListPrivateKeys lists all private keys
func (c *Client) ListPrivateKeys(ctx context.Context) ([]PrivateKey, error) {
	var keys []PrivateKey
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/security/keys", nil, &keys)
	return keys, err
}

// GetPrivateKey gets a private key by UUID
func (c *Client) GetPrivateKey(ctx context.Context, uuid string) (*PrivateKey, error) {
	path := fmt.Sprintf("/api/v1/security/keys/%s", uuid)
	var key PrivateKey
	err := c.doRequest(ctx, http.MethodGet, path, nil, &key)
	return &key, err
}

// CreatePrivateKey creates a new private key
func (c *Client) CreatePrivateKey(ctx context.Context, key PrivateKey) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/security/keys", key, &response)
	return &response, err
}

// DeletePrivateKey deletes a private key
func (c *Client) DeletePrivateKey(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/security/keys/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}
