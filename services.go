package cagc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListServices lists all services
func (c *Client) ListServices(ctx context.Context) ([]Service, error) {
	var services []Service
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/services", nil, &services)
	return services, err
}

// GetService gets a service by UUID
func (c *Client) GetService(ctx context.Context, uuid string) (*Service, error) {
	path := fmt.Sprintf("/api/v1/services/%s", uuid)
	var service Service
	err := c.doRequest(ctx, http.MethodGet, path, nil, &service)
	return &service, err
}

// CreateService creates a new one-click service
func (c *Client) CreateService(ctx context.Context, service Service) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/services", service, &response)
	return &response, err
}

// UpdateService updates an existing service
func (c *Client) UpdateService(ctx context.Context, uuid string, service Service) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, service, &response)
	return &response, err
}

// DeleteService deletes a service
func (c *Client) DeleteService(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s", uuid)
	query := url.Values{}
	query.Add("delete_configurations", fmt.Sprintf("%t", deleteConfigurations))
	query.Add("delete_volumes", fmt.Sprintf("%t", deleteVolumes))
	query.Add("docker_cleanup", fmt.Sprintf("%t", dockerCleanup))
	query.Add("delete_connected_networks", fmt.Sprintf("%t", deleteConnectedNetworks))

	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}

// StartService starts a service
func (c *Client) StartService(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/start", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// StopService stops a service
func (c *Client) StopService(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/stop", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// RestartService restarts a service
func (c *Client) RestartService(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/restart", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// ExecuteServiceCommand executes a command on a service's container
func (c *Client) ExecuteServiceCommand(ctx context.Context, uuid string, command string) (*CommandResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/execute", uuid)
	req := map[string]string{"command": command}
	var response CommandResponse
	err := c.doRequest(ctx, http.MethodPost, path, req, &response)
	return &response, err
}

// ListServiceEnvs lists all environment variables for a service
func (c *Client) ListServiceEnvs(ctx context.Context, uuid string) ([]EnvironmentVariable, error) {
	path := fmt.Sprintf("/api/v1/services/%s/envs", uuid)
	var envs []EnvironmentVariable
	err := c.doRequest(ctx, http.MethodGet, path, nil, &envs)
	return envs, err
}

// CreateServiceEnv creates a new environment variable for a service
func (c *Client) CreateServiceEnv(ctx context.Context, serviceUUID string, env EnvironmentVariable) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/envs", serviceUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, path, env, &response)
	return &response, err
}

// UpdateServiceEnv updates an environment variable for a service
func (c *Client) UpdateServiceEnv(ctx context.Context, serviceUUID string, env EnvironmentVariable) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/envs", serviceUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, env, &response)
	return &response, err
}

// DeleteServiceEnv deletes an environment variable for a service
func (c *Client) DeleteServiceEnv(ctx context.Context, serviceUUID string, envUUID string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/services/%s/envs/%s", serviceUUID, envUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}
