package common

// Country code following the ISO 3166-1 alpha-2 notation.
type CountryCode string

// A date in the ISO 8601 Complete Date Extended Format: YYYY-MM-DD. Example: 2019-09-13 for September 13th, 2019.
// JSON unmarshaling in Go does not allow the ISO 8601 Complete Date Extended Format to be represented as a time.Time type directly.
// Therefore, we use a String type for this purpose to ensure compatibility with the JSON format and avoid
// introducing additional complexity in the unmarshaling process.
type Date string

// Combination of a date and a time following ISO 8601 notation. Attributes : year, month, day, hour, minute, second, and timezone.
// Date and time fields MUST be represented as strings in the ISO 8601 format (YYYY-MM-DDThh:mm:ssZ).
// For example, 2023-10-01T12:00:00Z represents October 1st, 2023 at 12:00 PM UTC.
// The time zone is represented by the Z at the end of the string, indicating that the time is in UTC.
type DateTime string

// An email address. Example: example@example.com
type Email string

// ID - Should be represented as a string that identifies that particular entity. An ID:
//
//	MUST be unique within like fields (for example, station_id MUST be unique among stations)
//	Does not have to be globally unique, unless otherwise specified
//	MUST be in the ASCII printable character range, space excluded (0x21 to 0x7E) https://en.wikipedia.org/wiki/ASCII#Printable_characters (as of v3.0)
//	SHOULD be restricted to A-Z, a-z, 0-9 and .@:/_- (as of v3.0)
//	MUST be persistent for a given entity (station, plan, etc.). An exception is vehicle_id, which MUST NOT be persistent for privacy reasons (see vehicle_status.json). (as of v2.0)
type ID string

// An IETF BCP 47 language code. For an introduction to IETF BCP 47, refer to https://www.rfc-editor.org/rfc/bcp/bcp47.txt and https://www.w3.org/International/articles/language-tags/.
// Examples: en for English, en-US for American English, or de for German.
type Language string

// WGS84 latitude in decimal degrees. The value MUST be greater than or equal to -90.0 and less than or equal to 90.0. Example: 41.890169 for the Colosseum in Rome.
type Latitude float64

// WGS84 longitude in decimal degrees. The value MUST be greater than or equal to -180.0 and less than or equal to 180.0. Example: 12.492269 for the Colosseum in Rome.
type Longitude float64

// A JSON element representing a String value that has been translated into a specific language. (as of v3.0)
// The element consists of the following name-value pairs:
type LocalizedString struct {
	// The translated text.
	Text string `json:"text"`
	// IETF BCP 47 language code.
	Language Language `json:"language"`
}

// A JSON element representing the URL of a resource that has been translated into a specific language. (as of v3.0)
// The element consists of the following name-value pairs:
type LocalizedURL struct {
	// The URL of the resource.
	Text URL `json:"text"`
	// IETF BCP 47 language code.
	Language Language `json:"language"`
}

// Phone number in E.164 format. The phone number MUST start with a "+". The characters following the "+" MUST be integers and MUST NOT contain any hyphens, spaces or parentheses. (as of v3.0)
// For version 2.3 the same type is used despite not being defined in https://github.com/MobilityData/gbfs/blob/v2.3/gbfs.md#field-types although it is referenced in the spec (https://github.com/MobilityData/gbfs/blob/v2.3/gbfs.md#system_informationjson and https://github.com/MobilityData/gbfs/blob/v2.3/gbfs.md#station_informationjson)
// Also the type used in the examples for system_information.json and station_information.json is a string ensuring compatibility.
// See https://github.com/MobilityData/gbfs/issues/789 for more information.
type PhoneNumber string

// TZ timezone from the https://www.iana.org/time-zones. Timezone names never contain the space character but MAY contain an underscore.
// Refer to https://en.wikipedia.org/wiki/List_of_tz_zones for a list of valid values. Example: Asia/Tokyo, America/Los_Angeles or Africa/Cairo.
type Timezone string

// A fully qualified URI that includes the scheme (for example, com.example.android://). Any special characters in the URI MUST be correctly escaped.
// See the following https://www.w3.org/Addressing/URL/4_URI_Recommentations.html for a description of how to create fully qualified URI values. Note that URIs MAY be URLs.
type URI string

