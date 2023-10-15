package influxdb

import (
	"context"
	"crypto/tls"
	"fmt"

	influxdb "github.com/influxdata/influxdb-client-go"
	influxdbapi "github.com/influxdata/influxdb-client-go/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	client influxdb.Client
	api    influxdbapi.WriteAPI

	database, host, password, port, username string
)

func loadConfig() {
	viper.SetDefault("INFLUX_DATABASE", "test")
	viper.SetDefault("INFLUX_HOST", "localhost")
	viper.SetDefault("INFLUX_PORT", "8086")

	database = viper.GetString("INFLUX_DATABASE")
	host = viper.GetString("INFLUX_HOST")
	password = viper.GetString("INFLUX_PASSWORD")
	port = viper.GetString("INFLUX_PORT")
	username = viper.GetString("INFLUX_USERNAME")
}

func Connect(_ context.Context) error {
	log.Debug("influxdb: trying to connect")

	loadConfig()

	url := fmt.Sprintf("http://%s:%s", host, port)
	token := fmt.Sprintf("%s:%s", username, password)

	opts := influxdb.DefaultOptions().
		SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	client = influxdb.NewClientWithOptions(url, token, opts)

	hc, err := client.Health(context.Background())
	if err != nil {
		return err
	}
	if client == nil {
		return fmt.Errorf("influxdb: couldn't create a client")
	}

	log.
		WithField("status", hc.Status).
		WithField("message", *hc.Message).
		Info("influxdb: connected")

	api = client.WriteAPI("dummy-org", database)

	errorsCh := api.Errors()
	go func() {
		for err := range errorsCh {
			log.Errorf("influxdb: api write error - %s\n", err.Error())
		}
	}()

	return nil
}

func Send(_ string, payload interface{}) {
	log.Debug("influxdb: publishing")
	line := payload.([]byte)
	api.WriteRecord(string(line))
	api.Flush()
	log.
		WithField("payload", fmt.Sprintf("%s", payload)).
		Info("influxdb: sent (writed)")
}

func Disconnect(_ context.Context) {
	log.Debug("influxdb: disconnecting")
	client.Close()
	log.Info("influxdb: disconnected")
}
