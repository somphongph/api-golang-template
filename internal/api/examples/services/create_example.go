package services

import (
	"api-golang-template/internal/entities"
	"context"

	"github.com/tripgator/lib-golang-packages/xlogger"
)

func (s *service) AddExample(ctx context.Context) error {
	p := entities.Example{
		Name: "test",
	}

	err := s.storage.AddExample(ctx, &p)
	if err != nil {
		xlogger.Errorf("error adding example: %v", err)

		return err
	}

	return nil
}
