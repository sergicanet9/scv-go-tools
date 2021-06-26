package business

import (
	"github.com/scanet9/scv-mongo-framework/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

//Service struct
type Service struct {
	DB   *mongo.Database
	Repo infrastructure.MongoRepository
}
