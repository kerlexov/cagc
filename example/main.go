package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kerlexov/cgc"
)

func main() {
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
	client, err := cgc.NewClient(fmt.Printf("%s/api/v1", serverUrl), token)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

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

	// Example: Create a new PostgreSQL database
	fmt.Println("\nCreating a PostgreSQL database...")
	newDB := cgc.Database{
		ProjectUUID:      "your-project-uuid",
		ServerUUID:       "your-server-uuid",
		EnvironmentName:  "development",
		Name:             "example-postgres",
		PostgresUser:     "postgres",
		PostgresPassword: "secure-password",
		PostgresDB:       "exampledb",
		IsPublic:         false,
	}

	// Comment this out to actually create the database
	// resp, err := client.CreatePostgresDatabase(context.Background(), newDB)
	// if err != nil {
	//     log.Fatalf("Error creating database: %v", err)
	// }
	// fmt.Printf("Database created with UUID: %s\n", resp.UUID)

	fmt.Println("\nExample complete!")
}
