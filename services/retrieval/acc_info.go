package retrieval_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetAccountInfo(ctx context.Context, authHeader string) (models.AccountInfo, any, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.AccountInfo{}, nil, err
	}

	accountInfo, err := repo.GetAccountInfo(ctx, claims.ID)
	if err != nil {
		return models.AccountInfo{}, nil, err
	}

	var additionalInfo any

	switch claims.Permission {
	case permission.Patient.String():
		additionalInfo, err = GetPatientList(ctx, authHeader)
		if err != nil {
			return models.AccountInfo{}, nil, err
		}
	case permission.Receptionist.String(), permission.Doctor.String():
		additionalInfo, err = GetStaffInfo(ctx, authHeader, "")
		if err != nil {
			return models.AccountInfo{}, nil, err
		}
	}
	return accountInfo, additionalInfo, nil
}

func GetAccountList(ctx context.Context, authHeader string) ([]models.AccountInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	if claims.Permission != permission.Admin.String() {
		return nil, errs.ErrPermissionDenied
	}

	return repo.GetAccountList(ctx)
}
