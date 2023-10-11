package mqtt

import (
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var sorh = sync.Once{}

var reconnecthandler mqtt.ReconnectHandler = func(c mqtt.Client, options *mqtt.ClientOptions) {
	sorh.Do(func() {
		log.Warn("mqtt: client reconnected")
		time.Sleep(5 * time.Second)
	})
}
