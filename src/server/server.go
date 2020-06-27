package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/httptracer"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"monitoring/hello"
	"net/http"
)

type Server struct {
	Hello hello.Service

	logger zap.Logger
	router chi.Router
	tracer opentracing.Tracer
}

func New(hs hello.Service, logger zap.Logger, tracer opentracing.Tracer) *Server {
	s := &Server{Hello: hs, logger: logger}
	health := healthHandler{logger: logger}

	r := chi.NewRouter()
	r.Use(accessControl)
	r.Use(httptracer.Tracer(tracer, httptracer.Config{
		ServiceName:    "observability-demo",
		ServiceVersion: "0.0.0",
		SampleRate:     1,
		SkipFunc: func(r *http.Request) bool {
			return r.URL.Path == "/health"
		},
	}))

	r.Route("/hello", func(r chi.Router) {
		h := helloHandler{s.Hello, logger}
		r.Mount("/v1", h.router())
	})

	r.Mount("/health", health.router())

	r.Method("GET", "/metrics", promhttp.Handler())
	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case hello.ErrEmptyMessage:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
