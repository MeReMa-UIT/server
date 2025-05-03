package recovery

import (
	"context"
	"time"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(ctx context.Context, req models.PasswordResetRequest) error {
	secret, ok := otpSecrets[req.CitizenID]
	if !ok {
		return models.ErrExpiredOTP
	}
	if time.Now().After(secret.ExpirationTime) {
		delete(otpSecrets, req.CitizenID)
		return models.ErrExpiredOTP
	}
	if !secret.Verified {
		return models.ErrUnverifiedOTP
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.NewPassword = string(newPasswordHash)
	err = repo.UpdatePassword(ctx, req)

	delete(otpSecrets, req.CitizenID)
	return nil
}
