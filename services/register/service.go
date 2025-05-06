package register

import (
	"context"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"golang.org/x/crypto/bcrypt"
)

func CheckPermission(authHeader, role string) error {
	tokenString := auth.ExtractToken(authHeader)
	if tokenString == "" {
		return errors.ErrInvalidToken
	}
	permission, err := auth.ExtractPermissionFromToken(tokenString, auth.JWT_SECRET)
	if err != nil {
		return err
	}
	if permission != role {
		return errors.ErrPermissionDenied
	}
	return nil
}

func RegisterAccount(ctx context.Context, req models.AccountRegisterRequest) (int, error) {
	password_hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	req.Password = string(password_hash)
	createdAccID, err := repo.StoreAccountInfo(ctx, req)
	return createdAccID, err
}

func RegisterStaff(ctx context.Context, req models.PatientRegisterRequest, authHeader string) error {
	if err := CheckPermission(authHeader, permission.Admin.String()); err != nil {
		return err
	}
	return nil
}

func RegisterPatient(ctx context.Context, req models.PatientRegisterRequest, authHeader string) error {
	if err := CheckPermission(authHeader, permission.Receptionist.String()); err != nil {
		return err
	}
	accID, err := repo.GetAccIDByCitizenID(ctx, req.AccountRegisterRequest.CitizenID)
	if err != nil && err != errors.ErrAccountNotExist {
		return err
	}
	if err == errors.ErrAccountNotExist {
		accID, err = RegisterAccount(ctx, req.AccountRegisterRequest)
		if err != nil {
			return err
		}
	}
	err = repo.StorePatientInfo(ctx, req, accID)
	return err
}
