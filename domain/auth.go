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

type SendOTPRequestDTO struct {
	MobileNumber string `json:"mobile_number"`
}

type LoginWithOTPDTO struct {
	MobileNumber string `json:"mobile_number"`
	Code         string `json:"code"`
}

type RegisterRequestDTO struct {
	Name                string `json:"name"`
	LastName            string `json:"last_name"`
	Email               string `json:"email" validate:"required,email"`
	Password            string `json:"password"`
	ConfirmPassword     string `json:"confirm_password" validate:"required"`
	MobileNumber        string `json:"mobile_number"`
	MobileNumberCompany string `json:"mobile_number_company"`
	Position            string `json:"position"`
	CompanyName         string `json:"company_name"`
}

type ResetPasswordRequestDTO struct {
	UserID   int    `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required,email"`
}

type JwtCustomClaims struct {
	jwt.StandardClaims
	UserId int
	Role   uint
}

type AuthHttpHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type AuthUseCase interface {
	Login(c context.Context, dto LoginRequestDTO) (*string, error)
	Register(c context.Context, dto RegisterRequestDTO) error
	SendOTP(c context.Context, dto SendOTPRequestDTO) error
	LoginWithOTP(c context.Context, dto LoginWithOTPDTO) (*string, error)
	GenerateToken(claims JwtCustomClaims) (string, error)
	ResetPassword(ctx context.Context, id int, newPassword string) error
}
