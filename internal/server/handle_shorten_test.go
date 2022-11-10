package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/defer-panic/url-shortener-api/internal/server"
	"github.com/defer-panic/url-shortener-api/internal/shorten"
	"github.com/defer-panic/url-shortener-api/internal/storage/shortening"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleShorten(t *testing.T) {
	t.Run("returns shortened URL for a given URL", func(t *testing.T) {
		const payload = `{"url": "https://www.google.com"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = server.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)

		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, recorder.Code)

		var resp struct {
			ShortURL string `json:"short_url"`
		}

		require.NoError(t, json.NewDecoder(recorder.Body).Decode(&resp), &resp)
		assert.NotEmpty(t, resp.ShortURL)
	})

	t.Run("returns error if URL is invalid", func(t *testing.T) {
		const payload = `{"url": "invalid"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = server.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)

		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		var httpErr *echo.HTTPError
		require.ErrorAs(t, handler(c), &httpErr)
		assert.Equal(t, http.StatusBadRequest, httpErr.Code)
		assert.Contains(t, httpErr.Message, "Field validation for 'URL' failed")
	})

	t.Run("returns error if identifier is already taken", func(t *testing.T) {
		const payload = `{"url": "https://www.google.com", "identifier": "google"}`

		var (
			shortener = shorten.NewService(shortening.NewInMemory())
			handler   = server.HandleShorten(shortener)
			recorder  = httptest.NewRecorder()
			request   = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
			e         = echo.New()
			c         = e.NewContext(request, recorder)
		)

		e.Validator = server.NewValidator()
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		require.NoError(t, handler(c))
		assert.Equal(t, http.StatusOK, recorder.Code)

		request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recorder = httptest.NewRecorder()
		c = e.NewContext(request, recorder)

		var httpErr *echo.HTTPError
		require.ErrorAs(t, handler(c), &httpErr)
		assert.Equal(t, http.StatusConflict, httpErr.Code)
		assert.Contains(t, httpErr.Message, model.ErrIdentifierExists.Error())
	})
}
