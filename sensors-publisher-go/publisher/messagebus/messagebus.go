package messagebus

import (
	log "github.com/sirupsen/logrus"
)

type messagebus struct{}

func (mb *messagebus) Do(m interface{}) error {
	log.WithField("m", m).Debug("publishing")
	return nil
}

func New() *messagebus {
	return &messagebus{}
}
