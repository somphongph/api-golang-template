package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) PostExample(c echo.Context) error {
	err := h.service.AddExample(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"test": "post example"})
}
