package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/defer-panic/url-shortener-api/internal/config"
	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/defer-panic/url-shortener-api/internal/shorten"
	"github.com/labstack/echo/v4"
	. "github.com/samber/mo"
)

type shortener interface {
	Shorten(context.Context, model.ShortenInput) (*model.Shortening, error)
}

type shortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty" validate:"omitempty,alphanum"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Message  string `json:"message,omitempty"`
}

func HandleShorten(shortener shortener) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req shortenRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		if err := c.Validate(req); err != nil {
			return err
		}

		identifier := None[string]()
		if strings.TrimSpace(req.Identifier) != "" {
			identifier = Some(req.Identifier)
		}

		input := model.ShortenInput{
			RawURL:     req.URL,
			Identifier: identifier,
		}

		shortening, err := shortener.Shorten(c.Request().Context(), input)
		if err != nil {
			var (
				status int
				msg    = err.Error()
			)
			switch {
			case errors.Is(err, model.ErrInvalidURL):
				status = http.StatusBadRequest
			case errors.Is(err, model.ErrIdentifierExists):
				status = http.StatusConflict
			default:
				log.Printf("error shortening url %q: %v", req.URL, err)
				return err
			}

			return echo.NewHTTPError(status, msg)
		}

		shortURL, err := shorten.PrependBaseURL(config.Get().BaseURL, shortening.Identifier)
		if err != nil {
			log.Printf("error generating full url for %q: %v", shortening.Identifier, err)
			return err
		}

		return c.JSON(
			http.StatusOK,
			shortenResponse{ShortURL: shortURL},
		)
	}
}
