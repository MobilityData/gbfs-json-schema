package tests

import (
	"encoding/json"
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
)

func TestGbfs30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/gbfs.json", TestFixturesV30+"gbfs.json")
	var gbfsData gbfs.Gbfs
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGbfsVersions30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/gbfs_versions.json", TestFixturesV30+"gbfs_versions.json")
	var gbfsData gbfsversions.GbfsVersions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGeofencingZones30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/geofencing_zones.json", TestFixturesV30+"geofencing_zones.json")
	var gbfsData geofencingzones.GeofencingZones
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestManifest30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/manifest.json", TestFixturesV30+"manifest.json")
	var gbfsData manifest.Manifest
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationPhysicalLimitedHoursOfOperation30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/station_information.json", TestFixturesV30+"station_information_physical_station_limited_hours_of_operation.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationVirtual30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/station_information.json", TestFixturesV30+"station_information_virtual_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationStatus30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/station_status.json", TestFixturesV30+"station_status.json")
	var gbfsData stationstatus.StationStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemAlerts30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/system_alerts.json", TestFixturesV30+"system_alerts.json")
	var gbfsData systemalerts.SystemAlerts
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemInformation30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/system_information.json", TestFixturesV30+"system_information.json")
	var gbfsData systeminformation.SystemInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanA30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/system_pricing_plans.json", TestFixturesV30+"system_pricing_plans_a.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlan30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/system_pricing_plans.json", TestFixturesV30+"system_pricing_plans_b.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemRegions30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/system_regions.json", TestFixturesV30+"system_regions.json")
	var gbfsData systemregions.SystemRegions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleStatusCarsharing30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/vehicle_status.json", TestFixturesV30+"vehicle_status_carsharing.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleStatusMicromobility30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/vehicle_status.json", TestFixturesV30+"vehicle_status_micromobility.json")
	var gbfsData vehiclestatus.VehicleStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleTypes30(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v3.0/vehicle_types.json", TestFixturesV30+"vehicle_types.json")
	var gbfsData vehicletypes.VehicleTypes
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}
