package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/api/rest/request"
	"net/http"
)

func (h Handler) Login(c echo.Context) error {
	var req request.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrUnprocessableEntity
	}

	// TODO: check for auth
	// TODO: validate

	// TODO: make a transformer don't send bare domain to client
	d, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		// TODO: implement error catching.
		// TODO: implement response formatter
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, d)
}
