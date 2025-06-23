package main

import (
	"os"
	"web_clean_arch/internal/services"
)

func main() {
	asm := services.NewAppServiceManager(
		services.WithConfig(),
		services.WithLogger(),
	)
	asm.Logger.Infoln(
		asm.Config.Get("DB_NAME"),
	)
	asm.Logger.Warning(os.Getenv("DB_NAME"))
}
