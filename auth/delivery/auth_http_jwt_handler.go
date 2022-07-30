package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthJwtHttpHandler struct {
	AuthUseCase domain.AuthUseCase
}

// Register godoc
// @Summary  Register new user
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    message  body      domain.RegisterRequestDTO  true  "Registration data"
// @Success  200      {object}  common.ResponseDTO
// @Router   /auth/register [post]
func (ajh AuthJwtHttpHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	var p domain.RegisterRequestDTO

	if err := c.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err := ajh.AuthUseCase.Register(ctx, p)

	if errors.Is(common.ErrEmailDuplication, err) {
		return common.ErrHttpConflict(err)
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, common.ResponseDTO{Message: "The user is registered successfully"})
}

// Login godoc
// @Summary  Login with email
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    message  body      domain.LoginRequestDTO   true  "User credentials"
// @success  200      {object}  domain.LoginResponseDTO  "Login response model including access token"
// @Router   /auth/login [post]
func (ajh *AuthJwtHttpHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	p := domain.LoginRequestDTO{}

	if err := c.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	token, err := ajh.AuthUseCase.Login(ctx, p)

	if errors.Is(err, common.ErrInvalidCredential) {
		return common.ErrHttpUnauthorized(err)
	}

	if errors.Is(err, common.ErrEmailIsNotVerified) {
		return common.ErrHttpUnprocessableEntity(err)
	}

	if err != nil {
		return err
	}

	res := domain.LoginResponseDTO{
		Data: domain.LoginResponseData{
			AccessToken: *token,
		},
		Message: http.StatusText(http.StatusOK),
	}

	return c.JSON(http.StatusOK, res)
}

// ResetPassword godoc
// @Summary   Reset password
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    message  body      domain.ResetPasswordRequestDTO true  "User credentials"
// @success  200      {object}  common.ResponseDTO "Login response model including access token"
// @Router   /auth/reset-password [patch]
func (ajh *AuthJwtHttpHandler) ResetPassword(c echo.Context) error {
	ctx := c.Request().Context()
	p := domain.ResetPasswordRequestDTO{}
	if err := c.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err := ajh.AuthUseCase.ResetPassword(ctx, p.Email, p.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, common.ResponseDTO{Message: "the password is changed successfully"})
}

func NewAuthHttpHandler(e *echo.Echo, authUseCase domain.AuthUseCase) domain.AuthHttpHandler {
	handler := &AuthJwtHttpHandler{AuthUseCase: authUseCase}
	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)
	e.PATCH("/auth/reset-password", handler.ResetPassword)
	return handler
}
