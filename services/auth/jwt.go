package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "test"
const JWT_EXPIRY = 1 * time.Hour

type Claims struct {
	CitizenID  string `json:"citizen_id"`
	Permission string `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateJWT(citizenID string, role string, secret string, expiry time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiry)
	claims := &Claims{
		CitizenID:  citizenID,
		Permission: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJWT(tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func ExtractPermissionFromToken(tokenString, secret string) (string, error) {
	claims, err := ParseJWT(tokenString, secret)
	if err != nil {
		return "", err
	}
	return claims.Permission, nil
}

func ExtractToken(authHeader string) string {
	if authHeader == "" {
		return "" // No header found
	}

	// Split into ["Bearer", "token"]
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "" // Malformed header
	}

	return parts[1] // Return just the JWT token
}
