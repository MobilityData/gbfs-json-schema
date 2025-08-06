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

package system_regions

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes regions for a system that is broken up by geographic or political region.
type SystemRegions struct {
	header.HeaderStruct
	// Describe regions for a system that is broken up by geographic or political region.
	Data Data `json:"data"`
}

// Describe regions for a system that is broken up by geographic or political region.
type Data struct {
	// Array of regions.
	Regions []Region `json:"regions"`
}

type Region struct {
	// Public name for this region.
	Name string `json:"name"`
	// identifier of the region.
	RegionID common.ID `json:"region_id"`
}
