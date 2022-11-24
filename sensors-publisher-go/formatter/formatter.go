package formatter

import (
	"github.com/bernardolm/iot/sensors-publisher-go/message"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

type Formatter interface {
	Build(s sensor.Sensor) []message.Message
}
