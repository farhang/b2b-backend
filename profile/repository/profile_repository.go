package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
	"log"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p ProfileRepository) Update(ctx context.Context, profile domain.Profile) error {
	log.Println(profile.ID)
	return p.db.WithContext(ctx).Omit("user_id").Save(&profile).Error
}

func (p ProfileRepository) GetById(ctx context.Context, id int) (domain.Profile, error) {
	var profile domain.Profile
	log.Println("profileId", profile.ID)
	err := p.db.WithContext(ctx).First(&profile, id).Error
	return profile, err
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
