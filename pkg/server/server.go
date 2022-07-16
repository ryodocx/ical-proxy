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

	p := path.Join(c.Path)
	if len(c.Query.Encode()) > 0 {
		p += "?" + c.Query.Encode()
	}
	mux.HandleFunc(p, s.simpleIcal)
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
	w.WriteHeader(http.StatusNotImplemented) // TODO
}

func (s *Server) ListenAndServe() error {
	return s.s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.s.Shutdown(ctx)
}
