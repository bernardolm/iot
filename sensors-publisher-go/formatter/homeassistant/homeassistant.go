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
	statePayload      string
	stateTopic        string
}

func (ha *homeassistant) Availability() (string, string, error) {
	log.Debug("formatting availability")
	return "", "", nil
}

func (ha *homeassistant) Config() (string, string, error) {
	if ha.hasSentConfig {
		// just skip
		return "", "", nil
	}

	ha.hasSentConfig = true

	log.WithField("formatter", "homeassistant").WithField("topic", ha.configTopic).Debug("formatting config")
	return ha.configTopic, ha.configPayload, nil
}

func (ha *homeassistant) State(value interface{}) (string, interface{}, error) {
	if value == nil {
		return "", nil, fmt.Errorf("homeassistant.state: nil value received for topic %s", ha.stateTopic)
	}

	log.WithField("value", value).WithField("formatter", "homeassistant").WithField("topic", ha.configTopic).
		Debug("formatting state")
	return ha.stateTopic, fmt.Sprintf(ha.statePayload, value), nil
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
		statePayload:      buildStatePayload(s.DeviceClass()),
		stateTopic:        t.StateTopic,
	}
}
