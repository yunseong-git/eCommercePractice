package app

import "eCommerce/config"

type App struct {
	config *config.Config
}

func NewApp() *App {
	a := &App{
		config: config,
	}
	return a
}
