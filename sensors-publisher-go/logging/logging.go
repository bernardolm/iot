package logging

import (
	"github.com/bernardolm/iot/sensors-publisher-go/config"
	log "github.com/sirupsen/logrus"
)

func Init() {
	ll := config.Get[string]("LOG_LEVEL")

	var level log.Level = log.InfoLevel

	if ll != "" {
		level, _ = log.ParseLevel(ll)
	}

	log.SetLevel(level)
}
