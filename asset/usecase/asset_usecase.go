package usecase

import (
	"backend-core/domain"
	"context"
)

type AssetUseCase struct {
	ar  domain.AssetRepository
	upu domain.UserPlanUseCase
}

func (a AssetUseCase) DecreaseAmount(ctx context.Context, userId uint, amount float64) error {
	currentAmount, err := a.GetAmountByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if amount > currentAmount {
		return domain.ErrNotEnoughAssetAmount
	}

	decreasedAmount := currentAmount - amount

	return a.UpdateAmountByUserId(ctx, userId, decreasedAmount)
}

func (a AssetUseCase) IncreaseAmount(ctx context.Context, userId uint, amount float64) error {
	currentAmount, err := a.GetAmountByUserId(ctx, userId)
	if err != nil {
		return err
	}

	increasedAmount := currentAmount + amount

	return a.UpdateAmountByUserId(ctx, userId, increasedAmount)
}

func (a AssetUseCase) Store(ctx context.Context, asset domain.Asset) error {
	return a.ar.Store(ctx, asset)
}

func (a AssetUseCase) GetAmountByUserId(ctx context.Context, userId uint) (float64, error) {
	amount, err := a.upu.GetTotalAmountByUserId(ctx, userId)
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (a AssetUseCase) UpdateAmountByUserId(ctx context.Context, userId uint, amount float64) error {
	asset, err := a.ar.GetByUserId(ctx, userId)
	if err != nil {
		return err
	}
	asset.Amount = amount
	return a.ar.Update(ctx, asset)
}

func NewAssetUseCase(ar domain.AssetRepository, upu domain.UserPlanUseCase) domain.AssetUseCase {
	return &AssetUseCase{ar, upu}
}
