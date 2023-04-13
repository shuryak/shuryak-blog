package parser

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func ParseConfig[T any](path string) (*T, error) {
	cfg := new(T)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// Env variables have more priority than declared in YAML
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
