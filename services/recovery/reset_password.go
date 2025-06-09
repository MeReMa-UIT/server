package recovery_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(ctx context.Context, req models.PasswordResetRequest, authHeader string) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Recovery.String() {
		return errs.ErrPermissionDenied
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.NewPassword = string(newPasswordHash)
	err = repo.UpdatePassword(ctx, claims.ID, req.NewPassword)

	return nil
}
