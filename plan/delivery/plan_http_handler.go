package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PlanHttpHandler struct {
	pu domain.PlanUseCase
}

func (ph *PlanHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	plans, err := ph.pu.Fetch(c)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    plans,
		Message: "",
	})
}

// Store godoc
// @Summary   Add new plan
// @Tags     plan
// @Accept   json
// @Produce  json
// @Param    message  body      domain.PlanStoreRequestDTO true  "User credentials"
// @success  200      {object}  common.ResponseDTO "Login response model including access token"
// @Router   /plans/ [post]
func (ph *PlanHttpHandler) Store(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.PlanStoreRequestDTO
	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err := ph.pu.Store(c, p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: "the plan is created successfully",
	})
}

func (ph *PlanHttpHandler) Delete(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanHttpHandler(e *echo.Echo, pu domain.PlanUseCase) domain.PlanHttpHandler {
	handler := &PlanHttpHandler{pu}
	pg := e.Group("plans")
	pg.POST("/", handler.Store)
	pg.GET("/", handler.Fetch)

	return handler
}
