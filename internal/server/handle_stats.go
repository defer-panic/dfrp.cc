package server

import (
	"context"
	"log"

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
			log.Printf("failed to get shortening: %v", err)
			return echo.NewHTTPError(500, "failed to get shortening")
		}

		return c.JSON(200, shortening)
	}
}
