package auth

import (
	"time"

	"github.com/defer-panic/url-shortener-api/internal/config"
	"github.com/defer-panic/url-shortener-api/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func MakeJWT(user model.User) (string, error) {
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "defer panic",
			IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Get().Auth.JWTSecretKey))
}

func VerifyJWT(src string, c echo.Context) (any, error) {
	token, err := jwt.ParseWithClaims(
		src,
		&model.UserClaims{},
		func(_ *jwt.Token) (any, error) { return []byte(config.Get().Auth.JWTSecretKey), nil },
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, model.ErrInvalidToken
	}

	return token, nil
}
