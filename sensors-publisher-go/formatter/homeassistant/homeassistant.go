package homeassistant

import (
	log "github.com/sirupsen/logrus"
)

type homeassistant struct{}

func (ha *homeassistant) Do(m interface{}) (interface{}, error) {
	log.WithField("m", m).Debug("formatting")
	return nil, nil
}

func New() *homeassistant {
	return &homeassistant{}
}
