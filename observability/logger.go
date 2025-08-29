package observability

import (
	"log"
	"os"
	"sync"
)

var (
	logger *log.Logger
	once   sync.Once
)

// Logger returns the singleton logger instance
func Logger() *log.Logger {
	once.Do(func() {
		if logger == nil {
			logger = log.New(os.Stdout, "", log.Default().Flags())
		}
	})

	return logger
}
