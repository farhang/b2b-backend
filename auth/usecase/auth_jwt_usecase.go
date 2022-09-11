package authJwtUseCase

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/kavenegar/kavenegar-go"
	"os"
)

type AuthJwtUseCase struct {
	UserUseCase    domain.UserUseCase
	ProfileUseCase domain.ProfileUseCase
}

func (ac *AuthJwtUseCase) Register(c context.Context, authDTO domain.RegisterRequestDTO) error {
	validate := validator.New()
	err := validate.Struct(authDTO)

	if err != nil {
		return err
	}
	if isUserExist := ac.UserUseCase.CheckIsUserDuplicatedByEmail(c, authDTO.Email); isUserExist {
		return common.ErrEmailDuplication
	}

	authDTO.Password, _ = ac.UserUseCase.GeneratePasswordHash(authDTO.Password)

	return ac.UserUseCase.Register(c, authDTO)
}

func (ac *AuthJwtUseCase) SendOTP(c context.Context, sendOTPDTO domain.SendOTPRequestDTO) error {

	code, _ := ac.UserUseCase.GenerateVerificationCodeNumber(4)
	profile, err := ac.ProfileUseCase.GetByMobileNumber(c, sendOTPDTO.MobileNumber)
	if err != nil {
		return err
	}
	err = ac.UserUseCase.StoreVerificationCode(c, code, profile.UserID)
	if err != nil {
		return err
	}
	api := kavenegar.New(os.Getenv("KAVENEGAR_API_KEY"))
	receptor := sendOTPDTO.MobileNumber
	template := "42844"
	params := &kavenegar.VerifyLookupParam{}
	_, err = api.Verify.Lookup(receptor, template, code, params)

	if err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
		return err
	}

	return nil
}

func (ac *AuthJwtUseCase) LoginWithOTP(c context.Context, loginWithOTPDTO domain.LoginWithOTPDTO) (*string, error) {
	p, err := ac.ProfileUseCase.GetByMobileNumber(c, loginWithOTPDTO.MobileNumber)

	if err != nil {
		return nil, err
	}

	code, err := ac.UserUseCase.GetLatestVerificationCode(c, p.UserID)

	if err != nil {
		return nil, err
	}

	if code.Code != loginWithOTPDTO.Code {
		return nil, common.ErrNotFound
	}

	token, _ := ac.GenerateToken(domain.JwtCustomClaims{UserId: int(p.User.ID), Role: p.User.RoleID})
	return &token, nil
}

func (ac *AuthJwtUseCase) Login(c context.Context, loginUserDTO domain.LoginRequestDTO) (*string, error) {
	user, err := ac.UserUseCase.GetByEmail(c, loginUserDTO.Email)

	if errors.Is(err, common.ErrNotFound) {
		return nil, common.ErrInvalidCredential
	}

	isPasswordMatched := ac.UserUseCase.ComparePasswordHash(loginUserDTO.Password, user.Password)
	if !isPasswordMatched {
		return nil, common.ErrInvalidCredential
	}

	isEmailVerified := ac.UserUseCase.IsEmailVerified(c, loginUserDTO.Email)

	if !isEmailVerified {
		return nil, common.ErrEmailIsNotVerified
	}

	token, _ := ac.GenerateToken(domain.JwtCustomClaims{UserId: int(user.ID), Role: user.RoleID})
	return &token, nil

}

func (ac *AuthJwtUseCase) ResetPassword(ctx context.Context, id int, newPassword string) error {
	user, err := ac.UserUseCase.GetById(ctx, id)
	if errors.Is(common.ErrNotFound, err) {
		return common.ErrHttpNotFound(err)
	}
	hashNewPassword, _ := ac.UserUseCase.GeneratePasswordHash(newPassword)
	user.Password = hashNewPassword
	return ac.UserUseCase.Update(ctx, *user)
}

func (ac *AuthJwtUseCase) GenerateToken(claims domain.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	secret := []byte(os.Getenv("JWT_SECRET_KEY"))
	signedToken, err := token.SignedString(secret)
	return signedToken, err
}

func NewJwtAuthUseCase(userUseCase domain.UserUseCase, profileUseCase domain.ProfileUseCase) domain.AuthUseCase {
	return &AuthJwtUseCase{userUseCase, profileUseCase}
}
