package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileHandler struct {
	e  *echo.Echo
	pu domain.ProfileUseCase
}

// Fetch godoc
// @Summary   get profiles
// @Tags     profiles
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /profiles/ [get]
func (p ProfileHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	profiles, _ := p.pu.Fetch(c)
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    profiles,
		Message: http.StatusText(http.StatusOK),
	})
}

func (p ProfileHandler) Update(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileHttpHandler(e *echo.Echo, pu domain.ProfileUseCase) domain.ProfileDelivery {
	handler := &ProfileHandler{e, pu}
	e.GET("/profiles/", handler.Fetch)
	return handler
}
