package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/minskylab/assistservice"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type config struct {
	Port string `env:"PORT" envDefault:"3000"`
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Uri  string `env:"URI" envDefault:"mongodb://localhost:27017"`

	TwilioAccount string `env:"TWILIO_ACCOUNT"`
	TwilioAuthToken string `env:"TWILIO_AUTH_TOKEN"`
}

func extractConfigFromEnv() (*config, error) {
	config := new(config)
	if err := env.Parse(config); err != nil {
		return nil, errors.Wrap(err, "environment variables failed at try to parsing")
	}
	return config, nil
}

func boolToFloat(res bool) float64 {
	if res {
		return 1.0
	}
	return 0.0
}

func main() {
	config, err := extractConfigFromEnv()
	if err != nil {
		panic(err)
	}

	emitter, err := assistservice.NewEmitter(
		config.TwilioAccount,
		config.TwilioAuthToken,
	)
	if err != nil {
		panic(err)
	}

	engine := gin.New()

	engine.POST("/typeform-webhook", func(c *gin.Context) {
		newAutoGen := new(WebHookRequest)
		if err := c.BindJSON(newAutoGen); err != nil {
			panic(err)
		}

		answers := newAutoGen.FormResponse.Answers
		fmt.Printf("total answers: %d\n", len(answers))

		if len(answers) < 9 {
			c.String(http.StatusBadRequest, "invalid answer size")
			return
		}

		phone := answers[0].PhoneNumber
		name := "Jhon"

		score := assistservice.GetReport(
			assistservice.DiseasesPayload{
				Fiebre:                boolToFloat(answers[3].Boolean),
				CongestionNasal:       0.0,
				Nauseas:               0.0,
				Vomitos:               0.0,
				Escalofrios:           0.0,
				DolorDeCabeza:         0.0,
				DolorMuscular:         boolToFloat(answers[8].Boolean),
				DolorDeHuesos:         0.0,
				DolorDeGarganta:       boolToFloat(answers[5].Boolean),
				Cansancio:             boolToFloat(answers[6].Boolean),
				TosSeca:               boolToFloat(answers[4].Boolean),
				TosProductiva:         0.0,
				FaltaDeAireAlRespirar: boolToFloat(answers[7].Boolean),
			},
			assistservice.DefaultWeights,
		)

		log.WithField("score", score).Info("getting report from form")

		next, err := emitter.SendSMS(phone, name, score)
		if err != nil {
			panic(err)
		}

		go func() {
			<- next.C
			if err := emitter.SendRemember(phone, name); err != nil {
				panic(err)
			}
		}()
	})

	if err := engine.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		panic(err)
	}
}
