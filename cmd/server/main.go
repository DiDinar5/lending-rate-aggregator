package main

import (
	"github.com/DiDinar5/lending-rate-aggregator/config"
	"github.com/DiDinar5/lending-rate-aggregator/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app.Run(cfg)
}
