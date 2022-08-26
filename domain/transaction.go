package domain

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

var (
	ErrNotEnoughAssetAmount = errors.New("you want to withdraw more than you have")
)

const (
	RialDeposit    string = "RIAL_DEPOSIT"
	RialWithdraw   string = "RIAL_WITHDRAW"
	TetherDeposit  string = "TETHER_DEPOSIT"
	TetherWithdraw string = "WITHDRAW"
	TetherProfit   string = "PROFIT_TETHER"
	BuyTether      string = "BUY_TETHER"
	SellTether     string = "SELL_TETHER"
)

var TransactionTypes = []TransactionType{
	{ID: 1, Name: RialDeposit},
	{ID: 2, Name: RialWithdraw},
	{ID: 3, Name: TetherDeposit},
	{ID: 4, Name: TetherWithdraw},
	{ID: 5, Name: TetherProfit},
	{ID: 6, Name: BuyTether},
	{ID: 7, Name: SellTether},
}

type TransactionType struct {
	ID   uint
	Name string
}

func GetTransactionIdByName(name string) *uint {
	for _, tt := range TransactionTypes {
		if tt.Name == name {
			return &tt.ID
		}
	}
	return nil
}

func GetTransactionNameById(id uint) *string {
	for _, tt := range TransactionTypes {
		if tt.ID == id {
			return &tt.Name
		}
	}
	return nil
}

type Transaction struct {
	gorm.Model
	Amount            float64 `gorm:"check:amount >= 0"`
	TransactionType   TransactionType
	TransactionTypeID uint
	Description       string
	User              User
	UserId            uint
}

type UserPlanTransaction struct {
	gorm.Model
	Transaction   Transaction
	TransactionID uint
	UserPlan      UserPlan
	UserPlanID    uint
}

type TransactionResponseDTO struct {
	CreatedAt       time.Time `json:"created_at"`
	Amount          float64   `json:"amount"`
	TransactionType string    `json:"type"`
	Email           string    `json:"email,omitempty"`
}

type StoreTransactionRequestDTO struct {
	Amount            float64 `json:"amount"`
	Description       string  `json:"description"`
	TransactionTypeId uint    `json:"transaction_type_id"`
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
	MyTransactions(ctx echo.Context) error
}

type TransactionUseCase interface {
	Fetch(ctx context.Context) ([]Transaction, error)
	FetchByUserId(ctx context.Context, userId int) ([]Transaction, error)
	Store(ctx context.Context, userId uint, transaction StoreTransactionRequestDTO) error
	GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error)
}

type TransactionRepository interface {
	Store(ctx context.Context, transaction Transaction) error
	Fetch(ctx context.Context) ([]Transaction, error)
	FetchByUserId(ctx context.Context, userId int) ([]Transaction, error)
	GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error)
}
