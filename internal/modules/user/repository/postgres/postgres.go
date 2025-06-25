package postgres

import (
	"errors"
	"web_clean_arch/internal/modules/user/model"
	"web_clean_arch/internal/services"
	"web_clean_arch/internal/services/pgdb"
)

/*
Repo layer thực chất vẫn là layer bên ngoài, nằm ở tầng DB, Infra trong mô hình Clean Architecture
*/
type IUserPGRepo interface {
	CreateNewUser(u model.UserCreating) error
}

type UserPGRepo struct {
	pgdb *pgdb.PGDBService
}

func NewUserPGRepo(asm *services.AppServiceManager) (*UserPGRepo, error) {
	db := asm.GetService("pgdb")
	if db == nil {
		return nil, errors.New("No pgdb service found!")
	}
	svr, ok := db.(*pgdb.PGDBService)
	if !ok {
		return nil, errors.New("pgdb service seems not to be right")
	}

	return &UserPGRepo{pgdb: svr}, nil
}
