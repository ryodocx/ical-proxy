package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/ryodocx/ical-proxy/pkg/converter"
	"github.com/ryodocx/ical-proxy/pkg/feed"
)

type Config struct {
	Addr      string
	Path      string
	Query     url.Values
	Feed      feed.Feed
	Converter *converter.Converter
}

type Server struct {
	s *http.Server
	f feed.Feed
	c *converter.Converter
	q url.Values
}

func New(c *Config) (*Server, error) {
	mux := http.NewServeMux()
	s := &Server{
		s: &http.Server{
			Addr:    c.Addr,
			Handler: mux,
		},
		f: c.Feed,
		c: c.Converter,
	}
	mux.HandleFunc("/healthz", s.healthcheck)
	mux.HandleFunc(path.Join("/", c.Path), s.simpleIcal)
	return s, nil
}

func (s *Server) healthcheck(w http.ResponseWriter, req *http.Request) {
	if err := s.f.Healthcheck(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		jsonResp, _ := json.Marshal(map[string]string{
			"error": err.Error(),
		})
		w.Write(jsonResp)
	} else {
		w.WriteHeader(http.StatusOK)
		jsonResp, _ := json.Marshal(map[string]string{
			"msg": "ok",
		})
		w.Write(jsonResp)
	}
}

func (s *Server) simpleIcal(w http.ResponseWriter, req *http.Request) {

	// TODO: validate query params

	jsons, err := s.f.Get()
	if err != nil {
		w.Header().Add("REASON", "error occurred when Get()")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmp := []interface{}{}
	for _, j := range jsons {
		v := map[string]interface{}{}
		if err := json.Unmarshal([]byte(j), &v); err != nil {
			w.Header().Add("REASON", "error occurred when json.Unmarshal()")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmp = append(tmp, v)
	}

	ical, err := s.c.SimpleIcal(tmp)
	if err != nil {
		w.Header().Add("REASON", "error occurred when convert to iCal format")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(ical))
}

func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.s.Shutdown(ctx)
}
