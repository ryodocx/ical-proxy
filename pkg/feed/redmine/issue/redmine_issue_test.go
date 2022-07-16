package issue_test

import (
	"os"
	"testing"

	version "github.com/ryodocx/ical-proxy/pkg/feed/redmine/issue"
)

func TestXxx(t *testing.T) {
	feed, err := version.New(
		&version.Config{
			Url:        os.Getenv("ICALPROXY_REDMINE_URL"),
			ApiKey:     os.Getenv("ICALPROXY_REDMINE_APIKEY"),
			MaxEntries: 3,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := feed.Healthcheck(); err != nil {
		t.Error(err)
	}

	if entries, err := feed.Get(); err != nil {
		t.Error(err)
	} else {
		for _, e := range entries {
			t.Log(e)
		}
	}
}
