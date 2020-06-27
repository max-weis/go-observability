package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"monitoring/hello"
	"net/http"
)

var (
	helloCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "say_hello_total",
		Help: "The total number of said hellos",
	})

	messageCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "say_message_total",
		Help: "The total number of said messages",
	})
)

type helloHandler struct {
	s      hello.Service
	logger zap.Logger
}

func (h *helloHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.sayHello)
	r.Route("/{message}", func(r chi.Router) {
		r.Get("/", h.sayMessage)
	})
	return r
}

func (h *helloHandler) sayHello(w http.ResponseWriter, r *http.Request) {
	helloCounter.Inc()
	h.logger.Info("say hello")
	ctx := context.Background()

	message := h.s.SayHello()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		encodeError(ctx, err, w)
		return
	}
}

func (h *helloHandler) sayMessage(w http.ResponseWriter, r *http.Request) {
	messageCounter.Inc()
	ctx := context.Background()

	msg := chi.URLParam(r, "message")
	h.logger.Info(fmt.Sprintf("say message: %s", msg))

	message, err := h.s.SayMessage(msg)
	if err != nil {
		encodeError(ctx, err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		encodeError(ctx, err, w)
		return
	}
}