package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	gbfs "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/gbfs"
	gbfsversions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/gbfs_versions"
	geofencingzones "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/geofencing_zones"
	manifest "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/manifest"
	stationinformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/station_information"
	stationstatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/station_status"
	systemalerts "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/system_alerts"
	systeminformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/system_information"
	systempricingplans "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/system_pricing_plans"
	systemregions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/system_regions"
	vehicleavailability "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/vehicle_availability"
	vehiclestatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/vehicle_status"
	vehicletypes "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC2/vehicle_types"

	"github.com/xeipuuv/gojsonschema"
)

const pathToTestFixturesV31RC2 = "./../../../testFixtures/v3.1-RC2/"

func LoadSchemaAndFixtureV31RC2(t *testing.T, fileName string) (gojsonschema.JSONLoader, []byte) {

	pathToSchema := "./../../../v3.1-RC2/"
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
	jsonData, err2 := os.ReadFile(pathToTestFixturesV31RC2 + fileName)
	if err2 != nil {
		t.Error("Error opening JSON file:", err2)
		return nil, nil
	}
	return schemaLoader, jsonData
}

func ValidateSchemaToUnmarshalV31RC2(t *testing.T, schemaLoader gojsonschema.JSONLoader, gbfsData interface{}) {
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

func TestGbfsV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "gbfs.json")
	gbfsData, err := gbfs.UnmarshalGbfs(jsonData)
	if err != nil {
		t.Error("Error With Unmarshal:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestGbfsVersionsV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "gbfs_versions.json")
	gbfsData, err := gbfsversions.UnmarshalGbfsVersions(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestGeofencingZonesV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "geofencing_zones.json")
	gbfsData, err := geofencingzones.UnmarshalGeofencingZones(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestManifestV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "manifest.json")
	gbfsData, err := manifest.UnmarshalManifest(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestStationInformationV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "station_information.json")
	gbfsData, err := stationinformation.UnmarshalStationInformation(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestStationStatusV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "station_status.json")
	gbfsData, err := stationstatus.UnmarshalStationStatus(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestSystemAlertsV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "system_alerts.json")
	gbfsData, err := systemalerts.UnmarshalSystemAlerts(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestSystemInformationV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "system_information.json")
	gbfsData, err := systeminformation.UnmarshalSystemInformation(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "system_pricing_plans.json")
	gbfsData, err := systempricingplans.UnmarshalSystemPricingPlans(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestSystemRegionsV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "system_regions.json")
	gbfsData, err := systemregions.UnmarshalSystemRegions(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleStatusV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "vehicle_status.json")
	gbfsData, err := vehiclestatus.UnmarshalVehicleStatus(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleTypesV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "vehicle_types.json")
	gbfsData, err := vehicletypes.UnmarshalVehicleTypes(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleAvailabilityV31RC2(t *testing.T) {
	schemaLoader, jsonData := LoadSchemaAndFixtureV31RC2(t, "vehicle_availability.json")
	gbfsData, err := vehicleavailability.UnmarshalVehicleAvailability(jsonData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	ValidateSchemaToUnmarshalV31RC2(t, schemaLoader, gbfsData)
}
