// Api/Handler/Jugador/jugador.go
package jugador

import (
	"encoding/json"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
	"net/http"
)

type Handler struct {
	service port.JugadorService
}

func NewHandler(service port.JugadorService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CrearJugador(w http.ResponseWriter, r *http.Request) {
	var j domain.Jugador
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CrearJugador(&j); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(j)
}

// Implementar demás handlers...
