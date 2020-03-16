package assistservice

import (
	"fmt"
	"time"

	"github.com/sfreiberg/gotwilio"
	log "github.com/sirupsen/logrus"
)

type Emitter struct {
	twilio *gotwilio.Twilio
}

func NewEmitter(accountID, authToken string) (*Emitter, error) {
	twilio := gotwilio.NewTwilioClient(accountID, authToken)
	return &Emitter{twilio:twilio},nil
}


func (emitter *Emitter) SendSMS(to string, name string, score float64) error {
	from := "+16122840701"

	recommendation := ""
	if score > 0.5 {
		recommendation = "ve a un hospital e ingresa por emergencia para que te puedan ayudar"
	} else if score > 0.2 {
		recommendation = "te encontraras bien, solo te recomendamos reposar y tener un aislamiento en casa, mientras los sintomas te duren"
	} else {
		recommendation = "estas bien, no te preocupes, solo recuerda..."
	}

	next := time.Now().Add(5*time.Minute)

	nextString := next.Format("January 02, 2006 a las 15:04")

	nextMeasure := fmt.Sprintf("Tu siguiente medición, esta programada para %s ", nextString)

	message := fmt.Sprintf("Hola %s, según tus sintomas y tus factores de riesgo, %s. %s", name, recommendation, nextMeasure)

	log.Infof("message to send: %s", message)

	res, exp, err := emitter.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return err
	}

	if exp != nil {
		return exp
	}

	log.Infof("total price: %s", res.Price)
	return nil
}