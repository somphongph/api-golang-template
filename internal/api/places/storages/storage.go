package storages

import (
	"api-golang-template/configs"
	"api-golang-template/internal/constants"
	"api-golang-template/internal/entities"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storager interface {
	GetPlaceById(ctx context.Context) string
	AddPlace(ctx context.Context, place *entities.Place) error
}

type storage struct {
	mongoPlace *mongo.Collection
}

func NewStorages(mongoClient *mongo.Client) *storage {
	sect := configs.GetSecret()
	mongodb := mongoClient.Database(sect.Mongo.DbName)

	repo := storage{
		mongoPlace: mongodb.Collection(constants.CollectionNamePlaces),
	}

	return &repo
}
