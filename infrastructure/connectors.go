package infrastructure

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB connects to MongoDB
func ConnectMongoDB(ctx context.Context, name string, connection string) (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(connection)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	database := client.Database(name)

	return database, nil
}

//ConnectPostgresDB connect to PostgresDB
func ConnectPostgresDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}
	return db, nil
}
