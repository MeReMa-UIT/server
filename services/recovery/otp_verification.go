package recovery

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func sendOTPEmail(recipient, otp string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")

	// Email content
	subject := "Subject: Your OTP Code\n"
	mime := "MIME-Version: 1.0\nContent-Type: text/plain; charset=\"UTF-8\"\n\n"
	body := fmt.Sprintf("Your OTP code is: %s\n\nThis code will expire in 5 minutes.", otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
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
	token, err := auth.GenerateToken(fmt.Sprint(accID), permission.Recovery.String(), auth.JWT_RECOVERY_EXPIRY)
	if err != nil {
		return "", err
	}
	return token, nil
}
