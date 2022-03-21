package app_context

import (
	"github.com/alonelegion/notes_system/api_service/app/internal/config"
	"github.com/alonelegion/notes_system/api_service/app/pkg/logging"
	"sync"
)

type AppContext struct {
	Config *config.Config
}

var instance *AppContext
var once sync.Once

func GetInstance() *AppContext {
	logging.GetLogger().Println("initializing application context")
	once.Do(func() {
		instance = &AppContext{
			Config: config.GetConfig(),
		}
	})

	return instance
}
