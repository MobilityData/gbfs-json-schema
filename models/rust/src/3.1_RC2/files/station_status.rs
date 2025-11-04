#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(StationStatusFile, StationStatusData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per station in the system as defined below.
pub struct StationStatusData {
    pub stations: Vec<Station>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Station {
    /// Identifier of a station.
    pub station_id: StationID,
    /// Number of functional vehicles physically at the station that may be offered for rental. To know if the vehicles are available for rental, see [is_renting](Station::is_renting).
    pub num_vehicles_available: u32,
    /// Used to model the total number of each defined vehicle type available at a station.
    /// The total number of vehicles from each of these objects SHOULD add up to match the value specified in the [num_vehicles_available](Station::num_vehicles_available) field.
    pub vehicle_types_available: Option<Vec<VehicleTypeAvailable>>,
    /// Number of disabled vehicles of any type at the station.
    pub num_vehicles_disabled: Option<u32>,
    /// Number of functional docks physically at the station that are able to accept vehicles for return. To know if the docks are accepting vehicle returns, see [is_returning](Station::is_returning).
    pub num_docks_available: Option<u32>,
    /// Used to model the number of docks available for certain vehicle types.
    pub vehicle_docks_available: Option<Vec<VehicleDockAvailable>>,
    /// Number of disabled dock points at the station.
    pub num_docks_disabled: Option<u32>,
    /// Is the station currently on the street?
    ///
    /// In seasonal systems where equipment is removed during winter, boolean SHOULD be set to `false` during the off season.
    /// May also be set to `false` to indicate planned (future) stations which have not yet been installed.
    pub is_installed: bool,
    /// Is the station currently renting vehicles?
    ///
    /// If the station is temporarily taken out of service and not allowing rentals, this field MUST be set to `false`.
    ///
    /// If a station becomes inaccessible to users due to road construction or other factors this field SHOULD be set to `false`.
    /// Field SHOULD be set to `false` during hours or days when the system is not offering vehicles for rent.
    pub is_renting: bool,
    /// Is the station accepting vehicle returns?
    ///
    /// If the station is temporarily taken out of service and not allowing vehicle returns, this field MUST be set to `false`.
    /// If a station becomes inaccessible to users due to road construction or other factors, this field SHOULD be set to `false`.
    pub is_returning: bool,
    /// The last time this station reported its status to the operator's backend.
    pub last_reported: Option<Timestamp>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VehicleTypeAvailable {
    pub vehicle_type_id: VehicleTypeID,
    /// A number representing the total number of available vehicles of this type at this station.
    pub count: u32,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VehicleDockAvailable {
    pub vehicle_type_ids: Vec<VehicleTypeID>,
    /// A number representing the number of docks available for those vehicle types at this station.
    pub count: u32,
}
