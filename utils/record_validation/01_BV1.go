package record_validation

import (
	"encoding/json"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/utils"
	_ "github.com/santhosh-tekuri/jsonschema/v6"
)

func Validate01BV1(record *pgtype.JSONB, schemaPath string) (models.ExtractedRecordInfo, error) {
	schema, err := utils.LoadJSONFileToJSONB(schemaPath)
	if err != nil {
		return models.ExtractedRecordInfo{}, fmt.Errorf("Failed to load schema: %w", err)
	}

	if err := validateJSON(record, schema); err != nil {
		return models.ExtractedRecordInfo{}, err
	}

	var (
		additionalInfo models.ExtractedRecordInfo
		recordData     map[string]interface{}
	)

	json.Unmarshal(record.Bytes, &recordData)

	diagnosisPredictions := recordData["THÔNG TIN CHUNG"].(map[string]interface{})["Chẩn đoán"].(map[string]interface{})

	if pred, ok := diagnosisPredictions["Ra viện"].(map[string]interface{}); ok {
		if primary, ok := pred["Bệnh chính"].(string); ok {
			additionalInfo.PrimaryDiagnosis = primary
		}
		if secondary, ok := pred["Bệnh kèm theo"].(string); ok {
			additionalInfo.SecondaryDiagnosis = secondary
		}
	} else if pred, ok := diagnosisPredictions["Khi vào khoa điều trị"].(string); ok {
		additionalInfo.PrimaryDiagnosis = pred
	} else if pred, ok := diagnosisPredictions["KKB, Cấp cứu"].(string); ok {
		additionalInfo.PrimaryDiagnosis = pred
	} else if pred, ok := diagnosisPredictions["Nơi chuyển đến"].(string); ok {
		additionalInfo.PrimaryDiagnosis = pred
	}

	// primary diagnosis, secondary diagnosis
	return additionalInfo, nil
}
