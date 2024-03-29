package usecase

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	au             domain.AssetUseCase
	pu             domain.ProfileUseCase
	db             *gorm.DB
}

func (uc *userUseCase) GetLatestVerificationCode(ctx context.Context, userId uint) (*domain.VerificationCode, error) {
	return uc.userRepository.GetLatestVerificationCode(ctx, userId)
}

func (uc *userUseCase) VerifyEmail(ctx context.Context, email string) error {
	return uc.userRepository.VerifyEmail(ctx, email)
}

func (uc *userUseCase) GenerateVerificationCodeNumber(length int) (string, error) {
	const seed = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(seed)
	for i := 0; i < length; i++ {
		buffer[i] = seed[int(buffer[i])%otpCharsLength]
	}
	return string(buffer), nil

}

func (uc *userUseCase) StoreVerificationCode(ctx context.Context, code string, userId uint) error {
	emailVerification := domain.VerificationCode{
		UserId:    userId,
		Code:      code,
		ExpiresAt: time.Now().UTC().Add(5 * time.Minute),
	}
	return uc.userRepository.StoreVerificationCode(ctx, emailVerification)
}

func (uc *userUseCase) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return uc.userRepository.GetByEmail(ctx, email)
}

func (uc *userUseCase) Update(ctx context.Context, user domain.User) error {
	return uc.userRepository.Update(ctx, user)
}

func (uc *userUseCase) CheckIsUserDuplicatedByEmail(ctx context.Context, email string) bool {
	_, err := uc.GetByEmail(ctx, email)
	isUserNotFound := errors.Is(err, common.ErrNotFound)
	return !isUserNotFound
}

func (uc *userUseCase) Fetch(ctx context.Context) ([]domain.User, error) {
	return uc.userRepository.Fetch(ctx)
}
func (uc *userUseCase) Register(ctx context.Context, dto domain.RegisterRequestDTO) error {
	return uc.userRepository.Register(ctx, dto)
}

func (uc *userUseCase) GetById(ctx context.Context, id int) (*domain.User, error) {
	return uc.userRepository.GetById(ctx, id)
}

func (uc *userUseCase) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (uc *userUseCase) ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (uc *userUseCase) IsEmailVerified(c context.Context, email string) bool {
	user, err := uc.userRepository.GetByEmail(c, email)
	if err != nil {
		return false
	}
	return user.IsEmailVerified
}

func NewUserUseCase(userRepository domain.UserRepository, au domain.AssetUseCase, pu domain.ProfileUseCase, db *gorm.DB) domain.UserUseCase {
	return &userUseCase{
		userRepository,
		au,
		pu,
		db,
	}
}
