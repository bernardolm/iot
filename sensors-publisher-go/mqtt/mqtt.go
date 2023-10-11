package mqtt

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	client mqtt.Client
)

func Connect(_ context.Context) error {
	host := viper.GetString("MQTT_HOST")
	if host == "" {
		host = "localhost"
	}

	port := viper.GetInt("MQTT_PORT")
	if port == 0 {
		port = 1883
	}

	broker := fmt.Sprintf("tcp://%s:%d", host, port)
	clientID := "sensors-publisher-go"

	opts := mqtt.NewClientOptions().AddBroker(broker).
		SetAutoReconnect(true).
		SetCleanSession(false).
		SetClientID(clientID).
		SetKeepAlive(30 * time.Second).
		SetOrderMatters(true).
		SetPassword(viper.GetString("MQTT_PASSWORD")).
		SetPingTimeout(1 * time.Second).
		SetUsername(viper.GetString("MQTT_USER"))

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func Publish(topic string, payload interface{}) {
	token := client.Publish(topic, 0, true, payload)
	go func() {
		_ = token.Wait()
		if err := token.Error(); err != nil {
			log.WithError(err).Error("fail to mqtt publish")
		}
	}()
}

func Disconnect(_ context.Context) {
	log.Warn("mqtt: stopping")
	// client.Connect().Done()
	client.Disconnect(250)
}
