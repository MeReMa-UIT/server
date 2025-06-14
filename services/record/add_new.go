package record_services

import (
	"context"
	"mime/multipart"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
	"github.com/merema-uit/server/utils/record_validation"
)

func AddNewRecord(ctx context.Context, authHeader string, req *models.NewMedicalRecordRequest) (models.NewMedicalRecordResponse, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}
	if claims.Permission != permission.Doctor.String() {
		return models.NewMedicalRecordResponse{}, errs.ErrPermissionDenied
	}

	recordTypeInfo, err := repo.GetMedicalRecordType(ctx, req.TypeID)

	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	additionalInfo, err := record_validation.Validate01BV1(&req.RecordDetail, recordTypeInfo.SchemaPath)

	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	if additionalInfo.PrimaryDiagnosis == "" {
		return models.NewMedicalRecordResponse{}, errs.ErrPrimaryDiagnosisMissing
	}

	doctorID, _ := repo.GetStaffIDByAccID(ctx, claims.ID)

	return repo.StoreMedicalRecord(ctx, doctorID, req, additionalInfo)
}

func AddRecordAttachments(ctx context.Context, authHeader, recordID string, attachments []*multipart.FileHeader) error {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return err
	}
	if claims.Permission != permission.Doctor.String() {
		return errs.ErrPermissionDenied
	}

	prefix := utils.GetAttachmentPrefix(attachments[0].Filename)

	if prefix == "" {
		return errs.ErrInvalidAttachmentPrefix
	}

	return repo.StoreMedicalRecordAttachments(ctx, recordID, attachments, prefix)
}
