package services

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type IAppServiceManager interface {
	Start() error
	Stop() error
	GetService(srv_name string) *AppService
	GetEnvString(key string) string
	GetEnvInt(key string, default_value int) int
	GetEnvBool(key string, default_value bool) int
	GetEnvFloat(key string, default_value float32) float32
}
type AppServiceManager struct {
	err_chan          chan error
	service_container map[string]*AppService
	Logger            *logrus.Logger
	Config            *viper.Viper
}

func NewAppServiceManager(opts ...AppServiceManagerOption) *AppServiceManager {
	asm := &AppServiceManager{
		err_chan:          make(chan error),
		service_container: make(map[string]*AppService),
	}
	for _, opt := range opts {
		opt(asm)
	}
	return asm
}

func (asm *AppServiceManager) Start() error {
	return nil
}
func (asm *AppServiceManager) Stop() error {
	return nil
}
func (asm *AppServiceManager) GetService(svr_name string) *AppService {
	return nil
}
func (asm *AppServiceManager) GetEnvString(key string, default_value string) string {
	return ""
}
func (asm *AppServiceManager) GetEnvInt(key string, default_value int) int {
	return 0
}
func (asm *AppServiceManager) GetEnvBool(key string, default_value bool) bool {
	return false
}
func (asm *AppServiceManager) GetEnvFloat(key string, default_value float32) float32 {
	return 0
}

/* AppServiceManagerOption Definition */
type AppServiceManagerOption func(asm *AppServiceManager)

func WithLogger() AppServiceManagerOption {
	return func(asm *AppServiceManager) {
		asm.Logger = logrus.New()
	}
}
func WithConfig() AppServiceManagerOption {
	return func(asm *AppServiceManager) {
		v := viper.New()
		v.SetConfigName("config") // name of config file (without extension)
		v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		v.AddConfigPath(".")      // optionally look for config in the working directory
		err := v.ReadInConfig()   // Find and read the config file
		if err != nil {           // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		asm.Config = v
	}
}

func WithNewService(name, desc string, singleton bool) AppServiceManagerOption {
	return func(asm *AppServiceManager) {
		asm.service_container[name] = NewAppService(name, desc, singleton)
	}
}

/* AppService Definition */
type IAppService interface {
	GetInstance(asm *AppServiceManager) (*AppService, error)
	Start(asm *AppServiceManager) error
	Stop(asm *AppServiceManager) error
}
type AppService struct {
	Name         string
	Description  string
	instance     *AppService
	is_singleton bool
	locker       sync.Mutex
}

func NewAppService(name, desc string, singleton bool) *AppService {
	return &AppService{
		Name:         name,
		Description:  desc,
		is_singleton: singleton,
		instance:     nil,
	}
}
func (as *AppService) Start(asm *AppServiceManager) error {
	return nil
}

func (as *AppService) Stop(asm *AppServiceManager) error {
	return nil
}
func (as *AppService) GetInstance(asm *AppServiceManager) (*AppService, error) {
	return nil, nil
}
