package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	freebikestatus "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/free_bike_status"
	gbfs "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/gbfs"
	gbfsversions "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/gbfs_versions"
	geofencingzones "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/geofencing_zones"
	stationinformation "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/station_information"
	stationstatus "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/station_status"
	systemalerts "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_alerts"
	systemcalendar "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_calendar"
	systemhours "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_hours"
	systeminformation "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_information"
	systempricingplans "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_pricing_plans"
	systemregions "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_regions"

	"github.com/xeipuuv/gojsonschema"
)

const pathToTestFixtures23 = "./../../../testFixtures/v2.3/"

func loadSchemaAndFixture23(t *testing.T, schemaPath string, fixturePath string) (gojsonschema.JSONLoader, []byte) {
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

func validateSchemaToUnmarshal23(t *testing.T, schemaLoader gojsonschema.JSONLoader, gbfsData interface{}) {
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

func TestGbfs23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/gbfs.json", pathToTestFixtures23+"gbfs.json")
	var gbfsData gbfs.Gbfs
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error With Unmarshal:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestGbfsVersions23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/gbfs_versions.json", pathToTestFixtures23+"gbfs_versions.json")
	var gbfsData gbfsversions.GbfsVersions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestGeofencingZones23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/geofencing_zones.json", pathToTestFixtures23+"geofencing_zones.json")
	var gbfsData geofencingzones.GeofencingZones
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestStationInformationPhysical23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/station_information.json", pathToTestFixtures23+"station_information_physical_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestStationInformationVirtual23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/station_information.json", pathToTestFixtures23+"station_information_virtual_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestStationStatus23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/station_status.json", pathToTestFixtures23+"station_status.json")
	var gbfsData stationstatus.StationStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemAlerts23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_alerts.json", pathToTestFixtures23+"system_alerts.json")
	var gbfsData systemalerts.SystemAlerts
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemInformation23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_information.json", pathToTestFixtures23+"system_information.json")
	var gbfsData systeminformation.SystemInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanA23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_pricing_plans.json", pathToTestFixtures23+"system_pricing_plans_a.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanB23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_pricing_plans.json", pathToTestFixtures23+"system_pricing_plans_b.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemRegions23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_regions.json", pathToTestFixtures23+"system_regions.json")
	var gbfsData systemregions.SystemRegions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemCalendar23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_calendar.json", pathToTestFixtures23+"system_calendar.json")
	var gbfsData systemcalendar.SystemCalendar
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestSystemOperatingHours23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/system_hours.json", pathToTestFixtures23+"system_hours.json")
	var gbfsData systemhours.SystemHours
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestFreeBikeStatusCarsharing23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/free_bike_status.json", pathToTestFixtures23+"free_bike_status_carsharing.json")
	var gbfsData freebikestatus.FreeBikeStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}

func TestFreeBikeStatusMicromobility23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture23(t, "./../../../v2.3/free_bike_status.json", pathToTestFixtures23+"free_bike_status_micromobility.json")
	var gbfsData freebikestatus.FreeBikeStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal23(t, schemaLoader, gbfsData)
}
