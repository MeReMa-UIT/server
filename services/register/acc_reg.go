package register

import (
	"context"
	"fmt"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func InitRegistration(ctx context.Context, req models.InitRegistrationRequest, authHeader string) (string, int, error) {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return "", -1, err
	}

	var registrationType string
	switch claims.Permission {
	case permission.Admin.String():
		registrationType = permission.StaffRegistration.String()
	case permission.Receptionist.String():
		registrationType = permission.PatientRegistration.String()
	default:
		return "", -1, errs.ErrPermissionDenied
	}

	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err != nil {
		if err == errs.ErrAccountNotExist {
			token, _ := auth.GenerateJWT(claims.ID, registrationType, auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
			return token, -1, nil
		}
		return "", -1, err
	}
	token, _ := auth.GenerateJWT(claims.ID, registrationType, auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
	return token, accID, nil
}

func RegisterAccount(ctx context.Context, req models.AccountRegistrationRequest, authHeader string) (string, int, error) {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return "", -1, err
	}

	switch claims.Permission {
	case permission.PatientRegistration.String():
		if req.Role != permission.Patient.String() {
			return "", -1, errs.ErrPermissionDenied
		}
	case permission.StaffRegistration.String():
		if req.Role != permission.Doctor.String() && req.Role != permission.Receptionist.String() {
			return "", -1, errs.ErrPermissionDenied
		}
	default:
		return "", -1, errs.ErrPermissionDenied
	}

	password_hash, _ := bcrypt.GenerateFromPassword([]byte(req.Phone), bcrypt.DefaultCost)
	createdAccID, err := repo.StoreAccountInfo(ctx, req, string(password_hash))
	if err != nil {
		return "", -1, err
	}
	token, _ := auth.GenerateJWT(fmt.Sprint(createdAccID), permission.Patient.String(), auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
	return token, createdAccID, nil
}
