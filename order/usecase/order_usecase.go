package usecase

import (
	"backend-core/domain"
	"context"
	"time"
)

type OrderUseCase struct {
	or  domain.OrderRepository
	upu domain.UserPlanUseCase
}

func (o OrderUseCase) GetMe(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderUseCase) GetByUserId(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderUseCase) Accept(ctx context.Context) error {
	//return o.upu.Store(ctx, domain.StoreUserPlanRequestDTO{
	//	UserID:    1,
	//	PlanID:    1,
	//	Amount:    200,
	//	StartedAt: time.Now(),
	//	ExpiresAt: time.Now().Add(2),
	//}
	panic("")
}

func (o OrderUseCase) Fetch(ctx context.Context) ([]domain.Order, error) {
	return o.or.Fetch(ctx)
}

func (o OrderUseCase) Update(ctx context.Context, orderId uint, p domain.UpdateOrderDTO) error {
	order, err := o.or.GetById(ctx, orderId)
	if err != nil {
		return err
	}
	order.OrderStatusId = p.OrderStatusId
	isComplete := order.OrderStatusId == 2

	if isComplete {
		err = o.upu.Store(ctx, domain.StoreUserPlanRequestDTO{
			UserID:    order.UserId,
			PlanID:    order.PlanID,
			Amount:    0,
			StartedAt: time.Now(),
			ExpiresAt: time.Now(),
		})
		if err != nil {
			return err
		}
	}

	return o.or.Update(ctx, &order)
}

func (o OrderUseCase) Store(ctx context.Context, p domain.StoreOrderDTO) error {
	order := domain.Order{
		UserId:        p.UserID,
		PlanID:        p.PlanID,
		OrderStatusId: *domain.GetOrderStatusIdByName(domain.Processing),
	}
	return o.or.Store(ctx, order)
}

func NewOrderUseCase(or domain.OrderRepository, upu domain.UserPlanUseCase) domain.OrderUseCase {
	return OrderUseCase{or, upu}
}
