package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository struct of a mongo repository
type MongoRepository struct {
	collection  *mongo.Collection
	constructor func() interface{} // constructor of the repository target entity
}

// NewMongoRepository creates a mongodb repository
func NewMongoRepository(collection *mongo.Collection, constructor func() interface{}) *MongoRepository {
	return &MongoRepository{
		collection,
		constructor,
	}
}

// Create creates an entity in the repository's collection
func (r *MongoRepository) Create(ctx context.Context, entity interface{}) primitive.ObjectID {
	result, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}

// Get gets the documents mathing the filter in the repository's collection
func (r *MongoRepository) Get(ctx context.Context, filter primitive.M) []interface{} {
	var result []interface{}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		entry := r.constructor()
		if err := cur.Decode(entry); err != nil {
			panic(err)
		}
		result = append(result, entry)
	}

	return result
}

// GetByID get the document with the specified ID in the repository's collection
func (r *MongoRepository) GetByID(ctx context.Context, ID string) interface{} {
	result := r.constructor()

	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": _id}
	err = r.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		panic(err)
	}

	return result
}

// Update updates the document with the specified ID in the repository's collection
func (r *MongoRepository) Update(ctx context.Context, ID string, entity interface{}) {
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": _id}
	update := bson.M{"$set": entity}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	if result.ModifiedCount < 1 {
		panic(mongo.ErrNoDocuments)
	}
}

// Delete deletes the document with the specified ID in the repository's collection
func (r *MongoRepository) Delete(ctx context.Context, ID string) {
	_id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": _id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	if result.DeletedCount < 1 {
		panic(mongo.ErrNoDocuments)
	}
}
