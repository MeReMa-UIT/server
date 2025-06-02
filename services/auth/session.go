package auth

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/repo"
	"golang.org/x/crypto/bcrypt"
)

func NewSession(ctx context.Context, req models.LoginRequest) (string, error) {
	creds, err := repo.GetAccountCredentials(context.Background(), req.Identifier)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(req.Password)); err != nil {
		return "", errs.ErrPasswordIncorrect
	}
	token, err := GenerateToken(creds.AccID, creds.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
