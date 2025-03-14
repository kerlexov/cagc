package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListProjects lists all projects
func (c *Client) ListProjects(ctx context.Context) ([]Project, error) {
	var projects []Project
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/projects", nil, &projects)
	return projects, err
}

// GetProject gets a project by UUID
func (c *Client) GetProject(ctx context.Context, uuid string) (*Project, error) {
	path := fmt.Sprintf("/api/v1/projects/%s", uuid)
	var project Project
	err := c.doRequest(ctx, http.MethodGet, path, nil, &project)
	return &project, err
}

// CreateProject creates a new project
func (c *Client) CreateProject(ctx context.Context, project Project) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/projects", project, &response)
	return &response, err
}

// UpdateProject updates an existing project
func (c *Client) UpdateProject(ctx context.Context, uuid string, project Project) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/projects/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, project, &response)
	return &response, err
}

// DeleteProject deletes a project
func (c *Client) DeleteProject(ctx context.Context, uuid string) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/projects/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodDelete, path, nil, &response)
	return &response, err
}
