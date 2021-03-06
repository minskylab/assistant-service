package assistservice

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type config struct {
	Port string `env:"PORT" envDefault:"3000"`
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Uri  string `env:"URI" envDefault:"mongodb://localhost:27017"`
}

func extractConfigFromEnv() (*config, error) {
	config := new(config)
	if err := env.Parse(config); err != nil {
		return nil, errors.Wrap(err, "environment variables failed at try to parsing")
	}
	return config, nil
}
