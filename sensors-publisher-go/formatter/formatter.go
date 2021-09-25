package formatter

type Formatter interface {
	Do(interface{}) (interface{}, error)
}
