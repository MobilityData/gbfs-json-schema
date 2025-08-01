package common

import (
	"encoding/json"

	gbfs_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/gbfs"
	gbfs_versions_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/gbfs_versions"
	geofencing_zones_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/geofencing_zones"
	header_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
	manifest_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/manifest"
	station_information_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/station_information"
	station_status_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/station_status"
	system_alerts_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_alerts"
	system_information_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_information"
	system_pricing_plans_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_pricing_plans"
	system_regions_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/system_regions"
	vehicle_status_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/vehicle_status"
	vehicle_types_3_0 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/vehicle_types"

	free_bike_status_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/free_bike_status"
	gbfs_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/gbfs"
	gbfs_versions_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/gbfs_versions"
	geofencing_zones_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/geofencing_zones"
	header_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
	station_information_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/station_information"
	station_status_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/station_status"
	system_alerts_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_alerts"
	system_calendar_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_calendar"
	system_hours_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_hours"
	system_information_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_information"
	system_pricing_plans_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_pricing_plans"
	system_regions_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/system_regions"
	vehicle_types_2_3 "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/vehicle_types"
)

// feedStructs is a type that includes all the GBFS structs from both v2.3 and v3.0.
type feedStructs interface {
	gbfs_2_3.Gbfs |
		header_2_3.HeaderStruct |
		gbfs_versions_2_3.GbfsVersions |
		geofencing_zones_2_3.GeofencingZones |
		station_information_2_3.StationInformation |
		station_status_2_3.StationStatus |
		system_alerts_2_3.SystemAlerts |
		system_calendar_2_3.SystemCalendar |
		system_hours_2_3.SystemHours |
		system_information_2_3.SystemInformation |
		system_pricing_plans_2_3.SystemPricingPlans |
		system_regions_2_3.SystemRegions |
		free_bike_status_2_3.FreeBikeStatus |
		vehicle_types_2_3.VehicleTypes |

		gbfs_3_0.Gbfs |
		header_3_0.HeaderStruct |
		gbfs_versions_3_0.GbfsVersions |
		geofencing_zones_3_0.GeofencingZones |
		station_information_3_0.StationInformation |
		station_status_3_0.StationStatus |
		system_alerts_3_0.SystemAlerts |
		system_information_3_0.SystemInformation |
		system_pricing_plans_3_0.SystemPricingPlans |
		system_regions_3_0.SystemRegions |
		vehicle_status_3_0.VehicleStatus |
		vehicle_types_3_0.VehicleTypes |
		manifest_3_0.Manifest
}

// Generic helper functions for marshaling and unmarshaling structs. Requires Go 1.18+ for generics.
func UnmarshalStruct[T feedStructs](data []byte) (T, error) {
	var r T
	err := json.Unmarshal(data, &r)
	return r, err
}

func MarshalStruct[T feedStructs](v *T) ([]byte, error) {
	return json.Marshal(v)
}
