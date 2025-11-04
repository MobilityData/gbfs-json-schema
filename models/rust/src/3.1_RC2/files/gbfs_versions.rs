#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;
use crate::v3_0::urls::*;

file_struct!(GbfsVersionsFile, GbfsVersionsData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsVersionsData {
    /// Contains one object, as defined below, for each of the available versions of a feed.
    /// The array MUST be sorted by increasing MAJOR and MINOR version number.
    pub versions: Vec<GbfsVersion>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsVersion {
    /// The semantic version of the feed in the form `X.Y`.
    pub version: String,
    /// URL of the corresponding `gbfs.json` endpoint.
    pub url: GbfsFileUrl,
}
