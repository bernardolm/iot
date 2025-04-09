package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// Load env vars and config file to get app config
func Load() {
	log.SetLevel(log.DebugLevel)

	if err := godotenv.
		Load(
			"/sensors-publisher-go/.env",
			"./dev.env",
		); err != nil {
		log.WithError(err).Warn("error loading some env file")
	}
}

// Get value from env vars or config file
func Get[T any](key string) T {
	var out T
	var value any

	rawValue := os.Getenv(key)
	if rawValue == "" {
		log.Debugf("not found value for key %s or doesn't exist", key)
		return out
	}

	log.Debugf("value '%s' was recovered by ENV key '%s'", rawValue, key)

	switch any(out).(type) {
	case bool:
		value = cast.ToBool(rawValue)
	case float64:
		value = cast.ToFloat64(rawValue)
	case int:
		value = cast.ToInt(rawValue)
	case int64:
		value = cast.ToInt64(rawValue)
	case string:
		value = rawValue
	case time.Duration:
		value = cast.ToDuration(rawValue)
	default:
		value = *new(T)
		log.Errorf("unknown type to cast for key %s", key)
	}

	return value.(T)
}
