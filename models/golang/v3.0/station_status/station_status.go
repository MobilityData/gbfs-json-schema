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
//    stationStatus, err := UnmarshalStationStatus(bytes)
//    bytes, err = stationStatus.Marshal()

package station_status



import "encoding/json"

func UnmarshalStationStatus(data []byte) (StationStatus, error) {
	var r StationStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StationStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes the capacity and rental availability of the station
type StationStatus struct {
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
	// Is the station currently on the street?                                                                        
	IsInstalled                                                                               bool                    `json:"is_installed"`
	// Is the station currently renting vehicles?                                                                     
	IsRenting                                                                                 bool                    `json:"is_renting"`
	// Is the station accepting vehicle returns?                                                                      
	IsReturning                                                                               bool                    `json:"is_returning"`
	// The last time this station reported its status to the operator's backend in RFC3339                            
	// format.                                                                                                        
	LastReported                                                                              string               `json:"last_reported"`
	// Number of functional docks physically at the station.                                                          
	NumDocksAvailable                                                                         *int64                  `json:"num_docks_available,omitempty"`
	// Number of empty but disabled docks at the station.                                                             
	NumDocksDisabled                                                                          *int64                  `json:"num_docks_disabled,omitempty"`
	// Number of vehicles of any type physically available for rental at the station.                                 
	NumVehiclesAvailable                                                                      int64                   `json:"num_vehicles_available"`
	// Number of disabled vehicles of any type at the station.                                                        
	NumVehiclesDisabled                                                                       *int64                  `json:"num_vehicles_disabled,omitempty"`
	// Identifier of a station.                                                                                       
	StationID                                                                                 string                  `json:"station_id"`
	// Object displaying available docks by vehicle type (added in v2.1-RC).                                          
	VehicleDocksAvailable                                                                     []VehicleDocksAvailable `json:"vehicle_docks_available,omitempty"`
	// Array of objects displaying the total number of each vehicle type at the station (added                        
	// in v2.1-RC).                                                                                                   
	VehicleTypesAvailable                                                                     []VehicleTypesAvailable `json:"vehicle_types_available,omitempty"`
}

type VehicleDocksAvailable struct {
	// A number representing the total number of available docks for the defined vehicle type           
	// (added in v2.1-RC).                                                                              
	Count                                                                                      int64    `json:"count"`
	// An array of strings where each string represents a vehicle_type_id that is able to use a         
	// particular type of dock at the station (added in v2.1-RC).                                       
	VehicleTypeIDS                                                                             []string `json:"vehicle_type_ids"`
}

type VehicleTypesAvailable struct {
	// A number representing the total amount of this vehicle type at the station (added in       
	// v2.1-RC).                                                                                  
	Count                                                                                  int64  `json:"count"`
	// The vehicle_type_id of vehicle at the station (added in v2.1-RC).                          
	VehicleTypeID                                                                          string `json:"vehicle_type_id"`
}

type Version string

const (
	The30 Version = "3.0"
)
