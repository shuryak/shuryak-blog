package main

import (
	"log"
	"shuryak-blog/config"
	"shuryak-blog/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
