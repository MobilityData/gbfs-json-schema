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

const pathToTestFixtures31RC2 = "./../../../testFixtures/v3.1-RC2/"

func loadSchemaAndFixture31RC2(t *testing.T, schemaPath string, fixturePath string) (gojsonschema.JSONLoader, []byte) {
	schemaDataBytes, schemaErr := os.ReadFile(schemaPath)
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
	jsonData, err2 := os.ReadFile(fixturePath)
	if err2 != nil {
		t.Error("Error opening JSON file:", err2)
		return nil, nil
	}
	return schemaLoader, jsonData
}

func validateSchemaToUnmarshal31RC2(t *testing.T, schemaLoader gojsonschema.JSONLoader, gbfsData interface{}) {
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

func TestGbfs31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/gbfs.json", pathToTestFixtures31RC2+"gbfs.json")
	var gbfsData gbfs.Gbfs
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error With Unmarshal:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestGbfsVersions31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/gbfs_versions.json", pathToTestFixtures31RC2+"gbfs_versions.json")
	var gbfsData gbfsversions.GbfsVersions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestGeofencingZones31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/geofencing_zones.json", pathToTestFixtures31RC2+"geofencing_zones.json")
	var gbfsData geofencingzones.GeofencingZones
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestManifest31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/manifest.json", pathToTestFixtures31RC2+"manifest.json")
	var gbfsData manifest.Manifest
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestStationInformationPhysicalLimitedHoursOfOperation31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/station_information.json", pathToTestFixtures31RC2+"station_information_physical_station_limited_hours_of_operation.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestStationInformationVirtual31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/station_information.json", pathToTestFixtures31RC2+"station_information_virtual_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestStationStatus31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/station_status.json", pathToTestFixtures31RC2+"station_status.json")
	var gbfsData stationstatus.StationStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestSystemAlerts31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/system_alerts.json", pathToTestFixtures31RC2+"system_alerts.json")
	var gbfsData systemalerts.SystemAlerts
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestSystemInformation31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/system_information.json", pathToTestFixtures31RC2+"system_information.json")
	var gbfsData systeminformation.SystemInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanA31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/system_pricing_plans.json", pathToTestFixtures31RC2+"system_pricing_plans_a.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlan31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/system_pricing_plans.json", pathToTestFixtures31RC2+"system_pricing_plans_b.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestSystemRegions31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/system_regions.json", pathToTestFixtures31RC2+"system_regions.json")
	var gbfsData systemregions.SystemRegions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleStatusCarsharing31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/vehicle_status.json", pathToTestFixtures31RC2+"vehicle_status_carsharing.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleStatusMicromobility31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/vehicle_status.json", pathToTestFixtures31RC2+"vehicle_status_micromobility.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleTypes31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/vehicle_types.json", pathToTestFixtures31RC2+"vehicle_types.json")
	var gbfsData vehicletypes.VehicleTypes
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}

func TestVehicleAvailability31RC2(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture31RC2(t, "./../../../v3.1-RC2/vehicle_availability.json", pathToTestFixtures31RC2+"vehicle_availability.json")
	var gbfsData vehicleavailability.VehicleAvailability
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal31RC2(t, schemaLoader, gbfsData)
}
