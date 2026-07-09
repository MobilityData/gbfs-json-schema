#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v2_3::types::*;

file_struct!(SystemHoursFile, SystemHoursData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemHoursData {
    pub rental_hours: Vec<RentalHours>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RentalHours {
    pub user_types: Option<Vec<String>>,
    pub days: Option<Vec<String>>,
    pub start_time: Option<Time>,
    pub end_time: Option<Time>,
}
