package logging

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	ll := viper.GetString("LOG_LEVEL")

	var level log.Level = log.InfoLevel

	if ll != "" {
		level, _ = log.ParseLevel(ll)
	}

	log.SetLevel(level)
}
