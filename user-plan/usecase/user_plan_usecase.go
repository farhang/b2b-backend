package usecase

import (
	"backend-core/domain"
	"context"
	"time"
)

type userPlanUseCase struct {
	upr domain.UserPlanRepository
}

func (up userPlanUseCase) Store(ctx context.Context, userPlanDTO domain.StoreUserPlanRequestDTO) error {
	userPlan := &domain.UserPlan{
		UserID:           userPlanDTO.UserID,
		PlanID:           userPlanDTO.PlanID,
		Amount:           0,
		UserPlanStatusId: 1,
		StartedAt:        time.Now(),
		ExpiresAt:        time.Now(),
	}
	return up.upr.Store(ctx, userPlan)
}

func (up userPlanUseCase) StoreTransaction(ctx context.Context, userPlanTransactionDTO domain.StoreUserPlanTransactionDTO) error {
	userPlanTransaction := domain.UserPlanTransaction{
		Transaction: domain.Transaction{
			Amount:            0,
			TransactionTypeID: 0,
			Description:       "",
			UserId:            0,
		},
		TransactionID: 0,
		UserPlan:      domain.UserPlan{},
		UserPlanID:    0,
	}

	return up.upr.StoreTransaction(ctx, userPlanTransaction)
}

func NewUserPlanUseCase(upr domain.UserPlanRepository) domain.UserPlanUseCase {
	return &userPlanUseCase{upr}
}
