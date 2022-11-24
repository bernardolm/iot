package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
	"github.com/bernardolm/iot/sensors-publisher-go/formatter/homeassistant"
	"github.com/bernardolm/iot/sensors-publisher-go/logging"
	"github.com/bernardolm/iot/sensors-publisher-go/mqtt"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	publishermqtt "github.com/bernardolm/iot/sensors-publisher-go/publisher/mqtt"
	publisherstdout "github.com/bernardolm/iot/sensors-publisher-go/publisher/stdout"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor/ds18a20"
	sensormock "github.com/bernardolm/iot/sensors-publisher-go/sensor/mock"
	"github.com/bernardolm/iot/sensors-publisher-go/worker"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	logging.Init()

	if err := mqtt.Connect(); err != nil {
		log.Panic(err)
	}

	bridge := "sensors-publisher-go"

	msqt := publishermqtt.New()
	so := publisherstdout.New()
	w := worker.New()

	ds, err := ds18a20.New()
	if err != nil {
		log.Error(err)
	}

	for i := range ds {
		ha, err := homeassistant.New(bridge, ds[i])
		if err != nil {
			log.Error(err)
		}

		w.AddFlow(ds[i], ha, []publisher.Publisher{so, msqt})
	}

	if len(ds) == 0 {
		sm := sensormock.New()

		ham, err := homeassistant.New(bridge, sm)
		if err != nil {
			log.Error(err)
		}

		w.AddFlow(sm, ham, []publisher.Publisher{so, msqt})
	}

	w.Start()

	ec := make(<-chan error)
	sc := make(chan os.Signal)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM) // nolint

	select {
	case err := <-ec:
		panic(err)
	case <-sc:
		log.Info("Exiting")
	}
}
