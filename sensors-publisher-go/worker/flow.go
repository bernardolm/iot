package worker

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/formatter"
	"github.com/bernardolm/iot/sensors-publisher-go/publisher"
	"github.com/bernardolm/iot/sensors-publisher-go/sensor"
)

type flow struct {
	sensor     sensor.Sensor
	formatter  formatter.Formatter
	publishers []publisher.Publisher
}

func (a *flow) Start(ctx context.Context) {
	messages := a.formatter.Build(a.sensor)

	for _, p := range a.publishers {
		for _, m := range messages {
			if err := p.Publish(ctx, m.Topic, m.Body); err != nil {
				log.Error(err)
			}
		}
	}
}
