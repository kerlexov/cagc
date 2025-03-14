package cagc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListApplications lists all applications
func (c *Client) ListApplications(ctx context.Context) ([]Application, error) {
	var applications []Application
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/applications", nil, &applications)
	return applications, err
}

// GetApplication gets an application by UUID
func (c *Client) GetApplication(ctx context.Context, uuid string) (*Application, error) {
	path := fmt.Sprintf("/applications/%s", uuid)
	var application Application
	err := c.doRequest(ctx, http.MethodGet, path, nil, &application)
	return &application, err
}

// CreatePublicApplication creates a new application based on a public git repository
func (c *Client) CreatePublicApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/public", app, &response)
	return &response, err
}

// CreatePrivateGithubAppApplication creates a new application based on a private repo through Github App
func (c *Client) CreatePrivateGithubAppApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/private-github-app", app, &response)
	return &response, err
}

// CreatePrivateDeployKeyApplication creates a new application based on a private repo through Deploy Key
func (c *Client) CreatePrivateDeployKeyApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/private-deploy-key", app, &response)
	return &response, err
}

// CreateDockerfileApplication creates a new application based on a Dockerfile
func (c *Client) CreateDockerfileApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/dockerfile", app, &response)
	return &response, err
}

// CreateDockerImageApplication creates a new application based on a Docker image
func (c *Client) CreateDockerImageApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/dockerimage", app, &response)
	return &response, err
}

// CreateDockerComposeApplication creates a new application based on a docker-compose file
func (c *Client) CreateDockerComposeApplication(ctx context.Context, app Application) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/applications/dockercompose", app, &response)
	return &response, err
}

// UpdateApplication updates an existing application
func (c *Client) UpdateApplication(ctx context.Context, uuid string, app Application) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, app, &response)
	return &response, err
}

// DeleteApplication deletes an application
func (c *Client) DeleteApplication(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s", uuid)
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

// StartApplication starts an application
func (c *Client) StartApplication(ctx context.Context, uuid string, force, instantDeploy bool) (*DeploymentResponse, error) {
	path := fmt.Sprintf("/applications/%s/start", uuid)
	query := url.Values{}
	query.Add("force", fmt.Sprintf("%t", force))
	query.Add("instant_deploy", fmt.Sprintf("%t", instantDeploy))

	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	var response DeploymentResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// StopApplication stops an application
func (c *Client) StopApplication(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s/stop", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// RestartApplication restarts an application
func (c *Client) RestartApplication(ctx context.Context, uuid string) (*DeploymentResponse, error) {
	path := fmt.Sprintf("/applications/%s/restart", uuid)
	var response DeploymentResponse
	err := c.doRequest(ctx, http.MethodGet, path, nil, &response)
	return &response, err
}

// ExecuteCommand executes a command on an application's container
func (c *Client) ExecuteCommand(ctx context.Context, uuid string, command string) (*CommandResponse, error) {
	path := fmt.Sprintf("/applications/%s/execute", uuid)
	req := map[string]string{"command": command}
	var response CommandResponse
	err := c.doRequest(ctx, http.MethodPost, path, req, &response)
	return &response, err
}

// ListApplicationEnvs lists all environment variables for an application
func (c *Client) ListApplicationEnvs(ctx context.Context, uuid string) ([]EnvironmentVariable, error) {
	path := fmt.Sprintf("/applications/%s/envs", uuid)
	var envs []EnvironmentVariable
	err := c.doRequest(ctx, http.MethodGet, path, nil, &envs)
	return envs, err
}

// CreateApplicationEnv creates a new environment variable for an application
func (c *Client) CreateApplicationEnv(ctx context.Context, appUUID string, env EnvironmentVariable) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s/envs", appUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, path, env, &response)
	return &response, err
}

// UpdateApplicationEnv updates an environment variable for an application
func (c *Client) UpdateApplicationEnv(ctx context.Context, appUUID string, env EnvironmentVariable) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s/envs", appUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, env, &response)
	return &response, err
}

// DeleteApplicationEnv deletes an environment variable for an application
func (c *Client) DeleteApplicationEnv(ctx context.Context, appUUID string, envUUID string) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s/envs/%s", appUUID, envUUID)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}

// UpdateApplicationEnvsBulk updates multiple environment variables for an application
func (c *Client) UpdateApplicationEnvsBulk(ctx context.Context, appUUID string, envs []EnvironmentVariable) (*CreateResponse, error) {
	path := fmt.Sprintf("/applications/%s/envs/bulk", appUUID)
	req := map[string][]EnvironmentVariable{"data": envs}
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, req, &response)
	return &response, err
}
