package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListDeployments lists all currently running deployments
func (c *Client) ListDeployments(ctx context.Context) ([]Deployment, error) {
	var deployments []Deployment
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/deployments", nil, &deployments)
	return deployments, err
}

// GetDeployment gets a deployment by UUID
func (c *Client) GetDeployment(ctx context.Context, uuid string) (*Deployment, error) {
	path := fmt.Sprintf("/api/v1/deployments/%s", uuid)
	var deployment Deployment
	err := c.doRequest(ctx, http.MethodGet, path, nil, &deployment)
	return &deployment, err
}

// DeployByTagOrUUID deploys by tag or UUID
func (c *Client) DeployByTagOrUUID(ctx context.Context, tagOrUUID string) (*DeploymentResponse, error) {
	path := fmt.Sprintf("/api/v1/deployments/%s", tagOrUUID)
	var response DeploymentResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}
