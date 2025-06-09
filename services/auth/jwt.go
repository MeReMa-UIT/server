package auth_services

import (
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/utils"
)

var jwtSecret = utils.EnvVars["JWT_SECRET"]
var jwtSessionExpiration, _ = strconv.Atoi(utils.EnvVars["JWT_SESSION_EXPIRATION"])
var jwtRecoveryExpiration, _ = strconv.Atoi(utils.EnvVars["JWT_RECOVERY_EXPIRATION"])
var jwtRegistrationExpiration, _ = strconv.Atoi(utils.EnvVars["JWT_REGISTRATION_EXPIRATION"])

type Claims struct {
	ID         string `json:"id"` // Acc ID
	Permission string `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateToken(id string, perm string) (string, error) {
	var duration time.Duration

	switch perm {
	case permission.Recovery.String():
		duration = time.Duration(jwtRecoveryExpiration) * time.Minute
	case permission.PatientRegistration.String(), permission.StaffRegistration.String():
		duration = time.Duration(jwtRegistrationExpiration) * time.Minute
	default:
		duration = time.Duration(jwtSessionExpiration) * time.Minute
	}

	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		ID:         id,
		Permission: perm,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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
