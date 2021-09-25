package ds18a20

import (
	log "github.com/sirupsen/logrus"
)

type ds18a20 struct{}

func (d *ds18a20) Get() (interface{}, error) {
	log.Debug("getting values")
	return nil, nil
}

func (d *ds18a20) DeviceClass() string {
	return "temperature"
}

func (d *ds18a20) ID() string {
	return "some_random_device_id"
}

func (d *ds18a20) Manufacturer() string {
	return "Unknown"
}

func (d *ds18a20) Model() string {
	return "ds18a20"
}

func (d *ds18a20) UnitOfMeasurement() string {
	return "°C"
}

func New() *ds18a20 {
	return &ds18a20{}
}
