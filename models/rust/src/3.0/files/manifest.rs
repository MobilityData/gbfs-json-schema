#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use crate::v3_0::types::*;

use serde::{Deserialize, Serialize};

use crate::v3_0::files::gbfs_versions::GbfsVersion;

file_struct!(ManifestFile, ManifestData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ManifestData {
    pub datasets: Vec<System>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct System {
    /// The `system_id` from `system_information.json` for the corresponding data set(s).
    pub system_id: SystemID,
    /// Contains one object, as defined below, for each of the available versions of a feed.
    /// The array MUST be sorted by increasing MAJOR and MINOR version number.
    pub versions: Vec<GbfsVersion>,
}
