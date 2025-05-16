package influxdb

import (
	"fmt"

	"github.com/bernardolm/iot/sensors-publisher-go/message"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

type influxdb struct{}

func (a *influxdb) Build(s sensor.Sensor) ([]message.Message, error) {
	value, err := s.Get()
	if err != nil {
		return nil, err
	}

	line := fmt.Sprintf("%s,entity_id=%s %s=%f",
		s.DeviceClass(),
		s.UniqueID(),
		s.UnitOfMeasurement(),
		value)

	messages := []message.Message{{Body: []byte(line)}}

	return messages, nil
}

func New(s sensor.Sensor) (*influxdb, error) {
	f := influxdb{}
	return &f, nil
}
