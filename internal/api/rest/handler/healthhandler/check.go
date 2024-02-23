package healthhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func (h Handler) Health(c echo.Context) error {
	token := c.QueryParam("token")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	services := strings.Split(c.QueryParam("services"), ",")

	for _, s := range services {
		if err := h.healthSvc.HealthCheck(c.Request().Context(), token, s); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, echo.Map{
				"service": s,
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "all services are healthy",
	})
}
