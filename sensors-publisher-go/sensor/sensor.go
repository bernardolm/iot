package sensor

type Type string

const (
	DS18A20 Type = "DS18A20"
)

type Sensor interface {
	Get() (interface{}, error)
}
