package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
	formatterhomeassistant "github.com/bernardolm/iot/sensors-publisher-go/formatter/homeassistant"
	formatterinfluxdb "github.com/bernardolm/iot/sensors-publisher-go/formatter/influxdb"
	"github.com/bernardolm/iot/sensors-publisher-go/influxdb"
	"github.com/bernardolm/iot/sensors-publisher-go/logging"
	"github.com/bernardolm/iot/sensors-publisher-go/mqtt"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	publisherinfluxdb "github.com/bernardolm/iot/sensors-publisher-go/publisher/influxdb"
	publishermqtt "github.com/bernardolm/iot/sensors-publisher-go/publisher/mqtt"
	publisherstdout "github.com/bernardolm/iot/sensors-publisher-go/publisher/stdout"
	sensords18a20 "github.com/bernardolm/iot/sensors-publisher-go/sensor/ds18a20"
	sensormock "github.com/bernardolm/iot/sensors-publisher-go/sensor/mock"
	"github.com/bernardolm/iot/sensors-publisher-go/worker"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	logging.Init()

	ctx, ctxCancelFunc := context.WithCancel(context.Background())
	defer ctxCancelFunc()

	if err := mqtt.Connect(ctx); err != nil {
		log.Error(err)
	}

	if err := influxdb.Connect(ctx); err != nil {
		log.Error(err)
	}

	pbMqtt := publishermqtt.New()
	pbInfluxdb := publisherinfluxdb.New()
	pbStdout := publisherstdout.New()
	w := worker.New()

	ds, err := sensords18a20.New()
	if err != nil {
		log.Error(err)
	}

	for i := range ds {
		fha, err := formatterhomeassistant.New(ds[i])
		if err != nil {
			log.Error(err)
		}
		w.AddFlow(ds[i], fha, []publisher.Publisher{pbStdout, pbMqtt})

		fidb, err := formatterinfluxdb.New(ds[i])
		if err != nil {
			log.Error(err)
		}
		w.AddFlow(ds[i], fidb, []publisher.Publisher{pbStdout, pbInfluxdb})
	}

	if len(ds) == 0 { // entering in DEBUG mode
		sm := sensormock.New()

		fha, err := formatterhomeassistant.New(sm)
		if err != nil {
			log.Error(err)
		}
		w.AddFlow(sm, fha, []publisher.Publisher{pbStdout, pbMqtt})

		fidb, err := formatterinfluxdb.New(sm)
		if err != nil {
			log.Error(err)
		}
		w.AddFlow(sm, fidb, []publisher.Publisher{pbStdout, pbInfluxdb})
	}

	w.Start(ctx)

	ec := make(<-chan error)
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM) // nolint

	select {
	case err := <-ec:
		ctxCancelFunc()
		log.Error(err)
	case <-sc:
		log.Warn("cmd: shutdown requested")
	case <-ctx.Done():
		log.Warn("cmd: context done")
	}

	w.Stop(ctx)
	mqtt.Disconnect(ctx)
	influxdb.Disconnect(ctx)

	log.Info("cmd: graceful shutdown complete")
}
