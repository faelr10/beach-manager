package middleware

import (
	"context"
	"net/http"
	"strings"

	"go/beach-manager/internal/provider"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(jwtProvider *provider.JWTProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			userID, err := jwtProvider.DecodeToken(token)
			if err != nil {
				// Tentativa de renovar usando refresh token
				refreshToken := r.Header.Get("X-Refresh-Token")
				newAccessToken, newUserID, refreshErr := jwtProvider.RefreshAccessToken(refreshToken)
				if refreshErr != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				// Sucesso: define novo token no header da resposta
				w.Header().Set("X-New-Access-Token", newAccessToken)
				ctx := context.WithValue(r.Context(), UserIDKey, newUserID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Token válido
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
