pub mod files;
pub mod types;
pub mod urls;

#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsObjects {
    pub vehicle_status: Option<files::vehicle_status::VehicleStatusData>,
    pub station_status: Option<files::station_status::StationStatusData>,
    pub station_information: Option<files::station_information::StationInformationData>,
    pub vehicle_types: Option<files::vehicle_types::VehicleTypesData>,
    pub system_pricing_plans: Option<files::system_pricing_plans::SystemPricingPlansData>,
}
