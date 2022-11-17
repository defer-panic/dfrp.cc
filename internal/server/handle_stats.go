package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/labstack/echo/v4"
)

type shorteningProvider interface {
	Get(ctx context.Context, identifier string) (*model.Shortening, error)
}

func HandleStats(provider shorteningProvider) echo.HandlerFunc {
	return func(c echo.Context) error {
		identifier := c.Param("identifier")
		shortening, err := provider.Get(c.Request().Context(), identifier)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			log.Printf("failed to get shortening: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to get shortening")
		}

		return c.JSON(http.StatusOK, shortening)
	}
}
