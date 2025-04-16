package handlers

import (
	"encoding/json"
	"go/beach-manager/internal/dto"
	"go/beach-manager/internal/service"
	"net/http"
	"time"
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

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    output.RefreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode, // necessário para cross-domain
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})	
	

	// Retorna apenas os dados públicos e access_token no corpo
	response := map[string]string{
		"token": output.Token,
		"email": output.Email,
		"id":    output.UserID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Lê o cookie do refresh_token
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token ausente", http.StatusUnauthorized)
		return
	}

	refreshToken := cookie.Value

	// Gera novo access_token com base no refresh_token
	newAccessToken, _, err := h.authService.RefreshAccessToken(refreshToken)
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
