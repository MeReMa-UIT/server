package registration_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func InitRegistration(ctx context.Context, req models.InitRegistrationRequest, authHeader string) (models.AccountRegistrationResponse, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.AccountRegistrationResponse{}, err
	}

	var registrationType string
	switch claims.Permission {
	case permission.Admin.String():
		registrationType = permission.StaffRegistration.String()
	case permission.Receptionist.String():
		registrationType = permission.PatientRegistration.String()
	default:
		return models.AccountRegistrationResponse{}, errs.ErrPermissionDenied
	}

	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err != nil {
		if err == errs.ErrAccountNotExist {
			token, _ := auth_services.GenerateToken(claims.ID, registrationType)
			return models.AccountRegistrationResponse{Token: token, AccID: -1}, nil
		}
		return models.AccountRegistrationResponse{}, err
	}
	token, _ := auth_services.GenerateToken(claims.ID, registrationType)
	return models.AccountRegistrationResponse{Token: token, AccID: accID}, nil
}

func RegisterAccount(ctx context.Context, req models.AccountRegistrationRequest, authHeader string) (models.AccountRegistrationResponse, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.AccountRegistrationResponse{}, err
	}

	switch claims.Permission {
	case permission.PatientRegistration.String():
		if req.Role != permission.Patient.String() {
			return models.AccountRegistrationResponse{}, errs.ErrPermissionDenied
		}
	case permission.StaffRegistration.String():
		if req.Role != permission.Doctor.String() && req.Role != permission.Receptionist.String() {
			return models.AccountRegistrationResponse{}, errs.ErrPermissionDenied
		}
	default:
		return models.AccountRegistrationResponse{}, errs.ErrPermissionDenied
	}

	password_hash, _ := bcrypt.GenerateFromPassword([]byte(req.Phone), bcrypt.DefaultCost)
	createdAccID, err := repo.StoreAccountInfo(ctx, req, string(password_hash))
	if err != nil {
		return models.AccountRegistrationResponse{}, err
	}
	return models.AccountRegistrationResponse{Token: auth_services.ExtractToken(authHeader), AccID: createdAccID}, nil
}
