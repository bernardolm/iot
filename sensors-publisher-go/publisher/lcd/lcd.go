package lcd

import (
	log "github.com/sirupsen/logrus"
)

type lcd struct{}

func (a *lcd) Publish(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("message", string(message.([]byte))).
		WithField("publisher", "lcd").
		Debug("publisher: trying to publish")

	return nil
}

func New() *lcd {
	return &lcd{}
}
