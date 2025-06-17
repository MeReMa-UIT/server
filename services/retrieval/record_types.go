package retrieval_services

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

func GetMedicalRecordTypeList(ctx context.Context, authHeader string) ([]models.MedicalRecordType, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	return repo.GetMedicalRecordTypeList(ctx)
}

func GetMedicalRecordTemplate(ctx context.Context, authHeader, typeID string) (*pgtype.JSON, error) {
	_, err := auth_services.ParseToken(auth_services.ExtractToken(authHeader))
	if err != nil {
		return nil, err
	}

	// change from 01BV01 to 01/BV01
	typeID = typeID[:2] + "/" + typeID[2:]

	typeInfo, err := repo.GetMedicalRecordType(ctx, typeID)
	if err != nil {
		return nil, err
	}

	template, err := utils.LoadJSONFileToJSON(typeInfo.TemplatePath)
	if err != nil {
		return nil, err
	}
	return template, nil
}
