package usecase

import (
	"context"
	"web_clean_arch/internal/modules/user/model"
	"web_clean_arch/internal/modules/user/repository/postgres"
)

type CreateUserCase struct {
	pgrepo *postgres.UserPGRepo
}

func NewCreateUserCase(pgrepo *postgres.UserPGRepo) *CreateUserCase {
	return &CreateUserCase{pgrepo}
}
func (uc *CreateUserCase) CreateNewUser(ctx context.Context, u model.UserCreatingDTO) error {
	return uc.pgrepo.CreateNewUser(ctx, model.UserCreating{
		Fullname: u.Fullname,
		Email:    u.Email,
	})
}
