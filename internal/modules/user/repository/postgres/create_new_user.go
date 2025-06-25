package postgres

import (
	"context"
	"web_clean_arch/internal/modules/user/model"
)

func (repo *UserPGRepo) CreateNewUser(ctx context.Context, u model.UserCreating) error {
	if err := repo.pgdb.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}
