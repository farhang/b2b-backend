package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserPlanTransaction struct {
	gorm.Model
	Transaction   Transaction
	TransactionID uint
	UserPlan      UserPlan
	UserPlanID    uint
}

type StorePlanTransactionDTO struct {
	PlanID            uint    `json:"plan_id"`
	Amount            float64 `json:"amount"`
	TransactionTypeID uint    `json:"transaction_type_id"`
	Description       string  `json:"description"`
}

type UserPlanTransactionDelivery interface {
	Store(ctx echo.Context) error
}

type UserPlanTransactionUseCase interface {
	Store(ctx context.Context, transaction UserPlanTransaction) error
}

type UserPlanTransactionRepository interface {
	Store(ctx context.Context, transaction *UserPlanTransaction) error
}
