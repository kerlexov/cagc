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
	ID                             int     `json:"id,omitempty"`
	RepositoryProjectID            *int    `json:"repository_project_id,omitempty"`
	UUID                           string  `json:"uuid,omitempty"`
	Name                           string  `json:"name,omitempty"`
	Fqdn                           *string `json:"fqdn,omitempty"`
	ConfigHash                     string  `json:"config_hash,omitempty"`
	GitRepository                  string  `json:"git_repository,omitempty"`
	GitBranch                      string  `json:"git_branch,omitempty"`
	GitCommitSHA                   string  `json:"git_commit_sha,omitempty"`
	GitFullURL                     *string `json:"git_full_url,omitempty"`
	DockerRegistryImageName        *string `json:"docker_registry_image_name,omitempty"`
	DockerRegistryImageTag         *string `json:"docker_registry_image_tag,omitempty"`
	BuildPack                      string  `json:"build_pack,omitempty"`
	StaticImage                    string  `json:"static_image,omitempty"`
	InstallCommand                 string  `json:"install_command,omitempty"`
	BuildCommand                   string  `json:"build_command,omitempty"`
	StartCommand                   string  `json:"start_command,omitempty"`
	PortsExposes                   string  `json:"ports_exposes,omitempty"`
	PortsMappings                  *string `json:"ports_mappings,omitempty"`
	BaseDirectory                  string  `json:"base_directory,omitempty"`
	PublishDirectory               string  `json:"publish_directory,omitempty"`
	HealthCheckEnabled             bool    `json:"health_check_enabled,omitempty"`
	HealthCheckPath                string  `json:"health_check_path,omitempty"`
	HealthCheckPort                *string `json:"health_check_port,omitempty"`
	HealthCheckHost                *string `json:"health_check_host,omitempty"`
	HealthCheckMethod              string  `json:"health_check_method,omitempty"`
	HealthCheckReturnCode          int     `json:"health_check_return_code,omitempty"`
	HealthCheckScheme              string  `json:"health_check_scheme,omitempty"`
	HealthCheckResponseText        *string `json:"health_check_response_text,omitempty"`
	HealthCheckInterval            int     `json:"health_check_interval,omitempty"`
	HealthCheckTimeout             int     `json:"health_check_timeout,omitempty"`
	HealthCheckRetries             int     `json:"health_check_retries,omitempty"`
	HealthCheckStartPeriod         int     `json:"health_check_start_period,omitempty"`
	LimitsMemory                   string  `json:"limits_memory,omitempty"`
	LimitsMemorySwap               string  `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness         int     `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation        string  `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs                     string  `json:"limits_cpus,omitempty"`
	LimitsCPUSet                   *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares                int     `json:"limits_cpu_shares,omitempty"`
	Status                         string  `json:"status,omitempty"`
	PreviewURLTemplate             string  `json:"preview_url_template,omitempty"`
	DestinationType                string  `json:"destination_type,omitempty"`
	DestinationID                  int     `json:"destination_id,omitempty"`
	SourceID                       *int    `json:"source_id,omitempty"`
	PrivateKeyID                   *int    `json:"private_key_id,omitempty"`
	EnvironmentID                  int     `json:"environment_id,omitempty"`
	Dockerfile                     *string `json:"dockerfile,omitempty"`
	DockerfileLocation             string  `json:"dockerfile_location,omitempty"`
	CustomLabels                   *string `json:"custom_labels,omitempty"`
	DockerfileTargetBuild          *string `json:"dockerfile_target_build,omitempty"`
	ManualWebhookSecretGithub      *string `json:"manual_webhook_secret_github,omitempty"`
	ManualWebhookSecretGitlab      *string `json:"manual_webhook_secret_gitlab,omitempty"`
	ManualWebhookSecretBitbucket   *string `json:"manual_webhook_secret_bitbucket,omitempty"`
	ManualWebhookSecretGitea       *string `json:"manual_webhook_secret_gitea,omitempty"`
	DockerComposeLocation          string  `json:"docker_compose_location,omitempty"`
	DockerCompose                  *string `json:"docker_compose,omitempty"`
	DockerComposeRaw               *string `json:"docker_compose_raw,omitempty"`
	DockerComposeDomains           *string `json:"docker_compose_domains,omitempty"`
	DockerComposeCustomStartCmd    *string `json:"docker_compose_custom_start_command,omitempty"`
	DockerComposeCustomBuildCmd    *string `json:"docker_compose_custom_build_command,omitempty"`
	SwarmReplicas                  *int    `json:"swarm_replicas,omitempty"`
	SwarmPlacementConstraints      *string `json:"swarm_placement_constraints,omitempty"`
	CustomDockerRunOptions         *string `json:"custom_docker_run_options,omitempty"`
	PostDeploymentCommand          *string `json:"post_deployment_command,omitempty"`
	PostDeploymentCommandContainer *string `json:"post_deployment_command_container,omitempty"`
	PreDeploymentCommand           *string `json:"pre_deployment_command,omitempty"`
	PreDeploymentCommandContainer  *string `json:"pre_deployment_command_container,omitempty"`
	WatchPaths                     *string `json:"watch_paths,omitempty"`
	CustomHealthcheckFound         bool    `json:"custom_healthcheck_found,omitempty"`
	Redirect                       *string `json:"redirect,omitempty"`
	CreatedAt                      string  `json:"created_at,omitempty"`
	UpdatedAt                      string  `json:"updated_at,omitempty"`
	DeletedAt                      *string `json:"deleted_at,omitempty"`
	ComposeParsingVersion          string  `json:"compose_parsing_version,omitempty"`
	CustomNginxConfiguration       *string `json:"custom_nginx_configuration,omitempty"`
	Domains                        string  `json:"domains,omitempty"`
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
	ID                            int            `json:"id,omitempty"`
	UUID                          string         `json:"uuid,omitempty"`
	Name                          string         `json:"name,omitempty"`
	Description                   string         `json:"description,omitempty"`
	IP                            string         `json:"ip,omitempty"`
	User                          string         `json:"user,omitempty"` // Renamed from Username
	Port                          int            `json:"port,omitempty"`
	PrivateKeyUUID                string         `json:"private_key_uuid,omitempty"` // Added based on create/update
	ProxyType                     string         `json:"proxy_type,omitempty"`
	Proxy                         interface{}    `json:"proxy,omitempty"` // Using interface{} as type is 'object'
	HighDiskUsageNotificationSent bool           `json:"high_disk_usage_notification_sent,omitempty"`
	UnreachableNotificationSent   bool           `json:"unreachable_notification_sent,omitempty"`
	UnreachableCount              int            `json:"unreachable_count,omitempty"`
	ValidationLogs                string         `json:"validation_logs,omitempty"`
	LogDrainNotificationSent      bool           `json:"log_drain_notification_sent,omitempty"`
	SwarmCluster                  string         `json:"swarm_cluster,omitempty"`
	Settings                      *ServerSetting `json:"settings,omitempty"`
	IsBuildServer                 bool           `json:"is_build_server,omitempty"`  // Added based on create/update
	InstantValidate               bool           `json:"instant_validate,omitempty"` // Added based on create/update
	// Removed fields not in OpenAPI schema: IsValid, Status, EngineType, CoolifyPublicIp, CoolifyIsProxySetup, CPUs, Memory, Disk, SwapTotal, SwapUsed, SwapFree, CPULoad, MemoryUsed, DiskUsed, NetworkRX, NetworkTX, Password
}

// ServerSetting represents settings for a cagc server
type ServerSetting struct {
	ID                                int    `json:"id,omitempty"`
	ConcurrentBuilds                  int    `json:"concurrent_builds,omitempty"`
	DynamicTimeout                    int    `json:"dynamic_timeout,omitempty"`
	ForceDisabled                     bool   `json:"force_disabled,omitempty"`
	ForceServerCleanup                bool   `json:"force_server_cleanup,omitempty"`
	IsBuildServer                     bool   `json:"is_build_server,omitempty"`
	IsCloudflareTunnel                bool   `json:"is_cloudflare_tunnel,omitempty"`
	IsJumpServer                      bool   `json:"is_jump_server,omitempty"`
	IsLogdrainAxiomEnabled            bool   `json:"is_logdrain_axiom_enabled,omitempty"`
	IsLogdrainCustomEnabled           bool   `json:"is_logdrain_custom_enabled,omitempty"`
	IsLogdrainHighlightEnabled        bool   `json:"is_logdrain_highlight_enabled,omitempty"`
	IsLogdrainNewrelicEnabled         bool   `json:"is_logdrain_newrelic_enabled,omitempty"`
	IsMetricsEnabled                  bool   `json:"is_metrics_enabled,omitempty"`
	IsReachable                       bool   `json:"is_reachable,omitempty"`
	IsSentinelEnabled                 bool   `json:"is_sentinel_enabled,omitempty"`
	IsSwarmManager                    bool   `json:"is_swarm_manager,omitempty"`
	IsSwarmWorker                     bool   `json:"is_swarm_worker,omitempty"`
	IsUsable                          bool   `json:"is_usable,omitempty"`
	LogdrainAxiomAPIKey               string `json:"logdrain_axiom_api_key,omitempty"`
	LogdrainAxiomDatasetName          string `json:"logdrain_axiom_dataset_name,omitempty"`
	LogdrainCustomConfig              string `json:"logdrain_custom_config,omitempty"`
	LogdrainCustomConfigParser        string `json:"logdrain_custom_config_parser,omitempty"`
	LogdrainHighlightProjectID        string `json:"logdrain_highlight_project_id,omitempty"`
	LogdrainNewrelicBaseURI           string `json:"logdrain_newrelic_base_uri,omitempty"`
	LogdrainNewrelicLicenseKey        string `json:"logdrain_newrelic_license_key,omitempty"`
	SentinelMetricsHistoryDays        int    `json:"sentinel_metrics_history_days,omitempty"`
	SentinelMetricsRefreshRateSeconds int    `json:"sentinel_metrics_refresh_rate_seconds,omitempty"`
	SentinelToken                     string `json:"sentinel_token,omitempty"`
	DockerCleanupFrequency            string `json:"docker_cleanup_frequency,omitempty"`
	DockerCleanupThreshold            int    `json:"docker_cleanup_threshold,omitempty"`
	ServerID                          int    `json:"server_id,omitempty"`
	WildcardDomain                    string `json:"wildcard_domain,omitempty"`
	CreatedAt                         string `json:"created_at,omitempty"`
	UpdatedAt                         string `json:"updated_at,omitempty"`
	DeleteUnusedVolumes               bool   `json:"delete_unused_volumes,omitempty"`
	DeleteUnusedNetworks              bool   `json:"delete_unused_networks,omitempty"`
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
	ID                              int     `json:"id,omitempty"`
	UUID                            string  `json:"uuid,omitempty"`
	Name                            string  `json:"name,omitempty"`
	EnvironmentID                   int     `json:"environment_id,omitempty"`
	ServerID                        int     `json:"server_id,omitempty"`
	Description                     string  `json:"description,omitempty"`
	DockerComposeRaw                string  `json:"docker_compose_raw,omitempty"`
	DockerCompose                   string  `json:"docker_compose,omitempty"`
	DestinationType                 string  `json:"destination_type,omitempty"`
	DestinationID                   int     `json:"destination_id,omitempty"`
	ConnectToDockerNetwork          bool    `json:"connect_to_docker_network,omitempty"`
	IsContainerLabelEscapeEnabled   bool    `json:"is_container_label_escape_enabled,omitempty"`
	IsContainerLabelReadonlyEnabled bool    `json:"is_container_label_readonly_enabled,omitempty"`
	ConfigHash                      string  `json:"config_hash,omitempty"`
	ServiceType                     string  `json:"service_type,omitempty"`
	CreatedAt                       string  `json:"created_at,omitempty"`
	UpdatedAt                       string  `json:"updated_at,omitempty"`
	DeletedAt                       *string `json:"deleted_at,omitempty"`
}

// PrivateKey represents a cagc private key for SSH access
type PrivateKey struct {
	ID           int    `json:"id,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	PrivateKey   string `json:"private_key,omitempty"`
	IsGitRelated bool   `json:"is_git_related,omitempty"`
	TeamID       int    `json:"team_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// Project represents a cagc project
type Project struct {
	ID           int           `json:"id,omitempty"`
	UUID         string        `json:"uuid,omitempty"`
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
}

// Environment represents a project environment
type Environment struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ProjectID   int    `json:"project_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Description string `json:"description,omitempty"`
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

