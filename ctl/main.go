package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
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

func main() {
	config, err := extractConfigFromEnv()
	if err != nil {
		panic(err)
	}

	engine := gin.New()

	engine.POST("/typeform-webhook", func(c *gin.Context) {
		newAutoGen := new(*WebHookRequest)
		if err := c.BindJSON(newAutoGen); err != nil {
			panic(err)
		}

		fmt.Printf("%+v\n", newAutoGen)
		//fmt.Println(string(data))
	})

	if err := engine.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		panic(err)
	}
}
