package recovery

import (
	"context"
	"fmt"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func SendRecoveryEmail(ctx context.Context, req models.AccountRecoverRequest) error {

	err := repo.CheckEmailAndCitizenID(ctx, req)
	if err != nil {
		return err
	}
	accID, _ := repo.GetAccIDByCitizenID(ctx, req.CitizenID)

	otp := auth.GenerateOTP(fmt.Sprint(accID))
	return auth.SendOTPEmail(req.Email, otp)
}

func VerifyRecoveryOTP(ctx context.Context, req models.AccountRecoverConfirmRequest) (string, error) {
	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err != nil {
		return "", err
	}
	err = auth.ValidateOTP(fmt.Sprint(accID), req.OTP)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(fmt.Sprint(accID), permission.Recovery.String())
	if err != nil {
		return "", err
	}
	return token, nil
}
