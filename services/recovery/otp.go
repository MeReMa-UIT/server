package recovery

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/merema-uit/server/models"
)

type secret struct {
	OTP            string
	ExpirationTime time.Time
	Verified       bool
}

// temporary storage for otp, will be replaced with a database later
var otpSecrets = make(map[string]secret)

func generateOTP(citizenID string) string {
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	otpSecrets[citizenID] = secret{
		OTP:            otp,
		ExpirationTime: time.Now().Add(5 * time.Minute),
		Verified:       false,
	}
	return otp
}

func validateOTP(citizenID, otp string) error {

	secret, ok := otpSecrets[citizenID]

	if !ok {
		return models.ErrExpiredOTP
	}

	if time.Now().After(secret.ExpirationTime) {
		delete(otpSecrets, citizenID)
		return models.ErrExpiredOTP
	}

	if secret.OTP != otp {
		return models.ErrWrongOTP
	}

	secret.Verified = true
	otpSecrets[citizenID] = secret

	return nil

}
