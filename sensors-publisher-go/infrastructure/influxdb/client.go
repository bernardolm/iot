package influxdb

import (
	"context"
	"crypto/tls"

	influx "github.com/influxdata/influxdb-client-go"
	log "github.com/sirupsen/logrus"
)

var client influx.Client

func getClient(ctx context.Context) influx.Client {
	loadConfig()

	if client == nil {
		connect(ctx)
	}

	return client
}

func connect(_ context.Context) {
	log.Debug("influxdb: trying to connect")

	opts := influx.DefaultOptions().
		SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})

	log.Debugf("influxdb: connecting to %s with '%s'", url, token)

	client = influx.NewClientWithOptions(url, token, opts)

	if client == nil {
		log.Error("influxdb: couldn't create a client")
		return
	}

	healthcheck, err := client.Health(context.Background())
	if err != nil {
		log.WithError(err).Error("influxdb: health check failed")
		return
	}

	log.WithField("status", healthcheck.Status).
		WithField("message", *healthcheck.Message).
		Info("influxdb: connected")
}
