package homeassistant

import (
	"github.com/bernardolm/iot/sensors-publisher-go/message"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

const (
	bridge = "sensors_publisher_go"
	name   = "homeassistant"
)

type homeassistant struct {
	availabilityPayload string
	availabilityTopic   string
	bridge              string
	configPayload       string
	configTopic         string
	hasSentAvailability bool
	hasSentConfig       bool
	stateTopic          string
}

func (a *homeassistant) Build(s sensor.Sensor) ([]message.Message, error) {
	messages := []message.Message{}

	if !a.hasSentConfig {
		messages = append(messages, message.Message{
			Topic: a.configTopic,
			Body:  []byte(a.configPayload),
		})
		a.hasSentConfig = true
	}

	if !a.hasSentAvailability {
		messages = append(messages, message.Message{
			Topic: a.availabilityTopic,
			Body:  []byte(a.availabilityPayload),
		})
		a.hasSentAvailability = true
	}

	state, err := a.state(s)
	if err != nil {
		return nil, err
	}

	messages = append(messages, message.Message{
		Topic: a.stateTopic,
		Body:  []byte(state),
	})

	messages = append(messages, message.Message{
		Topic: a.stateTopic + "/availability",
		Body:  []byte("online"),
	})

	return messages, nil
}

func New(s sensor.Sensor) (*homeassistant, error) {
	ha := homeassistant{
		bridge: bridge,
	}

	ha.buildAvailability()
	ha.buildState(s)

	if err := ha.buildConfig(s); err != nil {
		return nil, err
	}

	return &ha, nil
}
