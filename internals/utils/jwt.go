package utils

import (
	"time"

	"github.com/aaryansinhaa/patient-management-system/internals/model"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		SecretKey:     secret,
		TokenDuration: duration,
	}
}

func (j *JWTManager) Generate(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(j.TokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}
