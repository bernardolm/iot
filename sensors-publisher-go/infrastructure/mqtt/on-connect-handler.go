package mqtt

import (
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var soch = sync.Once{}

var onConnectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
	soch.Do(func() {
		log.Info("mqtt: connected now")
		time.Sleep(5 * time.Second)
	})
}
