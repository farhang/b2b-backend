package domain

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	PlanWithdraw string = "PLAN_WITHDRAW"
	PlanRenewal  string = "PLAN_RENEWAL"
)

const (
	Processing string = "PROCESSING"
	Accepted   string = "ACCEPTED"
	Rejected   string = "REJECTED"
)

type RequestType struct {
	ID   uint
	Name string
}

var RequestTypes = []RequestType{
	{ID: 1, Name: PlanWithdraw},
	{ID: 2, Name: PlanRenewal},
}

var RequestStatuses = []RequestStatus{
	{ID: 1, Name: Processing},
	{ID: 2, Name: Accepted},
	{ID: 3, Name: Rejected},
}

type RequestStatus struct {
	ID   uint
	Name string
}

type Request struct {
	gorm.Model
	RequestType     RequestType
	RequestTypeID   uint
	RequestStatus   RequestStatus
	RequestStatusID uint
}

type PlanRequest struct {
	gorm.Model
	Request    Request
	RequestID  uint
	UserPlan   UserPlan
	UserPlanID uint
}
type StorePlanRequest struct {
	RequestTypeID uint `json:"request_type_id"`
}
type StorePlanRequestUseCaseDTO struct {
	UserPlanId    uint
	RequestTypeID uint
}

type PlanRequestDelivery interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
}

type PlanRequestUseCase interface {
	Store(ctx context.Context, request StorePlanRequestUseCaseDTO) error
	Fetch(ctx context.Context) ([]PlanRequest, error)
}

type PlanRequestRepository interface {
	Store(ctx context.Context, request PlanRequest) error
	Fetch(ctx context.Context) ([]PlanRequest, error)
}
