package messagebus

import (
	log "github.com/sirupsen/logrus"
)

type messagebus struct{}

func (mb *messagebus) Do(topic string, message interface{}) error {
	log.WithField("topic", topic).WithField("message", message).Debug("publishing")
	return nil
}

func New() *messagebus {
	return &messagebus{}
}
