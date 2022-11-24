package publisher

type Publisher interface {
	Publish(topic string, message interface{}) error
}
