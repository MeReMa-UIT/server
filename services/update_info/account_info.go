package info_update

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

var possibleFields = map[string]bool{
	"citizen_id": true,
	"phone":      true,
	"email":      true,
	"password":   true,
}

func UpdateAccountInfo(ctx context.Context, authHeader string, req models.UpdateAccountInfoRequest) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if _, ok := possibleFields[req.Field]; !ok {
		return errs.ErrInvalidField
	}
	creds, err := repo.GetAccountCredentials(ctx, claims.ID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(req.Password))
	if err != nil {
		return errs.ErrPasswordIncorrect
	}
	if req.Field == "password" {
		newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewValue), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.Field += "_hash"
		req.NewValue = string(newPasswordHash)
	}
	return repo.UpdateAccountInfo(ctx, claims.ID, req)
}
