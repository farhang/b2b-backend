package usecase

import (
	"backend-core/domain"
	"context"
)

type ProfileUseCase struct {
	pr domain.ProfileRepository
}

func (p ProfileUseCase) Store(ctx context.Context, profile domain.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (p ProfileUseCase) StoreEmptyProfileByUserId(ctx context.Context, userId int) error {
	ep := domain.Profile{
		Name:                "",
		LastName:            "",
		MobileNumber:        "",
		MobileNumberCompany: "",
		Position:            "",
		CompanyName:         "",
		UserID:              uint(userId),
		PlanId:              1,
	}
	return p.pr.Store(ctx, ep)
}

func (p ProfileUseCase) Fetch(ctx context.Context) ([]domain.Profile, error) {
	return p.pr.Fetch(ctx)
}

func NewProfileUseCase(pr domain.ProfileRepository) domain.ProfileUseCase {
	return &ProfileUseCase{pr}
}
