package services

import (
	"api-golang-template/internal/entities"
	"context"

	"github.com/somphongph/lib-golang-packages/xlogger"
)

func (s *service) AddPlace(ctx context.Context) error {
	p := entities.Place{
		Name: "test",
	}

	err := s.storage.AddPlace(ctx, &p)
	if err != nil {
		xlogger.Errorf("error adding place: %v", err)

		return err
	}

	return nil
}
