package common

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrNotFound                       = errors.New("your requested Item is not found")
	ErrEmailDuplication               = errors.New("the email is already exists")
	ErrEmailIsNotVerified             = errors.New("the email is not verified")
	ErrInvalidCredential              = errors.New("invalid credential")
	ErrEmailVerificationCodeIsInValid = errors.New("the verification code is invalid")
	ErrEmailIsNotExists               = errors.New("the email is not exists")
)

func ErrHttpBadRequest(err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, err)
}

func ErrHttpConflict(err error) error {
	return echo.NewHTTPError(http.StatusConflict, err)
}
func ErrHttpNotFound(err error) error {
	return echo.NewHTTPError(http.StatusNotFound, err)
}

func ErrHttpUnauthorized(err error) error {
	return echo.NewHTTPError(http.StatusUnauthorized, err)
}
func ErrHttpUnprocessableEntity(err error) error {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
}
