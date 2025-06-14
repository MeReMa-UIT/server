package record_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils/record_validation"
)

func UpdateMedicalRecord(ctx context.Context, authHeader, recordID string, req models.UpdateMedicalRecordRequest) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	recordTypeInfo, err := repo.GetMedicalRecordTypeByRecordID(ctx, recordID)
	if err != nil {
		return err
	}

	additionalInfo, err := record_validation.ValidateRecordDetail(&req.NewRecordDetail, recordTypeInfo.TypeID, recordTypeInfo.SchemaPath)
	if err != nil {
		return err
	}

	if additionalInfo.PrimaryDiagnosis == "" {
		return errs.ErrPrimaryDiagnosisMissing
	}

	err = repo.UpdateMedicalRecord(ctx, recordID, req, additionalInfo)

	if err != nil {
		return err
	}

	return nil
}
