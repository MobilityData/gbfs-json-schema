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

package system_information

import (
	"github.com/MobilityData/gbfs-json-schema/models/golang/common"
	"github.com/MobilityData/gbfs-json-schema/models/golang/v2.3/header"
)

// Details including system operator, system location, year implemented, URL, contact info,
// time zone.
type SystemInformation struct {
	header.HeaderStruct
	// Response data in the form of name:value pairs.
	Data Data `json:"data"`
}

// Response data in the form of name:value pairs.
type Data struct {
	// An object where each key defines one of the items listed below (added in v2.3-RC).
	BrandAssets *common.BrandAssets `json:"brand_assets,omitempty"`
	// This OPTIONAL field SHOULD contain a single contact email address actively monitored by the operator’s
	// customer service department. This email address SHOULD be a direct contact point where riders can reach
	// a customer service representative.
	Email *common.Email `json:"email,omitempty"`
	// This OPTIONAL field SHOULD contain a single contact email for feed consumers to report technical issues with the feed.
	// (added in v1.1).
	FeedContactEmail *common.Email `json:"feed_contact_email,omitempty"`
	// The language that will be used throughout the rest of the files. It must match the value in the gbfs.json file.
	Language common.Language `json:"language"`
	// A fully qualified URL of a page that defines the license terms for the GBFS data for this system,
	// as well as any other license terms the system would like to define (including the use of corporate trademarks, etc)
	LicenseURL *common.URL `json:"license_url,omitempty"`
	// Name of the system to be displayed to customers.
	Name string `json:"name"`
	// Hours and dates of operation for the system in OSM opening_hours format. (added in v3.0)
	OpeningHours string `json:"opening_hours"`
	// Name of the operator.
	Operator *string `json:"operator,omitempty"`
	// This OPTIONAL field SHOULD contain a single voice telephone number for the specified system’s customer service department.
	// It can and SHOULD contain punctuation marks to group the digits of the number. Dialable text
	// (for example, Capital Bikeshare’s "877-430-BIKE") is permitted, but the field MUST NOT contain any other descriptive text.
	PhoneNumber *common.PhoneNumber `json:"phone_number,omitempty"`
	// The date that the privacy policy provided at privacy_url was last updated (added in
	// v2.3-RC).
	PrivacyLastUpdated *string `json:"privacy_last_updated,omitempty"`
	// A fully qualified URL pointing to the privacy policy for the service (added in v2.3-RC)
	PrivacyURL *common.URL `json:"privacy_url,omitempty"`
	// URL where a customer can purchase a membership.
	PurchaseURL *common.URL `json:"purchase_url,omitempty"`
	// Contains rental app information in the android and ios JSON objects (added in v1.1).
	RentalApps *common.RentalApps `json:"rental_apps,omitempty"`
	// Optional abbreviation for a system.
	ShortName string `json:"short_name,omitempty"`
	// Date that the system began operations.
	StartDate *common.Date `json:"start_date,omitempty"`
	// This is a globally unique identifier for the vehicle share system. It is up to the publisher of the feed
	// to guarantee uniqueness and MUST be checked against existing system_id fields in systems.csv to ensure this.
	// This value is intended to remain the same over the life of the system.Each distinct system or geographic area
	// in which vehicles are operated SHOULD have its own system_id. System IDs SHOULD be recognizable as belonging
	// to a particular system - for example, bcycle_austin or biketown_pdx - as opposed to random strings.
	SystemID common.ID `json:"system_id"`
	// The date that the terms of service provided at terms_url were last updated (added in
	// v2.3-RC)
	TermsLastUpdated *common.Date `json:"terms_last_updated,omitempty"`
	// A fully qualified URL pointing to the terms of service (added in v2.3-RC)
	TermsURL *common.URL `json:"terms_url,omitempty"`
	// The time zone where the system is located.
	Timezone common.Timezone `json:"timezone"`
	// The URL of the vehicle share system.
	URL *common.URL `json:"url,omitempty"`
}
