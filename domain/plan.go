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
}

type PlanStoreRequestDTO struct {
	Title         string
	Description   string
	ProfitPercent int
	Duration      int
}
type PlanResponseDTO struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ProfitPercent int    `json:"profit"`
	Duration      int    `json:"duration"`
}

type PlanUseCase interface {
	Fetch(ctx context.Context) ([]Plan, error)
	Store(ctx context.Context, plan PlanStoreRequestDTO) error
	Delete(ctx context.Context) error
}

type PlanRepository interface {
	Fetch(ctx context.Context) ([]Plan, error)
	Store(ctx context.Context, plan Plan) error
	Delete(ctx context.Context) error
}

type PlanHttpHandler interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
