package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

var zeroLogger *zerolog.Logger
var zeroLoggerLevel = zerolog.InfoLevel

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	writer := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages\n", missed)
	})
	logger := zerolog.New(zerolog.ConsoleWriter{Out: writer}).
		Level(zeroLoggerLevel).
		With().
		Caller().
		Timestamp().
		Logger()
	zeroLogger = &logger
	zeroLogger = Logger()
}

// Logger returns a zerolog.Logger singleton
func Logger() *zerolog.Logger {
	return zeroLogger
}
