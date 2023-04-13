package config

type (
	Config struct {
		Service `yaml:"app"`
		HTTP    `yaml:"http"`
		Logger  `yaml:"logger"`
		Jaeger  `yaml:"jaeger"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	Service struct {
		Name    string `yaml:"name" env:"SERVICE_NAME"`
		Version string `yaml:"version" env:"SERVICE_VERSION"`
	}

	Logger struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Jaeger struct {
		URL string `yaml:"url" env:"JAEGER_URL"`
	}
)
