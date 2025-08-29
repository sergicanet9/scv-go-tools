package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/sergicanet9/scv-go-tools/v3/wrappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// ConnectMongoDB opens a connection to the MongoDB and ensures that the db is reachable
func ConnectMongoDB(ctx context.Context, dsn string, monitor *event.CommandMonitor) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(dsn)
	if monitor != nil {
		clientOptions.SetMonitor(monitor)
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("an unexpected error happened while opening the connection: %s", err)
	}

	return pingMongo(ctx, client, dsn)
}

func pingMongo(ctx context.Context, client *mongo.Client, dsn string) (*mongo.Database, error) {
	cs, err := connstring.ParseAndValidate(dsn)
	if err != nil {
		return nil, fmt.Errorf("DSN not valid: %s", err)
	}
	return client.Database(cs.Database), client.Ping(ctx, nil)
}

// MongoRepository struct of a mongo repository
// Implements the Repository interface
type MongoRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
	Target     interface{}
}

// Create creates an entity in the repository's collection
func (r *MongoRepository) Create(ctx context.Context, entity interface{}) (string, error) {
	result, err := r.Collection.InsertOne(ctx, entity)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Get gets the documents mathing the filter in the repository's collection
func (r *MongoRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var result []interface{}

	var skip64, take64 int64
	if skip != nil {
		skip64 = int64(*skip)
	}
	if take != nil {
		take64 = int64(*take)
	}
	cur, err := r.Collection.Find(ctx, filter, &options.FindOptions{Skip: &skip64, Limit: &take64})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		entry := reflect.New(reflect.TypeOf(r.Target)).Interface()
		if err := cur.Decode(entry); err != nil {
			cur.Close(ctx)
			return nil, err
		}
		result = append(result, entry)
	}

	if len(result) < 1 {
		return nil, wrappers.NewNonExistentErr(mongo.ErrNoDocuments)
	}

	return result, nil
}

// GetByID get the document with the specified ID in the repository's collection
func (r *MongoRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	result := reflect.New(reflect.TypeOf(r.Target)).Interface()

	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	err = r.Collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = wrappers.NewNonExistentErr(err)
		}
		return nil, err
	}

	return result, nil
}

// Update updates the document with the specified ID in the repository's collection
func (r *MongoRepository) Update(ctx context.Context, ID string, entity interface{}) error {
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": entity}
	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount < 1 && result.UpsertedCount < 1 {
		return wrappers.NewNonExistentErr(mongo.ErrNoDocuments)
	}
	return nil
}

// Delete deletes the document with the specified ID in the repository's collection
func (r *MongoRepository) Delete(ctx context.Context, ID string) error {
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}
	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return wrappers.NewNonExistentErr(mongo.ErrNoDocuments)
	}
	return nil
}
