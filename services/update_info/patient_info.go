package info_update

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func UpdatePatientInfo(ctx context.Context, authHeader, patientID string, updatedInfo models.PatientInfoUpdateRequest) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))

	if err != nil {
		return err
	}

	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Receptionist.String() {
		return errs.ErrPermissionDenied
	}

	return repo.UpdatePatientInfo(ctx, patientID, updatedInfo)
}