// A fully qualified URL that includes http:// or https://. Any special characters in the URL MUST be correctly escaped.
// See the following https://www.w3.org/Addressing/URL/4_URI_Recommentations.html for a description of how to create fully qualified URL values.
type URL string

// Timestamp fields MUST be represented as integers in POSIX time (representing the number of seconds since January 1st 1970 00:00:00 UTC). (before v3.0)
type PosixTimestamp int64

// Timestamp fields MUST be represented as strings in RFC3339 format, for example 2023-07-17T13:34:13+02:00. (as of v3.0)
type Timestamp string

// The semantic version of the feed in the form X.Y
type Version string

// Feed is a JSON element representing a feed in the GBFS specification.
type Feed struct {
	Name FeedName `json:"name"`
	URL  URL      `json:"url"`
}

// Key identifying the type of feed this is. The key MUST be the base file name defined in the spec for the corresponding feed type
// ( system_information for system_information.json file, station_information for station_information.json file).
type FeedName string

type VersionStruct struct {
	// GBFS version number to which the feed conforms, according to the versioning framework. (added in v1.1 -> therefore it is optional)
	Version *Version `json:"version,omitempty"`
}

type Header[T any] struct {
	LastUpdated T       `json:"last_updated"`
	Ttl         uint64  `json:"ttl"`
	Version     Version `json:"version"`
}

type VersionElement struct {
	// URL of the corresponding gbfs.json endpoint
	URL URL `json:"url"`
	// The semantic version of the feed in the form X.Y
	Version Version `json:"version"`
}

// Payment methods accepted at this station.
type RentalMethod string

// Vehicle equipment provided by the operator in addition to the accessories already
// provided in the vehicle (field vehicle_accessories of vehicle_types.json) but subject to more frequent updates.
type VehicleEquipment string

// Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
// in v1.1).
type RentalUris struct {
	// URI that can be passed to an Android app with an intent (added in v1.1).
	Android *URI `json:"android,omitempty"`
	// URI that can be used on iOS to launch the rental app for this station (added in v1.1).
	Ios *URI `json:"ios,omitempty"`
	// URL that can be used by a web browser to show more information about renting a vehicle at
	// this station (added in v1.1).
	Web *URI `json:"web,omitempty"`
}

// This indicates that this set of rental hours applies to either members or non-members only.
type UserType string

// Day of the week.
type Day string

// Type of parking station. Added in v2.3
type ParkingType string

// Type of system alert.
type Type string

// REQUIRED if the dataset is provided under a standard license. An identifier for a
// standard license from the SPDX License List. Provide license_id rather than license_url
// if the license is included in the SPDX License List. See the GBFS wiki for a comparison
// of a subset of standard licenses. If the license_id and license_url fields are blank or
// omitted, this indicates that the feed is provided under the Creative Commons Universal
// Public Domain Dedication.
type LicenseID string

// Contains rental app information in the android and ios JSON objects.
type RentalApps struct {
	Android *RentalApp `json:"android,omitempty"`
	IOS     *RentalApp `json:"ios,omitempty"`
}

// RentalApp contains the URIs for the store and discovery of the rental app.
type RentalApp struct {
	StoreUri     URI `json:"store_uri"`
	DiscoveryUri URI `json:"discovery_uri"`
}


// An object where each key defines one of the items listed below (added in v2.3-RC).
type BrandAssets struct {
	// A fully qualified URL pointing to the location of a graphic file representing the brand
	// for the service (added in v2.3-RC).
	BrandImageURL URL `json:"brand_image_url"`
	// A fully qualified URL pointing to the location of a graphic file representing the brand
	// for the service for use in dark mode (added in v2.3-RC).
	BrandImageURLDark *URL `json:"brand_image_url_dark,omitempty"`
	// Date that indicates the last time any included brand assets were updated (added in
	// v2.3-RC).
	BrandLastModified Date `json:"brand_last_modified"`
	// A fully qualified URL pointing to the location of a page that defines the license terms
	// of brand icons, colors or other trademark information (added in v2.3-RC).
	BrandTermsURL *URL `json:"brand_terms_url,omitempty"`
	// Color used to represent the brand for the service (added in v2.3-RC)
	Color *string `json:"color,omitempty"`
}

