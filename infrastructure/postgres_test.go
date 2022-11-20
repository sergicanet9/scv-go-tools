package infrastructure

import (
	"fmt"
	"log"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
)

const postgresInternalPort = "5432/tcp"

// TestConnectPostgresDB_Ok checks that ConnectPostgresDB does not return an error when a valid connection string is provided
func TestConnectPostgresDB_Ok(t *testing.T) {
	// Arrange
	connection := setupPostgres()

	// Act
	_, err := ConnectPostgresDB(connection)

	// Assert
	assert.Equal(t, nil, err)

}

// TestConnectPostgresDB_InvalidConnection checks that ConnectPostgresDB returns an error when an invalid connection string is provided
func TestConnectPostgresDB_InvalidConnection(t *testing.T) {
	// Arrange
	expectedError := "missing \"=\" after \"invalid-connection\" in connection info string\""

	// Act
	_, err := ConnectPostgresDB("invalid-connection")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

// TestOpenSQL_BadDriver checks that openSQL return an error when an invalid driver is provided
func TestOpenSQL_BadDriver(t *testing.T) {
	// Arrange
	expectedError := "sql: unknown driver \"bad-driver\" (forgotten import?)"

	// Act
	_, err := openSQL("bad-driver", "")

	// Assert
	assert.Equal(t, expectedError, err.Error())
}

func setupPostgres() string {
	// Uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	// Pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.6",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", "test"),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", "test"),
			fmt.Sprintf("POSTGRES_DB=%s", "test"),
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
	connectionString := fmt.Sprintf("host=localhost port=%s dbname=%s user=%s password=%s sslmode=disable", resource.GetPort(postgresInternalPort), "test", "test", "test")
	return connectionString
}
