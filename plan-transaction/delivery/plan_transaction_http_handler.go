package delivery

import (
	"backend-core/domain"
	"github.com/labstack/echo/v4"
)

type UserPlanTransactionHttpHandler struct {
	e *echo.Echo
}

func (u UserPlanTransactionHttpHandler) Store(ctx echo.Context) error {
	panic("")
}

func NewUserPlanTransactionHttpHandler(e *echo.Echo) domain.UserPlanTransactionDelivery {
	handler := &UserPlanTransactionHttpHandler{e}
	return handler
}
