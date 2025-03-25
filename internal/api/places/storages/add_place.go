package storages

import (
	"api-golang-template/internal/entities"
	"context"
	"time"
)

func (r *storage) AddPlace(ctx context.Context, e *entities.Place) error {

	e.CreatedBy = "test"
	e.CreatedAt = time.Now()
	e.UpdatedBy = "test"
	e.UpdatedAt = time.Now()

	if _, err := r.mongoPlace.InsertOne(ctx, e); err != nil {
		return err
	}

	return nil
}
