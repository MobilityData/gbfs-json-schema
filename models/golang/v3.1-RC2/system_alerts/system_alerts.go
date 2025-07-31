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
//    systemAlerts, err := UnmarshalSystemAlerts(bytes)
//    bytes, err = systemAlerts.Marshal()

package system_alerts



import "encoding/json"

func UnmarshalSystemAlerts(data []byte) (SystemAlerts, error) {
	var r SystemAlerts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SystemAlerts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes ad-hoc changes to the system.
type SystemAlerts struct {
	// Array that contains ad-hoc alerts for the system.                                                  
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

// Array that contains ad-hoc alerts for the system.
type Data struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	// Identifier for this alert.                                                       
	AlertID                                                               string        `json:"alert_id"`
	// Detailed description of the alert.                                               
	Description                                                           []Description `json:"description,omitempty"`
	// Indicates the last time the info for the alert was updated.                      
	LastUpdated                                                           *string    `json:"last_updated,omitempty"`
	// Array of identifiers of the regions for which this alert applies.                
	RegionIDS                                                             []string      `json:"region_ids,omitempty"`
	// Array of identifiers of the stations for which this alert applies.               
	StationIDS                                                            []string      `json:"station_ids,omitempty"`
	// A short summary of this alert to be displayed to the customer.                   
	Summary                                                               []Summary     `json:"summary"`
	// Array of objects indicating when the alert is in effect.                         
	Times                                                                 []Time        `json:"times,omitempty"`
	// Type of alert.                                                                   
	Type                                                                  Type          `json:"type"`
	// URL where the customer can learn more information about this alert.              
	URL                                                                   []URL         `json:"url,omitempty"`
}

type Description struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Summary struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Time struct {
	// End time of the alert.             
	End                        *string `json:"end,omitempty"`
	// Start time of the alert.           
	Start                      *string `json:"start,omitempty"`
}

type URL struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

// Type of alert.
type Type string

const (
	Other          Type = "other"
	StationClosure Type = "station_closure"
	StationMove    Type = "station_move"
	SystemClosure  Type = "system_closure"
)

type Version string

const (
	The31Rc2 Version = "3.1-RC2"
)
