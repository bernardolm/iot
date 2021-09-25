package mock

import (
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

type mock struct{}

func (m *mock) Get() (interface{}, error) {
	log.Debug("getting values")
	return fmt.Sprint(rand.Intn(100-20) + 20), nil
}

func (m *mock) DeviceClass() string {
	return "temperature"
}

func (m *mock) ID() string {
	return "some_random_device_id"
}

func (m *mock) Manufacturer() string {
	return "Unknown"
}

func (m *mock) Model() string {
	return "mock"
}

func (m *mock) UnitOfMeasurement() string {
	return "°F"
}

func New() *mock {
	return &mock{}
}
