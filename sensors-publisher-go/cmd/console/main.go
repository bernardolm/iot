package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
	"github.com/bernardolm/iot/sensors-publisher-go/formatter/homeassistant"
	"github.com/bernardolm/iot/sensors-publisher-go/logging"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher/messagebus"
	sensormock "github.com/bernardolm/iot/sensors-publisher-go/sensor/mock"
	"github.com/bernardolm/iot/sensors-publisher-go/worker"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	logging.Init()

	w := worker.New()

	ds := sensormock.New()
	w.AddSensor(ds)

	ha := homeassistant.New(ds)
	w.AddFormatter(ha)

	mb := messagebus.New()
	w.AddPublisher(mb)

	go func() {
		for {
			w.Do()
			time.Sleep(2 * time.Second)
		}
	}()

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
