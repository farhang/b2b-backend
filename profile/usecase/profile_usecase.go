package usecase

import (
	"backend-core/domain"
	"context"
)

type ProfileUseCase struct {
	pr domain.ProfileRepository
}

func (p ProfileUseCase) Fetch(ctx context.Context) ([]domain.Profile, error) {
	//TODO implement me
	panic("implement me")
}

func NewProfileUseCase(pr domain.ProfileRepository) domain.ProfileUseCase {
	return &ProfileUseCase{pr}
}
