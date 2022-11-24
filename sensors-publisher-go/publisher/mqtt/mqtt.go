package mqtt

import (
	mqttclient "github.com/bernardolm/iot/sensors-publisher-go/mqtt"
	log "github.com/sirupsen/logrus"
)

type mqtt struct{}

func (a *mqtt) Publish(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).WithField("message", message).WithField("publisher", "mqtt").
		Debug("publishing")

	mqttclient.Publish(topic, message)

	return nil
}

func New() *mqtt {
	return &mqtt{}
}
