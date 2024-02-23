package influxdb

import (
	"fmt"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
)

var database, password, url, username, token string

func loadConfig() {
	database = config.Get[string]("INFLUX_DATABASE")
	if database == "" {
		database = "test"
	}

	password = config.Get[string]("INFLUX_PASSWORD")

	url = config.Get[string]("INFLUX_URL")
	if url == "" {
		url = "localhost:8086"
	}

	username = config.Get[string]("INFLUX_USERNAME")

	token = fmt.Sprintf("%s:%s", username, password)
}