// Array of segments when the price is a function of distance traveled, displayed in kilometers.
// Total cost is the addition of price and all segments in per_km_pricing and per_min_pricing.
// If this array is not provided, there are no variable costs based on distance. (added in v2.2).
type PerKMPricing struct {
	// The kilometer at which the rate will no longer apply (exclusive) for example, if end is 20 the rate no longer applies at 20.00 km.
	// If this field is empty, the price issued for this segment is charged until the trip ends,
	// in addition to the cost of any subsequent segments. (added in v2.2).
	End *uint64 `json:"end,omitempty"`
	// REQUIRED if per_km_pricing is defined. Interval in kilometers at which the rate of this segment
	// is either reapplied indefinitely, or if defined, up until (but not including) end kilometer.
	// An interval of 0 indicates the rate is only charged once. (added in v2.2).
	Interval uint64 `json:"interval"`
	// REQUIRED if per_km_pricing is defined. Rate that is charged for each kilometer interval after the start.
	// Can be a negative number, which indicates that the traveler will receive a discount. (added in v2.2).
	Rate float64 `json:"rate"`
	// REQUIRED if per_km_pricing is defined. The kilometer at which this segment rate starts being charged (inclusive).
	// (added in v2.2).
	Start uint64 `json:"start"`
}

// Array of segments when the price is a function of time traveled, displayed in minutes.
// Total cost is the addition of price and all segments in per_km_pricing and per_min_pricing.
// If this array is not provided, there are no variable costs based on time. (added in v2.2).c
type PerMinPricing struct {
	// The minute at which the rate will no longer apply (exclusive) for example,
	// if end is 20 the rate no longer applies after 19:59. If this field is empty,
	// the price issued for this segment is charged until the trip ends,
	// in addition to the cost of any subsequent segments. (added in v2.2).
	End *uint64 `json:"end,omitempty"`
	// REQUIRED if per_min_pricing is defined. Interval in minutes at which the rate of this
	// segment is either reapplied indefinitely, or up until (but not including) the end minute,
	// if end is defined. An interval of 0 indicates the rate is only charged once. (added in v2.2).
	Interval uint64 `json:"interval"`
	// REQUIRED if per_min_pricing is defined. Rate that is charged for each minute interval after the start.
	// Can be a negative number, which indicates that the traveler will receive a discount. (added in v2.2).
	Rate float64 `json:"rate"`
	// REQUIRED if per_min_pricing is defined. The minute at which this segment rate starts being charged (inclusive). (added in v2.2).
	Start uint64 `json:"start"`
}

// Vehicle air quality certificate. Official anti-pollution certificate, based on the information on the vehicle's
// registration certificate, attesting to its level of pollutant emissions based on a defined standard. In Europe,
// for example, it is the European emission standard. The aim of this measure is to encourage the use of the least
// polluting vehicles by allowing them to drive during pollution peaks or in low emission zones.
type EcoLabel struct {
	// Country code following the ISO 3166-1 alpha-2 notation. Added in v2.3.
	CountryCode *CountryCode `json:"country_code,omitempty"`
	// Name of the eco label. Added in v2.3.
	EcoSticker *string `json:"eco_sticker,omitempty"`
}

// An object where each key defines one of the items listed below added in v2.3-RC.
type VehicleAssets struct {
	// Date that indicates the last time any included vehicle icon images were modified or
	// updated added in v2.3-RC.
	IconLastModified Date `json:"icon_last_modified"`
	// A fully qualified URL pointing to the location of a graphic icon file that MAY be used to
	// represent this vehicle type on maps and in other applications added in v2.3-RC.
	IconURL URL `json:"icon_url"`
	// A fully qualified URL pointing to the location of a graphic icon file to be used to
	// represent this vehicle type when in dark mode added in v2.3-RC.
	IconURLDark *URL `json:"icon_url_dark,omitempty"`
}

// The vehicle's general form factor.
type FormFactor string

// The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing
type PropulsionType string

// The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as
// return_type, and updated to return_constraint in v2.3.
type ReturnConstraint string

type VehicleAccessory string
