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
//    gbfs, err := UnmarshalGbfs(bytes)
//    bytes, err = gbfs.Marshal()

package gbfs



import "encoding/json"

func UnmarshalGbfs(data []byte) (Gbfs, error) {
	var r Gbfs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Gbfs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Auto-discovery file that links to all of the other files published by the system.
type Gbfs struct {
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
	// An array of all of the feeds that are published by the auto-discovery file. Each element       
	// in the array is an object with the keys below.                                                 
	Feeds                                                                                      []Feed `json:"feeds"`
}

type Feed struct {
	// Key identifying the type of feed this is. The key must be the base file name defined in       
	// the spec for the corresponding feed type.                                                     
	Name                                                                                      Name   `json:"name"`
	// URL for the feed.                                                                             
	URL                                                                                       string `json:"url"`
}

// Key identifying the type of feed this is. The key must be the base file name defined in
// the spec for the corresponding feed type.
type Name string

const (
	GbfsVersions        Name = "gbfs_versions"
	GeofencingZones     Name = "geofencing_zones"
	NameGbfs            Name = "gbfs"
	StationInformation  Name = "station_information"
	StationStatus       Name = "station_status"
	SystemAlerts        Name = "system_alerts"
	SystemInformation   Name = "system_information"
	SystemPricingPlans  Name = "system_pricing_plans"
	SystemRegions       Name = "system_regions"
	VehicleAvailability Name = "vehicle_availability"
	VehicleStatus       Name = "vehicle_status"
	VehicleTypes        Name = "vehicle_types"
)

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
