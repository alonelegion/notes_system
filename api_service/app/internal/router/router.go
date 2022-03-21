package router

import (
	"github.com/alonelegion/notes_system/api_service/app/pkg/logging"
	"github.com/alonelegion/notes_system/api_service/app/pkg/metric"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

func Init() {
	logger := logging.GetLogger()
	logger.Println("initialization application router")

	router := httprouter.New()
	// metrics
	router.GET(handler.TEST_URL, handler.Heartbeat)

	appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Fatal(err)
	}

	socketPath := path.Join(appDir, "app.sock")
	logger.Infof("socket path: %s", socketPath)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("create and listen unix socket")
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		logger.Fatal(err)
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		logger.Infof("Caught signal %s. Shutting down...", sig)
		err = unixListener.Close()
		if err != nil {
			logger.Error("can not close application unix socket")
		}
		// Here we can do graceful shutdown (close connection and etc)
	}(sign)

	logger.Fatal(server.Serve(unixListener))
}
