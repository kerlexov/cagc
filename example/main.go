package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kerlexov/cagc"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get API token from environment variable
	token := os.Getenv("COOLIFY_API_TOKEN")
	if token == "" {
		log.Fatal("COOLIFY_API_TOKEN environment variable not set")
	}

	serverUrl := os.Getenv("COOLIFY_API_HOST")
	if serverUrl == "" {
		log.Fatal("COOLIFY_API_HOST environment variable not set")
	}

	// Create a new client
	client, err := cagc.NewClient(serverUrl, token)
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

	// Example: Create a new PostgreSQL database (commented out to prevent accidental creation)
	/*
		fmt.Println("\nCreating a PostgreSQL database...")
		newDB := cagc.Database{
			ProjectUUID:      "your-project-uuid",
			ServerUUID:       "your-server-uuid",
			EnvironmentName:  "development",
			Name:             "example-postgres",
			PostgresUser:     "postgres",
			PostgresPassword: "secure-password",
			PostgresDB:       "exampledb",
			IsPublic:         false,
		}

		resp, err := client.CreatePostgresDatabase(context.Background(), newDB)
		if err != nil {
			log.Fatalf("Error creating database: %v", err)
		}
		fmt.Printf("Database created with UUID: %s\n", resp.UUID)
	*/

	// Example: Create a new project (commented out to prevent accidental creation)
	/*
		fmt.Println("\nCreating a new project...")
		newProject := cagc.Project{
			Name:        "Example Project",
			Description: "This is an example project created via the API",
		}

		projectResp, err := client.CreateProject(context.Background(), newProject)
		if err != nil {
			log.Fatalf("Error creating project: %v", err)
		}
		fmt.Printf("Project created with UUID: %s\n", projectResp.UUID)
	*/

	fmt.Println("\nExample complete!")
}
