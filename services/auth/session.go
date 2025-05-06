package auth

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	"golang.org/x/crypto/bcrypt"
)

func NewSession(ctx context.Context, req models.LoginRequest) (string, error) {
	creds, err := repo.GetAccountCredentials(context.Background(), req.Identifier)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(req.Password)); err != nil {
		return "", models.ErrPasswordIncorrect
	}
	token, err := GenerateJWT(creds.CitizenID, creds.Role, JWT_SECRET, JWT_EXPIRY)
	if err != nil {
		return "", err
	}
	return token, nil
}
