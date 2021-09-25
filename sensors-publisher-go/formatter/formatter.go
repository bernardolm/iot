package formatter

type Formatter interface {
	Availability() (string, string, error)
	Config() (string, string, error)
	State(interface{}) (string, interface{}, error)
}
