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
	"github.com/MobilityData/gbfs-json-schema/models/golang/v3.0/header"
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
	// If the feed license requires attribution, name of the organization to which attribution
	// should be provided. An array with one object per supported language with the following
	// keys:
	AttributionOrganizationName []common.LocalizedString `json:"attribution_organization_name,omitempty"`
	// URL of the organization to which attribution should be provided.
	AttributionURL *common.URL `json:"attribution_url,omitempty"`
	// An object where each key defines one of the items listed below (added in v2.3-RC).
	BrandAssets *common.BrandAssets `json:"brand_assets,omitempty"`
	// Email address actively monitored by the operator's customer service department.
	Email *common.Email `json:"email,omitempty"`
	// A single contact email address for consumers of this feed to report technical issues
	// (added in v1.1).
	FeedContactEmail common.Email `json:"feed_contact_email"`
	// List of languages used in translated strings. Each element in the list must be of type
	// Language.
	Languages []common.Language `json:"languages"`
	// REQUIRED if the dataset is provided under a standard license. An identifier for a
	// standard license from the SPDX License List. Provide license_id rather than license_url
	// if the license is included in the SPDX License List. See the GBFS wiki for a comparison
	// of a subset of standard licenses. If the license_id and license_url fields are blank or
	// omitted, this indicates that the feed is provided under the Creative Commons Universal
	// Public Domain Dedication.
	LicenseID *common.LicenseID `json:"license_id,omitempty"`
	// A fully qualified URL of a page that defines the license terms for the GBFS data for this
	// system.
	LicenseURL *common.URL `json:"license_url,omitempty"`
	// REQUIRED if the producer publishes datasets for more than one system geography, for
	// example Berlin and Paris. A fully qualified URL pointing to the manifest.json file for
	// the publisher.
	ManifestURL *common.URL `json:"manifest_url,omitempty"`
	// Name of the system to be displayed to customers. An array with one object per supported
	// language with the following keys:
	Name []common.LocalizedString `json:"name"`
	// Hours and dates of operation for the system in OSM opening_hours format. (added in v3.0)
	OpeningHours string `json:"opening_hours"`
	// Name of the system operator. An array with one object per supported language with the
	// following keys:
	Operator []common.LocalizedString `json:"operator,omitempty"`
	// This OPTIONAL field SHOULD contain a single voice telephone number for the specified
	// system’s customer service department. MUST be in E.164 format as defined in Field Types.
	PhoneNumber *common.PhoneNumber `json:"phone_number,omitempty"`
	// The date that the privacy policy provided at privacy_url was last updated (added in
	// v2.3-RC).
	PrivacyLastUpdated *common.Date `json:"privacy_last_updated,omitempty"`
	// A fully qualified URL pointing to the privacy policy for the service. An array with one
	// object per supported language with the following keys:
	PrivacyURL []common.LocalizedString `json:"privacy_url,omitempty"`
	// URL where a customer can purchase a membership.
	PurchaseURL *common.URL `json:"purchase_url,omitempty"`
	// Contains rental app information in the android and ios JSON objects (added in v1.1).
	RentalApps *common.RentalApps `json:"rental_apps,omitempty"`
	// Abbreviation for a system. An array with one object per supported language with the
	// following keys:
	ShortName []common.LocalizedString `json:"short_name,omitempty"`
	// Date that the system began operations.
	StartDate *common.Date `json:"start_date,omitempty"`
	// Identifier for this vehicle share system. This should be globally unique (even between
	// different systems).
	SystemID common.ID `json:"system_id"`
	// Date after which this data source will no longer be available to consuming applications.
	TerminationDate *common.Date `json:"termination_date,omitempty"`
	// The date that the terms of service provided at terms_url were last updated (added in v2.3-RC)
	TermsLastUpdated *common.Date `json:"terms_last_updated,omitempty"`
	// A fully qualified URL pointing to the terms of service (also often called "terms of use"
	// or "terms and conditions") for the service. An array with one object per supported
	// language with the following keys:
	TermsURL []common.LocalizedString `json:"terms_url,omitempty"`
	// The time zone where the system is located.
	Timezone common.Timezone `json:"timezone"`
	// The URL of the vehicle share system.
	URL *common.URL `json:"url,omitempty"`
}
