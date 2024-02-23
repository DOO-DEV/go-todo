package userhandler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-todo/internal/api/rest/request"
	cErr "go-todo/internal/errors"
	"net/http"
)

func (h Handler) Login(c echo.Context) error {
	var req request.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrUnprocessableEntity
	}
	// TODO: check for auth

	ctx := c.Request().Context()
	if err := h.validator.ValidateLoginUserRequest(ctx, &req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	// TODO: make a transformer don't send bare domain to client
	d, err := h.userSvc.Login(ctx, req)
	if err != nil {
		// TODO: implement error catching.
		// TODO: implement response formatter
		if errors.Is(err, cErr.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, d)
}
