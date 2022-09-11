package repository

import (
	"backend-core/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db *gorm.DB
}

func (o OrderRepository) Fetch(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.db.WithContext(ctx).Preload(clause.Associations).Find(&orders).Error
	return orders, err
}

func (o OrderRepository) Store(ctx context.Context, order domain.Order) error {
	return o.db.WithContext(ctx).Create(&order).Error
}
func (o OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	return o.db.WithContext(ctx).Save(&order).Error
}
func (o OrderRepository) GetById(ctx context.Context, id uint) (domain.Order, error) {
	var order domain.Order
	err := o.db.WithContext(ctx).First(&order, id).Error
	return order, err
}
func (o OrderRepository) Accept(ctx context.Context) error {
	panic("")
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	var result int64
	db.Table("order_statuses").Count(&result)
	if result == 0 {
		err := db.Create(domain.OrderStatuses).Error
		if err != nil {
			log.Error().Err(err)
		}
	}
	return &OrderRepository{db}
}
