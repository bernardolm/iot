package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Load env vars and config file to get app config
func Load() {
	viper.AutomaticEnv()

	viper.SetConfigFile("config.env")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError, *os.PathError:
			// NOTE: Need to log out to console regardless of log level
			log.Info("using config from env vars instead config file")
		default:
			log.WithError(err).Error("failed to load config using viper")
		}
	}

	viper.Debug()
}
