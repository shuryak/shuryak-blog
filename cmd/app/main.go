package main

import (
	"shuryak-blog/config"
	"shuryak-blog/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		// TODO: LOG
	}

	// Run
	app.Run(cfg)
}
