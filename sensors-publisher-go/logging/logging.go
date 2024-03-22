package logging

import (
	// "log/syslog"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	// lSyslog "github.com/sirupsen/logrus/hooks/syslog"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
)

func Init() {
	var level log.Level = log.InfoLevel

	if ll := config.Get[string]("LOG_LEVEL"); ll != "" {
		var err error
		if level, err = log.ParseLevel(ll); err != nil {
			log.WithError(err).Error("logging failed to set log level")
		}
	}

	log := logrus.New()
	log.SetLevel(level)

	// hook, err := lSyslog.
	// 	NewSyslogHook(
	// 		"udp",
	// 		config.Get[string]("SYSLOG_HOST"),
	// 		syslog.LOG_NOTICE,
	// 		"")
	// if err == nil {
	// 	log.Hooks.Add(hook)
	// }
}
