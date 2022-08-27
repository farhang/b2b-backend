package usecase

import (
	"backend-core/domain"
	"context"
)

type PlanTransactionUseCase struct {
	ptr domain.UserPlanTransactionRepository
}

func (p PlanTransactionUseCase) Store(ctx context.Context, transaction *domain.UserPlanTransaction) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanTransactionUseCase(ptr domain.UserPlanTransactionRepository) domain.UserPlanTransactionRepository {
	return PlanTransactionUseCase{ptr}
}
