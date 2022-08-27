package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	PROCESSING string = "PROCESSING"
	COMPLETE   string = "COMPLETE"
	CANCELED   string = "CANCELED"
	ACCEPTED   string = "ACCEPTED"
	REJECTED   string = "REJECTED"
)

type OrderStatus struct {
	ID   uint
	Name string
}

var OrderStatuses = []OrderStatus{
	{ID: 1, Name: PROCESSING},
	{ID: 2, Name: COMPLETE},
	{ID: 3, Name: CANCELED},
	{ID: 4, Name: ACCEPTED},
	{ID: 5, Name: REJECTED},
}

func GetOrderStatusIdByName(name string) *uint {
	for _, tt := range OrderStatuses {
		if tt.Name == name {
			return &tt.ID
		}
	}
	return nil
}

type Order struct {
	gorm.Model
	User          User
	UserId        uint
	Plan          Plan
	PlanID        uint
	OrderStatus   OrderStatus
	OrderStatusId uint
}

type StoreOrderDTO struct {
	PlanID uint
	UserID uint
}

type UpdateOrderDTO struct {
	OrderStatusId uint `json:"status_id"`
}

type StoreForAuthenticateUserDTO struct {
	PlanID uint `json:"plan_id"`
}

type OrderDelivery interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
}

type OrderRepository interface {
	Fetch(ctx context.Context) ([]Order, error)
	Store(ctx context.Context, order Order) error
	GetById(ctx context.Context, id uint) (Order, error)
	Update(ctx context.Context, order *Order) error
	Accept(ctx context.Context) error
}

type OrderUseCase interface {
	Fetch(ctx context.Context) ([]Order, error)
	GetMe(ctx context.Context) error
	Store(ctx context.Context, dto StoreOrderDTO) error
	Update(ctx context.Context, orderId uint, dto UpdateOrderDTO) error
	GetByUserId(ctx context.Context) error
}
