package services

import (
	"context"
	"template-api-please-replace/internal/api/examples/storages"
)

type Servicer interface {
	GetExample(ctx context.Context) string
	AddExample(ctx context.Context) error
}

type service struct {
	storage storages.Storager
}

func NewServices(storage storages.Storager) *service {
	return &service{storage: storage}
}
