package mqtt

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var client mqtt.Client

func Connect(_ context.Context) error {
	log.Debug("mqtt: trying to connect")

	if viper.GetBool("DEBUG") {
		mqtt.DEBUG = log.StandardLogger()
	}

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

	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetAutoReconnect(true).
		SetClientID(clientID).
		SetConnectionLostHandler(connectionLostHandler).
		SetConnectRetry(true).
		SetOnConnectHandler(onConnectHandler).
		SetPassword(viper.GetString("MQTT_PASSWORD")).
		SetReconnectingHandler(reconnecthandler).
		SetUsername(viper.GetString("MQTT_USER"))

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		log.Info("mqtt: connected")
	}

	return nil
}

func Publish(topic string, payload interface{}) {
	log.Debug("mqtt: publishing")
	token := client.Publish(topic, 0, true, payload)
	go func() {
		_ = token.Wait()
		if token.Error() != nil {
			log.WithError(token.Error()).Error("mqtt: fail to publish")
		}
		log.Debug("mqtt: published")
	}()
}

func Disconnect(_ context.Context) {
	log.Debug("mqtt: disconnecting")
	client.Disconnect(500)
	log.Info("mqtt: disconnected")
}
