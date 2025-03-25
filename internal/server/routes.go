package server

import (
	places "api-places/internal/api/places/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()

	secureConfig := middleware.SecureConfig{
		ContentTypeNosniff: "nosniff",
		ReferrerPolicy:     "same-origin",
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.SecureWithConfig(secureConfig),
	)

	// health
	healthz(e)

	apiRoot := e.Group("/v1")

	// places
	places.NewRoutes(apiRoot, s.mongo)

	return e
}
