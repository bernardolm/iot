package homeassistant

import (
	"fmt"

	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

const (
	stateTopicFormat = "%s/%s"
)

func (a *homeassistant) buildState(s sensor.Sensor) {
	a.stateTopic = fmt.Sprintf(stateTopicFormat, a.bridge, s.Name())
}

func (a *homeassistant) state(s sensor.Sensor) (string, error) {
	v, err := s.Get()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`{"%s":%v}`, s.DeviceClass(), v), nil
}
