package infrastructure

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

const mongoInternalPort = "27017/tcp"

// TestConnectMongoDB_Ok checks that ConnectMongoDB does not return an error when a valid connection string is provided
func TestConnectMongoDB_Ok(t *testing.T) {
	// Arrange
	connection := setupMongo()

	// Act
	_, err := ConnectMongoDB(context.Background(), "test", connection)

	// Assert
	assert.Equal(t, nil, err)

}

// TestConnectMongoDB_InvalidConnection checks that ConnectMongoDB returns an error when an invalid connection string is provided
func TestConnectMongoDB_InvalidConnection(t *testing.T) {
	// Arrange
	expectedError := "error parsing uri: scheme must be \"mongodb\" or \"mongodb+srv\""

	// Act
	_, err := ConnectMongoDB(context.Background(), "test", "invalid-connection")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

func setupMongo() string {
	// Uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	// Pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "3.0",
		Env: []string{
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("could not start resource: %s", err)
	}
	connectionString := fmt.Sprintf("mongodb://localhost:%s", resource.GetPort(mongoInternalPort))
	return connectionString
}
