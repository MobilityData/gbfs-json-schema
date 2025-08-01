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

package gbfsversions

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Lists all feed endpoints published according to version sof the GBFS documentation. (added in v1.1)
type GbfsVersions struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// Contains one object, as defined below, for each of the available versions of a feed. The
	// array must be sorted by increasing MAJOR and MINOR version number.
	Versions []VersionElement `json:"versions"`
}

type VersionElement struct {
	// URL of the corresponding gbfs.json endpoint
	URL string `json:"url"`
	// The semantic version of the feed in the form X.Y
	Version common.Version `json:"version"`
}
