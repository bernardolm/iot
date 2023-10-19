package influxdb

import (
	"context"

	log "github.com/sirupsen/logrus"

	influxdbclient "github.com/bernardolm/iot/sensors-publisher-go/infrastructure/influxdb"
)

type influxdb struct{}

func (a *influxdb) Publish(ctx context.Context, topic string, message interface{}) error {
	if message == nil {
		return nil
	}

	log.WithField("message", string(message.([]byte))).
		WithField("publisher", "influxdb").
		Debug("publisher: trying to publish")

	influxdbclient.Send(ctx, topic, message)

	return nil
}

func New() *influxdb {
	return &influxdb{}
}
