package publisher

import "context"

type Publisher interface {
	Publish(_ context.Context, topic string, message interface{}) error
}
