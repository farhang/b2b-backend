package delivery

import (
	"backend-core/domain"
	"github.com/labstack/echo/v4"
)

type RequestHttpHandler struct {
	e *echo.Echo
}

// Fetch godoc
// @Summary  Get plan's requests
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/requests/ [get]
func (r RequestHttpHandler) Fetch(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// Store godoc
// @Summary  Create request for plan
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/:id/requests/ [post]
func (r RequestHttpHandler) Store(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// Accept godoc
// @Summary  Accept a plan's request
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/requests/:id/accept [put]
func (r RequestHttpHandler) Accept(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// Reject godoc
// @Summary  Reject a plan's request
// @Tags     request,user,plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/plans/requests/:id/reject [put]
func (r RequestHttpHandler) Reject(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanRequestHttpHandler(e *echo.Echo) domain.PlanRequestDelivery {
	handler := RequestHttpHandler{e}
	return handler
}
