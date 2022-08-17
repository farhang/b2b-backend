package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHttpHandler struct {
	UserUseCase domain.UserUseCase
	pu          domain.PlanUseCase
}

func (uh *UserHttpHandler) DepositToUserPlan(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// VerifyEmail godoc
// @Summary  Send verification code to user
// @Tags      email
// @Accept    json
// @Produce   json
// @Param    email    path    string                   true  "User email"
// @Param    message  body    domain.VerifyRequestDTO  true  "Email verification data"
// @Success  200      {bool}  bool                     "true"
// @Router   /emails/{email}/verify [patch]
func (uh *UserHttpHandler) VerifyEmail(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.VerifyRequestDTO
	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	email := ctx.Param("email")
	latestEmailVerification, err := uh.UserUseCase.GetLatestEmailVerification(c, email)
	if errors.Is(common.ErrNotFound, err) {
		return common.ErrHttpNotFound(common.ErrEmailIsNotExists)
	}
	isCodeValid := p.Code == latestEmailVerification.Code
	if !isCodeValid {
		return common.ErrHttpUnprocessableEntity(common.ErrEmailVerificationCodeIsInValid)
	}

	err = uh.UserUseCase.VerifyEmail(c, email)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, isCodeValid)
}

// FetchUsers Store godoc
// @Summary  Add new user
// @Tags     user
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Security  ApiKeyAuth
// @Router   /users/ [get]
func (uh *UserHttpHandler) FetchUsers(c echo.Context) error {

	var ctx = c.Request().Context()
	users, _ := uh.UserUseCase.Fetch(ctx)
	usersResponse := make([]domain.UserResponseDTO, len(users))
	for i := range users {
		usersResponse[i] = domain.UserResponseDTO{
			ID:    users[i].ID,
			Email: users[i].Email,
		}
	}
	return c.JSON(http.StatusOK, common.ResponseDTO{
		Data: usersResponse,
	})
}

// Store godoc
// @Summary  Add new user
// @Tags     user
// @Accept   json
// @Produce  json
// @Success  200  {string}  string  "ok"
// @Router   /users [post]
//func (uh *UserHttpHandler) Store(c echo.Context) error {
//	var ctx = c.Request().Context()
//	user := domain.StoreUserRequestDTO{}
//	err := c.Bind(&user)
//	err = uh.UserUseCase.Store(ctx, user)
//
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusCreated, "ok")
//}

func (uh *UserHttpHandler) GetById(ctx echo.Context) error {
	var c = ctx.Request().Context()
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := uh.UserUseCase.GetById(c, id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

// GetMe godoc
// @Summary   Get user information
// @Tags     user
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO{data=domain.UserResponseDTO}
// @Router    /users/me [get]
func (uh *UserHttpHandler) GetMe(ctx echo.Context) error {
	c := ctx.Request().Context()
	id := ctx.Get("userID").(int)
	me, err := uh.UserUseCase.GetById(c, id)

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data: domain.UserResponseDTO{
			ID:    me.ID,
			Email: me.Email,
		},
		Message: "",
	})
}

// GetMyPlans godoc
// @Summary   Get user plans
// @Tags     user's plan
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO{data=[]domain.GetMyPlansDTO}
// @Router    /users/plans/me [get]
func (uh *UserHttpHandler) GetMyPlans(ctx echo.Context) error {
	var c = ctx.Request().Context()
	id := ctx.Get("userID").(int)
	p, err := uh.pu.GetByUserId(c, id)
	if err != nil {
		return err
	}
	pr := make([]domain.GetMyPlansDTO, len(p))
	for i := range pr {
		pr[i] = domain.GetMyPlansDTO{
			ID:            p[i].ID,
			PlanId:        p[i].PlanID,
			Title:         p[i].Plan.Title,
			Duration:      p[i].Plan.Duration,
			Description:   p[i].Plan.Description,
			ProfitPercent: p[i].Plan.ProfitPercent,
			Amount:        p[i].Amount,
			Status:        p[i].Status,
		}
	}

	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    pr,
		Message: http.StatusText(http.StatusOK),
	})
}

// SendEmailVerificationCode godoc
// @Summary  Send verification code to user
// @Tags     email
// @Accept   json
// @Produce  json
// @Param    email  path      string                        true  "User email"
// @success  200    {object}  common.ResponseDTO{data=int}  "Email verification code"
// @Router  /emails/{email}/send-verification-code [post]
func (uh *UserHttpHandler) SendEmailVerificationCode(ctx echo.Context) error {
	c := ctx.Request().Context()
	email := ctx.Param("email")

	err := uh.UserUseCase.StoreEmailVerificationCode(c, email)

	if err != nil {
		return err
	}

	latestEmailVerification, _ := uh.UserUseCase.GetLatestEmailVerification(c, email)

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: http.StatusText(http.StatusCreated),
		Data:    latestEmailVerification.Code,
	})
}

func NewUserHttpHandler(echo *echo.Echo, userUseCase domain.UserUseCase, pu domain.PlanUseCase) domain.UserHttpHandler {
	handler := &UserHttpHandler{
		UserUseCase: userUseCase,
		pu:          pu,
	}

	ug := echo.Group("users", common.AuthMiddleWare(), common.CASBINMiddleWare())
	ug.GET("/plans/me", handler.GetMyPlans)
	ug.GET("/", handler.FetchUsers)
	ug.GET("/me", handler.GetMe)
	eg := echo.Group("emails/:email/")
	eg.PATCH("verify", handler.VerifyEmail)
	eg.POST("send-verification-code", handler.SendEmailVerificationCode)

	return handler
}
