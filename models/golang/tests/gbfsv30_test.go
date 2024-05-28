package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	gbfs "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/gbfs"
	gbfsversions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/gbfs_versions"
	geofencingzones "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/geofencing_zones"
	manifest "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/manifest"
	stationinformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/station_information"
	stationstatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/station_status"
	systemalerts "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_alerts"
	systeminformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_information"
	systempricingplans "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_pricing_plans"
	systemregions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_regions"
	vehiclestatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/vehicle_status"
	vehicletypes "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/vehicle_types"

	"github.com/xeipuuv/gojsonschema"
)

const pathToTestFixtures = "./../../../testFixtures/v3.0/"

func LoadSchemaAndFixture(t *testing.T, fileName string) (gojsonschema.JSONLoader, []byte) {

	pathToSchema := "./../../../v3.0/"
	schemaDataBytes, schemaErr := os.ReadFile(pathToSchema + fileName)
	if schemaErr != nil {
		t.Error("Error opening JSON file:", schemaErr)
		return nil, nil
	}
	var schemaData map[string]interface{}
	schemaErr = json.Unmarshal(schemaDataBytes, &schemaData)
	if schemaErr != nil {
		t.Error("Error opening JSON file:", schemaErr)
		return nil, nil
	}
	schemaLoader := gojsonschema.NewGoLoader(schemaData)
	jsonData, err2 := os.ReadFile(pathToTestFixtures + fileName)
	if err2 != nil {
		t.Error("Error opening JSON file:", err2)
		return nil, nil
	}
	return schemaLoader, jsonData
}

func ValidateSchemaToUnmarshal(t *testing.T, schemaLoader gojsonschema.JSONLoader, gbfsData interface{}) {
	loader := gojsonschema.NewGoLoader(gbfsData)
	result, err3 := gojsonschema.Validate(schemaLoader, loader)
	if err3 != nil {
		t.Error("Error validating JSON file:", err3)
		return
	}

	if !result.Valid() {
		t.Error("The JSON is NOT valid:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func TestGbfs(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "gbfs.json")
	gbfsData, err := gbfs.UnmarshalGbfs(jsonData)
	if err != nil {
		t.Error("Error With Unmarshal:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGbfsVersions(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "gbfs_versions.json")
	gbfsData, err := gbfsversions.UnmarshalGbfsVersions(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGeofencingZones(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "geofencing_zones.json")
	gbfsData, err := geofencingzones.UnmarshalGeofencingZones(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestManifest(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "manifest.json")
	gbfsData, err := manifest.UnmarshalManifest(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformation(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "station_information.json")
	gbfsData, err := stationinformation.UnmarshalStationInformation(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationStatus(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "station_status.json")
	gbfsData, err := stationstatus.UnmarshalStationStatus(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemAlerts(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "system_alerts.json")
	gbfsData, err := systemalerts.UnmarshalSystemAlerts(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemInformation(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "system_information.json")
	gbfsData, err := systeminformation.UnmarshalSystemInformation(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlan(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "system_pricing_plans.json")
	gbfsData, err := systempricingplans.UnmarshalSystemPricingPlans(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemRegions(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "system_regions.json")
	gbfsData, err := systemregions.UnmarshalSystemRegions(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleStatus(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "vehicle_status.json")
	gbfsData, err := vehiclestatus.UnmarshalVehicleStatus(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleTypes(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixture(t, "vehicle_types.json")
	gbfsData, err := vehicletypes.UnmarshalVehicleTypes(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}
