package stdout

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type stdout struct{}

func (a *stdout) Publish(_ context.Context, topic string, message interface{}) error {
	if message == nil {
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
