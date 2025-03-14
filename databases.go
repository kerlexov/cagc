package cagc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListDatabases lists all databases
func (c *Client) ListDatabases(ctx context.Context) ([]Database, error) {
	var databases []Database
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/databases", nil, &databases)
	return databases, err
}

// GetDatabase gets a database by UUID
func (c *Client) GetDatabase(ctx context.Context, uuid string) (*Database, error) {
	path := fmt.Sprintf("/api/v1/databases/%s", uuid)
	var database Database
	err := c.doRequest(ctx, http.MethodGet, path, nil, &database)
	return &database, err
}

// CreatePostgresDatabase creates a new PostgreSQL database
func (c *Client) CreatePostgresDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/postgresql", db, &response)
	return &response, err
}

// CreateClickhouseDatabase creates a new Clickhouse database
func (c *Client) CreateClickhouseDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/clickhouse", db, &response)
	return &response, err
}

// CreateDragonflyDatabase creates a new DragonFly database
func (c *Client) CreateDragonflyDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/dragonfly", db, &response)
	return &response, err
}

// CreateRedisDatabase creates a new Redis database
func (c *Client) CreateRedisDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/redis", db, &response)
	return &response, err
}

// CreateKeyDBDatabase creates a new KeyDB database
func (c *Client) CreateKeyDBDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/keydb", db, &response)
	return &response, err
}

// CreateMariaDBDatabase creates a new MariaDB database
func (c *Client) CreateMariaDBDatabase(ctx context.Context, db Database) (*CreateResponse, error) {
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPost, "/api/v1/databases/mariadb", db, &response)
	return &response, err
}

// UpdateDatabase updates an existing database
func (c *Client) UpdateDatabase(ctx context.Context, uuid string, db Database) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/databases/%s", uuid)
	var response CreateResponse
	err := c.doRequest(ctx, http.MethodPatch, path, db, &response)
	return &response, err
}

// DeleteDatabase deletes a database
func (c *Client) DeleteDatabase(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error) {
	path := fmt.Sprintf("/api/v1/databases/%s", uuid)
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
