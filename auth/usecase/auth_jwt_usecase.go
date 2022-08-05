package authJwtUseCase

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"os"
)

type AuthJwtUseCase struct {
	UserUseCase domain.UserUseCase
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

	token, _ := ac.GenerateToken(domain.JwtCustomClaims{UserId: int(user.ID), Role: string(user.Role)})
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

func NewJwtAuthUseCase(userUseCase domain.UserUseCase) domain.AuthUseCase {
	return &AuthJwtUseCase{userUseCase}
}
