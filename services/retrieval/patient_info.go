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
	if claims.Permission != permission.Doctor.String() && claims.Permission != permission.Receptionist.String() {
		return nil, errs.ErrPermissionDenied
	}
	return repo.GetPatientList(ctx)
}
