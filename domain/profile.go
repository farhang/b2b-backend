package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name                string
	LastName            string
	MobileNumber        string
	MobileNumberCompany string
	Position            string
	CompanyName         string
	User                User
	UserID              uint
	Plan                Plan
	PlanId              uint
}

type ProfileResponseDTO struct {
	UserID              uint
	Name                string `json:"name"`
	LastName            string `json:"last_name"`
	MobileNumber        string `json:"mobile_number"`
	Position            string `json:"position"`
	CompanyName         string `json:"company_name"`
	MobileNumberCompany string `json:"mobile_number_company"`
}

type ProfileDelivery interface {
	Fetch(ctx echo.Context) error
	Update(ctx echo.Context) error
}

type ProfileUseCase interface {
	Fetch(ctx context.Context) ([]Profile, error)
	Store(ctx context.Context, profile Profile) error
	StoreEmptyProfileByUserId(ctx context.Context, userId int) error
}

type ProfileRepository interface {
	Fetch(ctx context.Context) ([]Profile, error)
	Store(ctx context.Context, profile Profile) error
}
