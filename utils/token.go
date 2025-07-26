package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func secretKey() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "secret"
	}
	return key
}

func GenerateToken(uid uint) (string, error) {
	expiry := time.Now().Add(24 * time.Hour).Unix()
	payload := fmt.Sprintf("%d:%d", uid, expiry)
	enc := base64.StdEncoding.EncodeToString([]byte(payload))
	sig := computeHMAC(enc, secretKey())
	return fmt.Sprintf("%s.%s", enc, sig), nil
}

func ValidateToken(token string) (uint, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return 0, errors.New("invalid token")
	}
	enc, sig := parts[0], parts[1]
	if computeHMAC(enc, secretKey()) != sig {
		return 0, errors.New("invalid token")
	}
	data, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	p := strings.Split(string(data), ":")
	if len(p) != 2 {
		return 0, errors.New("invalid token")
	}
	id, err := strconv.ParseUint(p[0], 10, 64)
	if err != nil {
		return 0, err
	}
	exp, err := strconv.ParseInt(p[1], 10, 64)
	if err != nil {
		return 0, err
	}
	if time.Now().Unix() > exp {
		return 0, errors.New("token expired")
	}
	return uint(id), nil
}

func computeHMAC(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
