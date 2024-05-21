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
//    systemRegions, err := UnmarshalSystemRegions(bytes)
//    bytes, err = systemRegions.Marshal()

package system_regions



import "encoding/json"

func UnmarshalSystemRegions(data []byte) (SystemRegions, error) {
	var r SystemRegions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SystemRegions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes regions for a system that is broken up by geographic or political region.
type SystemRegions struct {
	// Describe regions for a system that is broken up by geographic or political region.                 
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

// Describe regions for a system that is broken up by geographic or political region.
type Data struct {
	// Array of regions.         
	Regions             []Region `json:"regions"`
}

type Region struct {
	// Public name for this region.       
	Name                           []Name `json:"name"`
	// identifier of the region.          
	RegionID                       string `json:"region_id"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Version string

const (
	The30 Version = "3.0"
)
