package repository

import (
	"backend-core/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserPlanRepository struct {
	db *gorm.DB
}

func (upr *UserPlanRepository) GetByUserId(ctx context.Context, id uint) ([]domain.UserPlan, error) {
	var userPlans []domain.UserPlan
	err := upr.db.WithContext(ctx).Preload(clause.Associations).Find(&userPlans, domain.UserPlan{UserID: id}).Error

	return userPlans, err
}

func (upr *UserPlanRepository) Update(ctx context.Context, plan domain.UpdateUserPlanDTO, id uint) error {
	var userPlan domain.UserPlan
	err := upr.db.WithContext(ctx).First(&userPlan, id).Error
	if err != nil {
		return err
	}
	return upr.db.Model(&userPlan).Select("Amount", "UserPlanStatusId").Updates(domain.UserPlan{Amount: float64(plan.Amount), UserPlanStatusId: plan.UserPlanStatusId}).Error
}

func (upr *UserPlanRepository) GetById(ctx context.Context, id uint) (domain.UserPlan, error) {
	var userPlan domain.UserPlan
	err := upr.db.WithContext(ctx).First(&userPlan, id).Error
	return userPlan, err
}

func (upr *UserPlanRepository) Fetch(ctx context.Context) ([]domain.UserPlan, error) {
	var userPlans []domain.UserPlan
	err := upr.db.WithContext(ctx).Preload(clause.Associations).Find(&userPlans).Error
	return userPlans, err
}

func (upr *UserPlanRepository) Store(ctx context.Context, userPlan *domain.UserPlan) error {
	return upr.db.WithContext(ctx).Create(userPlan).Error
}

func (upr *UserPlanRepository) StoreTransaction(ctx context.Context, userPlanTransaction domain.UserPlanTransaction) error {
	if userPlanTransaction.Transaction.TransactionTypeID == 3 || userPlanTransaction.Transaction.TransactionTypeID == 5 {
		var userplan domain.UserPlan
		upr.db.WithContext(ctx).Preload(clause.Associations).First(&userplan, userPlanTransaction.UserPlanID)
		amount := userplan.Amount + userPlanTransaction.Transaction.Amount
		upr.db.Model(&userplan).Update("amount", amount)

		upr.db.Save(&userplan)
	}

	return upr.db.WithContext(ctx).Create(&userPlanTransaction).Error
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
