package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey []byte

type JWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func LoadJwtKey() {

	// Read the jwt secret key from the environment variable
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		log.Fatal("JWT_SECRET_KEY environment variable not set. Did you create a .env file?")
	}
	JwtKey = []byte(key)
}

func GenerateJwtToken(email string) (string, error) {
	claims := &JWTClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rydeapi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func ValidateJwtToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, nil
	}
	return nil, jwt.ErrTokenMalformed
}
