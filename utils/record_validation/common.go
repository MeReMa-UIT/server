package record_validation

import (
	"bytes"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/santhosh-tekuri/jsonschema"
)

func validateJSON(record, schema *pgtype.JSONB) error {
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("mem://schema.json",
		bytes.NewReader(schema.Bytes)); err != nil {
		return fmt.Errorf("Bad schema: %w", err)
	}
	compiled, err := compiler.Compile("mem://schema.json")
	if err != nil {
		return fmt.Errorf("Bad schema: %w", err)
	}

	if err := compiled.Validate(bytes.NewReader(record.Bytes)); err != nil {
		return errs.ErrInvalidMedicalRecordStructure
	}
	return nil
}

func ValidateRecordDetail(record *pgtype.JSONB, typeID, schemaPath string) (models.ExtractedRecordInfo, error) {
	switch typeID {
	case "01/BV1":
		return validate01BV1(record, schemaPath)
	default:
		return models.ExtractedRecordInfo{}, fmt.Errorf("Unsupported record type: %s", typeID)
	}
}
