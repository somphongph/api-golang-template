package handlers

import "api-golang-template/internal/api/examples/services"

type handlers struct {
	service services.Servicer
}

func NewHandlers(service services.Servicer) *handlers {
	return &handlers{
		service: service,
	}
}
