package main

import (
	"github.com/Rashad-j/edge-ai-framework/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})

	// load config
	_, err := config.ReadEnvConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("read env config failed")
	}

}
