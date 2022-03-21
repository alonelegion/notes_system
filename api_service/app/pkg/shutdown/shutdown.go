package shutdown

import (
	"github.com/alonelegion/notes_system/api_service/app/pkg/logging"
	"io"
	"os"
	"os/signal"
)

func Graceful(signals []os.Signal, closeItems ...io.Closer) {
	logger := logging.GetLogger()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign
	logger.Infof("Caught signal %s. Shutting down...", sig)

	// Here we can do graceful shutdown (close connection and etc)
	for _, closer := range closeItems {
		if err := closer.Close(); err != nil {
			logger.Errorf("failed to close %v: %v", closer, err)
		}
	}
}
