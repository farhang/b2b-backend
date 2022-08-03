package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p ProfileRepository) Store(ctx context.Context, profile domain.Profile) error {
	return p.db.WithContext(ctx).Create(&profile).Error
}

func (p ProfileRepository) Fetch(ctx context.Context) ([]domain.Profile, error) {
	var r []domain.Profile
	err := p.db.WithContext(ctx).Find(&r).Error
	return r, err
}

func NewProfileRepository(db *gorm.DB) domain.ProfileRepository {
	return &ProfileRepository{db}
}
