package ds18a20

import (
	log "github.com/sirupsen/logrus"
)

type ds18a20 struct{}

func (d *ds18a20) Get() (interface{}, error) {
	log.Debug("getting values")
	return nil, nil
}

func New() *ds18a20 {
	return &ds18a20{}
}
