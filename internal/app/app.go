package app

import (
	"github.com/DiDinar5/lending-rate-aggregator/config"
	"github.com/labstack/echo/v4"
)

func Run(cfg config.Config) error {
	e := echo.New()
	e.Start(cfg.HTTP.Port)
	return nil
}
