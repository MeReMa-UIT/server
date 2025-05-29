package comm_services

import (
	"context"

	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
)

func GetContactList(ctx context.Context, authHeader string) ([]models.ContactInfo, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}
	switch claims.Permission {
	case permission.Doctor.String():
		return repo.GetContactListDoctor(ctx, claims.ID)
	case permission.Patient.String():
		return repo.GetContactListPatient(ctx, claims.ID)
	default:
		return nil, errs.ErrPermissionDenied
	}
}
