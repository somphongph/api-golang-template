package routes

import (
	"api-golang-template/internal/api/examples/handlers"
	"api-golang-template/internal/api/examples/services"
	"api-golang-template/internal/api/examples/storages"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRoutes(apiRoot *echo.Group, mongo *mongo.Client) {
	repo := storages.NewStorages(mongo)
	service := services.NewServices(repo)

	// examples
	h := handlers.NewHandlers(service)
	g := apiRoot.Group("/examples")
	g.GET("", h.GetExampleList)
	g.GET("/:id", h.GetExampleItem)
	g.POST("", h.PostExample)
	g.PUT("/:id", h.PutExample)
}
