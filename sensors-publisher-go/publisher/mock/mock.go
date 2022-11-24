package mock

import (
	log "github.com/sirupsen/logrus"
)

type mock struct{}

func (a *mock) Do(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).WithField("message", message).WithField("publisher", "mock").
		Debug("publishing")
	return nil
}

func New() *mock {
	return &mock{}
}
