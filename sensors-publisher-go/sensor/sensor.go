package sensor

type Sensor interface {
	Get() (interface{}, error)
}
