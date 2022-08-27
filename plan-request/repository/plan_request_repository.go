package repository

import (
	"backend-core/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PlanRequestRepository struct {
	db *gorm.DB
}

func (p PlanRequestRepository) Fetch(ctx context.Context) ([]domain.PlanRequest, error) {
	var planRequests []domain.PlanRequest
	err := p.db.WithContext(ctx).Preload(clause.Associations).Find(&planRequests).Error
	return planRequests, err
}
func (p PlanRequestRepository) Store(ctx context.Context, request domain.PlanRequest) error {
	return p.db.WithContext(ctx).Create(&request).Error
}

func NewPlanRequestRepository(db *gorm.DB) domain.PlanRequestRepository {
	var requestTypesCount int64
	db.Table("request_types").Count(&requestTypesCount)
	if requestTypesCount == 0 {
		err := db.Create(domain.RequestTypes).Error
		if err != nil {
			log.Error().Err(err)
		}
	}
	var requestStatusesCount int64
	db.Table("request_statuses").Count(&requestStatusesCount)
	if requestStatusesCount == 0 {
		err := db.Create(domain.RequestStatuses).Error
		if err != nil {
			log.Error().Err(err)
		}
	}
	return &PlanRequestRepository{db}
}
