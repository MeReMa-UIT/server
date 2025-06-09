package prescription_services

import (
	"context"
	"slices"
	"strconv"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetPrescriptionList(ctx context.Context, authHeader string) ([]models.PrescriptionInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionInfo
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionList(ctx, nil)
	case permission.Patient.String():
		recordIDList, err := repo.GetRecordIDListByAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		list, err = repo.GetPrescriptionList(ctx, recordIDList)
	default:
		return nil, errs.ErrPermissionDenied
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetPrescriptionListByPatientID(ctx context.Context, authHeader, patientID string) ([]models.PrescriptionInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionInfo
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionListByPatientID(ctx, patientID)
	case permission.Patient.String():
		patientIDList, err := repo.GetPatientIDListByAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		id, _ := strconv.Atoi(patientID)
		if !slices.Contains(patientIDList, id) {
			return nil, errs.ErrPermissionDenied
		}
		list, err = repo.GetPrescriptionListByPatientID(ctx, patientID)
	default:
		return nil, errs.ErrPermissionDenied
	}
	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetPrescriptionInfoByRecordID(ctx context.Context, authHeader, recordID string) (models.PrescriptionInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return models.PrescriptionInfo{}, err
	}

	id, err := strconv.Atoi(recordID)

	if err != nil {
		return models.PrescriptionInfo{}, err
	}

	switch claims.Permission {
	case permission.Doctor.String():
	case permission.Patient.String():
		recordIDList, err := repo.GetRecordIDListByAccID(ctx, claims.ID)
		for _, recordID := range recordIDList {
			println(recordID)
		}
		if err != nil {
			return models.PrescriptionInfo{}, err
		}
		if !slices.Contains(recordIDList, id) {
			return models.PrescriptionInfo{}, errs.ErrPermissionDenied
		}
	default:
		return models.PrescriptionInfo{}, errs.ErrPermissionDenied
	}

	return repo.GetPrescriptionInfoByRecordID(ctx, id)
}

func GetPrescriptionDetails(ctx context.Context, authHeader, prescriptionID string) ([]models.PrescriptionDetailInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionDetailInfo
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionDetails(ctx, prescriptionID)
	case permission.Patient.String():
		id, _ := strconv.Atoi(prescriptionID)

		prescriptionIDList, err := repo.GetPrescriptionIDListByAccID(ctx, claims.ID)

		if err != nil {
			return nil, err
		}

		if !slices.Contains(prescriptionIDList, id) {
			return nil, errs.ErrPermissionDenied
		}
		list, err = repo.GetPrescriptionDetails(ctx, prescriptionID)
	default:
		return nil, errs.ErrPermissionDenied
	}
	if err != nil {
		return nil, err
	}

	return list, nil
}
