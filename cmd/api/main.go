package main

import (
	"web_clean_arch/internal/services"
	"web_clean_arch/internal/services/example"
)

func main() {
	asm := services.NewAppServiceManager(
		services.WithConfig(),
		services.WithLogger(),
		example.WithExampleService("exampleeeee"),
	)
	asm.Start()
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
