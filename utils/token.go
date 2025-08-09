package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func secretKey() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "secret"
	}
	return key
}

func GenerateToken(uid uint) (string, error) {
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey()))
}

func ValidateToken(token string) (uint, error) {
	claims := jwt.MapClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey()), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, errors.New("token expired")
		}
		return 0, errors.New("invalid token")
	}
	if !parsed.Valid {
		return 0, errors.New("invalid token")
	}
	uid, ok := claims["uid"].(float64)
	if !ok {
		return 0, errors.New("invalid token")
	}
	return uint(uid), nil
}
