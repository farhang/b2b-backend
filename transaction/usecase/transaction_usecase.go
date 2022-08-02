package usecase

import (
	"backend-core/domain"
	"context"
)

type TransactionUseCase struct {
	tr domain.TransactionRepository
	uu domain.UserUseCase
	au domain.AssetUseCase
}

func (t TransactionUseCase) GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error) {
	return t.tr.GetTotalProfitByUserId(ctx, userId)
}

func (t TransactionUseCase) FetchByUserId(ctx context.Context, userId int) ([]domain.Transaction, error) {
	return t.tr.FetchByUserId(ctx, userId)
}

func (t TransactionUseCase) Fetch(ctx context.Context) error {

	panic("implement me")
}

func (t TransactionUseCase) Deposit(ctx context.Context, deposit domain.DepositRequestDTO) error {
	user, err := t.uu.GetById(ctx, deposit.UserId)
	if err != nil {
		return err
	}

	tr := domain.Transaction{
		Amount:          deposit.Amount,
		TransactionType: domain.DEPOSIT,
		User:            *user,
	}

	return t.tr.Deposit(ctx, tr, int(user.ID))

}

func (t TransactionUseCase) WithDraw(ctx context.Context, withdraw domain.WithDrawRequestDTO) error {
	user, err := t.uu.GetById(ctx, withdraw.UserId)
	if err != nil {
		return err
	}

	tr := domain.Transaction{
		Amount:          withdraw.Amount,
		TransactionType: domain.WITHDRAW,
		User:            *user,
	}

	err = t.au.DecreaseAmount(ctx, int(user.ID), withdraw.Amount)

	if err != nil {
		return err
	}

	return t.tr.Store(ctx, tr)
}

func (t TransactionUseCase) Profit(ctx context.Context, profit domain.ProfitRequestDTO) error {
	user, err := t.uu.GetById(ctx, profit.UserId)
	if err != nil {
		return err
	}

	tr := domain.Transaction{
		Amount:          profit.Amount,
		TransactionType: domain.PROFIT,
		User:            *user,
	}

	err = t.au.IncreaseAmount(ctx, int(user.ID), profit.Amount)

	if err != nil {
		return err
	}

	return t.tr.Store(ctx, tr)
}

func NewTransactionUseCase(tr domain.TransactionRepository, uu domain.UserUseCase, au domain.AssetUseCase) domain.TransactionUseCase {
	return &TransactionUseCase{tr, uu, au}
}
