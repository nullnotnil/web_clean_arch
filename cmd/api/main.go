package main

import (
	"net/http"
	"web_clean_arch/internal/modules/user/controller"
	"web_clean_arch/internal/services"
	"web_clean_arch/internal/services/pgdb"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `gorm:"primarykey"`
	Fullname string `gorm:"column:fullname;"`
	Email    string `gorm:"column:email;"`
}

func (u *User) TableName() string { return "user" }

func main() {
	asm := services.NewAppServiceManager(
		services.WithConfig(),
		services.WithLogger(),
		// example.WithExampleService("exampleeeee"),
		pgdb.WithPGDBService("pgdb"),
	)
	asm.Start()

	// spin up the server
	router := gin.Default()

	v1_group := router.Group("v1")
	v1_group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1_group.POST("users", controller.NewUserController().CreateUserFunc(asm))

	router.Run(":3000")

}
