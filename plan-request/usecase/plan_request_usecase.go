package usecase

import (
	"backend-core/domain"
	"context"
)

type PlanRequestUseCase struct {
	prr domain.PlanRequestRepository
}

func (p PlanRequestUseCase) Update(ctx context.Context, requestId uint, dto domain.UpdatePlanRequestDTO) error {
	planRequest, err := p.prr.GetById(ctx, requestId)
	if err != nil {
		return err
	}
	planRequest.Request.RequestStatusID = dto.RequestStatusID
	return p.prr.Update(ctx, &planRequest)
}

func (p PlanRequestUseCase) Fetch(ctx context.Context) ([]domain.PlanRequest, error) {
	return p.prr.Fetch(ctx)
}

func (p PlanRequestUseCase) Store(ctx context.Context, planRequest domain.StorePlanRequestUseCaseDTO) error {
	request := domain.PlanRequest{
		Request: domain.Request{
			RequestTypeID:   planRequest.RequestTypeID,
			RequestStatusID: 1,
		},
		UserPlanID: planRequest.UserPlanId,
	}
	return p.prr.Store(ctx, request)
}

func NewPlanRequestUseCase(prr domain.PlanRequestRepository) domain.PlanRequestUseCase {

	return &PlanRequestUseCase{prr}
}
