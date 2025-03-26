package storages

import (
	"api-golang-template/configs"
	"api-golang-template/internal/constants"
	"api-golang-template/internal/entities"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storager interface {
	GetExampleById(ctx context.Context) string
	AddExample(ctx context.Context, example *entities.Example) error
}

type storage struct {
	mongoExample *mongo.Collection
}

func NewStorages(mongoClient *mongo.Client) *storage {
	sect := configs.GetSecret()
	mongodb := mongoClient.Database(sect.Mongo.DbName)

	repo := storage{
		mongoExample: mongodb.Collection(constants.CollectionNameExamples),
	}

	return &repo
}
