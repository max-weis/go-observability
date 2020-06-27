package server

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type healthHandler struct {
}

func (h *healthHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.health)
	return r
}

func (h *healthHandler) health(w http.ResponseWriter, r *http.Request) {
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
