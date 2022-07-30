package repository

import "gorm.io/gorm"

type PlanGormRepository struct {
	db *gorm.DB
}

func NewPlanGormRepository(db *gorm.DB) *PlanGormRepository {
	return &PlanGormRepository{db}
}
