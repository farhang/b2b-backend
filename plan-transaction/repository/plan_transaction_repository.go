package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type PlanTransactionRepository struct {
	db *gorm.DB
}

func (p PlanTransactionRepository) Store(ctx context.Context, transaction *domain.UserPlanTransaction) error {
	return p.db.WithContext(ctx).Create(transaction).Error
}

func NewPlanTransactionRepository(db *gorm.DB) domain.UserPlanTransactionRepository {
	return PlanTransactionRepository{db}
}