// Deployment represents a cagc deployment (maps to ApplicationDeploymentQueue in schema)
type Deployment struct {
	ID               int    `json:"id,omitempty"`              // Added
	ApplicationID    string `json:"application_id,omitempty"`  // Added
	DeploymentUUID   string `json:"deployment_uuid,omitempty"` // Renamed from UUID
	PullRequestID    int    `json:"pull_request_id,omitempty"` // Added
	ForceRebuild     bool   `json:"force_rebuild,omitempty"`   // Added
	Commit           string `json:"commit,omitempty"`          // Renamed from CommitInfo
	Status           string `json:"status,omitempty"`
	IsWebhook        bool   `json:"is_webhook,omitempty"` // Added
	IsAPI            bool   `json:"is_api,omitempty"`     // Added
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	Logs             string `json:"logs,omitempty"`               // Renamed from LogsURL
	CurrentProcessID string `json:"current_process_id,omitempty"` // Added
	RestartOnly      bool   `json:"restart_only,omitempty"`       // Added
	GitType          string `json:"git_type,omitempty"`           // Added
	ServerID         int    `json:"server_id,omitempty"`          // Added
	ApplicationName  string `json:"application_name,omitempty"`   // Added
	ServerName       string `json:"server_name,omitempty"`        // Added
	DeploymentURL    string `json:"deployment_url,omitempty"`     // Added
	DestinationID    string `json:"destination_id,omitempty"`     // Added
	OnlyThisServer   bool   `json:"only_this_server,omitempty"`   // Added
	Rollback         bool   `json:"rollback,omitempty"`           // Added
	CommitMessage    string `json:"commit_message,omitempty"`     // Added
	// Removed fields not in schema: ResourceUUID, ResourceType, Tag
}

// Team represents a cagc team
type Team struct {
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	PersonalTeam      bool   `json:"personal_team,omitempty"` // Added
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	ShowBoarding      bool   `json:"show_boarding,omitempty"`       // Added
	CustomServerLimit string `json:"custom_server_limit,omitempty"` // Added
	Members           []User `json:"members,omitempty"`             // Added
}

// User represents a Coolify user (part of Team schema)
type User struct {
	ID                   int     `json:"id,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Email                string  `json:"email,omitempty"`
	EmailVerifiedAt      *string `json:"email_verified_at,omitempty"` // Pointer for nullable
	CreatedAt            string  `json:"created_at,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	TwoFactorConfirmedAt *string `json:"two_factor_confirmed_at,omitempty"` // Pointer for nullable
	ForcePasswordReset   bool    `json:"force_password_reset,omitempty"`
	MarketingEmails      bool    `json:"marketing_emails,omitempty"`
}
