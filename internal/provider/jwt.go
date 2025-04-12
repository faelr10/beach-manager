package provider

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTProvider(secretKey string, duration time.Duration) *JWTProvider {
	return &JWTProvider{
		secretKey:     secretKey,
		tokenDuration: duration,
	}
}

func (p *JWTProvider) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(p.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(p.secretKey))
}
