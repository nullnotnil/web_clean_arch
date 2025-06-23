package pgdb

import (
	"web_clean_arch/internal/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDBService struct {
	*services.AppService
	DB *gorm.DB
}

func WithPGDBService(name string) services.AppServiceManagerOption {
	return func(asm *services.AppServiceManager) {
		asm.Service_container[name] = &PGDBService{
			services.NewAppService(
				name,
				"PostgresDB Connection",
				true,
			),
			nil,
		}
	}
}

func (as *PGDBService) Start(asm *services.AppServiceManager) error {
	dsn := "host=localhost user=psql169 password=psql169 dbname=psql169 port=54332 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		asm.Err_chan <- err
		return err
	}
	as.DB = db
	return nil
}

func (as *PGDBService) Stop(asm *services.AppServiceManager) error {
	return nil
}

/* func (as *ExampleService) GetInstance(asm *services.AppServiceManager) (*services.AppService, error) {
	return nil, nil
} */
