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

package system_alerts

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Describes ad-hoc changes to the system.
type SystemAlerts struct {
	header.HeaderStruct
	// Array that contains ad-hoc alerts for the system.
	Data Data `json:"data"`
}

// Array that contains ad-hoc alerts for the system.
type Data struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	// Identifier for this alert.
	AlertID common.ID `json:"alert_id"`
	// Detailed description of the alert.
	Description *string `json:"description,omitempty"`
	// Indicates the last time the info for the alert was updated.
	LastUpdated *common.PosixTimestamp `json:"last_updated,omitempty"`
	// If this system has regions, and if this alert only affects certain regions, include their ID(s). Otherwise, omit this field.
	// If both station_ids and region_ids are omitted, this alert affects the entire system.
	RegionIDS []common.ID `json:"region_ids,omitempty"`
	// If this is an alert that affects one or more stations, include their ID(s). Otherwise omit this field.
	// If both station_id and region_id are omitted, this alert affects the entire system.
	StationIDS []common.ID `json:"station_ids,omitempty"`
	// A short summary of this alert to be displayed to the customer.
	Summary string `json:"summary"`
	// Array of objects with the fields start and end indicating when the alert is in effect
	// (for example, when the system or station is actually closed, or when a station is scheduled to be moved).
	Times []Time `json:"times,omitempty"`
	// Type of alert.
	Type common.Type `json:"type"`
	// URL where the customer can learn more information about this alert.
	URL *common.URL `json:"url,omitempty"`
}

type Time struct {
	// End time of the alert. If there is currently no end time planned for the alert, this can be omitted.
	End *common.PosixTimestamp `json:"end,omitempty"`
	// REQUIRED if times array is defined. Start time of the alert.
	Start *common.PosixTimestamp `json:"start,omitempty"`
}
