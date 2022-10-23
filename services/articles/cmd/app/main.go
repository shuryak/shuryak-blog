package main

import (
	"github.com/shuryak/shuryak-blog/config"
	"github.com/shuryak/shuryak-blog/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
