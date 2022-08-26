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
	uu domain.UserUseCase
}

// GetById godoc
// @Summary   Get a profile
// @Tags     profile
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Param    id  path      int                        true  "Profile id"
// @Success  200  {object} common.ResponseDTO
// @Router    /profiles/{id}/ [get]
func (ph ProfileHandler) GetById(ctx echo.Context) error {
	c := ctx.Request().Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	p, err := ph.pu.GetById(c, id)

	if err != nil {
		return err
	}

	profileResponse := domain.ProfileResponseDTO{
		ID:                  p.ID,
		UserID:              p.UserID,
		Name:                p.Name,
		LastName:            p.LastName,
		MobileNumber:        p.MobileNumber,
		Position:            p.Position,
		CompanyName:         p.CompanyName,
		MobileNumberCompany: p.MobileNumberCompany,
	}

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    profileResponse,
		Message: http.StatusText(http.StatusOK),
	})
}

// GetMyProfile godoc
// @Summary  Get authenticated user profile
// @Tags     profile
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /user/profile/ [get]
func (ph ProfileHandler) GetMyProfile(ctx echo.Context) error {
	var c = ctx.Request().Context()
	id := ctx.Get("userID").(int)
	p, err := ph.pu.GetByUserId(c, id)

	if err != nil {
		return err
	}

	pr := domain.ProfileResponseDTO{
		ID:     p.ID,
		UserID: p.UserID,
		//PlanId:              p.PlanId,
		Name:                p.Name,
		LastName:            p.LastName,
		MobileNumber:        p.MobileNumber,
		Position:            p.Position,
		CompanyName:         p.CompanyName,
		MobileNumberCompany: p.MobileNumberCompany,
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    pr,
		Message: http.StatusText(http.StatusOK),
	})
}

// Fetch godoc
// @Summary   get profiles
// @Tags     profile
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
		//profilesResponse[i].PlanId = profiles[i].PlanId
		profilesResponse[i].UserID = profiles[i].UserID
		profilesResponse[i].Name = profiles[i].Name
		profilesResponse[i].LastName = profiles[i].LastName
		profilesResponse[i].MobileNumber = profiles[i].MobileNumber
		profilesResponse[i].Position = profiles[i].Position
		profilesResponse[i].CompanyName = profiles[i].CompanyName
		profilesResponse[i].MobileNumberCompany = profiles[i].MobileNumberCompany
		profilesResponse[i].MobileNumberCompany = profiles[i].MobileNumberCompany
		profilesResponse[i].IsActive = profiles[i].User.IsActive
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    profilesResponse,
		Message: http.StatusText(http.StatusOK),
	})
}

// Update godoc
// @Summary  Update a profile
// @Tags     profile
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
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

func NewProfileHttpHandler(e *echo.Echo, pu domain.ProfileUseCase, uu domain.UserUseCase) domain.ProfileDelivery {
	handler := &ProfileHandler{e, pu, uu}
	pg := e.Group("profiles", common.AuthMiddleWare())
	pg.GET("/", handler.Fetch)
	pg.GET("/me/", handler.GetMyProfile)
	pg.PUT("/:id/", handler.Update)
	pg.GET("/:id/", handler.GetById)
	return handler
}
