package version_test

import (
	"os"
	"testing"

	"github.com/ryodocx/ical-proxy/pkg/feed/redmine/version"
)

func TestXxx(t *testing.T) {
	feed, err := version.New(
		&version.Config{
			Project:    os.Getenv("REDMINE_PROJECT"),
			Url:        os.Getenv("REDMINE_URL"),
			ApiKey:     os.Getenv("REDMINE_APIKEY"),
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
