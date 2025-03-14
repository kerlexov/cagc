# Coolify API Go Client

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

	// You can also manage databases
	databases, err := client.ListDatabases(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, db := range databases {
		fmt.Printf("Database: %s (%s)\n", db.Name, db.UUID)
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

### Environment Variables

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

## License

MIT 