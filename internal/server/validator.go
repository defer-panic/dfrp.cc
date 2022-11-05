package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type requestValidator struct {
	v *validator.Validate
}

func NewValidator() *requestValidator {
	return &requestValidator{
		v: validator.New(),
	}
}

func (v *requestValidator) Validate(i any) error {
	if err := v.v.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
