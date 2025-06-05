package record_validation

import (
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/merema-uit/server/utils"
	_ "github.com/santhosh-tekuri/jsonschema/v6"
)

func Validate01BV1(record *pgtype.JSONB, schemaPath string) error {
	schema, err := utils.LoadJSONFileToJSONB(schemaPath)
	if err != nil {
		return fmt.Errorf("Failed to load schema: %w", err)
	}

	if err := validateJSON(record, schema); err != nil {
		return err
	}
	return nil
}
