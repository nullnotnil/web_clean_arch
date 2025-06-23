package main

import (
	"web_clean_arch/internal/services"
	"web_clean_arch/internal/services/pgdb"
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

	pg := asm.GetService("pgdb")
	srv := pg.(*pgdb.PGDBService)

	var user User
	srv.DB.First(&user)
	asm.Logger.Println(user)

	/* asm.Logger.Infoln(
		asm.Config.Get("DB_NAME"),
	)
	asm.Logger.Warning(os.Getenv("DB_NAME"))

	ex := asm.GetService("exampleeeee")
	if ex != nil {
		// interface to struct
		svr := ex.(*example.ExampleService)
		asm.Logger.Infoln(svr.Description)
	} */
}
