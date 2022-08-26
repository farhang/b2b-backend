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
	Amount           int
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

type StoreUserPlanTransactionDTO struct {
	Amount            float64
	TransactionTypeID uint
	Description       string
	UserId            uint
	UserPlanID        uint
}

type UserPlanDelivery interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
}
type UserPlanRepository interface {
	StoreTransaction(ctx context.Context, userPlanTransaction UserPlanTransaction) error
	Store(ctx context.Context, userPlan *UserPlan) error
}
type /**/ UserPlanUseCase interface {
	Store(ctx context.Context, userPlanDTO StoreUserPlanRequestDTO) error
	StoreTransaction(ctx context.Context, userPlanTransactionDTO StoreUserPlanTransactionDTO) error
}
