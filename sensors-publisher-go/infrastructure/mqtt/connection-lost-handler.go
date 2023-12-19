package mqtt

import (
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var socl = sync.Once{}

var connectionLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
	socl.Do(func() {
		log.WithError(err).Error("mqtt: connection lost")
		time.Sleep(5 * time.Second)
	})
}
