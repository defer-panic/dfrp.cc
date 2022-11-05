package server_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/defer-panic/url-shortener-api/internal/server"
	"github.com/defer-panic/url-shortener-api/internal/shorten"
	"github.com/defer-panic/url-shortener-api/internal/storage"
	"github.com/labstack/echo/v4"
	. "github.com/samber/mo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleRedirect(t *testing.T) {
	t.Run("redirects to original URL", func(t *testing.T) {
		const (
			url        = "https://www.google.com"
			identifier = "google"
		)

		var (
			redirecter = shorten.NewService(storage.NewInMemory())
			handler    = server.HandleRedirect(redirecter)
			recorder   = httptest.NewRecorder()
			request    = httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
			e          = echo.New()
			c          = e.NewContext(request, recorder)
		)

		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		_, err := redirecter.Shorten(
			context.Background(),
			model.ShortenInput{
				RawURL:     url,
				Identifier: Some(identifier),
			},
		)
		require.NoError(t, err)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusMovedPermanently, recorder.Code)
		assert.Equal(t, url, recorder.Header().Get("Location"))
	})

	t.Run("returns 404 if identifier is not found", func(t *testing.T) {
		const (
			url        = "https://www.google.com"
			identifier = "google"
		)

		var (
			redirecter = shorten.NewService(storage.NewInMemory())
			handler    = server.HandleRedirect(redirecter)
			recorder   = httptest.NewRecorder()
			request    = httptest.NewRequest(http.MethodGet, "/"+identifier, nil)
			e          = echo.New()
			c          = e.NewContext(request, recorder)
		)

		c.SetPath("/:identifier")
		c.SetParamNames("identifier")
		c.SetParamValues(identifier)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
