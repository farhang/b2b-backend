package usecase

import (
	"backend-core/domain"
	"context"
)

type ProfileUseCase struct {
	pr domain.ProfileRepository
}

func (pu ProfileUseCase) GetByUserId(ctx context.Context, id int) (domain.Profile, error) {
	return pu.pr.GetByUserId(ctx, id)
}

func (pu ProfileUseCase) Update(ctx context.Context, profile domain.UpdateProfileRequestDTO, id int) error {
	p, _ := pu.pr.GetById(ctx, id)

	p.Name = profile.Name
	p.LastName = profile.LastName
	p.MobileNumber = profile.MobileNumber
	p.MobileNumberCompany = profile.MobileNumberCompany
	p.Position = profile.Position
	p.PlanId = uint(profile.PlanId)
	p.CompanyName = profile.CompanyName
	p.User.IsActive = profile.IsActive
	p.Plan.ID = uint(profile.PlanId)

	return pu.pr.Update(ctx, p)
}

func (pu ProfileUseCase) GetById(ctx context.Context, id int) (domain.Profile, error) {
	return pu.pr.GetById(ctx, id)
}

func (pu ProfileUseCase) Store(ctx context.Context, profile domain.Profile) error {
	//TODO implement me
	panic("implement me")
}

func (pu ProfileUseCase) Fetch(ctx context.Context) ([]domain.Profile, error) {
	return pu.pr.Fetch(ctx)
}

func NewProfileUseCase(pr domain.ProfileRepository) domain.ProfileUseCase {
	return &ProfileUseCase{pr}
}
