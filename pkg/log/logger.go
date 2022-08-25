package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var log zerolog.Logger

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

func Info(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func Debug(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func Error(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func Panic(format string, v ...interface{}) {
	log.Panic().Msgf(format, v...)
}
