package cagc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client represents a cagc API client
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	token      string
}

// NewClient creates a new cagc API client
func NewClient(baseURL string, token string) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    parsedURL,
		httpClient: http.DefaultClient,
		token:      token,
	}, nil
}

// doRequest performs an HTTP request and decodes the response into v if provided
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, v interface{}) error {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return err
	}

	var req *http.Request
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		req, err = http.NewRequestWithContext(ctx, method, u.String(), bytes.NewBuffer(jsonBody))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(ctx, method, u.String(), nil)
		if err != nil {
			return err
		}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading error response: %v", err)
		}
		return fmt.Errorf("API error: %s, status code: %d", string(bodyBytes), resp.StatusCode)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}

// Error represents an API error response
type Error struct {
	Message string `json:"message"`
}

// Application represents a cagc application
type Application struct {
	UUID                           string   `json:"uuid,omitempty"`
	ProjectUUID                    string   `json:"project_uuid,omitempty"`
	ServerUUID                     string   `json:"server_uuid,omitempty"`
	EnvironmentName                string   `json:"environment_name,omitempty"`
	EnvironmentUUID                string   `json:"environment_uuid,omitempty"`
	GitRepository                  string   `json:"git_repository,omitempty"`
	GitBranch                      string   `json:"git_branch,omitempty"`
	BuildPack                      string   `json:"build_pack,omitempty"`
	PortsExposes                   string   `json:"ports_exposes,omitempty"`
	DestinationUUID                string   `json:"destination_uuid,omitempty"`
	Name                           string   `json:"name,omitempty"`
	Description                    string   `json:"description,omitempty"`
	Domains                        string   `json:"domains,omitempty"`
	GitCommitSHA                   string   `json:"git_commit_sha,omitempty"`
	DockerRegistryImageName        string   `json:"docker_registry_image_name,omitempty"`
	DockerRegistryImageTag         string   `json:"docker_registry_image_tag,omitempty"`
	IsStatic                       bool     `json:"is_static,omitempty"`
	StaticImage                    string   `json:"static_image,omitempty"`
	InstallCommand                 string   `json:"install_command,omitempty"`
	BuildCommand                   string   `json:"build_command,omitempty"`
	StartCommand                   string   `json:"start_command,omitempty"`
	PortsMappings                  string   `json:"ports_mappings,omitempty"`
	BaseDirectory                  string   `json:"base_directory,omitempty"`
	PublishDirectory               string   `json:"publish_directory,omitempty"`
	HealthCheckEnabled             bool     `json:"health_check_enabled,omitempty"`
	HealthCheckPath                string   `json:"health_check_path,omitempty"`
	HealthCheckPort                *string  `json:"health_check_port,omitempty"`
	HealthCheckHost                *string  `json:"health_check_host,omitempty"`
	HealthCheckMethod              string   `json:"health_check_method,omitempty"`
	HealthCheckReturnCode          int      `json:"health_check_return_code,omitempty"`
	HealthCheckScheme              string   `json:"health_check_scheme,omitempty"`
	HealthCheckResponseText        *string  `json:"health_check_response_text,omitempty"`
	HealthCheckInterval            int      `json:"health_check_interval,omitempty"`
	HealthCheckTimeout             int      `json:"health_check_timeout,omitempty"`
	HealthCheckRetries             int      `json:"health_check_retries,omitempty"`
	HealthCheckStartPeriod         int      `json:"health_check_start_period,omitempty"`
	LimitsMemory                   string   `json:"limits_memory,omitempty"`
	LimitsMemorySwap               string   `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness         int      `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation        string   `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs                     string   `json:"limits_cpus,omitempty"`
	LimitsCPUSet                   *string  `json:"limits_cpuset,omitempty"`
	LimitsCPUShares                int      `json:"limits_cpu_shares,omitempty"`
	CustomLabels                   string   `json:"custom_labels,omitempty"`
	CustomDockerRunOptions         string   `json:"custom_docker_run_options,omitempty"`
	PostDeploymentCommand          string   `json:"post_deployment_command,omitempty"`
	PostDeploymentCommandContainer string   `json:"post_deployment_command_container,omitempty"`
	PreDeploymentCommand           string   `json:"pre_deployment_command,omitempty"`
	PreDeploymentCommandContainer  string   `json:"pre_deployment_command_container,omitempty"`
	ManualWebhookSecretGithub      string   `json:"manual_webhook_secret_github,omitempty"`
	ManualWebhookSecretGitlab      string   `json:"manual_webhook_secret_gitlab,omitempty"`
	ManualWebhookSecretBitbucket   string   `json:"manual_webhook_secret_bitbucket,omitempty"`
	ManualWebhookSecretGitea       string   `json:"manual_webhook_secret_gitea,omitempty"`
	Redirect                       *string  `json:"redirect,omitempty"`
	InstantDeploy                  bool     `json:"instant_deploy,omitempty"`
	Dockerfile                     string   `json:"dockerfile,omitempty"`
	DockerComposeLocation          string   `json:"docker_compose_location,omitempty"`
	DockerComposeRaw               string   `json:"docker_compose_raw,omitempty"`
	DockerComposeCustomStartCmd    string   `json:"docker_compose_custom_start_command,omitempty"`
	DockerComposeCustomBuildCmd    string   `json:"docker_compose_custom_build_command,omitempty"`
	DockerComposeDomains           []string `json:"docker_compose_domains,omitempty"`
	WatchPaths                     string   `json:"watch_paths,omitempty"`
	UseBuildServer                 *bool    `json:"use_build_server,omitempty"`
	GitHubAppUUID                  string   `json:"github_app_uuid,omitempty"`
	PrivateKeyUUID                 string   `json:"private_key_uuid,omitempty"`
}

// Database represents a cagc database
type Database struct {
	UUID                    string `json:"uuid,omitempty"`
	ProjectUUID             string `json:"project_uuid,omitempty"`
	ServerUUID              string `json:"server_uuid,omitempty"`
	EnvironmentName         string `json:"environment_name,omitempty"`
	EnvironmentUUID         string `json:"environment_uuid,omitempty"`
	DestinationUUID         string `json:"destination_uuid,omitempty"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	Image                   string `json:"image,omitempty"`
	IsPublic                bool   `json:"is_public,omitempty"`
	PublicPort              int    `json:"public_port,omitempty"`
	LimitsMemory            string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         int    `json:"limits_cpu_shares,omitempty"`
	InstantDeploy           bool   `json:"instant_deploy,omitempty"`

	// PostgreSQL specific
	PostgresUser           string `json:"postgres_user,omitempty"`
	PostgresPassword       string `json:"postgres_password,omitempty"`
	PostgresDB             string `json:"postgres_db,omitempty"`
	PostgresInitdbArgs     string `json:"postgres_initdb_args,omitempty"`
	PostgresHostAuthMethod string `json:"postgres_host_auth_method,omitempty"`
	PostgresConf           string `json:"postgres_conf,omitempty"`

	// Clickhouse specific
	ClickhouseAdminUser     string `json:"clickhouse_admin_user,omitempty"`
	ClickhouseAdminPassword string `json:"clickhouse_admin_password,omitempty"`

	// DragonFly specific
	DragonflyPassword string `json:"dragonfly_password,omitempty"`

	// Redis specific
	RedisPassword string `json:"redis_password,omitempty"`
	RedisConf     string `json:"redis_conf,omitempty"`

	// KeyDB specific
	KeyDBPassword string `json:"keydb_password,omitempty"`
	KeyDBConf     string `json:"keydb_conf,omitempty"`

	// MariaDB specific
	MariaDBConf         string `json:"mariadb_conf,omitempty"`
	MariaDBRootPassword string `json:"mariadb_root_password,omitempty"`
	MariaDBUser         string `json:"mariadb_user,omitempty"`
	MariaDBPassword     string `json:"mariadb_password,omitempty"`
	MariaDBDatabase     string `json:"mariadb_database,omitempty"`

	// MongoDB specific
	MongoConf               string `json:"mongo_conf,omitempty"`
	MongoInitdbRootUsername string `json:"mongo_initdb_root_username,omitempty"`
	MongoInitdbRootPassword string `json:"mongo_initdb_root_password,omitempty"`
	MongoInitdbDatabase     string `json:"mongo_initdb_database,omitempty"`

	// MySQL specific
	MySQLRootPassword string `json:"mysql_root_password,omitempty"`
	MySQLPassword     string `json:"mysql_password,omitempty"`
	MySQLUser         string `json:"mysql_user,omitempty"`
	MySQLDatabase     string `json:"mysql_database,omitempty"`
	MySQLConf         string `json:"mysql_conf,omitempty"`
}

// EnvironmentVariable represents a cagc environment variable
type EnvironmentVariable struct {
	UUID        string `json:"uuid,omitempty"`
	Key         string `json:"key,omitempty"`
	Value       string `json:"value,omitempty"`
	IsPreview   bool   `json:"is_preview,omitempty"`
	IsBuildTime bool   `json:"is_build_time,omitempty"`
	IsLiteral   bool   `json:"is_literal,omitempty"`
	IsMultiline bool   `json:"is_multiline,omitempty"`
	IsShownOnce bool   `json:"is_shown_once,omitempty"`
}

// CreateResponse is a generic response for create operations
type CreateResponse struct {
	UUID    string `json:"uuid,omitempty"`
	Message string `json:"message,omitempty"`
}

// DeploymentResponse represents a response for deployment operations
type DeploymentResponse struct {
	Message        string `json:"message,omitempty"`
	DeploymentUUID string `json:"deployment_uuid,omitempty"`
}

// CommandResponse represents a response for command execution
type CommandResponse struct {
	Message  string `json:"message,omitempty"`
	Response string `json:"response,omitempty"`
}
