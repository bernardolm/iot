package influxdb

import (
	"fmt"

	"github.com/spf13/viper"
)

var database, password, url, username, token string

func loadConfig() {
	viper.SetDefault("INFLUX_DATABASE", "test")
	viper.SetDefault("INFLUX_URL", "http://localhost:8086")

	database = viper.GetString("INFLUX_DATABASE")
	password = viper.GetString("INFLUX_PASSWORD")
	url = viper.GetString("INFLUX_URL")
	username = viper.GetString("INFLUX_USERNAME")

	token = fmt.Sprintf("%s:%s", username, password)
}
