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

// Server represents a cagc server
type Server struct {
	UUID                string  `json:"uuid,omitempty"`
	Name                string  `json:"name,omitempty"`
	Description         string  `json:"description,omitempty"`
	IP                  string  `json:"ip,omitempty"`
	Port                int     `json:"port,omitempty"`
	Username            string  `json:"username,omitempty"`
	Password            string  `json:"password,omitempty"`
	PrivateKeyUUID      string  `json:"private_key_uuid,omitempty"`
	IsValid             bool    `json:"is_valid,omitempty"`
	Status              string  `json:"status,omitempty"`
	EngineType          string  `json:"engine_type,omitempty"`
	CoolifyPublicIp     string  `json:"coolify_public_ip,omitempty"`
	CoolifyIsProxySetup bool    `json:"coolify_is_proxy_setup,omitempty"`
	CPUs                int     `json:"cpus,omitempty"`
	Memory              int     `json:"memory,omitempty"`
	Disk                int     `json:"disk,omitempty"`
	SwapTotal           int     `json:"swap_total,omitempty"`
	SwapUsed            int     `json:"swap_used,omitempty"`
	SwapFree            int     `json:"swap_free,omitempty"`
	CPULoad             float64 `json:"cpu_load,omitempty"`
	MemoryUsed          int     `json:"memory_used,omitempty"`
	DiskUsed            int     `json:"disk_used,omitempty"`
	NetworkRX           int     `json:"network_rx,omitempty"`
	NetworkTX           int     `json:"network_tx,omitempty"`
}

// Resource represents a resource on a server
type Resource struct {
	ID        int    `json:"id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Status    string `json:"status,omitempty"`
}

// ServerDomain represents a domain configuration on a server
type ServerDomain struct {
	IP      string   `json:"ip,omitempty"`
	Domains []string `json:"domains,omitempty"`
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

// CommandResponse represents a command execution response
type CommandResponse struct {
	Message  string `json:"message,omitempty"`
	Response string `json:"response,omitempty"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message,omitempty"`
}

// Service represents a cagc service
type Service struct {
	UUID                    string `json:"uuid,omitempty"`
	ProjectUUID             string `json:"project_uuid,omitempty"`
	ServerUUID              string `json:"server_uuid,omitempty"`
	EnvironmentName         string `json:"environment_name,omitempty"`
	EnvironmentUUID         string `json:"environment_uuid,omitempty"`
	DestinationUUID         string `json:"destination_uuid,omitempty"`
	Type                    string `json:"type,omitempty"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	Image                   string `json:"image,omitempty"`
	Version                 string `json:"version,omitempty"`
	Configuration           string `json:"configuration,omitempty"`
	Domains                 string `json:"domains,omitempty"`
	LimitsMemory            string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         int    `json:"limits_cpu_shares,omitempty"`
	CustomLabels            string `json:"custom_labels,omitempty"`
	CustomDockerRunOptions  string `json:"custom_docker_run_options,omitempty"`
	PostDeploymentCommand   string `json:"post_deployment_command,omitempty"`
	PreDeploymentCommand    string `json:"pre_deployment_command,omitempty"`
	PublicPort              int    `json:"public_port,omitempty"`
	Volume                  string `json:"volume,omitempty"`
	InstantDeploy           bool   `json:"instant_deploy,omitempty"`
	PortsMappings           string `json:"ports_mappings,omitempty"`
	IsCustomCommand         bool   `json:"is_custom_command,omitempty"`
	CustomStartCommand      string `json:"custom_start_command,omitempty"`
	CustomBuildCommand      string `json:"custom_build_command,omitempty"`
	Commands                string `json:"commands,omitempty"`
	HealthCheckEnabled      bool   `json:"health_check_enabled,omitempty"`
	HealthCheckPath         string `json:"health_check_path,omitempty"`
	HealthCheckPort         string `json:"health_check_port,omitempty"`
	HealthCheckHost         string `json:"health_check_host,omitempty"`
	HealthCheckReturnCode   int    `json:"health_check_return_code,omitempty"`
	HealthCheckScheme       string `json:"health_check_scheme,omitempty"`
	HealthCheckMethod       string `json:"health_check_method,omitempty"`
	HealthCheckTimeout      int    `json:"health_check_timeout,omitempty"`
	HealthCheckInterval     int    `json:"health_check_interval,omitempty"`
	HealthCheckRetries      int    `json:"health_check_retries,omitempty"`
	HealthCheckStartPeriod  int    `json:"health_check_start_period,omitempty"`
	HealthCheckResponseText string `json:"health_check_response_text,omitempty"`
}

// PrivateKey represents a cagc private key for SSH access
type PrivateKey struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PrivateKey  string `json:"private_key,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// Project represents a cagc project
type Project struct {
	UUID        string `json:"uuid,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// Destination represents a cagc destination
type Destination struct {
	UUID          string `json:"uuid,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	ServerUUID    string `json:"server_uuid,omitempty"`
	EngineType    string `json:"engine_type,omitempty"`
	NetworkUUID   string `json:"network_uuid,omitempty"`
	NetworkName   string `json:"network_name,omitempty"`
	Engine        string `json:"engine,omitempty"`
	ResourceCount int    `json:"resource_count,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}
