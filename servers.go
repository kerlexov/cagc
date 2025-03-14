package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListServers lists all servers
func (c *Client) ListServers(ctx context.Context) ([]Server, error) {
	var servers []Server
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/servers", nil, &servers)
	return servers, err
}

// GetServer gets a server by UUID
func (c *Client) GetServer(ctx context.Context, uuid string) (*Server, error) {
	path := fmt.Sprintf("/api/v1/servers/%s", uuid)
	var server Server
	err := c.doRequest(ctx, http.MethodGet, path, nil, &server)
	return &server, err
}

// CreateServer creates a new server
func (c *Client) CreateServer(ctx context.Context, server Server) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/servers", server, &response)
	return &response, err
}

// UpdateServer updates an existing server
func (c *Client) UpdateServer(ctx context.Context, uuid string, server Server) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/servers/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, server, &response)
	return &response, err
}

// DeleteServer deletes a server
func (c *Client) DeleteServer(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/servers/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}

// ValidateServer validates a server by UUID
func (c *Client) ValidateServer(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/servers/%s/validate", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// GetServerResources gets resources by server UUID
func (c *Client) GetServerResources(ctx context.Context, uuid string) ([]Resource, error) {
	path := fmt.Sprintf("/api/v1/servers/%s/resources", uuid)
	var resources []Resource
	err := c.doRequest(ctx, http.MethodGet, path, nil, &resources)
	return resources, err
}

// GetServerDomains gets domains by server UUID
func (c *Client) GetServerDomains(ctx context.Context, uuid string) ([]ServerDomain, error) {
	path := fmt.Sprintf("/api/v1/servers/%s/domains", uuid)
	var domains []ServerDomain
	err := c.doRequest(ctx, http.MethodGet, path, nil, &domains)
	return domains, err
}
