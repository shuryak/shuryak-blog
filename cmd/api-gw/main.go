package main

import (
	"github.com/shuryak/shuryak-blog/internal/api-gw/app"
	"github.com/shuryak/shuryak-blog/internal/api-gw/config"
	"github.com/shuryak/shuryak-blog/pkg/parser"
	"log"
)

func main() {
	cfg, err := parser.ParseConfig[config.Config]("./config.yml")
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
