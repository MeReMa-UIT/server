package record_services

import (
	"bytes"
	"context"
	"os"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

func GetRecordList(ctx context.Context, authHeader string) ([]models.MedicalRecordBriefInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	switch claims.Permission {
	case permission.Doctor.String():
		return repo.GetMedicalRecordList(ctx, 0)
	case permission.Patient.String():
		patientIDList, err := repo.GetPatientIDListByAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		var list []models.MedicalRecordBriefInfo
		for _, patientID := range patientIDList {
			records, err := repo.GetMedicalRecordList(ctx, patientID)
			if err != nil {
				return nil, err
			}
			list = append(list, records...)
		}
		return list, nil
	default:
		return nil, errs.ErrPermissionDenied
	}
}

func GetRecordInfo(ctx context.Context, authHeader, recordID string) (models.MedicalRecordInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.MedicalRecordInfo{}, err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return models.MedicalRecordInfo{}, errs.ErrPermissionDenied
	}

	return repo.GetMedicalRecordInfo(ctx, recordID)
}

func GetRecordAttachments(ctx context.Context, authHeader, recordID string) (*bytes.Buffer, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Patient.String() {
		return nil, errs.ErrPermissionDenied
	}
	storagePath := os.ExpandEnv(utils.EnvVars["FILE_STORAGE_PATH"]) + "/records/" + recordID
	return utils.ZipDirectory(storagePath)
}
