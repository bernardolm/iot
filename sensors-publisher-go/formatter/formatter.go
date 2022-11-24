package formatter

type Formatter interface {
	Availability() (topic string, payload string, err error)
	Config() (topic string, payload string, err error)
	State(interface{}) (topic string, payload interface{}, err error)
}
