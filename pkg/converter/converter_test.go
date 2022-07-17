package converter_test

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/ryodocx/go-redmine/v2"
	"github.com/ryodocx/ical-proxy/pkg/converter"
)

func TestXxx(t *testing.T) {
	wd, _ := os.Getwd()
	c, err := converter.New(&converter.Config{
		RegoPaths: []string{path.Join(wd, "../../configs/")},
		RegoQuery: "data.ical.simple",
	})
	if err != nil {
		t.Error(err)
	}

	now := time.Now().Format("2006-01-02")
	tomorrow := time.Now().Add(time.Hour * 24).Format("2006-01-02")
	output, err := c.SimpleIcal([]interface{}{
		&redmine.Issue{
			ID:          1,
			Subject:     "redmine subject1",
			Description: "redmine description1",
			DueDate:     &now,
		},
		&redmine.Issue{
			ID:          2,
			Subject:     "redmine subject2",
			Description: "redmine description2",
			DueDate:     &tomorrow,
		},
		&redmine.Issue{
			ID:          3,
			Subject:     "redmine subject3",
			Description: "redmine description3",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("ical:\n%s", output))
}
