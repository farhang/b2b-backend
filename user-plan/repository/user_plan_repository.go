package repository

import (
	"backend-core/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type UserPlanRepository struct {
	db *gorm.DB
}

func (upr *UserPlanRepository) Store(ctx context.Context, userPlan *domain.UserPlan) error {
	return upr.db.WithContext(ctx).Create(userPlan).Error
}
func (upr *UserPlanRepository) StoreTransaction(ctx context.Context, userPlanTransaction domain.UserPlanTransaction) error {
	return upr.db.WithContext(ctx).Create(userPlanTransaction).Error
}

func NewUserPlanRepository(db *gorm.DB) domain.UserPlanRepository {
	var result int64
	db.Table("user_plan_statuses").Count(&result)
	if result == 0 {
		err := db.Create(domain.UserPlanStatuses).Error
		if err != nil {
			log.Error().Err(err)
		}
	}
	return &UserPlanRepository{db}
}
