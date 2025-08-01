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

package station_status

import (
	station_status_v30 "github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/station_status"
)

// Describes the capacity and rental availability of the station
type StationStatus struct {
	station_status_v30.StationStatus
}
