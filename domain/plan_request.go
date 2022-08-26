package domain

import (
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

var RequestStatuses = []RequestType{
	{ID: 1, Name: Processing},
	{ID: 2, Name: Accepted},
	{ID: 2, Name: Rejected},
}

type RequestStatus struct {
	ID   uint
	Name string
}

type Request struct {
	gorm.Model
	Type            RequestType
	TypeID          uint
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

type PlanRequestDelivery interface {
	Fetch(ctx echo.Context) error
	Store(ctx echo.Context) error
	Accept(ctx echo.Context) error
	Reject(ctx echo.Context) error
}

type PlanRequestRepository interface {
}
