package utils

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateAndValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	token, err := GenerateToken(42)
	if err != nil {
		t.Fatalf("generate token error: %v", err)
	}
	uid, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("validate token error: %v", err)
	}
	if uid != 42 {
		t.Fatalf("expected 42 got %d", uid)
	}
}

func TestValidateTokenInvalidSignature(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	token, _ := GenerateToken(1)
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("unexpected token format")
	}
	parts[2] = "tampered"
	if _, err := ValidateToken(strings.Join(parts, ".")); err == nil {
		t.Fatalf("expected error for invalid signature")
	}
}

func TestValidateTokenExpired(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	claims := jwt.MapClaims{
		"uid": 1,
		"exp": time.Now().Add(-time.Hour).Unix(),
	}
	expired := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := expired.SignedString([]byte(secretKey()))
	if _, err := ValidateToken(token); err == nil || err.Error() != "token expired" {
		t.Fatalf("expected token expired error, got %v", err)
	}
}
