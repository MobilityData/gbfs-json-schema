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

package station_information

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
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
	// Address (street number and name) where station is located. This MUST be a valid address,
	// not a free-form text description. Example: 1234 Main Street
	Address *string `json:"address,omitempty"`
	// Number of total docking points installed at this station, both available and unavailable, regardless of what
	// vehicle types are allowed at each dock. If this is a virtual station defined using the is_virtual_station field,
	// this number represents the total number of vehicles of all types that can be parked at the virtual station.
	// If the virtual station is defined by station_area, this is the number that can park within the station area.
	// If lat/lon are defined, this is the number that can park at those coordinates.
	Capacity *int64 `json:"capacity,omitempty"`
	// Contact phone of the station. Added in v2.3
	ContactPhone *common.PhoneNumber `json:"contact_phone,omitempty"`
	// Cross street or landmark where the station is located.
	CrossStreet *string `json:"cross_street,omitempty"`
	// Does the station support charging of electric vehicles?
	// true - Electric vehicle charging is available at this station.
	// false - Electric vehicle charging is not available at this station. (added in v2.3-RC)
	IsChargingStation *bool `json:"is_charging_station,omitempty"`
	// Are valet services provided at this station?
	// true - Valet services are provided at this station.
	// false - Valet services are not provided at this station.
	// If this field is empty, it is assumed that valet services are not provided at this station.
	// This field’s boolean SHOULD be set to true during the hours which valet service is provided at the station.
	// Valet service is defined as providing unlimited capacity at a station. (added in v2.1-RC)
	IsValetStation *bool `json:"is_valet_station,omitempty"`
	// Is this station a location with or without smart dock technology?
	// true - The station is a location without smart docking infrastructure. the station may be defined by a point (lat/lon) and/or station_area (below).
	// false - The station consists of smart docking infrastructure (docks).
	// This field SHOULD be published by mobility systems that have station locations without standard, internet connected
	// physical docking infrastructure. These may be racks or geofenced areas designated for rental and/or return of vehicles.
	// Locations that fit within this description SHOULD have the is_virtual_station boolean set to true. (added in v2.1-RC)
	IsVirtualStation *bool `json:"is_virtual_station,omitempty"`
	// The latitude of the station.
	Lat common.Latitude `json:"lat"`
	// The longitude fo the station.
	Lon common.Longitude `json:"lon"`
	// The public name of the station for display in maps, digital signage, and other text applications. Names SHOULD
	// reflect the station location through the use of a cross street or local landmark. Abbreviations
	// SHOULD NOT be used for names and other text (for example, "St." for "Street") unless a location is called by its
	// abbreviated name (for example, “JFK Airport”). See Text Fields and Naming.
	// Examples:
	//    Broadway and East 22nd Street
	//    Convention Center
	//    Central Park South
	Name string `json:"name"`
	// Are parking hoops present at this station?
	// true - Parking hoops are present at this station.
	// false - Parking hoops are not present at this station.
	// Parking hoops are lockable devices that are used to secure a parking space to prevent parking of unauthorized vehicles. Added in v2.3
	ParkingHoop *bool `json:"parking_hoop,omitempty"`
	// Type of parking station. Current valid values are:
	//    parking_lot (Off-street parking lot)
	//    street_parking (Curbside parking)
	//    underground_parking (Parking that is below street level, station may be non-communicating)
	//    sidewalk_parking (Park vehicle on sidewalk, out of the pedestrian right of way)
	//    other (Added in v2.3)
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
	ShortName *string `json:"short_name,omitempty"`
	// A GeoJSON MultiPolygon that describes the area of a virtual station. If station_area is supplied, then the record describes a virtual station.
	// If lat/lon and station_area are both defined, the lat/lon is the significant coordinate of the station (for example,
	// parking facility or valet drop-off and pick up point). The station_area takes precedence over any ride_allowed rules in overlapping geofencing_zones.
	StationArea *StationArea `json:"station_area,omitempty"`
	// Identifier of a station.
	StationID common.ID `json:"station_id"`
	// An object used to describe the parking capacity of virtual stations (defined using the is_virtual_station field),
	// where each key is a vehicle_type_id as described in vehicle_types.json, and the value is a number representing
	// the total number of vehicles of this type that can park within the virtual station. If the virtual station is
	// defined by station_area, this is the number that can park within the station area. If lat/lon is defined,
	// this is the number that can park at those coordinates. (added in v2.1-RC).
	VehicleCapacity map[string]int64 `json:"vehicle_capacity,omitempty"`
	// "An object where each key is a vehicle_type_id and the value is a number representing the total docking points installed at this station for each vehicle type (added in v2.1-RC).",
	VehicleTypesCapacity map[string]int64 `json:"vehicle_types_capacity,omitempty"`
}

type StationArea struct {
	Coordinates [][][][]float64 `json:"coordinates"`
	Type        Type            `json:"type"`
}

type Type string

const (
	MultiPolygon Type = "MultiPolygon"
)
