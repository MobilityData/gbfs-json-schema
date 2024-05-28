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
//    geofencingZones, err := UnmarshalGeofencingZones(bytes)
//    bytes, err = geofencingZones.Marshal()

package geofencing_zones



import "encoding/json"

func UnmarshalGeofencingZones(data []byte) (GeofencingZones, error) {
	var r GeofencingZones
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GeofencingZones) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Describes geofencing zones and their associated rules and attributes (added in v2.1-RC).
type GeofencingZones struct {
	// Array that contains geofencing information for the system.                                         
	Data                                                                                        Data      `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                      
	LastUpdated                                                                                 string `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should          
	// always be refreshed).                                                                              
	TTL                                                                                         int64     `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework.             
	Version                                                                                     Version   `json:"version"`
}

// Array that contains geofencing information for the system.
type Data struct {
	// Each geofenced zone and its associated rules and attributes is described as an object                     
	// within the array of features.                                                                             
	GeofencingZones                                                                         GeofencingZonesClass `json:"geofencing_zones"`
	// Array of Rule objects defining restrictions that apply globally in all areas as the                       
	// default restrictions, except where overridden with an explicit geofencing zone.                           
	GlobalRules                                                                             []GlobalRule         `json:"global_rules"`
}

// Each geofenced zone and its associated rules and attributes is described as an object
// within the array of features.
type GeofencingZonesClass struct {
	// Array of objects.                                          
	Features                                  []GeoJSONFeature    `json:"features"`
	// FeatureCollection as per IETF RFC 7946.                    
	Type                                      GeofencingZonesType `json:"type"`
}

type GeoJSONFeature struct {
	// A polygon that describes where rides might not be able to start, end, go through, or have                    
	// other limitations. Must follow the right-hand rule.                                                          
	Geometry                                                                                    GeoJSONMultiPolygon `json:"geometry"`
	// Describing travel allowances and limitations.                                                                
	Properties                                                                                  Properties          `json:"properties"`
	Type                                                                                        FeatureType         `json:"type"`
}

// A polygon that describes where rides might not be able to start, end, go through, or have
// other limitations. Must follow the right-hand rule.
type GeoJSONMultiPolygon struct {
	Coordinates [][][][]float64 `json:"coordinates"`
	Type        GeometryType    `json:"type"`
}

// Describing travel allowances and limitations.
type Properties struct {
	// End time of the geofencing zone in RFC3339 format.                                               
	End                                                                                      *string `json:"end,omitempty"`
	// Public name of the geofencing zone.                                                              
	Name                                                                                     []Name     `json:"name,omitempty"`
	// Array of Rule objects defining restrictions that apply within the area of the polygon.           
	Rules                                                                                    []Rule     `json:"rules,omitempty"`
	// Start time of the geofencing zone in RFC3339 format.                                             
	Start                                                                                    *string `json:"start,omitempty"`
}

type Name struct {
	// IETF BCP 47 language code.       
	Language                     string `json:"language"`
	// The translated text.             
	Text                         string `json:"text"`
}

type Rule struct {
	// What is the maximum speed allowed, in kilometers per hour?                                 
	MaximumSpeedKph                                                                      *int64   `json:"maximum_speed_kph,omitempty"`
	// Is the ride allowed to end in this zone?                                                   
	RideEndAllowed                                                                       bool     `json:"ride_end_allowed"`
	// Is the ride allowed to start in this zone?                                                 
	RideStartAllowed                                                                     bool     `json:"ride_start_allowed"`
	// Is the ride allowed to travel through this zone?                                           
	RideThroughAllowed                                                                   bool     `json:"ride_through_allowed"`
	// Vehicle MUST be parked at stations defined in station_information.json within this         
	// geofence zone                                                                              
	StationParking                                                                       *bool    `json:"station_parking,omitempty"`
	// Array of vehicle type IDs for which these restrictions apply.                              
	VehicleTypeIDS                                                                       []string `json:"vehicle_type_ids,omitempty"`
}

type GlobalRule struct {
	// What is the maximum speed allowed, in kilometers per hour?                                 
	MaximumSpeedKph                                                                      *int64   `json:"maximum_speed_kph,omitempty"`
	// Is the ride allowed to end in this zone?                                                   
	RideEndAllowed                                                                       bool     `json:"ride_end_allowed"`
	// Is the ride allowed to start in this zone?                                                 
	RideStartAllowed                                                                     bool     `json:"ride_start_allowed"`
	// Is the ride allowed to travel through this zone?                                           
	RideThroughAllowed                                                                   bool     `json:"ride_through_allowed"`
	// Vehicle MUST be parked at stations defined in station_information.json within this         
	// geofence zone                                                                              
	StationParking                                                                       *bool    `json:"station_parking,omitempty"`
	// Array of vehicle type IDs for which these restrictions apply.                              
	VehicleTypeIDS                                                                       []string `json:"vehicle_type_ids,omitempty"`
}

type GeometryType string

const (
	MultiPolygon GeometryType = "MultiPolygon"
)

type FeatureType string

const (
	Feature FeatureType = "Feature"
)

// FeatureCollection as per IETF RFC 7946.
type GeofencingZonesType string

const (
	FeatureCollection GeofencingZonesType = "FeatureCollection"
)

type Version string

const (
	The30 Version = "3.0"
)
