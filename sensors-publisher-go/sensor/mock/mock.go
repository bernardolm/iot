package mock

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type mock struct {
	id string
}

func (a *mock) Get() (interface{}, error) {
	value := aws.Float32(rand.Float32() * 85.96)

	if value == nil {
		return nil, fmt.Errorf("mock.get: fail to get value")
	}

	log.WithField("sensor", "mock").WithField("value", *value).Debug("getting values")

	return *value, nil
}

func (a *mock) DeviceClass() string {
	return "pressure"
}

func (a *mock) ID() string {
	return a.id
}

func (a *mock) Manufacturer() string {
	return "Unknown"
}

func (a *mock) Model() string {
	return "mock"
}

func (a *mock) Name() string {
	return fmt.Sprintf("%s %s sensor", a.Model(), a.DeviceClass())
}

func (a *mock) UniqueID() string {
	return fmt.Sprintf("%s_%s", a.ID(), a.DeviceClass())
}

func (a *mock) UnitOfMeasurement() string {
	return "hPa"
}

func New() *mock {
	return &mock{
		id: "mock_" + uuid.New().String(),
	}
}