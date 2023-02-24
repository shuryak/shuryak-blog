package main

import (
	"github.com/shuryak/shuryak-blog/internal/user/app"
	"github.com/shuryak/shuryak-blog/internal/user/config"
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
