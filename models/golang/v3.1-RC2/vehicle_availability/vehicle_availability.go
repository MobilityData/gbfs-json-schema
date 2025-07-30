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
//    vehicleAvailability, err := UnmarshalVehicleAvailability(bytes)
//    bytes, err = vehicleAvailability.Marshal()

package vehicle_availability



import "encoding/json"

func UnmarshalVehicleAvailability(data []byte) (VehicleAvailability, error) {
	var r VehicleAvailability
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VehicleAvailability) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes the vehicle availabilities of the system.
type VehicleAvailability struct {
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

type Data struct {
	// Contains one object per vehicle.          
	Vehicles                           []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	// Array of time slots during which the specified vehicle is available.               
	Availabilities                                                         []Availability `json:"availabilities"`
	// The plan_id of the pricing plan this vehicle is eligible for                       
	PricingPlanID                                                          *string        `json:"pricing_plan_id,omitempty"`
	// The id of the station where this vehicle is located when available                 
	StationID                                                              string         `json:"station_id"`
	// List of vehicle equipment provided by the operator                                 
	VehicleEquipment                                                       []string       `json:"vehicle_equipment,omitempty"`
	// Identifier of a vehicle                                                            
	VehicleID                                                              string         `json:"vehicle_id"`
	// Unique identifier of a vehicle type as defined in vehicle_types.json               
	VehicleTypeID                                                          string         `json:"vehicle_type_id"`
}

type Availability struct {
	// Start date and time of available time slot.           
	From                                          string  `json:"from"`
	// End date and time of available time slot.             
	Until                                         *string `json:"until,omitempty"`
}

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
