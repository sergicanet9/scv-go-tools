package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to mongo database
func ConnectDB(name string, connection string) *mongo.Database {

	clientOptions := options.Client().ApplyURI(connection)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err)
	}

	database := client.Database(name)

	return database
}
