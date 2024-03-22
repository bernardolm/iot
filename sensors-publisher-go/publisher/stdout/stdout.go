package stdout

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
)

type stdout struct{}

func (a *stdout) Publish(_ context.Context, topic string, message interface{}) error {
	if !config.Get[bool]("DEBUG") || message == nil {
		return nil
	}

	log.WithField("message", string(message.([]byte))).
		WithField("publisher", "stdout").
		Debug("publisher: trying to publish")

	return nil
}

func New() *stdout {
	return &stdout{}
}
