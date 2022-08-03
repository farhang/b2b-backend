package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Password        string
	Email           string
	IsEmailVerified bool
}

type EmailVerification struct {
	gorm.Model
	Email     string
	Code      int
	ExpiresAt time.Time
}

type VerifyRequestDTO struct {
	Code int `json:"code"`
}

type StoreUserRequestDTO struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponseDTO struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type UserUseCase interface {
	Fetch(ctx context.Context) ([]User, error)
	Store(ctx context.Context, user StoreUserRequestDTO) error
	GetById(ctx context.Context, id int) (*User, error)
	CheckIsUserDuplicatedByEmail(ctx context.Context, email string) bool
	GeneratePasswordHash(password string) (string, error)
	ComparePasswordHash(password string, hashedPassword string) bool
	GetByEmail(ctx context.Context, email string) (*User, error)
	VerifyEmail(ctx context.Context, email string) error
	GenerateVerificationCodeNumber(length int) (int, error)
	StoreEmailVerificationCode(ctx context.Context, email string) error
	GetLatestEmailVerification(ctx context.Context, email string) (*EmailVerification, error)
	IsEmailVerified(ctx context.Context, email string) bool
	Update(ctx context.Context, user User) error
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	Store(ctx context.Context, user *User) error
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	StoreEmailVerificationCode(ctx context.Context, emailVerification EmailVerification) error
	VerifyEmail(ctx context.Context, email string) error
	GetLatestEmailVerification(ctx context.Context, email string) (*EmailVerification, error)
	Update(ctx context.Context, user User) error
}

type UserHttpHandler interface {
	FetchUsers(ctx echo.Context) error
	Store(ctx echo.Context) error
	GetById(ctx echo.Context) error
	GetMe(ctx echo.Context) error
	VerifyEmail(ctx echo.Context) error
}
