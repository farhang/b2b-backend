package usecase

import (
	"backend-core/domain"
	"context"
)

type PlanUseCase struct {
	pr domain.PlanRepository
}

func (pu PlanUseCase) Fetch(ctx context.Context) ([]domain.Plan, error) {
	return pu.pr.Fetch(ctx)
}

func (pu PlanUseCase) Store(ctx context.Context, plan domain.PlanStoreRequestDTO) error {
	p := domain.Plan{
		Title:         plan.Title,
		Description:   plan.Description,
		ProfitPercent: plan.ProfitPercent,
		Duration:      plan.Duration,
	}
	return pu.pr.Store(ctx, p)
}

func (pu PlanUseCase) Delete(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanUseCase(pr domain.PlanRepository) domain.PlanUseCase {
	return &PlanUseCase{pr}
}
