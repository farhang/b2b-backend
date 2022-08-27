package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

const (
	PENDING  string = "PENDING"
	ACTIVE   string = "ACTIVE"
	INACTIVE string = "INACTIVE"
	EXPIRED  string = "EXPIRED"
	SETTLED  string = "SETTLED"
)

type UserPlanStatus struct {
	ID   uint
	Name string
}

var UserPlanStatuses = []UserPlanStatus{
	{ID: 1, Name: PENDING},
	{ID: 2, Name: EXPIRED},
	{ID: 3, Name: SETTLED},
	{ID: 4, Name: ACTIVE},
	{ID: 5, Name: INACTIVE},
}

type UserPlan struct {
	gorm.Model
	User             User
	UserID           uint
	Plan             Plan
	PlanID           uint
	Amount           float64
	UserPlanStatus   UserPlanStatus
	UserPlanStatusId uint
	StartedAt        time.Time
	ExpiresAt        time.Time
}

type StoreUserPlanRequestDTO struct {
	UserID    uint
	PlanID    uint
	Amount    int
	StartedAt time.Time
	ExpiresAt time.Time
}

type UpdateUserPlanDTO struct {
	Amount           int  `json:"amount"`
	UserPlanStatusId uint `json:"user_plan_status_id"`
}

type StoreUserPlanTransactionDTO struct {
	Amount            float64 `json:"amount"`
	TransactionTypeID uint    `json:"transaction_type_id"`
	Description       string  `json:"description"`
}

type UserPlanDelivery interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
}
type UserPlanRepository interface {
	StoreTransaction(ctx context.Context, userPlanTransaction UserPlanTransaction) error
	Store(ctx context.Context, userPlan *UserPlan) error
	Fetch(ctx context.Context) ([]UserPlan, error)
	GetById(ctx context.Context, id uint) (UserPlan, error)
	Update(ctx context.Context, plan UpdateUserPlanDTO, id uint) error
}
type UserPlanUseCase interface {
	Fetch(ctx context.Context) ([]UserPlan, error)
	Store(ctx context.Context, userPlanDTO StoreUserPlanRequestDTO) error
	StoreTransaction(ctx context.Context, userPlanTransactionDTO StoreUserPlanTransactionDTO, planId uint) error
	Update(ctx context.Context, dto UpdateUserPlanDTO, id uint) error
}
