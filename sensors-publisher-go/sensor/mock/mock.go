package mock

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	log "github.com/sirupsen/logrus"
)

type mock struct{}

func (m *mock) Get() (interface{}, error) {
	var value *int
	value = aws.Int(rand.Intn(100-20) + 20)

	if value == nil {
		return nil, fmt.Errorf("mock.get: fail to get value")
	}

	log.WithField("sensor", "mock").WithField("value", *value).Debug("getting values")
	return *value, nil

}

func (m *mock) DeviceClass() string {
	return "pressure"
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
