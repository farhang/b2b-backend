package domain

import (
	"context"
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Amount int `json:"tether_balance"`
	UserID int
	User   User
}

type AssetRepository interface {
	Store(ctx context.Context, asset Asset) error
	GetByUserId(ctx context.Context, UserId int) (Asset, error)
	Fetch(ctx context.Context) error
	Update(ctx context.Context, asset Asset) error
}

type AssetUseCase interface {
	Store(ctx context.Context, asset Asset) error
	GetAmountByUserId(ctx context.Context, userId int) (int, error)
	IncreaseAmount(ctx context.Context, userId int, amount int) error
	DecreaseAmount(ctx context.Context, userId int, amount int) error
}
