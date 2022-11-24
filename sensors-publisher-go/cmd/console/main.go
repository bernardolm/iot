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

	// publishermock "github.com/bernardolm/iot/sensors-publisher-go/publisher/mock"
	sensords18a20 "github.com/bernardolm/iot/sensors-publisher-go/sensor/ds18a20"
	sensormock "github.com/bernardolm/iot/sensors-publisher-go/sensor/mock"
	"github.com/bernardolm/iot/sensors-publisher-go/worker"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	logging.Init()

	w := worker.New()

	sm := sensormock.New()
	w.AddSensor(sm)

	ham := homeassistant.New(sm)
	w.AddFormatter(ham)

	sd := sensords18a20.New()
	w.AddSensor(sd)

	had := homeassistant.New(sd)
	w.AddFormatter(had)

	mb := messagebus.New()
	w.AddPublisher(mb)

	// pm := publishermock.New()
	// w.AddPublisher(pm)

	go func() {
		for {
			w.Do()
			time.Sleep(20 * time.Second)
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
