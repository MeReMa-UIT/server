package retrieval_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetDiagnosisList(ctx context.Context, authHeader string) ([]models.DiagnosisInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	if claims.Permission != permission.Doctor.String() {
		return nil, errs.ErrPermissionDenied
	}
	return repo.GetDiagnosisList(ctx)
}

func GetDiagnosisInfo(ctx context.Context, authHeader, icdCode string) (models.DiagnosisInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.DiagnosisInfo{}, err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return models.DiagnosisInfo{}, errs.ErrPermissionDenied
	}
	return repo.GetDiagnosisInfo(ctx, icdCode)
}
