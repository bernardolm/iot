package sensor

type Type string

const (
	DS18A20 Type = "DS18A20"
)

type Sensor interface {
	DeviceClass() string
	Get() (interface{}, error)
	ID() string
	Manufacturer() string
	Model() string
	UnitOfMeasurement() string
}
