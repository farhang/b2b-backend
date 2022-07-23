package domain

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginRequestDTO struct {
	Email    string
	Password string
}
type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

type LoginResponseDTO struct {
	Data    LoginResponseData `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
}

type RegisterRequestDTO struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type ResetPasswordRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,email"`
}

type JwtCustomClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthHttpHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type AuthUseCase interface {
	Login(c context.Context, loginUserDTO LoginRequestDTO) (*string, error)
	Register(c context.Context, registerUserDTO RegisterRequestDTO) error
	GenerateToken(claims JwtCustomClaims) (string, error)
	ResetPassword(ctx context.Context, email string, newPassword string) error
}
