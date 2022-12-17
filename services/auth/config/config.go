package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   `yaml:"auth"`
		GRPC  `yaml:"grpc"`
		Log   `yaml:"logger"`
		PG    `yaml:"postgres"`
		Certs `yaml:"certs"`
	}

	App struct {
		Name    string `yaml:"name" env:"APP_ENV"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	GRPC struct {
		Port string `yaml:"port" env:"GRPC_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `yaml:"url" env:"PG_URL"`
	}

	Certs struct {
		PrivateKeyPEMPath string `yaml:"private_key_pem_path" env:"PRIVATE_KEY_PEM_PATH"`
		PublicKeyPEMPath  string `yaml:"public_key_pem_path" env:"PUBLIC_KEY_PEM_PATH"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
