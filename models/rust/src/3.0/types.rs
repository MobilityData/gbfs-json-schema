use serde::{Deserialize, Serialize};

#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

#[cfg(doc)]
use crate::v3_0::files::gbfs::GbfsFile;
#[cfg(doc)]
use crate::v3_0::files::gbfs_versions::GbfsVersionsFile;
#[cfg(doc)]
use crate::v3_0::files::geofencing_zones::GeofencingZonesFile;
#[cfg(doc)]
use crate::v3_0::files::station_information::StationInformationFile;
#[cfg(doc)]
use crate::v3_0::files::station_status::StationStatusFile;
#[cfg(doc)]
use crate::v3_0::files::system_alerts::SystemAlertsFile;
#[cfg(doc)]
use crate::v3_0::files::system_information::SystemInformationFile;
#[cfg(doc)]
use crate::v3_0::files::system_pricing_plans::SystemPricingPlansFile;
#[cfg(doc)]
use crate::v3_0::files::system_regions::SystemRegionsFile;
#[cfg(doc)]
use crate::v3_0::files::vehicle_status::VehicleStatusFile;
#[cfg(doc)]
use crate::v3_0::files::vehicle_types::VehicleTypesFile;

#[cfg(doc)]
use crate::v3_0::files::manifest::System;
#[cfg(doc)]
use crate::v3_0::files::station_information::Station;
#[cfg(doc)]
use crate::v3_0::files::system_alerts::SystemAlert;
#[cfg(doc)]
use crate::v3_0::files::system_information::SystemInformationData;
#[cfg(doc)]
use crate::v3_0::files::system_pricing_plans::SystemPricingPlan;
#[cfg(doc)]
use crate::v3_0::files::system_regions::Region;
#[cfg(doc)]
use crate::v3_0::files::vehicle_status::Vehicle;
#[cfg(doc)]
use crate::v3_0::files::vehicle_types::VehicleType;

/// Should be represented as a string that identifies that particular entity. An ID:
/// - MUST be unique within like fields (for example, `station_id` MUST be unique among stations).
/// - Does not have to be globally unique, unless otherwise specified.
/// - MUST be in the ASCII printable character range, space excluded (`0x21` to `0x7E`).
/// - SHOULD be restricted to `A-Z`, `a-z`, `0-9` and `.@:/_-`
/// - MUST be persistent for a given entity (station, plan, etc). An exception is floating bike bike_id, which MUST NOT be persistent for privacy reasons (see free_bike_status.json). (as of v2.0)
pub type ID = String;

// A date in the ISO 8601 Complete Date Extended Format: YYYY-MM-DD . Example: `2019-09-13` for September 13th, 2019.
pub type Date = String;

// Datetime (added in v2.3)- Combination of a date and a time following ISO 8601 notation. Attributes : year, month, day, hour, minute, second, and timezone.
pub type Datetime = String;

/// Service time in the HH:MM:SS format for the time zone indicated in `system_information.json` (00:00:00 - 47:59:59). Time can stretch up to one additional day in the future to accommodate situations where, for example, a system was open from 11:30pm - 11pm the next day (23:30:00-47:00:00).
pub type Time = String;

/// An email address. Example: `example@example.com`
pub type Email = String;

/// Phone number in E.164 format. The phone number MUST start with a "+". The characters following the "+" MUST be integers and MUST NOT contain any hyphens, spaces or parentheses.
pub type PhoneNumber = String;

/// A fully qualified URL that includes `http://` or `https://`. Any special characters in the URL MUST be correctly escaped. See the following <https://www.w3.org/Addressing/URL/4_URI_Recommentations.html> for a description of how to create fully qualified URL values.
pub type URL = String; // url::URL;

/// A fully qualified URI that includes the scheme (for example, com.example.android://). Any special characters in the URI MUST be correctly escaped. See the following <https://www.w3.org/Addressing/URL/4_URI_Recommentations.html> for a description of how to create fully qualified URI values. Note that URIs MAY be URLs.
pub type URI = String; // http::Uri;

/// Country code following the ISO 3166-1 alpha-2 notation.
pub type CountryCode = String;

/// An IETF BCP 47 language code. For an introduction to IETF BCP 47, refer to <https://www.rfc-editor.org/rfc/bcp/bcp47.txt> and <https://www.w3.org/International/articles/language-tags/>. Examples: `en` for English, `en-US`
pub type Language = String;

