package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/vleukhin/telegram-proxy/internal"
)

func main() {
	cfg := internal.Config{}
	if err := cfg.Parse(); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	app, err := internal.NewApp(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create app")
	}
	errChan := make(chan error)
	sigChan := make(chan os.Signal, 1)

	go app.Run(errChan)

	signal.Ignore(syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigChan:
		log.Info().Msg("Terminating...")
		os.Exit(0)
	case err := <-errChan:
		log.Error().Msg("App error: " + err.Error())
		os.Exit(1)
	}

}
