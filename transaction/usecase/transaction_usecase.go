package usecase

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	tr domain.TransactionRepository
	uu domain.UserUseCase
	au domain.AssetUseCase
	db *gorm.DB
}

func (t TransactionUseCase) GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error) {
	return t.tr.GetTotalProfitByUserId(ctx, userId)
}

func (t TransactionUseCase) Store(ctx context.Context, userId uint, trDTO domain.StoreTransactionRequestDTO) error {
	tr := domain.Transaction{
		Amount:            trDTO.Amount,
		TransactionTypeID: trDTO.TransactionTypeId,
		Description:       trDTO.Description,
		UserId:            userId,
	}

	return t.tr.Store(ctx, tr)
}

func (t TransactionUseCase) FetchByUserId(ctx context.Context, userId int) ([]domain.Transaction, error) {
	return t.tr.FetchByUserId(ctx, userId)
}

func (t TransactionUseCase) Fetch(ctx context.Context) ([]domain.Transaction, error) {
	return t.tr.Fetch(ctx)
}

func NewTransactionUseCase(tr domain.TransactionRepository, uu domain.UserUseCase, au domain.AssetUseCase, db *gorm.DB) domain.TransactionUseCase {
	return &TransactionUseCase{tr, uu, au, db}
}
