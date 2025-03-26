package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tripgator/lib-golang-packages/xres"
)

func (h *handlers) GetExampleList(c echo.Context) error {
	resp := h.service.GetExample(c.Request().Context())

	return c.JSON(http.StatusOK, xres.Success(resp))
}
