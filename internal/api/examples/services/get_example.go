package services

import "context"

func (s *service) GetExample(ctx context.Context) string {
	return s.storage.GetExampleById(ctx)
}
