package messagebus

import (
	log "github.com/sirupsen/logrus"
)

type messagebus struct{}

func (a *messagebus) Do(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).WithField("message", message).WithField("publisher", "messagebus").
		Debug("publishing")
	return nil
}

func New() *messagebus {
	return &messagebus{}
}
