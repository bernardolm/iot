package influxdb

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/message"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

type influxdb struct{}

func (a *influxdb) Build(s sensor.Sensor) []message.Message {
	value, err := s.Get()
	if err != nil {
		log.Panic(err)
	}

	line := fmt.Sprintf("%s,entity_id=%s %s=%f",
		s.UnitOfMeasurement(),
		s.UniqueID(),
		s.DeviceClass(),
		value)

	messages := []message.Message{{Body: []byte(line)}}

	return messages
}

func New(s sensor.Sensor) (*influxdb, error) {
	f := influxdb{}
	return &f, nil
}
