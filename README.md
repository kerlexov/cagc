# Coolify API Go Client

## Version 0.0.2

This is a Go client library for the [Coolify](https://coolify.io/) API, allowing you to programmatically manage your Coolify resources.

## Installation

```bash
go get github.com/kerlexov/cagc
```

## Authentication

To use the Coolify API, you need an API token. You can get one from the Coolify dashboard under "Keys & Tokens" / "API tokens".

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kerlexov/cagc"
)

func main() {
	// Create a new client
	client, err := cagc.NewClient("https://coolify.example.com", "your-api-token")
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Get API version
	version, err := client.GetVersion(context.Background())
	if err != nil {
		log.Fatalf("Error getting API version: %v", err)
	}
	fmt.Printf("Coolify API Version: %f\n\n", version)

	// List all applications
	apps, err := client.ListApplications(context.Background())
	if err != nil {
		log.Fatalf("Error listing applications: %v", err)
	}

	// Print application details
	fmt.Printf("Found %d applications:\n", len(apps))
	for _, app := range apps {
		fmt.Printf("- UUID: %s, Name: %s\n", app.UUID, app.Name)
	}

	// List all servers
	servers, err := client.ListServers(context.Background())
	if err != nil {
		log.Fatalf("Error listing servers: %v", err)
	}

	// Print server details
	fmt.Printf("\nFound %d servers:\n", len(servers))
	for _, server := range servers {
		fmt.Printf("- UUID: %s, Name: %s, Status: %s\n", server.UUID, server.Name, server.Status)
	}

	// List all services
	services, err := client.ListServices(context.Background())
	if err != nil {
		log.Fatalf("Error listing services: %v", err)
	}

	// Print service details
	fmt.Printf("\nFound %d services:\n", len(services))
	for _, service := range services {
		fmt.Printf("- UUID: %s, Name: %s, Type: %s\n", service.UUID, service.Name, service.Type)
	}

	// List all projects
	projects, err := client.ListProjects(context.Background())
	if err != nil {
		log.Fatalf("Error listing projects: %v", err)
	}

	// Print project details
	fmt.Printf("\nFound %d projects:\n", len(projects))
	for _, project := range projects {
		fmt.Printf("- UUID: %s, Name: %s\n", project.UUID, project.Name)
	}

	// List all databases
	dbs, err := client.ListDatabases(context.Background())
	if err != nil {
		log.Fatalf("Error listing databases: %v", err)
	}

	// Print database details
	fmt.Printf("\nFound %d databases:\n", len(dbs))
	for _, db := range dbs {
		fmt.Printf("- UUID: %s, Name: %s\n", db.UUID, db.Name)
	}

	// List all resources
	resources, err := client.ListResources(context.Background())
	if err != nil {
		log.Fatalf("Error listing resources: %v", err)
	}

	// Print resource details
	fmt.Printf("\nFound %d resources:\n", len(resources))
	for _, resource := range resources {
		fmt.Printf("- UUID: %s, Name: %s, Type: %s\n", resource.UUID, resource.Name, resource.Type)
	}

	// List all private keys
	keys, err := client.ListPrivateKeys(context.Background())
	if err != nil {
		log.Fatalf("Error listing private keys: %v", err)
	}

	// Print private key details
	fmt.Printf("\nFound %d private keys:\n", len(keys))
	for _, key := range keys {
		fmt.Printf("- UUID: %s, Name: %s\n", key.UUID, key.Name)
	}

	// List all deployments
	deployments, err := client.ListDeployments(context.Background())
	if err != nil {
		log.Fatalf("Error listing deployments: %v", err)
	}

	// Print deployment details
	fmt.Printf("\nFound %d deployments:\n", len(deployments))
	for _, deployment := range deployments {
		fmt.Printf("- UUID: %s, Status: %s\n", deployment.UUID, deployment.Status)
	}

	// List all teams
	teams, err := client.ListTeams(context.Background())
	if err != nil {
		log.Fatalf("Error listing teams: %v", err)
	}

	// Print team details
	fmt.Printf("\nFound %d teams:\n", len(teams))
	for _, team := range teams {
		fmt.Printf("- ID: %d, Name: %s\n", team.ID, team.Name)
	}

}
```

## Available Operations

### Applications

- `ListApplications(ctx context.Context) ([]Application, error)`
- `GetApplication(ctx context.Context, uuid string) (*Application, error)`
- `CreatePublicApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `CreatePrivateGithubAppApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `CreatePrivateDeployKeyApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `CreateDockerfileApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `CreateDockerImageApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `CreateDockerComposeApplication(ctx context.Context, app Application) (*CreateResponse, error)`
- `UpdateApplication(ctx context.Context, uuid string, app Application) (*CreateResponse, error)`
- `DeleteApplication(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error)`
- `StartApplication(ctx context.Context, uuid string, force, instantDeploy bool) (*DeploymentResponse, error)`
- `StopApplication(ctx context.Context, uuid string) (*CreateResponse, error)`
- `RestartApplication(ctx context.Context, uuid string) (*DeploymentResponse, error)`
- `ExecuteCommand(ctx context.Context, uuid string, command string) (*CommandResponse, error)`

### Application Environment Variables

