package utils

import (
	"fmt"
	"os"

	"github.com/jackc/pgtype"
)

func LoadJSONFileToJSON(path string) (*pgtype.JSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read JSON file: %w", err)
	}

	var schema pgtype.JSON
	if err := schema.UnmarshalJSON(data); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON to JSONB: %w", err)
	}

	return &schema, nil
}
