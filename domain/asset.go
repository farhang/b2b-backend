package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Amount float64 `gorm:"check:amount >= 0" json:"tether_balance"`
	UserID int
	User   User
}
type AssetResponseDTO struct {
	Amount      float64 `json:"amount"`
	TotalProfit float64 `json:"total_profit"`
}

type AssetDelivery interface {
	GetByUserId(ctx echo.Context) error
}

type AssetRepository interface {
	Store(ctx context.Context, asset Asset) error
	GetByUserId(ctx context.Context, UserId int) (Asset, error)
	Fetch(ctx context.Context) error
	Update(ctx context.Context, asset Asset) error
}

type AssetUseCase interface {
	Store(ctx context.Context, asset Asset) error
	GetAmountByUserId(ctx context.Context, userId int) (float64, error)
	IncreaseAmount(ctx context.Context, userId int, amount float64) error
	DecreaseAmount(ctx context.Context, userId int, amount float64) error
}
