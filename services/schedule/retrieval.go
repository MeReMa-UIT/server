package schedule_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
)

func GetScheduleList(ctx context.Context, authHeader string, req models.GetScheduleListRequest) ([]models.ScheduleInfo, error) {
	claims, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	switch claims.Permission {
	case permission.Patient.String():
		return repo.GetScheduleList(ctx, &claims.ID, req)
	case permission.Receptionist.String():
		return repo.GetScheduleList(ctx, nil, req)
	default:
		return nil, errs.ErrPermissionDenied
	}
}
