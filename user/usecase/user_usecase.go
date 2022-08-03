package usecase

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	au             domain.AssetUseCase
	pu             domain.ProfileUseCase
	db             *gorm.DB
}

func (uc *userUseCase) GetLatestEmailVerification(ctx context.Context, email string) (*domain.EmailVerification, error) {
	return uc.userRepository.GetLatestEmailVerification(ctx, email)
}

func (uc *userUseCase) VerifyEmail(ctx context.Context, email string) error {
	return uc.userRepository.VerifyEmail(ctx, email)
}

func (uc *userUseCase) GenerateVerificationCodeNumber(length int) (int, error) {
	const seed = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0, err
	}

	otpCharsLength := len(seed)
	for i := 0; i < length; i++ {
		buffer[i] = seed[int(buffer[i])%otpCharsLength]
	}
	code := string(buffer)
	return strconv.Atoi(code)
}

func (uc *userUseCase) StoreEmailVerificationCode(ctx context.Context, email string) error {
	code, _ := uc.GenerateVerificationCodeNumber(6)
	emailVerification := domain.EmailVerification{
		Email:     email,
		Code:      code,
		ExpiresAt: time.Now().UTC().Add(5 * time.Minute),
	}
	return uc.userRepository.StoreEmailVerificationCode(ctx, emailVerification)
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

func (uc *userUseCase) Store(ctx context.Context, userDTO domain.StoreUserRequestDTO) error {
	hashPassword, _ := uc.GeneratePasswordHash(userDTO.Password)
	user := domain.User{
		Password: hashPassword,
		Email:    userDTO.Email,
	}

	asset := domain.Asset{
		Amount: 0,
		User:   user,
	}

	err := uc.au.Store(ctx, asset)

	if err != nil {
		return err
	}

	return uc.userRepository.Store(ctx, &user)
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
