// api/handler/jugador/jugador.go
package jugador

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hexagonal/internal/services/jugador"

	"github.com/gorilla/mux"
)

type JugadorHandler struct {
	createHandler  *CreateHandler
	getByIDHandler *GetByIDHandler
	updateHandler  *UpdateHandler
	deleteHandler  *DeleteHandler
	getAllHandler  *GetAllHandler
}

func NewJugadorHandler(service jugador.Service) *JugadorHandler {
	return &JugadorHandler{
		createHandler:  NewCreateHandler(service),
		getByIDHandler: NewGetByIDHandler(service),
		updateHandler:  NewUpdateHandler(service),
		deleteHandler:  NewDeleteHandler(service),
		getAllHandler:  NewGetAllHandler(service),
	}
}

// Métodos adicionales para registrar rutas
func (h *JugadorHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/jugadores", h.getAllHandler.Handle).Methods("GET")
	router.HandleFunc("/jugadores", h.createHandler.Handle).Methods("POST")
	router.HandleFunc("/jugadores/{id}", h.getByIDHandler.Handle).Methods("GET")
	router.HandleFunc("/jugadores/{id}", h.updateHandler.Handle).Methods("PUT")
	router.HandleFunc("/jugadores/{id}", h.deleteHandler.Handle).Methods("DELETE")
}

// GetByIDHandler
type GetByIDHandler struct {
	service jugador.Service
}

func NewGetByIDHandler(service jugador.Service) *GetByIDHandler {
	return &GetByIDHandler{
		service: service,
	}
}

func (h *GetByIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jugador, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jugador)
}

// Implementaciones similares para UpdateHandler, DeleteHandler y GetAllHandler
