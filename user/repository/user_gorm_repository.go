package repository

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func (ur *UserGormRepository) GetLatestEmailVerification(ctx context.Context, email string) (*domain.EmailVerification, error) {
	emailVerification := domain.EmailVerification{}
	err := ur.db.WithContext(ctx).Last(&emailVerification, "email = ?", email).Error
	isNotFoundError := errors.Is(gorm.ErrRecordNotFound, err)

	if isNotFoundError {
		return nil, common.ErrNotFound
	}

	return &emailVerification, err
}

func (ur *UserGormRepository) StoreEmailVerificationCode(ctx context.Context, emailVerification domain.EmailVerification) error {
	result := ur.db.WithContext(ctx).Create(&emailVerification)
	return result.Error
}

func (ur *UserGormRepository) VerifyEmail(ctx context.Context, email string) error {
	user, err := ur.GetByEmail(ctx, email)

	if err != nil {
		return err
	}
	user.IsEmailVerified = true
	return ur.db.Save(&user).Error
}

func (ur *UserGormRepository) Update(ctx context.Context, user domain.User) error {
	return ur.db.WithContext(ctx).Model(&user).Save(&user).Error
}

func (ur *UserGormRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).Where(domain.User{Email: email}).First(&user).Error
	isNotFoundError := errors.Is(gorm.ErrRecordNotFound, err)

	if isNotFoundError {
		return nil, common.ErrNotFound
	}

	return &user, err
}

func (ur *UserGormRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := ur.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (ur *UserGormRepository) Register(ctx context.Context, registerDTO domain.RegisterRequestDTO) error {
	var user = &domain.User{
		Password: registerDTO.Password,
		Email:    registerDTO.Email,
		Role:     domain.MEMBER,
	}
	return ur.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
			return err
		}

		asset := domain.Asset{
			Amount: 0,
			UserID: int(user.ID),
		}

		if err := tx.WithContext(ctx).Create(&asset).Error; err != nil {
			return err
		}

		p := domain.Profile{
			Name:                registerDTO.Name,
			LastName:            registerDTO.LastName,
			MobileNumber:        registerDTO.MobileNumber,
			MobileNumberCompany: registerDTO.MobileNumberCompany,
			Position:            registerDTO.Position,
			CompanyName:         registerDTO.CompanyName,
			UserID:              user.ID,
		}

		if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
			return err
		}
		var pl domain.Plan

		if err := tx.FirstOrCreate(&pl, domain.Plan{
			Title:         "Basic",
			Description:   "",
			ProfitPercent: 0,
			Duration:      0,
		}).Error; err != nil {
			return err
		}

		pp := domain.ProfilePlan{
			Profile: p,
			Plan:    pl,
		}

		if err := tx.WithContext(ctx).Create(&pp).Error; err != nil {
			return err
		}

		return nil
	})
}

func (ur *UserGormRepository) GetById(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	err := ur.db.WithContext(ctx).First(&user, id).Error
	isNotFoundError := errors.Is(gorm.ErrRecordNotFound, err)

	if isNotFoundError {
		return nil, common.ErrNotFound
	}

	return &user, err
}

func NewGormUserRepository(db *gorm.DB) domain.UserRepository {
	db.Exec("DROP TYPE IF EXISTS role;CREATE TYPE role AS ENUM ('ADMIN', 'MEMBER');")
	return &UserGormRepository{db}
}
