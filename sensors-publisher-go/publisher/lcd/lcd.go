package lcd

import (
	log "github.com/sirupsen/logrus"
)

type lcd struct{}

func (a *lcd) Do(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).WithField("message", message).WithField("publisher", "lcd").
		Debug("publishing")
	return nil
}

func New() *lcd {
	return &lcd{}
}
