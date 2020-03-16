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

func (emitter *Emitter) SendRemember(to string, name string) error {
	from := "+16122840701"
	message := fmt.Sprintf("Hola %s, te toca tu siguiente medición, entra a este link para realizar tu evaluación: https://bregymalpartida.typeform.com/to/w2V19L", name)

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

func (emitter *Emitter) SendSMS(to string, name string, score float64) (*time.Timer, error) {
	from := "+16122840701"

	recommendation := ""
	if score > 0.5 {
		recommendation = "ve a un hospital e ingresa por emergencia para que te puedan ayudar"
	} else if score > 0.2 {
		recommendation = "te encontraras bien, solo te recomendamos reposar y tener un aislamiento en casa, mientras los sintomas te duren"
	} else {
		recommendation = "estas bien, no te preocupes, solo recuerda..."
	}

	next := time.Now().Add(5*time.Minute).Add(-5*time.Hour)
	nextTimer := time.NewTimer(5*time.Minute)

	nextString := next.Format("January 02, 2006 a las 15:04")

	nextMeasure := fmt.Sprintf("Tu siguiente medición, esta programada para %s, te estaremos avisando. ", nextString)

	message := fmt.Sprintf("Hola %s, según tus sintomas y tus factores de riesgo, %s. %s", name, recommendation, nextMeasure)

	log.Infof("message to send: %s", message)

	res, exp, err := emitter.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return nil, err
	}

	if exp != nil {
		return nil, exp
	}

	log.Infof("total price: %s", res.Price)
	return nextTimer, nil
}