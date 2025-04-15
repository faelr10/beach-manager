package provider

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secretKey        string
	refreshSecretKey string
	tokenDuration    time.Duration
	refreshDuration  time.Duration
}

func NewJWTProvider(secretKey string, tokenDuration time.Duration, refreshSecret string, refreshDuration time.Duration) *JWTProvider {
	return &JWTProvider{
		secretKey:        secretKey,
		tokenDuration:    tokenDuration,
		refreshSecretKey: refreshSecret,
		refreshDuration:  refreshDuration,
	}
}

// GenerateToken gera o access token
func (p *JWTProvider) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(p.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(p.secretKey))
}

// DecodeToken valida o access token
func (p *JWTProvider) DecodeToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token")
	}

	return userID, nil
}

// GenerateRefreshToken cria um refresh token
func (p *JWTProvider) GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(p.refreshDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(p.refreshSecretKey))
}

// ValidateRefreshToken valida o refresh token
func (p *JWTProvider) ValidateRefreshToken(tokenStr string) (string, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.refreshSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in refresh token")
	}

	return userID, nil
}

// RefreshAccessToken gera um novo access token a partir do refresh token
func (p *JWTProvider) RefreshAccessToken(refreshToken string) (string, string, error) {
	userID, err := p.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	newAccessToken, err := p.GenerateToken(userID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, userID, nil
}
