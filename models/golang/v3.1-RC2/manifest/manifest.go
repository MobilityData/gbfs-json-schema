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
//    manifest, err := UnmarshalManifest(bytes)
//    bytes, err = manifest.Marshal()

package manifest



import "encoding/json"

func UnmarshalManifest(data []byte) (Manifest, error) {
	var r Manifest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Manifest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
// instance of this file should be published at a single stable URL, for example:
// https://example.com/gbfs/manifest.json.
type Manifest struct {
	Data                                                                                        Data            `json:"data"`
	// Last time the data in the feed was updated in RFC3339 format.                                            
	LastUpdated                                                                                 string       `json:"last_updated"`
	// Number of seconds before the data in the feed will be updated again (0 if the data should                
	// always be refreshed).                                                                                    
	TTL                                                                                         int64           `json:"ttl"`
	// GBFS version number to which the feed conforms, according to the versioning framework                    
	// (added in v1.1).                                                                                         
	Version                                                                                     ManifestVersion `json:"version"`
}

type Data struct {
	// An array of objects, each containing the keys below.          
	Datasets                                               []Dataset `json:"datasets"`
}

type Dataset struct {
	// A GeoJSON MultiPolygon that describes the operating area.                                                
	Area                                                                                       *Area            `json:"area,omitempty"`
	// The ISO 3166-1 alpha-2 country code of the operating area.                                               
	CountryCode                                                                                *string          `json:"country_code,omitempty"`
	// The system_id from system_information.json for the corresponding data set(s).                            
	SystemID                                                                                   string           `json:"system_id"`
	// Contains one object, as defined below, for each of the available versions of a feed. The                 
	// array MUST be sorted by increasing MAJOR and MINOR version number.                                       
	Versions                                                                                   []VersionElement `json:"versions"`
}

// A GeoJSON MultiPolygon that describes the operating area.
type Area struct {
	Coordinates [][][][]float64 `json:"coordinates"`
	Type        Type            `json:"type"`
}

type VersionElement struct {
	// URL of the corresponding gbfs.json endpoint                    
	URL                                                string         `json:"url"`
	// The semantic version of the feed in the form X.Y               
	Version                                            VersionVersion `json:"version"`
}

type Type string

const (
	MultiPolygon Type = "MultiPolygon"
)

// The semantic version of the feed in the form X.Y
type VersionVersion string

const (
	Purple31RC2 VersionVersion = "3.1-RC2"
	The10       VersionVersion = "1.0"
	The11       VersionVersion = "1.1"
	The20       VersionVersion = "2.0"
	The21       VersionVersion = "2.1"
	The22       VersionVersion = "2.2"
	The23       VersionVersion = "2.3"
	The30       VersionVersion = "3.0"
)

type ManifestVersion string

const (
	Fluffy31RC2 ManifestVersion = "3.1-RC2"
)
