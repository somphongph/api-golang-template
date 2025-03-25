package routes

import (
	"api-places/internal/api/places/handlers"
	"api-places/internal/api/places/services"
	"api-places/internal/api/places/storages"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRoutes(apiRoot *echo.Group, mongo *mongo.Client) {
	repo := storages.NewStorages(mongo)
	service := services.NewServices(repo)

	// places
	h := handlers.NewHandlers(service)
	g := apiRoot.Group("/places")
	g.GET("", h.GetPlaceList)
	g.GET("/:id", h.GetPlaceItem)
	g.POST("", h.PostPlace)
	g.PUT("/:id", h.PutPlace)
}
