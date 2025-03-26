package handlers

import "template-api-please-replace/internal/api/examples/services"

type handlers struct {
	service services.Servicer
}

func NewHandlers(service services.Servicer) *handlers {
	return &handlers{
		service: service,
	}
}
