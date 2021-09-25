package publisher

type Publisher interface {
	Do(string, interface{}) error
}
