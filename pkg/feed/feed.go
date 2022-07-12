package feed

type Entry struct {
}

type Feed interface {
	Get() ([]*Entry, error)
	Healthcheck() error
}
