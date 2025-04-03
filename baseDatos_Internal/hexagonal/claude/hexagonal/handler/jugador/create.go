// api/handler/jugador/create.go
package jugador

import (
	"encoding/json"
	"net/http"

	"hexagonal/internal/domain"
	"hexagonal/internal/services/jugador"
)

type CreateHandler struct {
	service jugador.Service
}

func NewCreateHandler(service jugador.Service) *CreateHandler {
	return &CreateHandler{
		service: service,
	}
}

func (h *CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var jugadorDTO domain.JugadorDTO
	err := json.NewDecoder(r.Body).Decode(&jugadorDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	createdJugador, err := h.service.Create(r.Context(), jugadorDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdJugador)
}
