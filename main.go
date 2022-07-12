package main

import (
	"fmt"

	"github.com/ryodocx/ical-proxy/pkg/converter"
	"github.com/ryodocx/ical-proxy/pkg/feed/redmine"
	"github.com/ryodocx/ical-proxy/pkg/server"
)

func main() {
	redmine.New(nil)
	converter.New(nil)
	server.New(nil)
	fmt.Println("hello world")
}
