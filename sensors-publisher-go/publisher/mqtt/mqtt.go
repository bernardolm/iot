package mqtt

import (
	log "github.com/sirupsen/logrus"

	mqttclient "github.com/bernardolm/iot/sensors-publisher-go/infrastructure/mqtt"
)

type mqtt struct{}

func (a *mqtt) Publish(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).
		WithField("message", string(message.([]byte))).
		WithField("publisher", "mqtt").
		Debug("publishing")

	mqttclient.Publish(topic, message)

	return nil
}

func New() *mqtt {
	return &mqtt{}
}
