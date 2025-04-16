package handlers

import (
	"encoding/json"
	"go/beach-manager/internal/dto"
	"go/beach-manager/internal/service"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.authService.Login(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// ❌ Não salva mais o refresh_token no cookie
	// http.SetCookie(...) foi removido

	// ✅ Retorna o refresh_token no corpo da resposta
	response := map[string]string{
		"token":         output.Token,
		"refresh_token": output.RefreshToken,
		"email":         output.Email,
		"id":            output.UserID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Estrutura para decodificar o JSON recebido
	var reqBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	// Decodifica o corpo JSON
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil || reqBody.RefreshToken == "" {
		http.Error(w, "Refresh token ausente ou inválido", http.StatusBadRequest)
		return
	}

	// Gera novo access_token com base no refresh_token
	newAccessToken, _, err := h.authService.RefreshAccessToken(reqBody.RefreshToken)
	if err != nil {
		http.Error(w, "Refresh token inválido", http.StatusUnauthorized)
		return
	}

	// Retorna novo access token no corpo da resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": newAccessToken,
	})
}
