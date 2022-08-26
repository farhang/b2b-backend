package repository

import (
	"backend-core/domain"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type PlanRequestRepository struct {
}

func NewPlanRequestRepository(db *gorm.DB) domain.PlanRequestRepository {
	var result int64
	db.Table("request_types").Count(&result)
	if result == 0 {
		err := db.Create(domain.RequestTypes).Error
		if err != nil {
			log.Error().Err(err)
		}
	}
	return &PlanRequestRepository{}
}
