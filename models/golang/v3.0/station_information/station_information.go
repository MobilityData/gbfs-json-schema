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

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    stationInformation, err := UnmarshalStationInformation(bytes)
//    bytes, err = stationInformation.Marshal()

package station_information



import "encoding/json"

func UnmarshalStationInformation(data []byte) (StationInformation, error) {
	var r StationInformation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StationInformation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// List of all stations, their capacities and locations. REQUIRED of systems utilizing docks.
type StationInformation struct {
	// Array that contains one object per station as defined below.                                       
	Data                                                                                        Data      `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                      
	LastUpdated                                                                                 string `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should          
	// always be refreshed).                                                                              
	TTL                                                                                         int64     `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework              
	// (added in v1.1).                                                                                   
	Version                                                                                     Version   `json:"version"`
}

// Array that contains one object per station as defined below.
type Data struct {
	Stations []Station `json:"stations"`
}

type Station struct {
	// Address where station is located.                                                                               
	Address                                                                                     *string                `json:"address,omitempty"`
	// Number of total docking points installed at this station, both available and unavailable.                       
	Capacity                                                                                    *int64                 `json:"capacity,omitempty"`
	// Contact phone of the station. Added in v2.3                                                                     
	ContactPhone                                                                                *string                `json:"contact_phone,omitempty"`
	// Cross street or landmark where the station is located.                                                          
	CrossStreet                                                                                 *string                `json:"cross_street,omitempty"`
	// Does the station support charging of electric vehicles? (added in v2.3-RC)                                      
	IsChargingStation                                                                           *bool                  `json:"is_charging_station,omitempty"`
	// Are valet services provided at this station? (added in v2.1-RC)                                                 
	IsValetStation                                                                              *bool                  `json:"is_valet_station,omitempty"`
	// Is this station a location with or without physical infrastructure? (added in v2.1-RC)                          
	IsVirtualStation                                                                            *bool                  `json:"is_virtual_station,omitempty"`
	// The latitude of the station.                                                                                    
	Lat                                                                                         float64                `json:"lat"`
	// The longitude fo the station.                                                                                   
	Lon                                                                                         float64                `json:"lon"`
	// Public name of the station.                                                                                     
	Name                                                                                        []Name                 `json:"name"`
	// Are parking hoops present at this station? Added in v2.3                                                        
	ParkingHoop                                                                                 *bool                  `json:"parking_hoop,omitempty"`
	// Type of parking station. Added in v2.3                                                                          
	ParkingType                                                                                 *ParkingType           `json:"parking_type,omitempty"`
	// Postal code where station is located.                                                                           
	PostCode                                                                                    *string                `json:"post_code,omitempty"`
	// Identifier of the region where the station is located.                                                          
	RegionID                                                                                    *string                `json:"region_id,omitempty"`
	// Payment methods accepted at this station.                                                                       
	RentalMethods                                                                               []RentalMethod         `json:"rental_methods,omitempty"`
	// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added                       
	// in v1.1).                                                                                                       
	RentalUris                                                                                  *RentalUris            `json:"rental_uris,omitempty"`
	// Short name or other type of identifier.                                                                         
	ShortName                                                                                   []ShortName            `json:"short_name,omitempty"`
	// A multipolygon that describes the area of a virtual station (added in v2.1-RC).                                 
	StationArea                                                                                 *StationArea           `json:"station_area,omitempty"`
	// Identifier of a station.                                                                                        
	StationID                                                                                   string                 `json:"station_id"`
	// Hours of operation for the station in OSM opening_hours format.                                                 
	StationOpeningHours                                                                         *string                `json:"station_opening_hours,omitempty"`
	// This field's value is an array of objects containing the keys vehicle_type_ids and count                        
	// defined below. These objects are used to model the total docking capacity of a station,                         
	// both available and unavailable, for each type of vehicle that may dock at this station.                         
	VehicleDocksCapacity                                                                        []VehicleDocksCapacity `json:"vehicle_docks_capacity,omitempty"`
	// This field's value is an array of objects containing the keys vehicle_type_ids and count                        
	// defined below. These objects are used to model the parking capacity of virtual stations                         
	// (defined using the is_virtual_station field) for each vehicle type that can be returned                         
	// to this station.                                                                                                
	VehicleTypesCapacity                                                                        []VehicleTypesCapacity `json:"vehicle_types_capacity,omitempty"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
// in v1.1).
type RentalUris struct {
	// URI that can be passed to an Android app with an intent (added in v1.1).                         
	Android                                                                                     *string `json:"android,omitempty"`
	// URI that can be used on iOS to launch the rental app for this station (added in v1.1).           
	Ios                                                                                         *string `json:"ios,omitempty"`
	// URL that can be used by a web browser to show more information about renting a vehicle at        
	// this station (added in v1.1).                                                                    
	Web                                                                                         *string `json:"web,omitempty"`
}

type ShortName struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// A multipolygon that describes the area of a virtual station (added in v2.1-RC).
type StationArea struct {
	Coordinates [][][][]float64 `json:"coordinates"`
	Type        Type            `json:"type"`
}

type VehicleDocksCapacity struct {
	// A number representing the total number of docks at the station, both available and               
	// unavailable, that may accept the vehicle types specified by vehicle_type_ids.                    
	Count                                                                                      int64    `json:"count"`
	// An array of strings where each string represents a vehicle_type_id that is able to use a         
	// particular type of dock at the station.                                                          
	VehicleTypeIDS                                                                             []string `json:"vehicle_type_ids"`
}

type VehicleTypesCapacity struct {
	// A number representing the total number of vehicles of the specified vehicle_type_ids that         
	// can park within the virtual station.                                                              
	Count                                                                                       int64    `json:"count"`
	// The vehicle_type_ids, as defined in vehicle_types.json, that may park at the virtual              
	// station.                                                                                          
	VehicleTypeIDS                                                                              []string `json:"vehicle_type_ids"`
}

// Type of parking station. Added in v2.3
type ParkingType string

const (
	Other              ParkingType = "other"
	ParkingLot         ParkingType = "parking_lot"
	SidewalkParking    ParkingType = "sidewalk_parking"
	StreetParking      ParkingType = "street_parking"
	UndergroundParking ParkingType = "underground_parking"
)

type RentalMethod string

const (
	Accountnumber RentalMethod = "accountnumber"
	Androidpay    RentalMethod = "androidpay"
	Applepay      RentalMethod = "applepay"
	Creditcard    RentalMethod = "creditcard"
	Key           RentalMethod = "key"
	Paypass       RentalMethod = "paypass"
	Phone         RentalMethod = "phone"
	Transitcard   RentalMethod = "transitcard"
)

type Type string

const (
	MultiPolygon Type = "MultiPolygon"
)

type Version string

const (
	The30 Version = "3.0"
)
