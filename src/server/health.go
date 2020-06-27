package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
	"net/http"
)

var (
	healthCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "health_total",
		Help: "The total number of health checks",
	})
)

type healthHandler struct {
	logger zap.Logger
}

func (h *healthHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.health)
	return r
}

func (h *healthHandler) health(w http.ResponseWriter, r *http.Request) {
	healthCounter.Inc()
	h.logger.Info("up")
	ctx := context.Background()

	message := struct {
		Up bool `json:"up"`
	}{
		true,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		encodeError(ctx, err, w)
		return
	}
}
