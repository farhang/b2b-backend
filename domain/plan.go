package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Title         string
	Description   string
	ProfitPercent int
	Duration      int
	Profiles      []ProfilePlan
}

type PlanStoreRequestDTO struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProfitPercent int    `json:"profit_percent"`
	Duration      int    `json:"duration"`
}
type PlanResponseDTO struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProfitPercent int    `json:"profit"`
	Duration      int    `json:"duration"`
}

type GetMyPlansDTO struct {
	ID            uint   `json:"id"`
	PlanId        int    `json:"plan_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProfitPercent int    `json:"profit"`
	Duration      int    `json:"duration"`
	Amount        int    `json:"amount"`
	Status        int    `json:"status"`
}

type PlanUseCase interface {
	GetByUserId(ctx context.Context, id int) ([]ProfilePlan, error)
	Fetch(ctx context.Context) ([]Plan, error)
	Store(ctx context.Context, plan PlanStoreRequestDTO) error
	Delete(ctx context.Context) error
}

type PlanRepository interface {
	GetByUserId(ctx context.Context, id int) ([]ProfilePlan, error)
	Fetch(ctx context.Context) ([]Plan, error)
	Store(ctx context.Context, plan Plan) error
	Delete(ctx context.Context) error
}

type PlanHttpHandler interface {
	GetMyPlan(ctx echo.Context) error
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