- `ListApplicationEnvs(ctx context.Context, uuid string) ([]EnvironmentVariable, error)`
- `CreateApplicationEnv(ctx context.Context, appUUID string, env EnvironmentVariable) (*CreateResponse, error)`
- `UpdateApplicationEnv(ctx context.Context, appUUID string, env EnvironmentVariable) (*CreateResponse, error)`
- `DeleteApplicationEnv(ctx context.Context, appUUID string, envUUID string) (*CreateResponse, error)`
- `UpdateApplicationEnvsBulk(ctx context.Context, appUUID string, envs []EnvironmentVariable) (*CreateResponse, error)`

### Databases

- `ListDatabases(ctx context.Context) ([]Database, error)`
- `GetDatabase(ctx context.Context, uuid string) (*Database, error)`
- `CreatePostgresDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `CreateClickhouseDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `CreateDragonflyDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `CreateRedisDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `CreateKeyDBDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `CreateMariaDBDatabase(ctx context.Context, db Database) (*CreateResponse, error)`
- `UpdateDatabase(ctx context.Context, uuid string, db Database) (*CreateResponse, error)`
- `DeleteDatabase(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error)`

### Deployments

- `ListDeployments(ctx context.Context) ([]Deployment, error)`
- `GetDeployment(ctx context.Context, uuid string) (*Deployment, error)`
- `DeployByTagOrUUID(ctx context.Context, tagOrUUID string) (*DeploymentResponse, error)`

### Projects

- `ListProjects(ctx context.Context) ([]Project, error)`
- `GetProject(ctx context.Context, uuid string) (*Project, error)`
- `CreateProject(ctx context.Context, project Project) (*CreateResponse, error)`
- `UpdateProject(ctx context.Context, uuid string, project Project) (*CreateResponse, error)`
- `DeleteProject(ctx context.Context, uuid string) (*CreateResponse, error)`

### Resources

- `ListResources(ctx context.Context) ([]Resource, error)`
- `ListDestinations(ctx context.Context) ([]Destination, error)` (maintained for backward compatibility)
- `GetDestination(ctx context.Context, uuid string) (*Destination, error)` (maintained for backward compatibility)
- `CreateDestination(ctx context.Context, destination Destination) (*CreateResponse, error)` (maintained for backward compatibility)
- `UpdateDestination(ctx context.Context, uuid string, destination Destination) (*CreateResponse, error)` (maintained for backward compatibility)
- `DeleteDestination(ctx context.Context, uuid string) (*CreateResponse, error)` (maintained for backward compatibility)

### Servers

- `ListServers(ctx context.Context) ([]Server, error)`
- `GetServer(ctx context.Context, uuid string) (*Server, error)`
- `CreateServer(ctx context.Context, server Server) (*CreateResponse, error)`
- `UpdateServer(ctx context.Context, uuid string, server Server) (*CreateResponse, error)`
- `DeleteServer(ctx context.Context, uuid string) (*CreateResponse, error)`
- `ValidateServer(ctx context.Context, uuid string) (*CreateResponse, error)`
- `GetServerResources(ctx context.Context, uuid string) ([]Resource, error)`
- `GetServerDomains(ctx context.Context, uuid string) ([]ServerDomain, error)`

### Services

- `ListServices(ctx context.Context) ([]Service, error)`
- `GetService(ctx context.Context, uuid string) (*Service, error)`
- `CreateService(ctx context.Context, service Service) (*CreateResponse, error)`
- `UpdateService(ctx context.Context, uuid string, service Service) (*CreateResponse, error)`
- `DeleteService(ctx context.Context, uuid string, deleteConfigurations, deleteVolumes, dockerCleanup, deleteConnectedNetworks bool) (*CreateResponse, error)`
- `StartService(ctx context.Context, uuid string) (*CreateResponse, error)`
- `StopService(ctx context.Context, uuid string) (*CreateResponse, error)`
- `RestartService(ctx context.Context, uuid string) (*CreateResponse, error)`
- `ExecuteServiceCommand(ctx context.Context, uuid string, command string) (*CommandResponse, error)`

### Service Environment Variables

- `ListServiceEnvs(ctx context.Context, uuid string) ([]EnvironmentVariable, error)`
- `CreateServiceEnv(ctx context.Context, serviceUUID string, env EnvironmentVariable) (*CreateResponse, error)`
- `UpdateServiceEnv(ctx context.Context, serviceUUID string, env EnvironmentVariable) (*CreateResponse, error)`
- `DeleteServiceEnv(ctx context.Context, serviceUUID string, envUUID string) (*CreateResponse, error)`

### Teams

- `ListTeams(ctx context.Context) ([]Team, error)`
- `GetTeam(ctx context.Context, id string) (*Team, error)`
- `GetTeamMembers(ctx context.Context, id string) ([]TeamMember, error)`
- `GetCurrentTeam(ctx context.Context) (*Team, error)`
- `GetCurrentTeamMembers(ctx context.Context) ([]TeamMember, error)`

### Private Keys

- `ListPrivateKeys(ctx context.Context) ([]PrivateKey, error)`
- `GetPrivateKey(ctx context.Context, uuid string) (*PrivateKey, error)`
- `CreatePrivateKey(ctx context.Context, key PrivateKey) (*CreateResponse, error)`
- `DeletePrivateKey(ctx context.Context, uuid string) (*CreateResponse, error)`

### API Management

- `GetVersion(ctx context.Context) (string, error)`
- `EnableAPI(ctx context.Context) (*MessageResponse, error)`
- `DisableAPI(ctx context.Context) (*MessageResponse, error)`

## License

MIT 