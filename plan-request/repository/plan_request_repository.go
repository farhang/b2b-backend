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
	err := p.db.WithContext(ctx).Preload(clause.Associations).Preload("Request.RequestType").Preload("Request.RequestStatus").Preload("UserPlan.User").Preload("UserPlan.UserPlanStatus").Preload("UserPlan.Plan").Find(&planRequests).Error
	return planRequests, err
}

func (p PlanRequestRepository) GetById(ctx context.Context, id uint) (domain.PlanRequest, error) {
	var planRequest domain.PlanRequest
	err := p.db.WithContext(ctx).Preload("Request").First(&planRequest, id).Error
	return planRequest, err
}

func (p PlanRequestRepository) Update(ctx context.Context, request *domain.PlanRequest) error {
	return p.db.WithContext(ctx).Model(domain.Request{}).Where("id = ?", request.RequestID).Update("request_status_id", request.Request.RequestStatusID).Error
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
