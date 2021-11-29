package main

import (
	"os"

	"github.com/rs/zerolog"
)

var log = zerolog.New(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	TimeFormat: "15:04:05",
}).With().Timestamp().Logger()
