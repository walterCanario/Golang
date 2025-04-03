package child

import (
	"encoding/json"
	"net/http"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/service"
	"strconv"
)

type ChildHandler struct {
	childService *service.ChildService
}

func NewChildHandler(childService *service.ChildService) *ChildHandler {
	return &ChildHandler{childService: childService}
}

func (h *ChildHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /children", h.CreateChild)
	mux.HandleFunc("GET /children", h.GetAllChildren)
	mux.HandleFunc("GET /children/{id}", h.GetChildByID)
	mux.HandleFunc("PUT /children", h.UpdateChild)
	mux.HandleFunc("DELETE /children", h.DeleteChild)
	mux.HandleFunc("GET /children/user/{userId}", h.GetChildrenByUserID)
	mux.HandleFunc("GET /children/siblings/{childId}", h.GetSiblings)
	mux.HandleFunc("GET /children/{id}/details", h.GetChildWithUserDetails) // Nueva ruta
}

func (h *ChildHandler) CreateChild(w http.ResponseWriter, r *http.Request) {
	var child domain.Child
	err := json.NewDecoder(r.Body).Decode(&child)
	if err != nil {
		http.Error(w, "Entrada no válida", http.StatusBadRequest)
		return
	}

	err = h.childService.CreateChild(&child)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) GetChildByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	child, err := h.childService.GetChildByID(id)
	if err != nil {
		http.Error(w, "Hijo no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) UpdateChild(w http.ResponseWriter, r *http.Request) {
	var child domain.Child
	err := json.NewDecoder(r.Body).Decode(&child)
	if err != nil {
		http.Error(w, "Entrada no válida", http.StatusBadRequest)
		return
	}

	err = h.childService.UpdateChild(&child)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(child)
}

func (h *ChildHandler) DeleteChild(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.childService.DeleteChild(id)
	if err != nil {
		http.Error(w, "Error al eliminar el hijo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChildHandler) GetAllChildren(w http.ResponseWriter, r *http.Request) {
	children, err := h.childService.GetAllChildren()
	if err != nil {
		http.Error(w, "Error al obtener los hijos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(children)
}

func (h *ChildHandler) GetChildrenByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.PathValue("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	children, err := h.childService.GetChildrenByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(children)
}

func (h *ChildHandler) GetSiblings(w http.ResponseWriter, r *http.Request) {
	childIDStr := r.PathValue("childId")
	childID, err := strconv.Atoi(childIDStr)
	if err != nil {
		http.Error(w, "ID de hijo inválido", http.StatusBadRequest)
		return
	}

	siblings, err := h.childService.GetSiblings(childID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(siblings)
}

// GetChildWithUserDetails maneja la solicitud para obtener los detalles de un hijo con información del usuario padre.
func (h *ChildHandler) GetChildWithUserDetails(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	childWithDetails, err := h.childService.GetChildWithUserDetails(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(childWithDetails)
}
