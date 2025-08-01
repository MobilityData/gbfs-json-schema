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

package stationstatus

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes the capacity and rental availability of the station
type StationStatus struct {
	header.HeaderStruct
	// Array that contains one object per station as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per station as defined below.
type Data struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	// Is the station currently on the street?
	// true - Station is installed on the street.
	// false - Station is not installed on the street.
	// Boolean SHOULD be set to true when equipment is present on the street. In seasonal systems where
	// equipment is removed during winter, boolean SHOULD be set to false during the off season.
	// May also be set to false to indicate planned (future) stations which have not yet been installed.
	IsInstalled bool `json:"is_installed"`
	// Is the station currently renting vehicles?
	// true - Station is renting vehicles. Even if the station is empty, if it would otherwise allow rentals, this value MUST be true.
	// false - Station is not renting vehicles.
	// If the station is temporarily taken out of service and not allowing rentals, this field MUST be set to false.
	// If a station becomes inaccessible to users due to road construction or other factors this field SHOULD be set to false.
	// Field SHOULD be set to false during hours or days when the system is not offering vehicles for rent.
	IsRenting bool `json:"is_renting"`
	// Is the station accepting vehicle returns?
	// true - Station is accepting vehicle returns. Even if the station is full, if it would otherwise allow vehicle returns, this value MUST be true.
	// false - Station is not accepting vehicle returns.
	// If the station is temporarily taken out of service and not allowing vehicle returns, this field MUST be set to false.
	// If a station becomes inaccessible to users due to road construction or other factors, this field SHOULD be set to false.
	IsReturning bool `json:"is_returning"`
	// The last time this station reported its status to the operator's backend in POSIX time.
	LastReported common.PosixTimestamp `json:"last_reported"`
	// REQUIRED except for stations that have unlimited docking capacity (for example, valet stations) (as of v2.0).
	// Number of functional docks physically at the station that are able to accept vehicles for return.
	// To know if the docks are accepting vehicle returns, see is_returning.
	// If is_returning = true, this is the number of docks that are currently available to accept vehicle returns.
	// If is_returning = false, this is the number of docks that would be available if the station were set to allow returns.
	NumDocksAvailable *uint64 `json:"num_docks_available,omitempty"`
	// Number of empty but disabled docks at the station.
	NumDocksDisabled *uint64 `json:"num_docks_disabled,omitempty"`
	// Number of functional vehicles physically at the station that may be offered for rental. To know if the vehicles are available for rental, see is_renting.
	// If is_renting = true, this is the number of vehicles that are currently available for rent.
	// If is_renting =false, this is the number of vehicles that would be available for rent if the station were set to allow rentals.
	NumBikesAvailable int64 `json:"num_bikes_available"`
	// Number of disabled vehicles of any type at the station. Vendors who do not want to publicize the number of
	// disabled vehicles or docks in their system can opt to omit station capacity (in station_information.json,
	// num_bikes_disabled, and num_docks_disabled (as of v2.0). If station capacity is published, then broken
	// docks/vehicles can be inferred (though not specifically whether the decreased capacity is a broken vehicle or dock).
	NumBikesDisabled *int64 `json:"num_bikes_disabled,omitempty"`
	// Identifier of a station.
	StationID common.ID `json:"station_id"`
	// Object displaying available docks by vehicle type (added in v2.1-RC).
	VehicleDocksAvailable []VehicleDocksAvailable `json:"vehicle_docks_available,omitempty"`
	// Array of objects displaying the total number of each vehicle type at the station (added
	// in v2.1-RC).
	VehicleTypesAvailable []VehicleTypesAvailable `json:"vehicle_types_available,omitempty"`
}

type VehicleDocksAvailable struct {
	// REQUIRED if vehicle_docks_available is defined. An array of strings where each string represents a vehicle_type_id
	// that is able to use a particular type of dock at the station. (added in v2.1)
	Count uint64 `json:"count"`
	// REQUIRED if vehicle_docks_available is defined. The total number of available docks at the station,
	// that can accept vehicles of the corresponding vehicle_type_id, in the vehicle_type_ids array. (added in v2.1)
	VehicleTypeIDS []string `json:"vehicle_type_ids"`
}

// REQUIRED if the vehicle_types.json file has been defined. This field's value is an array of objects.
// Each of these objects is used to model the total number of each defined vehicle type available at a station.
// The total number of vehicles from each of these objects SHOULD add up to match the value specified in the num_bikes_available field.
type VehicleTypesAvailable struct {
	// A number representing the total amount of this vehicle type at the station
	// (added in v2.1-RC).
	Count uint64 `json:"count"`
	// REQUIRED if the vehicle_types_available is defined. The vehicle_type_id of each vehicle type at the station
	// as described in vehicle_types.json. (added in v2.1-RC).
	VehicleTypeID common.ID `json:"vehicle_type_id"`
}
