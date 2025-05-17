package schedule_services

import (
	"context"
	"fmt"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetScheduleList(ctx context.Context, authHeader string, req models.GetScheduleListRequest) ([]models.ScheduleInfo, error) {
	token := auth.ExtractToken(authHeader)
	claims, err := auth.ParseJWT(token, auth.JWT_SECRET)
	if err != nil {
		return nil, err
	}
	switch claims.Permission {
	case permission.Patient.String():
		patientID, err := repo.GetPatientID(ctx, claims.ID)
		fmt.Println("patientID: ", patientID)
		if err != nil {
			return nil, err
		}
		return repo.GetScheduleList(ctx, &patientID, req)
	case permission.Receptionist.String():
		return repo.GetScheduleList(ctx, nil, req)
	default:
		return nil, errs.ErrPermissionDenied
	}
}
