package controller

import (
	"net/http"
	"web_clean_arch/internal/modules/user/model"
	"web_clean_arch/internal/modules/user/repository/postgres"
	"web_clean_arch/internal/modules/user/usecase"
	"web_clean_arch/internal/services"

	"github.com/gin-gonic/gin"
)

func (uc UserController) CreateUserFunc(asm *services.AppServiceManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user_data model.UserCreatingDTO
		if err := c.Bind(&user_data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		// Cái này có nên inject vào không ???
		repo, err := postgres.NewUserPGRepo(asm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		// Cái này có nên inject vào không ???
		uc := usecase.NewCreateUserCase(repo)
		err = uc.CreateNewUser(c.Request.Context(), user_data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "User creating failed!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Message": "Success"})

	}
}
