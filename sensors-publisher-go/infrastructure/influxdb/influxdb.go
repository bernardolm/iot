package influxdb

import (
	"context"
	"crypto/tls"
	"fmt"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	influxdbapi "github.com/influxdata/influxdb-client-go/v2/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	api    influxdbapi.WriteAPI
	client influxdb.Client
)

func Connect(_ context.Context) error {
	log.Debug("influxdb: trying to connect")

	host := viper.GetString("INFLUXDB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := viper.GetInt("INFLUXDB_PORT")
	if port == 0 {
		port = 8086
	}

	url := fmt.Sprintf("http://%s:%d", host, port)

	token := viper.GetString("INFLUXDB_TOKEN")

	opts := influxdb.DefaultOptions().SetTLSConfig(&tls.Config{
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

	database := viper.GetString("INFLUXDB_DATABASE")
	if database == "" {
		database = "test/autogen"
	}

	api = client.WriteAPI("my-org", database)

	errorsCh := api.Errors()
	go func() {
		for err := range errorsCh {
			log.Errorf("influxdb: api write error - %s\n", err.Error())
		}
	}()

	return nil
}

func Publish(topic string, payload interface{}) {
	log.Debug("influxdb: publishing")
	line := payload.([]byte)
	api.WriteRecord(string(line))
	api.Flush()
	log.Debug("influxdb: published")
}

func Disconnect(_ context.Context) {
	log.Debug("influxdb: disconnecting")
	client.Close()
	log.Info("influxdb: disconnected")
}
