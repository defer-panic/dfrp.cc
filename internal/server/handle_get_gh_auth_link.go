package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type gitHubAuthLinkProvider interface {
	GitHubAuthLink() string
}

func HandleGetGitHubAuthLink(provider gitHubAuthLinkProvider) echo.HandlerFunc {
	return func(c echo.Context) error {
		link := provider.GitHubAuthLink()
		return c.JSON(http.StatusOK, map[string]string{"link": link})
	}
}
