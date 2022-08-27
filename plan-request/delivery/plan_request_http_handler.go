package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RequestHttpHandler struct {
	e   *echo.Echo
	pru domain.PlanRequestUseCase
}

// Fetch godoc
// @Summary  Get plan's requests
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/requests [get]
func (r RequestHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	planRequests, err := r.pru.Fetch(c)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, planRequests)
}

// Update godoc
// @Summary  Create a request for a plan
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    id    path    string                   true  "User plan Id"
// @Param    message  body      domain.UpdatePlanRequestDTO  true  "Data"
// @Router   /users/plans/requests/{id} [patch]
func (r RequestHttpHandler) Update(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.UpdatePlanRequestDTO
	requestId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return err
	}

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err = r.pru.Update(c, uint(requestId), p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "OK")
}

// Store godoc
// @Summary  Create a request for a plan
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    id    path    string                   true  "User plan Id"
// @Param    message  body      domain.StorePlanRequest  true  "Data"
// @Router   /users/plans/{id}/requests [post]
func (r RequestHttpHandler) Store(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.StorePlanRequest
	userPlanId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return err
	}

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err = r.pru.Store(c, domain.StorePlanRequestUseCaseDTO{
		UserPlanId:    uint(userPlanId),
		RequestTypeID: p.RequestTypeID,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "OK")
}

func NewPlanRequestHttpHandler(e *echo.Echo, pru domain.PlanRequestUseCase) domain.PlanRequestDelivery {
	handler := RequestHttpHandler{e, pru}
	e.POST("/users/plans/:id/requests", handler.Store, common.AuthMiddleWare())
	e.GET("/users/plans/requests", handler.Fetch, common.AuthMiddleWare())
	e.PATCH("/users/plans/requests/:id", handler.Update, common.AuthMiddleWare())

	return handler
}
