package converter

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/open-policy-agent/opa/rego"
	"github.com/ryodocx/ical-proxy/pkg/util"
)

func init() {
	rego.New()
}

type Config struct {
	RegoPaths []string
	RegoQuery string
}

type Converter struct {
	rego *rego.PreparedEvalQuery
}

func New(c *Config) (*Converter, error) {

	r := rego.New(
		rego.Load(c.RegoPaths, nil),
		rego.Query(c.RegoQuery),
	)

	if _, err := r.Compile(context.Background()); err != nil {
		return nil, util.WrapError(err, "compile error")
	}

	pq, err := r.PrepareForEval(context.Background())
	if err != nil {
		return nil, util.WrapError(err, "prepare error")
	}

	return &Converter{
		rego: &pq,
	}, nil
}

func (s *Converter) SimpleIcal(input []interface{}) (string, error) {
	output := "BEGIN:VCALENDAR"

	for _, i := range input {
		rs, err := s.rego.Eval(context.Background(), rego.EvalInput(i))
		if err != nil {
			return "", util.WrapError(err, "rego eval error")
		}

		jsonBytes, err := json.Marshal(rs)
		if err != nil {
			return "", util.WrapError(err, "json.Marshal error")
		}

		v := []struct {
			Expressions []struct {
				Value struct {
					Allowed bool              `json:"allowed"`
					Event   map[string]string `json:"event"`
				} `json:"value"`
			} `json:"expressions"`
		}{}
		if err := json.Unmarshal(jsonBytes, &v); err != nil {
			return "", util.WrapError(err, "json.Unmarshal error")
		}

		if !v[0].Expressions[0].Value.Allowed {
			continue
		}

		output = fmt.Sprintf("%s\nBEGIN:VEVENT", output)
		for k, v := range v[0].Expressions[0].Value.Event {
			output = fmt.Sprintf("%s\n%s:%s", output, k, strings.Trim(fmt.Sprintf("%q", v), `"`))
		}
		output = fmt.Sprintf("%s\nEND:VEVENT", output)
	}

	output = fmt.Sprintf("%s\nEND:VCALENDAR\n", output)
	return output, nil
}
