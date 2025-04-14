package handlers

import (
	"encoding/json"
	"go/beach-manager/internal/dto"
	"go/beach-manager/internal/service"
	"go/beach-manager/internal/web/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AgendaHandler struct {
	agendaService *service.AgendaService
}

func NewAgendaHandler(agendaService *service.AgendaService) *AgendaHandler {
	return &AgendaHandler{
		agendaService: agendaService,
	}
}

func (h *AgendaHandler) CreateAgenda(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CreateAgendaInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// atribui o ID do usuário ao input
	input.UserID = userID.(string)

	output, err := h.agendaService.CreateAgenda(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

}

func (h *AgendaHandler) GetAgendaByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	output, err := h.agendaService.GetAgendaByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (h *AgendaHandler) GetAllAgendas(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	outputs, err := h.agendaService.GetAllAgendas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputs)
}
