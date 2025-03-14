package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListResources lists all resources
func (c *Client) ListResources(ctx context.Context) ([]Resource, error) {
	var resources []Resource
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/resources", nil, &resources)
	return resources, err
}

// ListDestinations lists all destinations (keeping this for backward compatibility)
func (c *Client) ListDestinations(ctx context.Context) ([]Destination, error) {
	var destinations []Destination
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/destinations", nil, &destinations)
	return destinations, err
}

// GetDestination gets a destination by UUID (keeping this for backward compatibility)
func (c *Client) GetDestination(ctx context.Context, uuid string) (*Destination, error) {
	path := fmt.Sprintf("/api/v1/destinations/%s", uuid)
	var destination Destination
	err := c.doRequest(ctx, http.MethodGet, path, nil, &destination)
	return &destination, err
}

// CreateDestination creates a new destination (keeping this for backward compatibility)
func (c *Client) CreateDestination(ctx context.Context, destination Destination) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/destinations", destination, &response)
	return &response, err
}

// UpdateDestination updates an existing destination (keeping this for backward compatibility)
func (c *Client) UpdateDestination(ctx context.Context, uuid string, destination Destination) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/destinations/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, destination, &response)
	return &response, err
}

// DeleteDestination deletes a destination (keeping this for backward compatibility)
func (c *Client) DeleteDestination(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/destinations/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}
