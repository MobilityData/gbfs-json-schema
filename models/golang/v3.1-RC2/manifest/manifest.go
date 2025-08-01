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

package manifest

import (
	manifest_v31RC "github.com/MobilityData/gbfs-json-schema/models/golang/v3.1-RC/manifest"
)

// An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
// instance of this file should be published at a single stable URL, for example:
// https://example.com/gbfs/manifest.json.
type Manifest struct {
	manifest_v31RC.Manifest
}
