package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
// @Router   /users/plans [get]
func (u UserPlanDelivery) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	userPlans, err := u.upu.Fetch(c)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, userPlans)
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
	planId, err := strconv.Atoi(ctx.Param("id"))
	var p domain.StoreUserPlanTransactionDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err = u.upu.StoreTransaction(c, p, uint(planId))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "OK")
}

// Store godoc
// @Summary  Create plan for user
// @Tags     plan,user
// @Security  ApiKeyAuth
// @Accept   json
// @Param    user_id    path    string true  "user id"
// @Param    message  body   domain.StoreUserPlanRequestDTO data "payload"
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/{user_id}/plans [post]
func (u UserPlanDelivery) Store(ctx echo.Context) error {
	c := ctx.Request().Context()
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return err
	}

	var p domain.StoreUserPlanRequestDTO
	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err = u.upu.Store(c, uint(userId), p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Message: http.StatusText(http.StatusOK),
	})
}

// GetByUserId godoc
// @Summary  Get a user plans
// @Tags     plan,user
// @Security  ApiKeyAuth
// @Accept   json
// @Param    user_id    path    string                   true  "user id"
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/{user_id}/plans [get]
func (u UserPlanDelivery) GetByUserId(ctx echo.Context) error {
	c := ctx.Request().Context()
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	userPlans, err := u.upu.GetByUserId(c, uint(userId))
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    userPlans,
		Message: http.StatusText(http.StatusOK),
	})
}

// Update godoc
// @Summary  update user's plan
// @Tags     user,plan
// @Security  ApiKeyAuth
// @Param    id    path    string                   true  "user plan id"
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    message  body      domain.UpdateUserPlanDTO  true  "Update User plan"
// @Router   /users/plans/{id} [patch]
func (u UserPlanDelivery) Update(ctx echo.Context) error {
	c := ctx.Request().Context()
	planId, err := strconv.Atoi(ctx.Param("id"))
	var p domain.UpdateUserPlanDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err = u.upu.Update(c, p, uint(planId))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, "OK")

}

func NewUserPlanDelivery(e *echo.Echo, upu domain.UserPlanUseCase) domain.UserPlanDelivery {
	handler := &UserPlanDelivery{e, upu}
	e.POST("/users/plans/:id/transactions", handler.StoreUserPlanTransaction, common.AuthMiddleWare())
	e.GET("/users/plans", handler.Fetch, common.AuthMiddleWare())
	e.PATCH("/users/plans/:id", handler.Update, common.AuthMiddleWare())
	e.GET("/users/:id/plans", handler.GetByUserId, common.AuthMiddleWare())
	e.POST("/users/:id/plans", handler.Store, common.AuthMiddleWare())
	return handler
}
