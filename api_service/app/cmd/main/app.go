package main

import (
	"github.com/alonelegion/notes_system/api_service/app/internal/router"
	"github.com/alonelegion/notes_system/api_service/app/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	defer router.Init()

	logger.Print("application initialized and started")
}
