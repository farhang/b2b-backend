package domain

import (
	"context"
	"database/sql/driver"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionType string

var (
	ErrNotEnoughAssetAmount = errors.New("you want to withdraw more than you have")
)

const (
	DEPOSIT  TransactionType = "DEPOSIT"
	WITHDRAW TransactionType = "WITHDRAW"
	PROFIT   TransactionType = "PROFIT"
)

func (tt *TransactionType) Scan(value interface{}) error {
	*tt = TransactionType(value.(string))
	return nil
}

func (tt TransactionType) Value() (driver.Value, error) {
	return string(tt), nil
}

type Transaction struct {
	gorm.Model
	Amount          float64         `gorm:"check:amount >= 0"`
	TransactionType TransactionType `sql:"transaction_type"`
	User            User
	UserId          uint
}

type TransactionResponseDTO struct {
	Amount          float64         `json:"amount"`
	TransactionType TransactionType `json:"type"`
}

type DepositRequestDTO struct {
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type WithDrawRequestDTO struct {
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type ProfitRequestDTO struct {
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
}

type TransactionHttpHandler interface {
	Fetch(ctx echo.Context) error
	Deposit(ctx echo.Context) error
	WithDraw(ctx echo.Context) error
	Profit(ctx echo.Context) error
}

type TransactionUseCase interface {
	Fetch(ctx context.Context) error
	FetchByUserId(ctx context.Context, userId int) ([]Transaction, error)
	Deposit(ctx context.Context, deposit DepositRequestDTO) error
	WithDraw(ctx context.Context, withdraw WithDrawRequestDTO) error
	Profit(ctx context.Context, profit ProfitRequestDTO) error
	GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error)
}

type TransactionRepository interface {
	Store(ctx context.Context, transaction Transaction) error
	Fetch(ctx context.Context) ([]Transaction, error)
	FetchByUserId(ctx context.Context, userId int) ([]Transaction, error)
	GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error)
	Deposit(ctx context.Context, transaction Transaction, userId int) error
}
