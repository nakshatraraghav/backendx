package lib

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/rs/zerolog"
)

var (
	log  zerolog.Logger
	once sync.Once
)

func GetLogger() zerolog.Logger {
	once.Do(func() {
		path := filepath.Dir("/var/log/backendx/")
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(err)
		}

		file, err := os.OpenFile("/var/log/backendx/backendx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		writer := zerolog.MultiLevelWriter(file, zerolog.ConsoleWriter{Out: os.Stdout})

		log = zerolog.New(writer).With().Timestamp().Logger()
	})

	return log
}
