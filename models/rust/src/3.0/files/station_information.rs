#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(StationInformationFile, StationInformationData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per station in the system as defined below.
pub struct StationInformationData {
    pub stations: Vec<Station>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Station {
    /// Identifier of a station.
    pub station_id: StationID,
    /// The public name of the station for display in maps, digital signage, and other text applications.
    /// Names SHOULD reflect the station location through the use of a cross street or local landmark.
    /// Abbreviations SHOULD NOT be used for names and other text (for example, "St." for "Street") unless a location is called by its abbreviated name (for example, “JFK Airport”).
    ///
    /// Examples:
    ///
    /// - Broadway and East 22nd Street
    /// - Convention Center
    /// - Central Park South
    pub name: Option<Vec<LocalizedString>>,
    /// Short name or other type of identifier.
    pub short_name: Option<Vec<LocalizedString>>,
    /// Latitude of the station in decimal degrees.
    /// This field SHOULD have a precision of 6 decimal places (0.000001).
    pub lat: Latitude,
    /// Longitude of the station in decimal degrees.
    /// This field SHOULD have a precision of 6 decimal places (0.000001).
    pub lon: Longitude,
    /// Address (street number and name) where station is located.
    /// This MUST be a valid address, not a free-form text description.
    ///
    /// Example: 1234 Main Street
    pub address: Option<String>,
    /// Cross street or landmark where the station is located.
    pub cross_street: Option<String>,
    /// Identifier of the region where station is located.
    pub region_id: Option<RegionID>,
    /// Postal code where station is located.
    pub post_code: Option<String>,
    /// Hours of operation for the station.
    /// If `station_opening_hours` is defined it overrides any `opening_hours` defined in `system_information.json` for the station for which it is defined.
    pub station_opening_hours: Option<OSMOpeningHours>,
    /// Payment methods accepted at this station.
    pub rental_methods: Option<Vec<String>>,
    /// Is this station a location with or without smart dock technology?
    ///
    /// - `true`: The station is a location without smart docking infrastructure. the station may be defined by a point (lat/lon) and/or station_area (below).
    /// - `false`: The station consists of smart docking infrastructure (docks).
    pub is_virtual_station: Option<bool>,
    /// A GeoJSON MultiPolygon that describes the area of a virtual station.
    ///
    /// If `station_area` is supplied, then the record describes a virtual station.
    ///
    /// If `lat`/`lon` and `station_area` are both defined, the `lat`/`lon` is the significant coordinate of the station (for example, parking facility or valet drop-off and pick up point).
    /// The station_area takes precedence over any `ride_start_allowed` and `ride_end_allowed` rules in overlapping `geofencing_zones`.
    pub station_area: Option<geo_json::MultiPolygon>,
    /// Type of parking station.
    pub parking_type: Option<String>,
    /// Are parking hoops present at this station?
    pub parking_hoop: Option<bool>,
    /// Contact phone of the station.
    pub contact_phone: Option<PhoneNumber>,
    /// Number of total docking points installed at this station, both available and unavailable, regardless of what vehicle types are allowed at each dock.
    pub capacity: Option<u32>,
    /// Used to model the parking capacity of virtual stations (defined using the is_virtual_station field) for each vehicle type that can be returned to this station.
    /// The total number of vehicles from each of these objects SHOULD add up to match the value specified in the `capacity` field.
    pub vehicle_types_capacity: Option<Vec<VehicleTypesCapacity>>,
    /// These objects are used to model the total docking capacity of a station, both available and unavailable, for each type of vehicle that may dock at this station.
    /// The total number of docks from each of these objects SHOULD add up to match the value specified in the `capacity` field.
    pub vehicle_docks_capacity: Option<Vec<VehicleDocksCapacity>>,
    /// Are valet services provided at this station?
    pub is_valet_station: Option<bool>,
    /// Does the station support charging of electric vehicles?
    pub is_charging_station: Option<bool>,
    /// Contains rental URIs for Android, iOS, and web
    pub rental_uris: Option<RentalUris>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VehicleTypesCapacity {
    /// Vehicles types that may park at the virtual station.
    pub vehicle_type_ids: Vec<VehicleTypeID>,
    /// If the virtual station is defined by [station_area](Station::station_area), this is the number that can park within the station area.
    ///
    /// If [lat](Station::lat)/[lon](Station::lon) is defined, this is the number that can park at those coordinates.
    pub count: u32,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VehicleDocksCapacity {
    /// Vehicles types that able to use a particular type of dock at the station.
    pub vehicle_type_ids: Vec<VehicleTypeID>,
    //  number representing the total number of docks at the station, both available and unavailable, that may accept that vehicle types
    pub count: u32,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RentalUris {
    /// URI that can be passed to an Android app with an `android.intent.action.VIEW` Android intent to support Android Deep Links (<https://developer.android.com/training/app-links/deep-linking>).
    pub android: Option<URI>,
    /// URI that can be used on iOS to launch the rental app for this station.
    pub ios: Option<URI>,
    /// URL that can be used by a web browser to show more information about renting a vehicle at this station.
    pub web: Option<URL>,
}
