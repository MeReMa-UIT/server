package retrieval

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetMedicationList(ctx context.Context, authHeader string) ([]models.MedicationInfo, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	if claims.Permission != permission.Doctor.String() {
		return nil, errs.ErrPermissionDenied
	}
	return repo.GetMedicationList(ctx)
}

func GetMedicationInfo(ctx context.Context, authHeader string, medicationID string) (models.MedicationInfo, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return models.MedicationInfo{}, err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return models.MedicationInfo{}, errs.ErrPermissionDenied
	}
	medicationInfo, err := repo.GetMedicationInfo(ctx, medicationID)
	if err != nil {
		return models.MedicationInfo{}, err
	}
	return medicationInfo, nil
}
