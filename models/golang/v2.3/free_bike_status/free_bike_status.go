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

package freebikestatus

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes the vehicles that are available for rent (as of v2.3).
type FreeBikeStatus struct {
	header.HeaderStruct
	// Array that contains one object per bike as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per bike as defined below.
type Data struct {
	Bikes []Bike `json:"bikes"`
}

type Bike struct {
	// Rotating (as of v2.0) identifier of a vehicle.
	BikeID common.ID `json:"bike_id"`
	// The latitude of the vehicle.
	Lat *common.Latitude `json:"lat,omitempty"`
	// The longitude of the vehicle.
	Lon *common.Longitude `json:"lon,omitempty"`
	// Is the vehicle currently reserved?
	IsReserved bool `json:"is_reserved"`
	// Is the vehicle currently disabled (broken)?
	IsDisabled bool `json:"is_disabled"`
	// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added in v1.1).
	RentalUris *common.RentalUris `json:"rental_uris,omitempty"`
	// The vehicle_type_id of this vehicle (added in v2.1-RC).
	VehicleTypeID *common.ID `json:"vehicle_type_id,omitempty"`
	// The last time this vehicle reported its status to the operator's backend in POSIX time (added in v2.1-RC).
	// With PR #171 (https://github.com/MobilityData/gbfs-json-schema/pull/171), the type of LastReported was changed from number to integer to be consistent with the POSIX time format.
	// As there is no version bump, many GBFS feeds will still have LastReported as a number.
	// Therefore, the type is defined as a float64 to accommodate both cases. In the future, it is recommended to use an integer for LastReported.
	// But for now, the float64 type is used to ensure compatibility with existing feeds.
	LastReported *common.PosixTimestamp `json:"last_reported,omitempty"`
	// The furthest distance in meters that the vehicle can travel without recharging or refueling with the vehicle's current charge or fuel (added in v2.1-RC).
	CurrentRangeMeters *float64 `json:"current_range_meters,omitempty"`
	// This value represents the current percentage, expressed from 0 to 1, of fuel or battery power remaining in the vehicle. Added in v2.3-RC.
	CurrentFuelPercent *float64 `json:"current_fuel_percent,omitempty"`
	// Identifier referencing the station_id if the vehicle is currently at a station (added in v2.1-RC2).
	StationID *common.ID `json:"station_id,omitempty"`
	// The station_id of the station this vehicle must be returned to (added in v2.3-RC).
	HomeStationID *common.ID `json:"home_station_id,omitempty"`
	// The plan_id of the pricing plan this vehicle is eligible for (added in v2.2).
	PricingPlanID *common.ID `json:"pricing_plan_id,omitempty"`
	// List of vehicle equipment provided by the operator in addition to the accessories already provided in the vehicle. Added in v2.3.
	VehicleEquipment []common.VehicleEquipment `json:"vehicle_equipment,omitempty"`
	// The date and time when any rental of the vehicle must be completed. Added in v2.3.
	AvailableUntil *common.DateTime `json:"available_until,omitempty"`
}
