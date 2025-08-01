// Copyright 2024 MobilityData
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package geofencing_zones

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
	geojson "github.com/paulmach/orb/geojson"
)

// Describes geofencing zones and their associated rules and attributes (added in v2.1-RC).
type GeofencingZones struct {
	header.HeaderStruct
	// Array that contains geofencing information for the system.
	Data Data `json:"data"`
}

// Array that contains geofencing information for the system.
type Data struct {
	// Each geofenced zone and its associated rules and attributes is described as an object
	// within the array of features.
	GeofencingZones *geojson.FeatureCollection `json:"geofencing_zones"`
	// Array of Rule objects defining restrictions that apply globally in all areas as the
	// default restrictions, except where overridden with an explicit geofencing zone.
	GlobalRules []GlobalRule `json:"global_rules"`
}

type GlobalRule struct {
	// What is the maximum speed allowed, in kilometers per hour?
	MaximumSpeedKph *uint64 `json:"maximum_speed_kph,omitempty"`
	// Is the ride allowed to end in this zone?
	RideEndAllowed bool `json:"ride_end_allowed"`
	// Is the ride allowed to start in this zone?
	RideStartAllowed bool `json:"ride_start_allowed"`
	// Is the ride allowed to travel through this zone?
	RideThroughAllowed bool `json:"ride_through_allowed"`
	// Vehicle MUST be parked at stations defined in station_information.json within this
	// geofence zone
	StationParking *bool `json:"station_parking,omitempty"`
	// Array of vehicle type IDs for which these restrictions apply.
	VehicleTypeIDS []common.ID `json:"vehicle_type_ids,omitempty"`
}
