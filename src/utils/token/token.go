package token

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	SecretKey string
}

func NewToken() *Token {
	// Inisialisasi struct Token dengan kunci rahasia
	secretKey := os.Getenv("JWT_SECRET") // Gantilah dengan kunci rahasia Anda
	return &Token{SecretKey: secretKey}
}

func (t *Token) GenerateToken(userID string) (string, error) {
	// Buat token dengan claims
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign dan mendapatkan token sebagai string
	tokenString, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Fungsi untuk memverifikasi token JWT
func (t *Token) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode penandatanganan benar (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(t.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
