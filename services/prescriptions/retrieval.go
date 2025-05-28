package prescription_services

import (
	"context"
	"slices"
	"strconv"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetPrescriptionListWithRecordID(ctx context.Context, authHeader, recordID string) ([]models.PrescriptionInfo, error) {
	claims, err := auth.ParseJWT(auth.ExtractToken(authHeader), auth.JWT_SECRET)
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionInfo
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionListWithRecordID(ctx, recordID)
	case permission.Patient.String():
		recordIDList, err := repo.GetRecordIDListByAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		id, _ := strconv.Atoi(recordID)
		if !slices.Contains(recordIDList, id) {
			return nil, errs.ErrPermissionDenied
		}
		list, err = repo.GetPrescriptionListWithRecordID(ctx, recordID)
	default:
		return nil, errs.ErrPermissionDenied
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetPrescriptionListWithPatientID(ctx context.Context, authHeader, patientID string) ([]models.PrescriptionInfo, error) {
	claims, err := auth.ParseJWT(auth.ExtractToken(authHeader), auth.JWT_SECRET)
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionInfo
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionListWithPatientID(ctx, patientID)
	case permission.Patient.String():
		patientIDList, err := repo.GetPatientIDListByAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		id, _ := strconv.Atoi(patientID)
		if !slices.Contains(patientIDList, id) {
			return nil, errs.ErrPermissionDenied
		}
		list, err = repo.GetPrescriptionListWithPatientID(ctx, patientID)
	default:
		return nil, errs.ErrPermissionDenied
	}
	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetPrescriptionDetails(ctx context.Context, authHeader, prescriptionID string) ([]models.PrescriptionDetail, error) {
	claims, err := auth.ParseJWT(auth.ExtractToken(authHeader), auth.JWT_SECRET)
	if err != nil {
		return nil, err
	}

	var list []models.PrescriptionDetail
	switch claims.Permission {
	case permission.Doctor.String():
		list, err = repo.GetPrescriptionDetails(ctx, prescriptionID)
	case permission.Patient.String():
		prescriptionIDList, err := repo.GetPrescriptionIDListWithAccID(ctx, claims.ID)
		if err != nil {
			return nil, err
		}
		id, _ := strconv.Atoi(prescriptionID)
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
