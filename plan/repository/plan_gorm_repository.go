package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type PlanGormRepository struct {
	db *gorm.DB
}

func (p PlanGormRepository) GetByUserId(ctx context.Context, id int) ([]domain.ProfilePlan, error) {
	//var pr domain.Profile
	//var pl domain.Plan
	//var pp []domain.ProfilePlan

	//err := p.db.WithContext(ctx).Joins("Profile").Joins("Profile.User").Model(&domain.ProfilePlan{}).Scan(&pp).Error
	var pp []domain.ProfilePlan
	err := p.db.Joins("JOIN profiles p on p.id = profile_plans.profile_id").
		Joins("JOIN users u on u.id = p.user_id").
		Where("u.id = ?", id).
		Preload("Plan").
		Find(&pp).Error

	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (p PlanGormRepository) Fetch(ctx context.Context) ([]domain.Plan, error) {
	var plans []domain.Plan
	err := p.db.WithContext(ctx).Find(&plans).Error
	return plans, err
}

func (p PlanGormRepository) Store(ctx context.Context, plan domain.Plan) error {
	return p.db.WithContext(ctx).Create(&plan).Error
}

func (p PlanGormRepository) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanGormRepository(db *gorm.DB) domain.PlanRepository {
	return &PlanGormRepository{db}
}
