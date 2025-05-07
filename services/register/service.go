package register

import (
	"context"
	"strconv"

	"fmt"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func InitRegistration(ctx context.Context, req models.InitRegistrationRequest, authHeader string) (string, int, error) {
	tokenString := auth.ExtractToken(authHeader)
	perm, err := auth.ExtractPermissionFromToken(tokenString, auth.JWT_SECRET)
	if err != nil {
		return "", -1, err
	}

	var registrationType string
	switch perm {
	case permission.Admin.String():
		registrationType = permission.StaffRegistration.String()
	case permission.Receptionist.String():
		registrationType = permission.PatientRegistration.String()
	default:
		return "", -1, errors.ErrPermissionDenied
	}

	accID, err := repo.GetAccIDByCitizenID(ctx, req.CitizenID)
	if err == errors.ErrAccountNotExist {
		token, _ := auth.GenerateJWT(req.CitizenID, registrationType, auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
		return token, -1, nil
	}
	if err != nil {
		return "", -1, err
	}
	token, _ := auth.GenerateJWT(fmt.Sprint(accID), registrationType, auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
	return token, accID, nil
}

func RegisterAccount(ctx context.Context, req models.AccountRegistrationRequest, authHeader string) (string, error) {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return "", err
	}

	if len(claims.ID) != 12 {
		return "", errors.ErrInvalidToken
	}

	switch claims.Permission {
	case permission.PatientRegistration.String():
		if req.Role != permission.Patient.String() {
			return "", errors.ErrPermissionDenied
		}
	case permission.StaffRegistration.String():
		if req.Role != permission.Doctor.String() && req.Role != permission.Receptionist.String() {
			return "", errors.ErrPermissionDenied
		}
	default:
		return "", errors.ErrPermissionDenied
	}

	password_hash, _ := bcrypt.GenerateFromPassword([]byte(req.Phone), bcrypt.DefaultCost)
	createdAccID, err := repo.StoreAccountInfo(ctx, req, claims.ID, string(password_hash))
	if err != nil {
		return "", err
	}
	token, _ := auth.GenerateJWT(fmt.Sprint(createdAccID), permission.Patient.String(), auth.JWT_SECRET, auth.JWT_REGISTRATION_EXPIRY)
	return token, nil
}

func RegisterStaff(ctx context.Context, req models.StaffRegistrationRequest, authHeader string) error {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.StaffRegistration.String() {
		return errors.ErrPermissionDenied
	}
	accID, err := strconv.Atoi(claims.ID)
	if err != nil {
		return errors.ErrInvalidToken
	}
	err = repo.StoreStaffInfo(ctx, req, accID)
	if err != nil {
		return err
	}
	return nil
}

func RegisterPatient(ctx context.Context, req models.PatientRegistrationRequest, authHeader string) error {
	tokenString := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(tokenString, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if claims.Permission != permission.PatientRegistration.String() {
		return errors.ErrPermissionDenied
	}
	accID, err := strconv.Atoi(claims.ID)
	if err != nil {
		return errors.ErrInvalidToken
	}
	err = repo.StorePatientInfo(ctx, req, accID)
	if err != nil {
		return err
	}

	return nil
}
