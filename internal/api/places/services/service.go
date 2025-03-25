package services

import (
	"api-golang-template/internal/api/places/storages"
	"context"
)

type Servicer interface {
	GetPlace(ctx context.Context) string
	AddPlace(ctx context.Context) error
}

type service struct {
	storage storages.Storager
}

func NewServices(storage storages.Storager) *service {
	return &service{storage: storage}
}
