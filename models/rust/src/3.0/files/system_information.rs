#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;
use crate::v3_0::urls::*;

file_struct!(SystemInformationFile, SystemInformationData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemInformationData {
    /// This is a globally unique identifier for the vehicle share system.
    /// Each distinct system or geographic area in which vehicles are operated MUST have its own unique system_id.
    /// It is up to the publisher of the feed to guarantee uniqueness and MUST be checked against existing system_id fields in `systems.csv` to ensure this.
    /// This value is intended to remain the same over the life of the system.
    ///
    /// System IDs SHOULD be recognizable as belonging to a particular system as opposed to random strings - for example, bcycle_austin or biketown_pdx.
    pub system_id: SystemID,
    /// List of languages used in translated strings.
    pub languages: Option<Vec<Language>>,
    /// Name of the system to be displayed to customers.
    pub name: Option<Vec<LocalizedString>>,
    /// Hours and dates of operation for the system in [OSM opening_hours](https://wiki.openstreetmap.org/wiki/Key:opening_hours) format.
    pub opening_hours: Option<OSMOpeningHours>,
    /// Abbreviation for a system.
    pub short_name: Option<Vec<LocalizedString>>,
    /// Name of the system operator.
    pub operator: Option<Vec<LocalizedString>>,
    /// The URL of the vehicle share system.
    pub url: Option<URL>,
    /// URL where a customer can purchase a membership.
    pub purchase_url: Option<URL>,
    /// Date that the system began operations.
    pub start_date: Option<Date>,
    /// Date after which this data source will no longer be available to consuming applications.
    ///
    /// This OPTIONAL field SHOULD be used to notify 3rd party data consumers when a service is planning a permanent (non-seasonal) shutdown.
    /// Publishers SHOULD include this date in their feeds as soon as they know of an impending shutdown.
    /// Publishers SHOULD continue to publish feeds for 30 days following a permanent shutdown after which they SHOULD return a helpful http status code and text, for example `410 service no longer available ...` for all feed endpoint URLs.
    pub termination_date: Option<Date>,
    /// This OPTIONAL field SHOULD contain a single voice telephone number for the specified system’s customer service department.
    pub phone_number: Option<PhoneNumber>,
    /// This OPTIONAL field SHOULD contain a single contact email address actively monitored by the operator’s customer service department.
    /// This email address SHOULD be a direct contact point where riders can reach a customer service representative.
    pub email: Option<Email>,
    /// This field MUST contain a single contact email for feed consumers to report issues with the feed.
    /// This email address SHOULD point to a stable email address, that does not correspond to an individual but rather the team or company that manages GBFS feeds.
    pub feed_contact_email: Option<Email>,
    // REQUIRED if the producer publishes datasets for more than one system geography, for example Berlin and Paris. A fully qualified URL pointing to the manifest.json file for the publisher.
    pub manifest_url: Option<ManifestFileUrl>,
    /// The time zone where the system is located.
    pub timezone: Option<Timezone>,
    /// An identifier for a standard license from the SPDX License List. Provide `license_id` rather than `license_url` if the license is included in the SPDX License List.
    pub license_id: Option<LicenseID>,
    /// A fully qualified URL of a page that defines the license terms for the GBFS data for this system.
    pub license_url: Option<URL>,
    /// If the feed license requires attribution, name of the organization to which attribution should be provided.
    pub attribution_organization_name: Option<Vec<LocalizedString>>,
    /// URL of the organization to which attribution should be provided.
    pub attribution_url: Option<URL>,
    pub brand_assets: Option<BrandAssets>,
    /// A fully qualified URL pointing to the terms of service (also often called "terms of use" or "terms and conditions") for the service.
    pub terms_url: Option<Vec<LocalizedUrl>>,
    /// The date that the terms of service provided at `terms_url` were last updated.
    pub terms_last_updated: Option<Date>,
    /// A fully qualified URL pointing to the privacy policy for the service.
    pub privacy_url: Option<Vec<LocalizedUrl>>,
    /// The date that the privacy policy provided at `privacy_url` was last updated.
    pub privacy_last_updated: Option<Date>,
    /// Contains rental app information for android and ios.
    pub rental_apps: Option<RentalApps>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct BrandAssets {
    /// Date that indicates the last time any included brand assets were updated or modified.
    pub brand_last_modified: Option<Date>,
    /// A fully qualified URL pointing to the location of a page that defines the license terms of brand icons, colors, or other trademark information.
    /// This field MUST NOT take the place of license_url or license_id.
    pub brand_terms_url: Option<URL>,
    /// A fully qualified URL pointing to the location of a graphic file representing the brand for the service.
    /// File MUST be in SVG V1.1 format and MUST be either square or round.
    pub brand_image_url: URL,
    /// A fully qualified URL pointing to the location of a graphic file representing the brand for the service for use in dark mode applications.
    /// File MUST be in SVG V1.1 format and MUST be either square or round.
    pub brand_image_url_dark: Option<URL>,
    /// Color used to represent the brand for the service expressed as a 6 digit hexadecimal color code in the form #000000.
    pub color: Option<Color>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RentalApps {
    /// Contains rental app download and app discovery information for the Android platform in the `store_uri` and `discovery_uri` fields.
    pub android: Option<AndroidRentalApp>,
    /// Contains rental information for the iOS platform in the store_uri and discovery_uri fields.
    pub ios: Option<IosRentalApp>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AndroidRentalApp {
    /// URI where the rental Android app can be downloaded from. Typically this will be a URI to an app store, such as Google Play.
    /// If the URI points to an app store, the URI SHOULD follow Android best practices so the viewing app can directly open
    ///
    /// Example value: `https://play.google.com/store/apps/details?id=com.example.android`
    pub store_uri: URI,
    /// URI that can be used to discover if the rental Android app is installed on the device (for example, using `PackageManager.queryIntentActivities()`).
    /// This intent is used by viewing apps to prioritize rental apps for a particular user based on whether they already have a particular rental app installed.
    ///
    /// Example value: `com.example.android://`
    pub discovery_uri: URI,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct IosRentalApp {
    /// URI where the rental iOS app can be downloaded from.
    /// Typically this will be a URI to an app store, such as the Apple App Store.
    /// If the URI points to an app store, the URI SHOULD follow iOS best practices so the viewing app can directly open the URI to the native app store app instead of a website.
    ///
    /// Example value: `https://apps.apple.com/app/apple-store/id123456789`
    pub store_uri: URI,
    /// URI that can be used to discover if the rental iOS app is installed on the device (for example, using `UIApplication canOpenURL:`).
    /// This intent is used by viewing apps to prioritize rental apps for a particular user based on whether they already have a particular rental app installed.
    ///
    /// Example value: `com.example.ios://`
    pub discovery_uri: URI,
}
