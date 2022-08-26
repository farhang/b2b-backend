package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserPlanDelivery struct {
	e   *echo.Echo
	upu domain.UserPlanUseCase
}

// Fetch godoc
// @Summary  get user's plans
// @Tags     plan,user
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/ [get]
func (u UserPlanDelivery) Fetch(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// StoreUserPlanTransaction godoc
// @Summary  Create transaction for a plan
// @Tags     transaction,plan,user
// @Security  ApiKeyAuth
// @Param    plan_id    path    string true "Plan id"
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    message  body   domain.StoreUserPlanTransactionDTO data "payload"
// @Router   /users/plans/{plan_id}/transactions [post]
func (u *UserPlanDelivery) StoreUserPlanTransaction(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.StoreUserPlanTransactionDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	u.upu.StoreTransaction(c, p)

	return ctx.JSON(http.StatusOK, "")
}

// Store godoc
// @Summary  Create plan for user
// @Tags     plan,user
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/ [post]
func (u UserPlanDelivery) Store(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// FetchTransaction godoc
// @Summary  Get a plan's transactions
// @Tags     transaction,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/:id/transactions [get]
func (u UserPlanDelivery) FetchTransaction(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewUserPlanDelivery(e *echo.Echo, upu domain.UserPlanUseCase) domain.UserPlanDelivery {
	handler := &UserPlanDelivery{e, upu}
	e.POST("/users/plans/:id/transactions", handler.StoreUserPlanTransaction, common.AuthMiddleWare())
	return handler
}
