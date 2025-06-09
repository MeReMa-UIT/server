package recovery_services

import (
	"context"
	"fmt"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func SendRecoveryEmail(ctx context.Context, req models.AccountRecoverRequest) error {

	err := repo.CheckEmailAndCitizenID(ctx, req)
	if err != nil {
		return err
	}
	accID, _ := repo.GetAccIDByCitizenID(ctx, req.CitizenID)

	otp := auth_services.GenerateOTP(fmt.Sprint(accID))
	return auth_services.SendOTPEmail(req.Email, otp)
}

func VerifyRecoveryOTP(ctx context.Context, req models.AccountRecoverConfirmRequest) (string, error) {
	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err != nil {
		return "", err
	}
	err = auth_services.ValidateOTP(fmt.Sprint(accID), req.OTP)
	if err != nil {
		return "", err
	}
	token, err := auth_services.GenerateToken(fmt.Sprint(accID), permission.Recovery.String())
	if err != nil {
		return "", err
	}
	return token, nil
}
