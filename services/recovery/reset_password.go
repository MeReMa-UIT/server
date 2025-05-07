package recovery

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(ctx context.Context, req models.PasswordResetRequest, authHeader string) error {
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.Recovery.String() {
		return errors.ErrPermissionDenied
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.NewPassword = string(newPasswordHash)
	err = repo.UpdatePassword(ctx, claims.ID, req.NewPassword)

	return nil
}
