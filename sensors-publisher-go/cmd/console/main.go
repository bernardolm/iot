package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
	formatterhomeassistant "github.com/bernardolm/iot/sensors-publisher-go/formatter/homeassistant"
	formatterinfluxdb "github.com/bernardolm/iot/sensors-publisher-go/formatter/influxdb"
	"github.com/bernardolm/iot/sensors-publisher-go/infrastructure/influxdb"
	"github.com/bernardolm/iot/sensors-publisher-go/infrastructure/mqtt"
	"github.com/bernardolm/iot/sensors-publisher-go/logging"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	publisherinfluxdb "github.com/bernardolm/iot/sensors-publisher-go/publisher/influxdb"
	publishermqtt "github.com/bernardolm/iot/sensors-publisher-go/publisher/mqtt"
	publisherstdout "github.com/bernardolm/iot/sensors-publisher-go/publisher/stdout"
	sensords18a20 "github.com/bernardolm/iot/sensors-publisher-go/sensor/ds18a20"
	sensormock "github.com/bernardolm/iot/sensors-publisher-go/sensor/mock"
	"github.com/bernardolm/iot/sensors-publisher-go/worker"
)

func main() {
	config.Load()
	logging.Init()

	ctx := context.Background()
	ctx, ctxCancelFunc := context.WithCancel(ctx)
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

	ec := make(chan error)
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM) // nolint

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	log.Warn("cmd: calling context cancel func")
	// 	ctxCancelFunc()
	// }()

	// go func() {
	// 	time.Sleep(90 * time.Second)
	// 	log.Warn("cmd: throwing error to test")
	// 	// panic("my error test")
	// 	ec <- errors.New("cmd: some dummy error")
	// }()

	// time.Sleep(3 * time.Second)
	// ctxCancelFunc()

	select {
	case err := <-ec:
		log.Warn("cmd: received message on error channel")
		// ctxCancelFunc()
		log.Error(err)
	case <-sc:
		log.Warn("cmd: shutdown requested")
	case <-ctx.Done():
		log.Warn("cmd: context done, context cancel func called")
	}

	if err := w.Stop(ctx); err != nil {
		log.Error(err)
	}

	mqtt.Disconnect(ctx)
	influxdb.Disconnect(ctx)

	log.Info("cmd: graceful shutdown complete")
}
