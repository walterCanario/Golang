package jugador

import (
	"encoding/json"
	"hexagonal/internal/domain"
	"hexagonal/internal/services"
	"net/http"
	"strconv"
)

type JugadorHandler struct {
	service *services.JugadorService
}

func NewJugadorHandler(service *services.JugadorService) *JugadorHandler {
	return &JugadorHandler{service: service}
}

func (h *JugadorHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var j domain.Jugador
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		http.Error(w, "Error en el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if err := h.service.RegistrarJugador(&j); err != nil {
		http.Error(w, "Error al registrar jugador", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(j)
}

func (h *JugadorHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	jugador, err := h.service.ObtenerPorID(id)
	if err != nil {
		http.Error(w, "Jugador no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(jugador)
}

func (h *JugadorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	jugadores, err := h.service.ObtenerTodos()
	if err != nil {
		http.Error(w, "Error al obtener jugadores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(jugadores)
}

func (h *JugadorHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var j domain.Jugador
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		http.Error(w, "Error en el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if err := h.service.ActualizarJugador(&j); err != nil {
		http.Error(w, "Error al actualizar jugador", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(j)
}

func (h *JugadorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.EliminarJugador(id); err != nil {
		http.Error(w, "Error al eliminar jugador", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
