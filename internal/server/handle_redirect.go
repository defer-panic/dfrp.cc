package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/labstack/echo/v4"
)

type redirectResponse struct {
	Message string `json:"message,omitempty"`
}

type redirecter interface {
	Redirect(ctx context.Context, identifier string) (string, error)
}

func HandleRedirect(redirecter redirecter) echo.HandlerFunc {
	return func(c echo.Context) error {
		identifier := c.Param("identifier")

		redirectURL, err := redirecter.Redirect(c.Request().Context(), identifier)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return c.JSON(
					http.StatusNotFound,
					redirectResponse{Message: err.Error()},
				)
			}

			log.Printf("error getting redirect url for %q: %v", identifier, err)
			return err
		}

		return c.Redirect(http.StatusMovedPermanently, redirectURL)
	}
}
