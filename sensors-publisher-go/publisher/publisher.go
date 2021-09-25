package publisher

type Publisher interface {
	Do(interface{}) error
}
