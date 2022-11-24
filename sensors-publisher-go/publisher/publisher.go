package publisher

type Publisher interface {
	Do(topic string, message interface{}) error
}
