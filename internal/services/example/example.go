package example

import (
	"errors"
	"time"
	"web_clean_arch/internal/services"
)

type ExampleService struct {
	*services.AppService
}

func WithExampleService(name string) services.AppServiceManagerOption {
	return func(asm *services.AppServiceManager) {
		asm.Service_container[name] = &ExampleService{
			services.NewAppService(
				name,
				"This is just an example of implementation of AppService",
				true,
			),
		}
	}
}

func (as *ExampleService) Start(asm *services.AppServiceManager) error {
	asm.Logger.Warningln("ExampleService is starting up...")
	time.Sleep(time.Second * 3)
	as.GetInstance(asm)
	asm.Logger.Infof("[ExampleService] DB_NAME value: %s \n", asm.Config.Get("DB_NAME"))
	asm.Logger.Warningln("ExampleService is DONE starting up!")
	asm.Err_chan <- errors.New("Example Service errorrrrrrr")
	return nil
}

func (as *ExampleService) Stop(asm *services.AppServiceManager) error {
	return nil
}

/* func (as *ExampleService) GetInstance(asm *services.AppServiceManager) (*services.AppService, error) {
	return nil, nil
} */
