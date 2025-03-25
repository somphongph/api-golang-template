package services

import "context"

func (s *service) GetPlace(ctx context.Context) string {
	return s.storage.GetPlaceById(ctx)
}
