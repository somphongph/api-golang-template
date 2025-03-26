package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) PutExample(c echo.Context) error {
	// resp := h.service.GetHello(c.Request().Context())

	return c.JSON(http.StatusOK, map[string]string{"test": "put example"})
}
