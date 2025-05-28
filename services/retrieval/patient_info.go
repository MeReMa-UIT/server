package retrieval

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetPatientList(ctx context.Context, authHeader string) ([]models.PatientBriefInfo, error) {
	claims, err := auth.ParseJWT(auth.ExtractToken(authHeader), auth.JWT_SECRET)
	print("claims:", claims.Permission)
	if err != nil {
		return nil, err
	}

	switch claims.Permission {
	case permission.Patient.String():
		return repo.GetPatientList(ctx, &claims.ID)
	case permission.Receptionist.String(), permission.Doctor.String():
		return repo.GetPatientList(ctx, nil)
	default:
		return nil, errs.ErrPermissionDenied
	}
}

func GetPatientInfo(ctx context.Context, authHeader string, patientID string) (models.PatientInfo, error) {
	claims, err := auth.ParseJWT(auth.ExtractToken(authHeader), auth.JWT_SECRET)
	if err != nil {
		return models.PatientInfo{}, err
	}

	var info models.PatientInfo
	switch claims.Permission {
	case permission.Doctor.String(), permission.Receptionist.String():
		info, err = repo.GetPatientInfo(ctx, patientID, "0")
	case permission.Patient.String():
		info, err = repo.GetPatientInfo(ctx, patientID, claims.ID)
	default:
		return models.PatientInfo{}, errs.ErrPermissionDenied
	}

	if err != nil {
		return models.PatientInfo{}, err
	}

	return info, nil
}
