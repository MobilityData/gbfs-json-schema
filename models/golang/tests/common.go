package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

// Test fixture path constants for different GBFS versions
const (
	TestFixturesV23   = "./../../../testFixtures/v2.3/"
	TestFixturesV30   = "./../../../testFixtures/v3.0/"
	TestFixturesV31RC = "./../../../testFixtures/v3.1-RC/"
	TestFixturesV31RC2 = "./../../../testFixtures/v3.1-RC2/"
)

// loadSchemaAndFixture loads a JSON schema and test fixture file,
// returning a schema loader and the fixture data for validation testing.
func loadSchemaAndFixture(t *testing.T, schemaPath string, fixturePath string) (gojsonschema.JSONLoader, []byte) {
	schemaDataBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("Failed to read schema file %s: %v", schemaPath, err)
	}
	
	var schemaData map[string]interface{}
	if err := json.Unmarshal(schemaDataBytes, &schemaData); err != nil {
		t.Fatalf("Failed to unmarshal schema file %s: %v", schemaPath, err)
	}
	
	schemaLoader := gojsonschema.NewGoLoader(schemaData)
	
	jsonData, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf("Failed to read fixture file %s: %v", fixturePath, err)
	}
	
	return schemaLoader, jsonData
}

// validateSchemaToUnmarshal validates that the provided GBFS data
// conforms to the given JSON schema, reporting any validation errors.
func validateSchemaToUnmarshal(t *testing.T, schemaLoader gojsonschema.JSONLoader, gbfsData interface{}) {
	loader := gojsonschema.NewGoLoader(gbfsData)
	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		t.Fatalf("Failed to validate JSON schema: %v", err)
	}

	if !result.Valid() {
		t.Error("The JSON is NOT valid:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
