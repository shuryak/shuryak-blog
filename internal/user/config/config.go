package config

type (
	Config struct {
		Service `yaml:"service"`
		Logger  `yaml:"logger"`
		PG      `yaml:"postgres"`
	}

	Service struct {
		Name    string `yaml:"name" env:"SERVICE_NAME"`
		Version string `yaml:"version" env:"SERVICE_VERSION"`
	}

	Logger struct {
		Level string `yaml:"level" env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env:"PG_URL"`
	}
)
