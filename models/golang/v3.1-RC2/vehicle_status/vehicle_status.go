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

// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    vehicleStatus, err := UnmarshalVehicleStatus(bytes)
//    bytes, err = vehicleStatus.Marshal()

package vehicle_status



import "encoding/json"

func UnmarshalVehicleStatus(data []byte) (VehicleStatus, error) {
	var r VehicleStatus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VehicleStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes the vehicles that are available for rent (as of v3.0, formerly
// free_bike_status).
type VehicleStatus struct {
	// Array that contains one object per vehicle as defined below.                                       
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

// Array that contains one object per vehicle as defined below.
type Data struct {
	Vehicles []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	// The date and time when any rental of the vehicle must be completed. Added in v2.3.                          
	AvailableUntil                                                                              *string            `json:"available_until,omitempty"`
	// This value represents the current percentage, expressed from 0 to 1, of fuel or battery                     
	// power remaining in the vehicle. Added in v2.3-RC.                                                           
	CurrentFuelPercent                                                                          *float64           `json:"current_fuel_percent,omitempty"`
	// The furthest distance in meters that the vehicle can travel without recharging or                           
	// refueling with the vehicle's current charge or fuel (added in v2.1-RC).                                     
	CurrentRangeMeters                                                                          *float64           `json:"current_range_meters,omitempty"`
	// The station_id of the station this vehicle must be returned to (added in v2.3-RC).                          
	HomeStationID                                                                               *string            `json:"home_station_id,omitempty"`
	// Is the vehicle currently disabled (broken)?                                                                 
	IsDisabled                                                                                  bool               `json:"is_disabled"`
	// Is the vehicle currently reserved?                                                                          
	IsReserved                                                                                  bool               `json:"is_reserved"`
	// The last time this vehicle reported its status to the operator's backend in RFC3339                         
	// format (added in v2.1-RC).                                                                                  
	LastReported                                                                                *string         `json:"last_reported,omitempty"`
	// The latitude of the vehicle.                                                                                
	Lat                                                                                         *float64           `json:"lat,omitempty"`
	// The longitude of the vehicle.                                                                               
	Lon                                                                                         *float64           `json:"lon,omitempty"`
	// The plan_id of the pricing plan this vehicle is eligible for (added in v2.2).                               
	PricingPlanID                                                                               *string            `json:"pricing_plan_id,omitempty"`
	// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added                   
	// in v1.1).                                                                                                   
	RentalUris                                                                                  *RentalUris        `json:"rental_uris,omitempty"`
	// Identifier referencing the station_id if the vehicle is currently at a station (added in                    
	// v2.1-RC2).                                                                                                  
	StationID                                                                                   *string            `json:"station_id,omitempty"`
	// List of vehicle equipment provided by the operator in addition to the accessories already                   
	// provided in the vehicle. Added in v2.3.                                                                     
	VehicleEquipment                                                                            []VehicleEquipment `json:"vehicle_equipment,omitempty"`
	// Rotating (as of v2.0) identifier of a vehicle.                                                              
	VehicleID                                                                                   string             `json:"vehicle_id"`
	// The vehicle_type_id of this vehicle (added in v2.1-RC).                                                     
	VehicleTypeID                                                                               *string            `json:"vehicle_type_id,omitempty"`
}

// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
// in v1.1).
type RentalUris struct {
	// URI that can be passed to an Android app with an intent (added in v1.1).                         
	Android                                                                                     *string `json:"android,omitempty"`
	// URI that can be used on iOS to launch the rental app for this vehicle (added in v1.1).           
	Ios                                                                                         *string `json:"ios,omitempty"`
	// URL that can be used by a web browser to show more information about renting this vehicle        
	// (added in v1.1).                                                                                 
	Web                                                                                         *string `json:"web,omitempty"`
}

type VehicleEquipment string

const (
	ChildSeatA  VehicleEquipment = "child_seat_a"
	ChildSeatB  VehicleEquipment = "child_seat_b"
	ChildSeatC  VehicleEquipment = "child_seat_c"
	SnowChains  VehicleEquipment = "snow_chains"
	WinterTires VehicleEquipment = "winter_tires"
)

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
