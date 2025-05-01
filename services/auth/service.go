package auth

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
)

func Auth(ctx context.Context, req models.LoginRequest) (string, error) {
	password_hash, err := repo.GetCredentialsByUsername(context.Background(), req.Username)
	if err != nil {
		return "", err
	}
	if password_hash != req.Password {
		return "wrong", nil
	}
	token, err := GenerateJWT(req.Username, JWT_SECRET, JWT_EXPIRY)
	if err != nil {
		return "", err
	}
	return token, nil
}
