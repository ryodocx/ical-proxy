package issue

import (
	"encoding/json"
	"net/url"

	redmineClient "github.com/ryodocx/go-redmine"
)

type Config struct {
	Url        string
	Query      url.Values
	ApiKey     string
	MaxEntries int
}

type redmine struct {
	config *Config
	client *redmineClient.Client
}

func New(c *Config) (*redmine, error) {
	newClient, err := redmineClient.NewClient(c.Url, c.ApiKey)
	if err != nil {
		return nil, err
	}
	return &redmine{
		config: c,
		client: newClient,
	}, nil
}

func (s *redmine) Get() (entries []string, err error) {
	tmp, err := s.client.GetIssues(s.config.Query, s.config.MaxEntries)
	if err != nil {
		return nil, err
	}
	for _, v := range tmp {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		entries = append(entries, string(b))
	}
	return entries, nil
}

func (s *redmine) Healthcheck() error {
	return s.client.HealthCheck()
}