///  WGS84 latitude in decimal degrees. The value MUST be greater than or equal to -90.0 and less than or equal to 90.0. Example: `41.890169` for the Colosseum in Rome.
pub type Latitude = f64;

/// WGS84 longitude in decimal degrees. The value MUST be greater than or equal to -180.0 and less than or equal to 180.0. Example: `12.492269` for the Colosseum in Rome.
pub type Longitude = f64;

// A Geometry Object as described by the IETF RFC <https://tools.ietf.org/html/rfc7946#section-3.1.7>.
//pub type MultiPolygon = geo_types::geometry::MultiPolygon;

/// Timestamp fields MUST be represented as strings in RFC3339 format, for example 2023-07-17T13:34:13+02:00.
pub type Timestamp = String;

/// TZ timezone from the <https://www.iana.org/time-zones>. Timezone names never contain the space character but MAY contain an underscore. Refer to <https://en.wikipedia.org/wiki/List_of_tz_zones> for a list of valid values. Example: `Asia/Tokyo`, `America/Los_Angeles` or `Africa/Cairo`.
pub type Timezone = String;

/// Float (added in v2.1) - A 32-bit floating point number.
pub type Float = f64;

/// A 32-bit floating point number greater than or equal to 0.
pub type NonNegativeFloat = f64;

/// An integer greater than or equal to 0.
pub type NonNegativeInteger = u32;

/// An identifier for a standard license from the [SPDX License List](https://spdx.org/licenses/).
pub type LicenseID = String;

/// A 6 digit hexadecimal color code in the form #000000.
pub type Color = String;

/// ISO 4217 code: <https://en.wikipedia.org/wiki/ISO_4217> (for example, CAD for Canadian dollars, EUR for euros, or JPY for Japanese yen.)
pub type Currency = String;

/// Unique identifier of a vehicle type.
/// View [VehicleType] for more information.
pub type VehicleTypeID = ID;
/// Unique identifier of a vehicle.
/// View [Vehicle] for more information.
pub type VehicleID = ID;
/// Unique identifier of a station.
/// View [Station] for more information.
pub type StationID = ID;
/// Unique identifier of a region.
/// View [Region] for more information.
pub type RegionID = ID;
/// Unique identifier of a system.
/// View [SystemInformationData] and [System] for more information.
pub type SystemID = ID;
/// Unique identifier of a pricing plan.
/// View [SystemPricingPlan] for more information.
pub type PricingPlanID = ID;
/// Unique identifier of a system alert.
/// View [SystemAlert] for more information.
pub type AlertID = ID;

/// Type of a GBFS feed.
/// Current values are :
/// - `gbfs` for [GbfsFile],
/// - `gbfs_versions` for [GbfsVersionsFile],
/// - `system_information` for [SystemInformationFile],
/// - `vehicle_types` for [VehicleTypesFile],
/// - `station_information` for [StationInformationFile],
/// - `station_status` for [StationStatusFile],
/// - `vehicle_status` for [VehicleStatusFile],
/// - `system_regions` for [SystemRegionsFile],
/// - `system_pricing_plans` for [SystemPricingPlansFile],
/// - `system_alerts` for [SystemAlertsFile],
/// - `geofencing_zones` for [GeofencingZonesFile]
pub type FeedType = String;

/// Opening hours in the [OSM opening_hours](https://wiki.openstreetmap.org/wiki/Key:opening_hours).
pub type OSMOpeningHours = String;

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocalizedString {
    /// The translated text.
    pub text: String,
    /// The language code.
    pub language: Language,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocalizedUrl {
    pub text: URL,
    /// The language code.
    pub language: Language,
}

pub mod geo_json {
    use serde::{Deserialize, Serialize};

    #[cfg(feature = "pyo3")]
    use pyo3::prelude::*;

    #[cfg(feature = "napi")]
    use napi_derive::napi;

    pub mod coordinates {
        pub type Point = Vec<f64>;
        pub type LineString = Vec<Point>;
        pub type Polygon = Vec<LineString>;
        pub type MultiPolygon = Vec<Polygon>;
    }

    #[cfg_attr(feature = "napi", napi(object))]
    #[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
    #[serde_with::skip_serializing_none]
    #[derive(Serialize, Deserialize, Debug, Clone)]
    pub struct MultiPolygon {
        pub r#type: String,
        pub coordinates: coordinates::MultiPolygon,
    }
}
