package worker

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
	"github.com/bernardolm/iot/sensors-publisher-go/formatter"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

type worker struct {
	delta time.Duration
	flows []flow
	// sc chan os.Signal
}

func (w *worker) AddFlow(s sensor.Sensor, f formatter.Formatter, p []publisher.Publisher) {
	w.flows = append(w.flows, flow{
		sensor:     s,
		formatter:  f,
		publishers: p,
	})
}

func (w *worker) Start(ctx context.Context) {
	log.Debug("worker: starting")
	go func() {
		for {
			for _, flow := range w.flows {
				flow.Start(ctx)
			}
			log.Debugf("worker: waiting %s", w.delta)
			time.Sleep(w.delta)
		}
	}()
}

func (w *worker) Stop(_ context.Context) error {
	log.Info("worker: stopped")
	return nil
}

func New() *worker {
	w := worker{
		delta: config.Get[time.Duration]("WORKER_DELTA"),
	}

	if w.delta == 0 {
		w.delta = 5 * time.Second
	}

	return &w
}
