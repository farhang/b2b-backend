package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (p ProfileRepository) GetByUserId(ctx context.Context, id int) (domain.Profile, error) {
	var pr domain.Profile
	err := p.db.WithContext(ctx).Where(domain.Profile{UserID: uint(id)}).First(&pr).Error
	return pr, err
}

func (p ProfileRepository) Update(ctx context.Context, profile domain.Profile) error {
	p.db.Model(&profile.User).Updates(domain.User{IsActive: profile.User.IsActive})
	return p.db.WithContext(ctx).Omit("user_id").Save(&profile).Error
}

func (p ProfileRepository) GetById(ctx context.Context, id int) (domain.Profile, error) {
	var profile domain.Profile
	log.Println("profileId", profile.ID)
	err := p.db.WithContext(ctx).Preload(clause.Associations).First(&profile, id).Error
	return profile, err
}

func (p ProfileRepository) Store(ctx context.Context, profile domain.Profile) error {
	return p.db.WithContext(ctx).Create(&profile).Error
}

func (p ProfileRepository) Fetch(ctx context.Context) ([]domain.Profile, error) {
	var r []domain.Profile
	err := p.db.WithContext(ctx).Preload(clause.Associations).Find(&r).Error
	return r, err
}

func NewProfileRepository(db *gorm.DB) domain.ProfileRepository {
	return &ProfileRepository{db}
}
