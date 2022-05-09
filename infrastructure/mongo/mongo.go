package infrastructure

import (
	"context"

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
