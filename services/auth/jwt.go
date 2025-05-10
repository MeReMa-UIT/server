package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	errs "github.com/merema-uit/server/models/errors"
)

const JWT_SECRET = "MySuperSecretKeyForMeReMa"
const JWT_SESSION_EXPIRY = 3 * time.Hour
const JWT_RECOVERY_EXPIRY = 5 * time.Minute
const JWT_REGISTRATION_EXPIRY = 15 * time.Minute

type Claims struct {
	ID         string `json:"id"` // Acc ID
	Permission string `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string, permission string, secret string, expiry time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiry)
	claims := &Claims{
		ID:         id,
		Permission: permission,
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
		switch err {
		case jwt.ErrTokenMalformed:
			return nil, errs.ErrMalformedToken
		case jwt.ErrTokenExpired:
			return nil, errs.ErrExpiredToken
		default:
			return nil, errs.ErrInvalidToken
		}
	}

	if !token.Valid {
		return nil, errs.ErrInvalidToken
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
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
