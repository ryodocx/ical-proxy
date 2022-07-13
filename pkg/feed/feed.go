package feed

type Feed interface {
	Get() (jsonEntries []string, err error)
	Healthcheck() error
}
