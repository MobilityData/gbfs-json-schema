package tests

import (
	"encoding/json"
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
	vehicle_types "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/vehicle_types"
)


func TestGbfs23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/gbfs.json", TestFixturesV23+"gbfs.json")
	var gbfsData gbfs.Gbfs
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGbfsVersions23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/gbfs_versions.json", TestFixturesV23+"gbfs_versions.json")
	var gbfsData gbfsversions.GbfsVersions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestGeofencingZones23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/geofencing_zones.json", TestFixturesV23+"geofencing_zones.json")
	var gbfsData geofencingzones.GeofencingZones
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationPhysical23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/station_information.json", TestFixturesV23+"station_information_physical_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationInformationVirtual23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/station_information.json", TestFixturesV23+"station_information_virtual_station.json")
	var gbfsData stationinformation.StationInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestStationStatus23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/station_status.json", TestFixturesV23+"station_status.json")
	var gbfsData stationstatus.StationStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemAlerts23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_alerts.json", TestFixturesV23+"system_alerts.json")
	var gbfsData systemalerts.SystemAlerts
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemInformation23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_information.json", TestFixturesV23+"system_information.json")
	var gbfsData systeminformation.SystemInformation
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanA23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_pricing_plans.json", TestFixturesV23+"system_pricing_plans_a.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemPricingPlanB23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_pricing_plans.json", TestFixturesV23+"system_pricing_plans_b.json")
	var gbfsData systempricingplans.SystemPricingPlans
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemRegions23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_regions.json", TestFixturesV23+"system_regions.json")
	var gbfsData systemregions.SystemRegions
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemCalendar23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_calendar.json", TestFixturesV23+"system_calendar.json")
	var gbfsData systemcalendar.SystemCalendar
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestSystemOperatingHours23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/system_hours.json", TestFixturesV23+"system_hours.json")
	var gbfsData systemhours.SystemHours
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestFreeBikeStatusCarsharing23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/free_bike_status.json", TestFixturesV23+"free_bike_status_carsharing.json")
	var gbfsData freebikestatus.FreeBikeStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestFreeBikeStatusMicromobility23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/free_bike_status.json", TestFixturesV23+"free_bike_status_micromobility.json")
	var gbfsData freebikestatus.FreeBikeStatus
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}

func TestVehicleTypes23(t *testing.T) {
	schemaLoader, jsonData := loadSchemaAndFixture(t, "./../../../v2.3/vehicle_types.json", TestFixturesV23+"vehicle_types.json")
	var gbfsData vehicle_types.VehicleTypes
	err := json.Unmarshal(jsonData, &gbfsData)
	if err != nil {
		t.Error("Error UnmarshalGbfs:", err)
		return
	}
	validateSchemaToUnmarshal(t, schemaLoader, gbfsData)
}