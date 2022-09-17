package usecase

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
)

type userPlanUseCase struct {
	upr domain.UserPlanRepository
}

func (up userPlanUseCase) GetTotalAmountByUserId(ctx context.Context, userId uint) (float64, error) {
	return up.upr.GetTotalAmountByUserId(ctx, userId)
}

func (up userPlanUseCase) GetByUserId(ctx context.Context, userId uint) ([]domain.UserPlan, error) {
	return up.upr.GetByUserId(ctx, userId)
}

func (up userPlanUseCase) Update(ctx context.Context, dto domain.UpdateUserPlanDTO, id uint) error {
	return up.upr.Update(ctx, dto, id)
}

func (up userPlanUseCase) Fetch(ctx context.Context) ([]domain.UserPlansRes, error) {
	return up.upr.Fetch(ctx)
}

func (up userPlanUseCase) Store(ctx context.Context, userId uint, userPlanDTO domain.StoreUserPlanRequestDTO) error {
	userPlan := &domain.UserPlan{
		UserID:           userId,
		PlanID:           userPlanDTO.PlanID,
		Amount:           userPlanDTO.Amount,
		UserPlanStatusId: 1,
		StartedAt:        common.ConvertTimeStampToTime(userPlanDTO.StartedAt),
		ExpiresAt:        common.ConvertTimeStampToTime(userPlanDTO.ExpiresAt),
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
