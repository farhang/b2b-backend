package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

const (
	ADMIN  string = "ADMIN"
	MEMBER string = "MEMBER"
)

type UserRole struct {
	ID   uint
	Name string
}

var UserRoles = []UserRole{
	{ID: 1, Name: ADMIN},
	{ID: 2, Name: MEMBER},
}

type User struct {
	gorm.Model
	Password        string
	Email           string
	IsEmailVerified bool `gorm:"default:true"`
	Role            UserRole
	RoleID          uint
	IsActive        bool `gorm:"default:true"`
}

type VerificationCode struct {
	gorm.Model
	User      User
	UserId    uint
	Code      string
	ExpiresAt time.Time
}

type VerifyRequestDTO struct {
	Code string `json:"code"`
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
	GetById(ctx context.Context, id int) (*User, error)
	CheckIsUserDuplicatedByEmail(ctx context.Context, email string) bool
	GeneratePasswordHash(password string) (string, error)
	ComparePasswordHash(password string, hashedPassword string) bool
	Register(ctx context.Context, registerDTO RegisterRequestDTO) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	VerifyEmail(ctx context.Context, email string) error
	GenerateVerificationCodeNumber(length int) (string, error)
	StoreVerificationCode(ctx context.Context, code string, userId uint) error
	GetLatestVerificationCode(ctx context.Context, userId uint) (*VerificationCode, error)
	IsEmailVerified(ctx context.Context, email string) bool
	Update(ctx context.Context, user User) error
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	Register(ctx context.Context, registerDTO RegisterRequestDTO) error
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	StoreVerificationCode(ctx context.Context, emailVerification VerificationCode) error
	VerifyEmail(ctx context.Context, email string) error
	GetLatestVerificationCode(ctx context.Context, userId uint) (*VerificationCode, error)
	Update(ctx context.Context, user User) error
}

type UserHttpHandler interface {
	DepositToUserPlan(ctx echo.Context) error
	FetchUsers(ctx echo.Context) error
	GetById(ctx echo.Context) error
	GetMe(ctx echo.Context) error
	VerifyEmail(ctx echo.Context) error
}
