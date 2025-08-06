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

package systemhours

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes the system hours of operation.
type SystemHours struct {
	header.HeaderStruct
	// Array that contains system hours of operations.
	Data Data `json:"data"`
}

type Data struct {
	RentalHours []RentalHour `json:"rental_hours"`
}

type RentalHour struct {
	// Array of member and nonmember value(s) indicating that this set of rental hours applies to either members or non-members only.
	UserTypes []common.UserType `json:"user_types"`
	// An array of abbreviations (first 3 letters) of English names of the days of the week for which this object applies.
	Days []common.Day `json:"days"`
	// Start time for the hours of operation of the system.
	StartTime string `json:"start_time"`
	// End time for the hours of operation of the system.
	EndTime string `json:"end_time"`
}
