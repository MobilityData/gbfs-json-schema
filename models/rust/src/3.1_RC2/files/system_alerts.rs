#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(SystemAlertsFile, SystemAlertsData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemAlertsData {
    pub alerts: Vec<SystemAlert>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemAlert {
    /// Identifier for this alert.
    pub alert_id: AlertID,
    pub r#type: String,
    /// Indicating when the alert is in effect (for example, when the system or station is actually closed, or when a station is scheduled to be moved).
    pub times: Vec<AlertTime>,
    /// If this is an alert that affects one or more stations, include their ID(s). Otherwise omit this field.
    /// If both [station_ids][SystemAlert::station_ids] and [region_ids][SystemAlert::region_ids] are omitted, this alert affects the entire system.
    pub station_ids: Option<Vec<StationID>>,
    /// If this system has regions, and if this alert only affects certain regions, include their ID(s).
    pub region_ids: Option<Vec<RegionID>>,
    /// URL where the customer can learn more information about this alert.
    pub url: Option<Vec<LocalizedUrl>>,
    /// A short summary of this alert to be displayed to the customer.
    pub summary: Option<Vec<LocalizedString>>,
    /// Detailed description of the alert.
    pub description: Option<Vec<LocalizedString>>,
    /// Indicates the last time the info for the alert was updated.
    pub last_updated: Option<Timestamp>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AlertTime {
    /// Start time of the alert.
    pub start: Timestamp,
    /// End time of the alert. If there is currently no end time planned for the alert, this can be omitted.
    pub end: Option<Timestamp>,
}
