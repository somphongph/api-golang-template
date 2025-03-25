package handlers

import "api-places/internal/api/places/services"

type handlers struct {
	service services.Servicer
}

func NewHandlers(service services.Servicer) *handlers {
	return &handlers{
		service: service,
	}
}
