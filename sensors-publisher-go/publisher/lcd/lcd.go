package lcd

import (
	log "github.com/sirupsen/logrus"
)

type lcd struct{}

func (a *lcd) Publish(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).
		WithField("message", string(message.([]byte))).
		WithField("publisher", "lcd").
		Debug("publishing")

	return nil
}

func New() *lcd {
	return &lcd{}
}
