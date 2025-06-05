package record_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils/record_validation"
)

func AddNewRecord(ctx context.Context, authHeader string, req *models.NewMedicalRecordRequest) error {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	recordTypeInfo, err := repo.GetMedicalRecordType(ctx, req.TypeID)

	if err != nil {
		return err
	}

	err = record_validation.Validate01BV1(&req.RecordDetail, recordTypeInfo.SchemaPath)

	if err != nil {
		return err
	}

	// return repo.StoreMedicalRecord(ctx)
	return nil
}
