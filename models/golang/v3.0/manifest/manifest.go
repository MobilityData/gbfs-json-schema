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

package manifest

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
)

// An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
// instance of this file should be published at a single stable URL, for example:
// https://example.com/gbfs/manifest.json.
type Manifest struct {
	header.HeaderStruct
	Data Data `json:"data"`
}

type Data struct {
	// An array of objects, each containing the keys below.
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	// The system_id from system_information.json for the corresponding data set(s).
	SystemID common.ID `json:"system_id"`
	// Contains one object, as defined below, for each of the available versions of a feed. The
	// array MUST be sorted by increasing MAJOR and MINOR version number.
	Versions []common.VersionElement `json:"versions"`
}
