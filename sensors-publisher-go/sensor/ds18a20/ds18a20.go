package ds18a20

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yryz/ds18b20"
)

type ds18a20 struct {
	id string
}

func (d *ds18a20) Get() (interface{}, error) {
	var value *float64

	if len(d.id) == 0 {
		return nil, fmt.Errorf("none ds18a20 sensor found")
	}

	t, err := ds18b20.Temperature(d.id)
	if err == nil {
		log.WithField("sensor", "ds18a20").
			WithField("id", d.id).
			WithField("value", t).
			Debug("sensor retrieved value")
		value = &t
	}

	if value == nil {
		return nil, fmt.Errorf("ds18a20.get: fail to get value")
	}

	log.WithField("sensor", "ds18a20").
		WithField("value", *value).
		Debug("getting values")

	return *value, nil
}

func (d *ds18a20) DeviceClass() string {
	return "temperature"
}

func (d *ds18a20) ID() string {
	return d.id
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

func New() ([]*ds18a20, error) {
	sensorIDs, err := ds18b20.Sensors()
	if err != nil {
		return nil, err
	}

	if len(sensorIDs) == 0 {
		log.Debug("sensors ds18a20 not found")
		return nil, nil
	}

	log.WithField("sensors", strings.Join(sensorIDs, ",")).
		Debugf("%d sensors ds18a20 found", len(sensorIDs))

	sensors := []*ds18a20{}

	for _, sensor := range sensorIDs {
		sensors = append(sensors, &ds18a20{
			id: sensor,
		})
	}

	return sensors, nil
}
