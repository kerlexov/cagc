# Coolify API Go Client

## Version 0.0.1

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
	client, err := coolify.NewClient("https://app.coolify.io/api/v1", "your-api-token")
	if err != nil {
		log.Fatal(err)
	}

	// List all applications
	applications, err := client.ListApplications(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, app := range applications {
		fmt.Printf("Application: %s (%s)\n", app.Name, app.UUID)
	}

	// You can also manage servers, services, databases, and more
	servers, err := client.ListServers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Printf("Server: %s (%s)\n", server.Name, server.UUID)
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