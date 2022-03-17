package api

type Publisher interface {
	Publish(topic string, message string) error
}
