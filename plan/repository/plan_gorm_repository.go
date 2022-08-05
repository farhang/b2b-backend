package repository

import (
	"backend-core/domain"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type PlanGormRepository struct {
	db *gorm.DB
}

func (p PlanGormRepository) GetByUserId(ctx context.Context, id int) (domain.Plan, error) {
	var pr domain.Profile
	var pl domain.Plan

	err := p.db.WithContext(ctx).Where(domain.Profile{UserID: uint(id)}).First(&pr).Error
	if err != nil {
		return domain.Plan{}, err
	}
	fmt.Println(pr.PlanId)
	err = p.db.WithContext(ctx).First(&pl, pr.PlanId).Error

	if err != nil {
		return pl, nil
	}
	return pl, nil
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
