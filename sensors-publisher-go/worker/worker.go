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
	for sensor := range w.sensor {
		value, err := w.sensor[sensor].Get()
		if err != nil {
			log.Error(err)
			continue
		}

		for formatter := range w.formatter {
			configTopic, configMessage, err := w.formatter[formatter].Config()
			if err != nil {
				log.Error(err)
				continue
			}

			if configTopic != "" && configMessage != "" {
				for publisher := range w.publisher {
					err := w.publisher[publisher].Do(configTopic, configMessage)
					if err != nil {
						log.Error(err)
					}
				}
			}

			stateTopic, stateMessage, err := w.formatter[formatter].State(value)
			if err != nil {
				log.Error(err)
				continue
			}

			if stateTopic != "" && stateMessage != "" {
				for publisher := range w.publisher {
					err := w.publisher[publisher].Do(stateTopic, stateMessage)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
	}
}

func New() *worker {
	return &worker{}
}
