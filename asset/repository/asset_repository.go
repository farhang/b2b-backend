package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func (a AssetRepository) Store(ctx context.Context, asset domain.Asset) error {
	return a.db.WithContext(ctx).Create(&asset).Error
}

func (a AssetRepository) GetByUserId(ctx context.Context, UserId int) (domain.Asset, error) {
	asset := domain.Asset{}
	err := a.db.WithContext(ctx).Where(domain.Asset{UserID: UserId}).First(&asset).Error
	return asset, err
}

func (a AssetRepository) Fetch(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AssetRepository) Update(ctx context.Context, asset domain.Asset) error {
	return a.db.Save(&asset).Error

}

func NewAssetRepository(db *gorm.DB) domain.AssetRepository {
	return &AssetRepository{db}
}
