package record_services

import (
	"context"
	"mime/multipart"
	"strconv"

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

	additionalInfo, err := record_validation.ValidateRecordDetail(&req.RecordDetail, recordTypeInfo.TypeID, recordTypeInfo.SchemaPath)

	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	if additionalInfo.PrimaryDiagnosis == "" {
		return models.NewMedicalRecordResponse{}, errs.ErrPrimaryDiagnosisMissing
	}

	doctorID, _ := repo.GetStaffIDByAccID(ctx, claims.ID)

	res, err := repo.StoreMedicalRecord(ctx, doctorID, req, additionalInfo)
	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	patientAccID, err := repo.GetAccIDByPatientID(ctx, req.PatientID)
	if err != nil {
		return models.NewMedicalRecordResponse{}, err
	}

	doctorAccID, _ := strconv.ParseInt(claims.ID, 10, 64)

	if err := repo.CheckConversationExists(ctx, patientAccID, doctorAccID); err != nil {
		if err := repo.AddNewConversation(ctx, min(patientAccID, doctorAccID), max(patientAccID, doctorAccID)); err != nil {
			return models.NewMedicalRecordResponse{}, err
		}
	}

	return res, nil
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
