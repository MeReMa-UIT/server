package record_validation

import (
	"bytes"
	"fmt"

	"github.com/jackc/pgtype"
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
