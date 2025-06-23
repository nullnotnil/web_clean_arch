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
	GetService(srv_name string) IAppService
	GetEnvString(key string) string
	GetEnvInt(key string, default_value int) int
	GetEnvBool(key string, default_value bool) int
	GetEnvFloat(key string, default_value float32) float32
}
type AppServiceManager struct {
	Err_chan          chan error
	Service_container map[string]IAppService
	Logger            *logrus.Logger
	Config            *viper.Viper
}

func NewAppServiceManager(opts ...AppServiceManagerOption) *AppServiceManager {
	asm := &AppServiceManager{
		Err_chan:          make(chan error),
		Service_container: make(map[string]IAppService),
	}
	for _, opt := range opts {
		opt(asm)
	}
	return asm
}

func (asm *AppServiceManager) Start() error {
	if asm.Logger == nil {
		panic("Logger is not initialized")
	}
	if asm.Config == nil {
		panic("Config is not loaded")
	}

	asm.Logger.Warningln("App services are starting up...")
	var wg sync.WaitGroup
	for _, svr := range asm.Service_container {
		wg.Add(1)
		go func() {
			defer wg.Done()
			svr.Start(asm)
		}()
	}
	go func() {
		wg.Wait()
		close(asm.Err_chan)
	}()

	for err := range asm.Err_chan {
		asm.Logger.Errorln(err)
	}

	return nil
}
func (asm *AppServiceManager) Stop() error {
	return nil
}
func (asm *AppServiceManager) GetService(svr_name string) IAppService {
	s, ok := asm.Service_container[svr_name]
	if !ok {
		return nil
	}
	return s
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
		asm.Service_container[name] = NewAppService(name, desc, singleton)
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
	if as.instance != nil && as.is_singleton {
		return as.instance, nil
	}
	srv, ok := asm.Service_container[as.Name]
	if ok {
		as.instance = srv.(*AppService)
	}
	return as.instance, nil
}
