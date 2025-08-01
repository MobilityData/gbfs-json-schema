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

package station_status

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
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
	IsInstalled bool `json:"is_installed"`
	// Is the station currently renting vehicles?
	IsRenting bool `json:"is_renting"`
	// Is the station accepting vehicle returns?
	IsReturning bool `json:"is_returning"`
	// The last time this station reported its status to the operator's backend in RFC3339
	// format.
	LastReported common.Timestamp `json:"last_reported"`
	// Number of functional docks physically at the station.
	NumDocksAvailable *uint64 `json:"num_docks_available,omitempty"`
	// Number of empty but disabled docks at the station.
	NumDocksDisabled *uint64 `json:"num_docks_disabled,omitempty"`
	// Number of vehicles of any type physically available for rental at the station.
	NumVehiclesAvailable uint64 `json:"num_vehicles_available"`
	// Number of disabled vehicles of any type at the station.
	NumVehiclesDisabled *uint64 `json:"num_vehicles_disabled,omitempty"`
	// Identifier of a station.
	StationID common.ID `json:"station_id"`
	// Object displaying available docks by vehicle type (added in v2.1-RC).
	VehicleDocksAvailable []VehicleDocksAvailable `json:"vehicle_docks_available,omitempty"`
	// Array of objects displaying the total number of each vehicle type at the station (added
	// in v2.1-RC).
	VehicleTypesAvailable []VehicleTypesAvailable `json:"vehicle_types_available,omitempty"`
}

type VehicleDocksAvailable struct {
	// A number representing the total number of available docks for the defined vehicle type
	// (added in v2.1-RC).
	Count uint64 `json:"count"`
	// An array of strings where each string represents a vehicle_type_id that is able to use a
	// particular type of dock at the station (added in v2.1-RC).
	VehicleTypeIDs []common.ID `json:"vehicle_type_ids"`
}

type VehicleTypesAvailable struct {
	// A number representing the total amount of this vehicle type at the station (added in
	// v2.1-RC).
	Count uint64 `json:"count"`
	// The vehicle_type_id of vehicle at the station (added in v2.1-RC).
	VehicleTypeID common.ID `json:"vehicle_type_id"`
}
