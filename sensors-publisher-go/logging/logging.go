package logging

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	// set log level
	log.SetLevel(log.DebugLevel)
}
