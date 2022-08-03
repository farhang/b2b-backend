package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
func (ph ProfileHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	profiles, _ := ph.pu.Fetch(c)
	var profilesResponse = make([]domain.ProfileResponseDTO, len(profiles))

	for i := range profiles {
		profilesResponse[i].ID = profiles[i].ID
		profilesResponse[i].PlanId = profiles[i].PlanId
		profilesResponse[i].UserID = profiles[i].UserID
		profilesResponse[i].Name = profiles[i].Name
		profilesResponse[i].LastName = profiles[i].LastName
		profilesResponse[i].MobileNumber = profiles[i].MobileNumber
		profilesResponse[i].Position = profiles[i].Position
		profilesResponse[i].CompanyName = profiles[i].CompanyName
		profilesResponse[i].MobileNumberCompany = profiles[i].MobileNumberCompany
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    profilesResponse,
		Message: http.StatusText(http.StatusOK),
	})
}

// Update godoc
// @Summary  update profile
// @Tags     profiles
// @Accept   json
// @Produce  json
// @Param    id  path      int                        true  "Profile id"
// @success  200    {object}  common.ResponseDTO  "Email verification code"
// @Param    message  body    domain.UpdateProfileRequestDTO  true  "Profile"
// @Router  /profiles/{id}/ [put]
func (ph ProfileHandler) Update(ctx echo.Context) error {
	c := ctx.Request().Context()
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return err
	}

	var p domain.UpdateProfileRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	return ph.pu.Update(c, p, id)
}

func NewProfileHttpHandler(e *echo.Echo, pu domain.ProfileUseCase) domain.ProfileDelivery {
	handler := &ProfileHandler{e, pu}
	e.GET("/profiles/", handler.Fetch)
	e.PUT("/profiles/:id/", handler.Update)
	return handler
}
