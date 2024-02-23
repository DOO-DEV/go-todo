package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/api/rest/request"
	"net/http"
)

func (h Handler) Register(c echo.Context) error {
	var req request.RegisterUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrUnprocessableEntity
	}
	// TODO: check for auth

	ctx := c.Request().Context()
	if err := h.validator.ValidateRegisterUserRequest(ctx, &req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	// TODO: make a transformer don't send bare domain to client
	d, err := h.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		// TODO: implement error catching.
		// TODO: implement response formatter
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, d)
}
