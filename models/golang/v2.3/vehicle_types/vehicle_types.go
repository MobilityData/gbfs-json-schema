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

package vehicle_types

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes the types of vehicles that System operator has available for rent (added in
// v2.1-RC).
type VehicleTypes struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// Array that contains one object per vehicle type in the system as defined below.
	VehicleTypes []VehicleType `json:"vehicle_types"`
}

type VehicleType struct {
	// The capacity of the vehicle cargo space (excluding passengers), expressed in kilograms.
	CargoLoadCapacity *uint64 `json:"cargo_load_capacity,omitempty"`
	// Cargo volume available in the vehicle, expressed in liters. For cars, it corresponds to the space between the boot
	// floor, including the storage under the hatch, to the rear shelf in the trunk.
	CargoVolumeCapacity *uint64 `json:"cargo_volume_capacity,omitempty"`
	// The color of the vehicle. Added in v2.3
	Color *string `json:"color,omitempty"`
	// A plan_id as defined in system_pricing_plans.json added in v2.3-RC.
	DefaultPricingPlanID *common.ID `json:"default_pricing_plan_id,omitempty"`
	// Maximum time in minutes that a vehicle can be reserved before a rental begins added in
	// v2.3-RC.
	DefaultReserveTime *uint64 `json:"default_reserve_time,omitempty"`
	// Vehicle air quality certificate. added in v2.3.
	EcoLabel []common.EcoLabel `json:"eco_label,omitempty"`
	// The vehicle's general form factor.
	FormFactor common.FormFactor `json:"form_factor"`
	// Maximum quantity of CO2, in grams, emitted per kilometer, according to the WLTP. Added in
	// v2.3
	GCO2KM *uint64 `json:"g_CO2_km,omitempty"`
	// The name of the vehicle manufacturer. Added in v2.3
	Make *string `json:"make,omitempty"`
	// The maximum speed in kilometers per hour this vehicle is permitted to reach in accordance
	// with local permit and regulations. Added in v2.3
	MaxPermittedSpeed *uint64 `json:"max_permitted_speed,omitempty"`
	// The furthest distance in meters that the vehicle can travel without recharging or
	// refueling when it has the maximum amount of energy potential.
	MaxRangeMeters *float64 `json:"max_range_meters,omitempty"`
	// The name of the vehicle model. Added in v2.3
	Model *string `json:"model,omitempty"`
	// The public name of this vehicle type. An array with one object per supported language
	// with the following keys:
	Name *string `json:"name,omitempty"`
	// Array of all pricing plan IDs as defined in system_pricing_plans.json added in v2.3-RC.
	PricingPlanIDS []common.ID `json:"pricing_plan_ids,omitempty"`
	// The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing
	PropulsionType common.PropulsionType `json:"propulsion_type"`
	// The rated power of the motor for this vehicle type in watts. Added in v2.3
	RatedPower *uint64 `json:"rated_power,omitempty"`
	// The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as
	// return_type, and updated to return_constraint in v2.3.
	ReturnConstraint *common.ReturnConstraint `json:"return_constraint,omitempty"`
	// The number of riders (driver included) the vehicle can legally accommodate
	RiderCapacity *uint64 `json:"rider_capacity,omitempty"`
	// Description of accessories available in the vehicle.
	VehicleAccessories []common.VehicleAccessory `json:"vehicle_accessories,omitempty"`
	// An object where each key defines one of the items listed below added in v2.3-RC.
	VehicleAssets *common.VehicleAssets `json:"vehicle_assets,omitempty"`
	// URL to an image that would assist the user in identifying the vehicle. JPEG or PNG. Added
	// in v2.3
	VehicleImage *common.URL `json:"vehicle_image,omitempty"`
	// Unique identifier of a vehicle type.
	VehicleTypeID string `json:"vehicle_type_id"`
	// Number of wheels this vehicle type has. Added in v2.3
	WheelCount *uint64 `json:"wheel_count,omitempty"`
}
