package main

import (
	"fmt"

	"github.com/ryodocx/ical-proxy/pkg/converter"
	"github.com/ryodocx/ical-proxy/pkg/feed/redmine/issue"
	"github.com/ryodocx/ical-proxy/pkg/feed/redmine/version"
	"github.com/ryodocx/ical-proxy/pkg/server"
)

func main() {
	issue.New(nil)
	version.New(nil)
	converter.New(nil)
	server.New(nil)
	fmt.Println("hello world")
}
