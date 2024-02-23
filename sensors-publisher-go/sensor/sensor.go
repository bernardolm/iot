package sensor

type Type string

const (
	DS18A20 Type = "DS18A20"
	Mock    Type = "Mock"
)

type Sensor interface {
	DeviceClass() string
	Get() (value interface{}, err error)
	ID() string
	Manufacturer() string
	Model() string
	Name() string
	UniqueID() string
	UnitOfMeasurement() string
}
