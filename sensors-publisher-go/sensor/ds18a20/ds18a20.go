package ds18a20

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/yryz/ds18b20"
)

type ds18a20 struct {
	sensors []string
}

func (d *ds18a20) Get() (interface{}, error) {
	var value *float64

	for _, sensor := range d.sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			log.WithField("sensor", "ds18a20").WithField("id", sensor).WithField("value", t).
				Debug("sensor retrieved value")
			value = &t
			break
		}
	}

	if value == nil {
		return nil, fmt.Errorf("ds18a20.get: fail to get value")
	}

	log.WithField("sensor", "ds18a20").WithField("value", *value).Debug("getting values")
	return *value, nil
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
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	return &ds18a20{
		sensors: sensors,
	}
}
