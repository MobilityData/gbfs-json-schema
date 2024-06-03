#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

file_struct!(GeofencingZonesFile, GeofencingZonesData);

use crate::v3_0::types::*;

#[cfg(doc)]
use crate::v3_0::files::geofencing_zones;
#[cfg(doc)]
use crate::v3_0::files::station_information;

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GeofencingZonesData {
    pub geofencing_zones: GeofencingZone,
    /// restrictions that apply globally in all areas as the default restrictions, except where overridden with an explicit geofencing zone
    pub global_rules: Option<Vec<Rule>>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// geofenced zone and its associated rules and attributes
pub struct GeofencingZone {
    /// `FeatureCollection` (as per IETF RFC 7946).
    pub r#type: String,
    pub features: Vec<Feature>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Feature {
    /// `Feature` (as per IETF RFC 7946).
    pub r#type: String,
    /// A polygon that describes where rides may or may not be able to start, end, go through, or have other limitations or affordances.
    /// Rules may only apply to the interior of a polygon.
    /// All geofencing zones contained in this list are public (meaning they can be displayed on a map for public use).
    pub geometry: geo_json::MultiPolygon,
    /// Properties describing travel allowances and limitations.
    pub properties: Properties,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Properties {
    /// Public name of the geofencing zone.
    pub name: Vec<LocalizedString>,
    /// Start time of the geofencing zone. If the geofencing zone is always active, this can be omitted.
    pub start: Option<Timestamp>,
    /// End time of the geofencing zone. If the geofencing zone is always active, this can be omitted.
    pub end: Option<Timestamp>,
    /// Restrictions that apply within the area of the polygon.
    pub rules: Option<Vec<Rule>>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Geofencing Rule Precedence
///
/// Geofencing Rule objects are specified within arrays for the rules and [global_rules](GeofencingZonesData::global_rules) fields of [geofencing_zones] to allow for different restrictions for different vehicle types.
/// When multiple rules in the same array apply to a particular vehicle type, per the semantics of the vehicle_type_ids field, then the earlier rule (in order of the JSON file) takes precedence for that vehicle type.
///
/// When multiple overlapping polygons define rules that apply to a particular vehicle type,
/// then the rules from the earlier polygon (in order of the JSON file) takes precedence for that vehicle type in the overlapping area.
/// Polygons with inactive time ranges should be excluded from consideration when considering precedence.
///
/// When a polygon and the [global_rules](GeofencingZonesData::global_rules) field define rules that apply to a particular vehicle type, then the rules from the polygon take precedence for that vehicle type in the area of the polygon.
pub struct Rule {
    /// Array of IDs of vehicle types for which any restrictions SHOULD be applied.
    /// If vehicle type IDs are not specified, then restrictions apply to all vehicle types.
    pub vehicle_type_ids: Option<Vec<VehicleTypeID>>,
    /// Is the ride allowed to start in this zone?
    pub ride_start_allowed: bool,
    /// Is the ride allowed to end in this zone?
    pub ride_end_allowed: bool,
    /// Is the ride allowed to travel through this zone?
    pub ride_through_allowed: bool,
    /// What is the maximum speed allowed, in kilometers per hour?
    ///
    /// If there is no maximum speed to observe, this can be omitted.
    pub maximum_speed_kph: Option<i32>,
    /// Can vehicles only be parked at stations defined in [station_information] within this geofence zone?
    pub station_parking: Option<bool>,
}
