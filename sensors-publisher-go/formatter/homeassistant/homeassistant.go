package homeassistant

import (
	"fmt"

	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
	log "github.com/sirupsen/logrus"
)

const (
	availabilityTopic string = "%s/bridge/state"
	configTopic       string = "homeassistant/sensor/%s/%s/config"
	stateTopic        string = "%s/%s"
)

type homeassistant struct {
	availabilityTopic string
	configPayload     string
	configTopic       string
	hasSentConfig     bool
	stateTopic        string
}

func (ha *homeassistant) Do(value interface{}) (interface{}, error) {
	log.WithField("value", value).
		Debug("formatting")

	return nil, nil
}

func New(s sensor.Sensor) *homeassistant {
	t := configPayloadTemplate{
		AvailabilityTopic: fmt.Sprintf(availabilityTopic, s.Model()),
		ConfigTopic:       fmt.Sprintf(configTopic, s.ID(), s.Model()),
		DeviceClass:       s.DeviceClass(),
		ID:                s.ID(),
		Manufacturer:      s.Manufacturer(),
		Model:             s.Model(),
		Name:              fmt.Sprintf("%s %s sensor", s.Model(), s.DeviceClass()),
		StateTopic:        fmt.Sprintf(stateTopic, s.Model(), s.ID()),
		Unique:            fmt.Sprintf("%s_%s_sensor_%s", s.Model(), s.DeviceClass(), s.ID()),
		UnitOfMeasurement: s.UnitOfMeasurement(),
	}
	return &homeassistant{
		availabilityTopic: t.AvailabilityTopic,
		configPayload:     buildConfigPayload(t),
		configTopic:       t.ConfigTopic,
		stateTopic:        t.StateTopic,
	}
}
