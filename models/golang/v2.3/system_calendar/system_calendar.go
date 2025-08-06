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

package systemcalendar

import "github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"

// Describes the operating calendar for a system.
type SystemCalendar struct {
	header.HeaderStruct
	// Array that contains operations calendar for the system.
	Data Data `json:"data"`
}

type Data struct {
	Calendars []Calendar `json:"calendars"`
}

type Calendar struct {
	// Starting month for the system operations.
	StartMonth uint64 `json:"start_month"`
	// Starting day for the system operations.
	StartDay uint64 `json:"start_day"`
	// Starting year for the system operations.
	StartYear *uint64 `json:"start_year,omitempty"`
	// End month for the system operations.
	EndMonth uint64 `json:"end_month"`
	// End day for the system operations.
	EndDay uint64 `json:"end_day"`
	// End year for the system operations.
	EndYear *uint64 `json:"end_year,omitempty"`
}
