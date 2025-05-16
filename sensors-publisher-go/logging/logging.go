package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"

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

	log.SetLevel(level)
	log.SetOutput(os.Stdout)

	// hook, err := lSyslog.
	// 	NewSyslogHook(
	// 		"udp",
	// 		config.Get[string]("SYSLOG_HOST"),
	// 		syslog.LOG_NOTICE,
	// 		"")
	// if err == nil {
	// 	log.Hooks.Add(hook)
	// }

	lokiOpts := lokirus.NewLokiHookOptions().
		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"service": "sensors-publisher-go",
		})
	lokiHook := lokirus.NewLokiHookWithOpts(
		config.Get[string]("LOKI_HOST"), lokiOpts)
	log.AddHook(lokiHook)
}
