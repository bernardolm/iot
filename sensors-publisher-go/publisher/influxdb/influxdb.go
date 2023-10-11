package influxdb

import (
	log "github.com/sirupsen/logrus"

	influxdbclient "github.com/bernardolm/iot/sensors-publisher-go/influxdb"
)

type influxdb struct{}

func (a *influxdb) Publish(topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("topic", topic).
		WithField("message", string(message.([]byte))).
		WithField("publisher", "influxdb").
		Debug("publishing")

	influxdbclient.Publish(topic, message)

	return nil
}

func New() *influxdb {
	return &influxdb{}
}
