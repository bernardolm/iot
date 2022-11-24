package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var (
	client mqtt.Client
)

var debugger mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Connect() error {
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
		SetClientID(clientID).
		SetUsername(viper.GetString("MQTT_USER")).
		SetPassword(viper.GetString("MQTT_PASSWORD"))

	opts.SetDefaultPublishHandler(debugger)

	// opts.SetKeepAlive(60 * time.Second)
	// opts.SetPingTimeout(1 * time.Second)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func Publish(topic string, payload interface{}) {
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
}

func Disconnect() {
	client.Disconnect(250)
	time.Sleep(1 * time.Second)
}
