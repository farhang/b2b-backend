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

// GetMyPlan godoc
// @Summary   get my plan
// @Tags     plan
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Router   /plans/me/ [get]
// @Success  200  {object} common.ResponseDTO
func (ph *PlanHttpHandler) GetMyPlan(ctx echo.Context) error {
	var c = ctx.Request().Context()
	id := ctx.Get("userID").(int)
	p, err := ph.pu.GetByUserId(c, id)
	if err != nil {
		return err
	}

	pr := domain.PlanResponseDTO{
		ID:            p.ID,
		Title:         p.Title,
		Description:   p.Description,
		ProfitPercent: p.ProfitPercent,
		Duration:      p.Duration,
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    pr,
		Message: http.StatusText(http.StatusOK),
	})
}

// Fetch godoc
// @Summary  Get plans
// @Tags     plan
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /plans/ [get]
func (ph *PlanHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	plans, err := ph.pu.Fetch(c)
	var plansResponse = make([]domain.PlanResponseDTO, len(plans))
	for i := range plans {
		plansResponse[i].ID = plans[i].ID
		plansResponse[i].Title = plans[i].Title
		plansResponse[i].Duration = plans[i].Duration
		plansResponse[i].Description = plans[i].Description
		plansResponse[i].ProfitPercent = plans[i].ProfitPercent
	}

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    plansResponse,
		Message: http.StatusText(http.StatusOK),
	})
}

// Store godoc
// @Summary   Add new plan
// @Tags     plan
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
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
	pg := e.Group("/plans", common.AuthMiddleWare(), common.CASBINMiddleWare())
	pg.POST("/", handler.Store)
	pg.GET("/me/", handler.GetMyPlan)
	pg.GET("/", handler.Fetch)

	return handler
}
