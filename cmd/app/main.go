package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/rewebcan/ys-memoli/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := app.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err)
	}

	app.Run(cfg)
}
