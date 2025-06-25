package controller

import (
	"context"
	"web_clean_arch/internal/modules/user/model"
)

/*
Viết interface với struct loằng ngoằng thế này là để hỗ trợ dependency injection tốt hơn,
đặc biệt là với thư viện google/wire
*/

type IUserController interface {
	CreateNewUser(ctx context.Context, user model.UserCreating) error
}

type UserController struct {
	// createUserCase *usecase.CreateUserCase
}

func NewUserController() UserController {
	return UserController{}
}
