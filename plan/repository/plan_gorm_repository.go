package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlanGormRepository struct {
	db *gorm.DB
}

func (p PlanGormRepository) GetByUserId(ctx context.Context, id int) ([]domain.UserPlan, error) {
	var userPlans []domain.UserPlan
	err := p.db.Where("user_id", id).Preload(clause.Associations).Find(&userPlans).Error
	return userPlans, err
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
