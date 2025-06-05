package retrieval

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/models/permission"
	"github.com/merema-uit/server/repo"
	"github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

func GetMedicalRecordTypeList(ctx context.Context, authHeader string) ([]models.MedicalRecordType, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	if claims.Permission != permission.Doctor.String() {
		return nil, errs.ErrPermissionDenied
	}

	return repo.GetMedicalRecordTypeList(ctx)
}

func GetMedicalRecordTemplate(ctx context.Context, authHeader, typeID string) (*pgtype.JSONB, error) {
	claims, err := auth.ParseToken(auth.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	if claims.Permission != permission.Doctor.String() {
		return nil, errs.ErrPermissionDenied
	}

	// change from 01BV01 to 01/BV01
	typeID = typeID[:2] + "/" + typeID[2:]

	typeInfo, err := repo.GetMedicalRecordType(ctx, typeID)
	if err != nil {
		return nil, err
	}

	template, err := utils.LoadJSONFileToJSONB(typeInfo.TemplatePath)
	if err != nil {
		return nil, err
	}
	return template, nil
}
