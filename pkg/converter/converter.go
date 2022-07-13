package converter

import "github.com/open-policy-agent/opa/rego"

func init() {
	rego.New()
}

type Config struct {
}

type Converter struct {
}

func New(c *Config) (*Converter, error) {
	return nil, nil // TODO
}

func (s *Converter) Convert() {
	// TODO
}
