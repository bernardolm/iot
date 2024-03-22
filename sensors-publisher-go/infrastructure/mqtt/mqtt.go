package mqtt

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"

	"github.com/bernardolm/iot/sensors-publisher-go/config"
)

var client mqtt.Client

func Connect(_ context.Context) error {
	log.Debug("mqtt: trying to connect")

	if config.Get[bool]("DEBUG") {
		mqtt.DEBUG = log.StandardLogger()
	}

	host := config.Get[string]("MQTT_HOST")
	if host == "" {
		host = "localhost"
	}

	port := config.Get[int]("MQTT_PORT")
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
		SetPassword(config.Get[string]("MQTT_PASSWORD")).
		SetReconnectingHandler(reconnecthandler).
		SetUsername(config.Get[string]("MQTT_USERNAME"))

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		log.Info("mqtt: connected")
	}

	return nil
}

func Send(_ context.Context, topic string, payload interface{}) {
	log.Debug("mqtt: publishing")
	token := client.Publish(topic, 0, true, payload)
	go func() {
		_ = token.Wait()
		if token.Error() != nil {
			log.WithError(token.Error()).Error("mqtt: fail to publish")
		}
		log.WithField("topic", topic).
			WithField("payload", fmt.Sprintf("%s", payload)).
			Info("mqtt: sent")
	}()
}

func Close(_ context.Context) {
	log.Debug("mqtt: disconnecting")
	client.Disconnect(500)
	log.Info("mqtt: disconnected")
}
