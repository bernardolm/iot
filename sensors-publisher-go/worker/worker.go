package worker

import (
	"github.com/bernardolm/iot/sensors-publisher-go/formatter"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
	log "github.com/sirupsen/logrus"
)

type worker struct {
	formatter []formatter.Formatter
	publisher []publisher.Publisher
	sensor    []sensor.Sensor
}

func (w *worker) AddSensor(s sensor.Sensor) {
	w.sensor = append(w.sensor, s)
}

func (w *worker) AddPublisher(p publisher.Publisher) {
	w.publisher = append(w.publisher, p)
}

func (w *worker) AddFormatter(f formatter.Formatter) {
	w.formatter = append(w.formatter, f)
}

func (w *worker) Do() {
	for s := range w.sensor {
		m, err := w.sensor[s].Get()
		if err != nil {
			log.Error(err)
			continue
		}

		for f := range w.formatter {
			mf, err := w.formatter[f].Do(m)
			if err != nil {
				log.Error(err)
				continue
			}

			for p := range w.publisher {
				err := w.publisher[p].Do(mf)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}

func New() *worker {
	return &worker{}
}
