package redmine

import "github.com/ryodocx/ical-proxy/pkg/feed"

type Config struct {
}

type redmine struct {
}

func New(c *Config) (*redmine, error) {
	return nil, nil // TODO
}

func (s *redmine) Get() ([]*feed.Entry, error) {
	return nil, nil // TODO
}

func (s *redmine) Healthcheck() error {
	return nil // TODO
}
