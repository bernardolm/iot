package worker

import (
	"time"

	"github.com/bernardolm/iot/sensors-publisher-go/formatter"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type worker struct {
	delta time.Duration
	flows []flow
}

func (w *worker) AddFlow(s sensor.Sensor, f formatter.Formatter, p []publisher.Publisher) {
	w.flows = append(w.flows, flow{
		sensor:     s,
		formatter:  f,
		publishers: p,
	})
}

func (w *worker) Start() {
	go func() {
		for {
			for _, flow := range w.flows {
				flow.Start()
			}
			log.Debug("worker waiting...")
			time.Sleep(w.delta)
		}
	}()
}

func New() *worker {
	w := worker{
		delta: viper.GetDuration("WORKER_DELTA"),
	}

	if w.delta == 0 {
		w.delta = 5 * time.Second
	}

	return &w
}
