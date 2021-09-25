package mock

type mock struct{}

func (m *mock) Do(T string) error {
	return nil
}
