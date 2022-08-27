package usecase

import (
	"backend-core/domain"
	"context"
	"time"
)

type userPlanUseCase struct {
	upr domain.UserPlanRepository
}

func (up userPlanUseCase) Fetch(ctx context.Context) ([]domain.UserPlan, error) {
	return up.upr.Fetch(ctx)
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

func (up userPlanUseCase) StoreTransaction(ctx context.Context, userPlanTransactionDTO domain.StoreUserPlanTransactionDTO, planId uint) error {
	userPlan, err := up.upr.GetById(ctx, planId)
	if err != nil {
		return err
	}
	userPlanTransaction := domain.UserPlanTransaction{
		Transaction: domain.Transaction{
			Amount:            userPlanTransactionDTO.Amount,
			TransactionTypeID: userPlanTransactionDTO.TransactionTypeID,
			Description:       userPlanTransactionDTO.Description,
			UserId:            userPlan.UserID,
		},
		UserPlanID: planId,
	}

	return up.upr.StoreTransaction(ctx, userPlanTransaction)
}

func NewUserPlanUseCase(upr domain.UserPlanRepository) domain.UserPlanUseCase {
	return &userPlanUseCase{upr}
}
