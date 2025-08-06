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

package station_information

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
	"github.com/paulmach/orb/geojson"
)

// List of all stations, their capacities and locations. REQUIRED of systems utilizing docks.
type StationInformation struct {
	header.HeaderStruct
	// Array that contains one object per station as defined below.
	Data Data `json:"data"`
}

// Array that contains one object per station as defined below.
type Data struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	// Address where station is located.
	Address *string `json:"address,omitempty"`
	// Number of total docking points installed at this station, both available and unavailable.
	Capacity *uint64 `json:"capacity,omitempty"`
	// Contact phone of the station. Added in v2.3
	ContactPhone *common.PhoneNumber `json:"contact_phone,omitempty"`
	// Cross street or landmark where the station is located.
	CrossStreet *string `json:"cross_street,omitempty"`
	// Does the station support charging of electric vehicles? (added in v2.3-RC)
	IsChargingStation *bool `json:"is_charging_station,omitempty"`
	// Are valet services provided at this station? (added in v2.1-RC)
	IsValetStation *bool `json:"is_valet_station,omitempty"`
	// Is this station a location with or without physical infrastructure? (added in v2.1-RC)
	IsVirtualStation *bool `json:"is_virtual_station,omitempty"`
	// The latitude of the station.
	Lat common.Latitude `json:"lat"`
	// The longitude fo the station.
	Lon common.Longitude `json:"lon"`
	// Public name of the station.
	Name []common.LocalizedString `json:"name"`
	// Are parking hoops present at this station? Added in v2.3
	ParkingHoop *bool `json:"parking_hoop,omitempty"`
	// Type of parking station. Added in v2.3
	ParkingType *common.ParkingType `json:"parking_type,omitempty"`
	// Postal code where station is located.
	PostCode *string `json:"post_code,omitempty"`
	// Identifier of the region where the station is located.
	RegionID *common.ID `json:"region_id,omitempty"`
	// Payment methods accepted at this station.
	RentalMethods []common.RentalMethod `json:"rental_methods,omitempty"`
	// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
	// in v1.1).
	RentalUris *common.RentalUris `json:"rental_uris,omitempty"`
	// Short name or other type of identifier.
	ShortName []common.LocalizedString `json:"short_name,omitempty"`
	// A multipolygon that describes the area of a virtual station (added in v2.1-RC).
	StationArea *geojson.MultiPolygon `json:"station_area,omitempty"`
	// Identifier of a station.
	StationID common.ID `json:"station_id"`
	// Hours of operation for the station in OSM opening_hours format.
	StationOpeningHours *string `json:"station_opening_hours,omitempty"`
	// This field's value is an array of objects containing the keys vehicle_type_ids and count
	// defined below. These objects are used to model the total docking capacity of a station,
	// both available and unavailable, for each type of vehicle that may dock at this station.
	VehicleDocksCapacity []VehicleDocksCapacity `json:"vehicle_docks_capacity,omitempty"`
	// This field's value is an array of objects containing the keys vehicle_type_ids and count
	// defined below. These objects are used to model the parking capacity of virtual stations
	// (defined using the is_virtual_station field) for each vehicle type that can be returned
	// to this station.
	VehicleTypesCapacity []VehicleTypesCapacity `json:"vehicle_types_capacity,omitempty"`
}

type VehicleDocksCapacity struct {
	// A number representing the total number of docks at the station, both available and
	// unavailable, that may accept the vehicle types specified by vehicle_type_ids.
	Count uint64 `json:"count"`
	// An array of strings where each string represents a vehicle_type_id that is able to use a
	// particular type of dock at the station.
	VehicleTypeIDs []common.ID `json:"vehicle_type_ids"`
}

type VehicleTypesCapacity struct {
	// A number representing the total number of vehicles of the specified vehicle_type_ids that
	// can park within the virtual station.
	Count uint64 `json:"count"`
	// The vehicle_type_ids, as defined in vehicle_types.json, that may park at the virtual
	// station.
	VehicleTypeIDs []common.ID `json:"vehicle_type_ids"`
}
