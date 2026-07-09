#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v2_3::types::*;

file_struct!(SystemCalendarFile, SystemCalendarData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemCalendarData {
    pub calendars: Vec<Calendar>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Calendar {
    pub start_month: Option<u8>,
    pub start_day: Option<u8>,
    pub start_year: Option<u16>,
    pub end_month: Option<u8>,
    pub end_day: Option<u8>,
    pub end_year: Option<u16>,
}
