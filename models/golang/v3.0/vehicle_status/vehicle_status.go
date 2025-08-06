// Copyright 2025 MobilityData
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

package vehicle_status

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
)

// Describes the vehicles that are available for rent (as of v3.0, formerly
// free_bike_status).
type VehicleStatus struct {
	header.HeaderStruct
	// Array that contains one object per vehicle as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per vehicle as defined below.
type Data struct {
	Vehicles []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	// The date and time when any rental of the vehicle must be completed. Added in v2.3.
	AvailableUntil *common.DateTime `json:"available_until,omitempty"`
	// This value represents the current percentage, expressed from 0 to 1, of fuel or battery
	// power remaining in the vehicle. Added in v2.3-RC.
	CurrentFuelPercent *float64 `json:"current_fuel_percent,omitempty"`
	// The furthest distance in meters that the vehicle can travel without recharging or
	// refueling with the vehicle's current charge or fuel (added in v2.1-RC).
	CurrentRangeMeters *float64 `json:"current_range_meters,omitempty"`
	// The station_id of the station this vehicle must be returned to (added in v2.3-RC).
	HomeStationID *common.ID `json:"home_station_id,omitempty"`
	// Is the vehicle currently disabled (broken)?
	IsDisabled bool `json:"is_disabled"`
	// Is the vehicle currently reserved?
	IsReserved bool `json:"is_reserved"`
	// The last time this vehicle reported its status to the operator's backend in RFC3339
	// format (added in v2.1-RC).
	LastReported *common.Timestamp `json:"last_reported,omitempty"`
	// The latitude of the vehicle.
	Lat *common.Latitude `json:"lat,omitempty"`
	// The longitude of the vehicle.
	Lon *common.Longitude `json:"lon,omitempty"`
	// The plan_id of the pricing plan this vehicle is eligible for (added in v2.2).
	PricingPlanID *common.ID `json:"pricing_plan_id,omitempty"`
	// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
	// in v1.1).
	RentalUris *common.RentalUris `json:"rental_uris,omitempty"`
	// Identifier referencing the station_id if the vehicle is currently at a station (added in
	// v2.1-RC2).
	StationID *common.ID `json:"station_id,omitempty"`
	// List of vehicle equipment provided by the operator in addition to the accessories already
	// provided in the vehicle. Added in v2.3.
	VehicleEquipment []common.VehicleEquipment `json:"vehicle_equipment,omitempty"`
	// Rotating (as of v2.0) identifier of a vehicle.
	VehicleID common.ID `json:"vehicle_id"`
	// The vehicle_type_id of this vehicle (added in v2.1-RC).
	VehicleTypeID *common.ID `json:"vehicle_type_id,omitempty"`
}
