package auth

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
)

func NewSession(ctx context.Context, req models.LoginRequest) (string, error) {
	creds, err := repo.GetCredentialsByUsername(context.Background(), req.Username)
	if err != nil {
		return "", err
	}
	if creds.PasswordHash != req.Password {
		return "", models.ErrPasswordIncorrect
	}
	token, err := GenerateJWT(creds.Username, creds.PasswordHash, JWT_SECRET, JWT_EXPIRY)
	if err != nil {
		return "", err
	}
	return token, nil
}
