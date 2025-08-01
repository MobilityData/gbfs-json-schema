package tests

import (
	"encoding/json"
	"testing"

	gbfs "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/gbfs"
	gbfsversions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/gbfs_versions"
	geofencingzones "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/geofencing_zones"
	manifest "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/manifest"
	stationinformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/station_information"
	stationstatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/station_status"
	systemalerts "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/system_alerts"
	systeminformation "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/system_information"
	systempricingplans "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/system_pricing_plans"
	systemregions "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/system_regions"
	vehiclestatus "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/vehicle_status"
	vehicletypes "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/vehicle_types"
)

func TestGbfs31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/gbfs.json", TestFixturesV31RC+"gbfs.json")
	var gbfsData gbfs.Gbfs
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGbfsVersions31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/gbfs_versions.json", TestFixturesV31RC+"gbfs_versions.json")
	var gbfsData gbfsversions.GbfsVersions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGeofencingZones31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/geofencing_zones.json", TestFixturesV31RC+"geofencing_zones.json")
	var gbfsData geofencingzones.GeofencingZones
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestManifest31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/manifest.json", TestFixturesV31RC+"manifest.json")
	var gbfsData manifest.Manifest
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationPhysicalLimitedHoursOfOperation31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/station_information.json", TestFixturesV31RC+"station_information_physical_station_limited_hours_of_operation.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationVirtual31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/station_information.json", TestFixturesV31RC+"station_information_virtual_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationStatus31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/station_status.json", TestFixturesV31RC+"station_status.json")
	var gbfsData stationstatus.StationStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemAlerts31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/system_alerts.json", TestFixturesV31RC+"system_alerts.json")
	var gbfsData systemalerts.SystemAlerts
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemInformation31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/system_information.json", TestFixturesV31RC+"system_information.json")
	var gbfsData systeminformation.SystemInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanA31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/system_pricing_plans.json", TestFixturesV31RC+"system_pricing_plans_a.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlan31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/system_pricing_plans.json", TestFixturesV31RC+"system_pricing_plans_b.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemRegions31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/system_regions.json", TestFixturesV31RC+"system_regions.json")
	var gbfsData systemregions.SystemRegions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleStatusCarsharing31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/vehicle_status.json", TestFixturesV31RC+"vehicle_status_carsharing.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleStatusMicromobility31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/vehicle_status.json", TestFixturesV31RC+"vehicle_status_micromobility.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleTypes31RC(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.1-RC/vehicle_types.json", TestFixturesV31RC+"vehicle_types.json")
	var gbfsData vehicletypes.VehicleTypes
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}
