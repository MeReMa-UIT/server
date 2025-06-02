package recovery

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

var gmailUsername, _ = utils.EnvVars["GMAIL_USERNAME"]
var gmailPassword, _ = utils.EnvVars["GMAIL_PASSWORD"]

func sendOTPEmail(recipient, otp string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Your OTP Code\n"
	mime := "MIME-Version: 1.0\nContent-Type: text/plain; charset=\"UTF-8\"\n\n"
	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire in 5 minutes.", otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", gmailUsername, gmailPassword, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		gmailUsername,
		[]string{recipient},
		message,
	)

	return err
}

func SendRecoveryEmail(ctx context.Context, req models.AccountRecoverRequest) error {

	err := repo.CheckEmailAndCitizenID(ctx, req)
	if err != nil {
		return err
	}
	accID, _ := repo.GetAccIDByCitizenID(ctx, req.CitizenID)

	otp := generateOTP(fmt.Sprint(accID))
	return sendOTPEmail(req.Email, otp)
}

func VerifyRecoveryOTP(ctx context.Context, req models.AccountRecoverConfirmRequest) (string, error) {
	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err != nil {
		return "", err
	}
	err = validateOTP(fmt.Sprint(accID), req.OTP)
	if err != nil {
		return "", err
	}
	token, err := auth.GenerateToken(fmt.Sprint(accID), permission.Recovery.String())
	if err != nil {
		return "", err
	}
	return token, nil
}
