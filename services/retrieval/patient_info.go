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
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
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
