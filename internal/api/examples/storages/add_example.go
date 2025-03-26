package storages

import (
	"api-golang-template/internal/entities"
	"context"
	"time"
)

func (r *storage) AddExample(ctx context.Context, e *entities.Example) error {

	e.CreatedBy = "test"
	e.CreatedAt = time.Now()
	e.UpdatedBy = "test"
	e.UpdatedAt = time.Now()

	if _, err := r.mongoExample.InsertOne(ctx, e); err != nil {
		return err
	}

	return nil
}
