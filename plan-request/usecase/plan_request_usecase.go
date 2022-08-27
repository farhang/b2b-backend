package usecase

import (
	"backend-core/domain"
	"context"
)

type PlanRequestUseCase struct {
	prr domain.PlanRequestRepository
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
