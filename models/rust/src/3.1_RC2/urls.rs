#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};
use url;

use super::files::*;

macro_rules! file_url {
    ( $ty: ident, $body: ident ) => {
        #[cfg_attr(feature = "napi", napi(object))]
        #[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
        #[derive(Debug, Clone)]
        pub struct $ty {
            pub value: String,
        }

        impl $ty {
            pub fn new(value: &str) -> Self {
                Self {
                    value: value.to_string(),
                }
            }

            pub fn get(&self) -> url::Url {
                self.value.parse().unwrap()
            }

            #[cfg(feature = "reqwest_blocking")]
            pub fn fetch(&self) -> Result<$body, reqwest::Error> {
                reqwest::blocking::get(self.get())?.json()
            }

            pub async fn fetch_async(&self) -> Result<$body, reqwest::Error> {
                reqwest::get(self.get()).await?.json().await
            }
        }

        impl From<String> for $ty {
            fn from(value: String) -> Self {
                $ty { value }
            }
        }

        impl<'de> Deserialize<'de> for $ty {
            fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                let val: String = Deserialize::deserialize(deserializer)?;

                val.try_into().map_err(serde::de::Error::custom)
            }
        }

        impl Serialize for $ty {
            fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
            where
                S: serde::Serializer,
            {
                self.value.serialize(serializer)
            }
        }
    };
}

file_url!(ManifestFileUrl, ManifestFile);
file_url!(GbfsFileUrl, GbfsFile);

file_url!(GbfsVersionsFileUrl, GbfsVersionsFile);
file_url!(GeofencingZonesFileUrl, GeofencingZonesFile);
file_url!(StationInformationFileUrl, StationInformationFile);
file_url!(StationStatusFileUrl, StationStatusFile);
file_url!(SystemAlertsFileUrl, SystemAlertsFile);
file_url!(SystemInformationFileUrl, SystemInformationFile);
file_url!(SystemPricingPlansFileUrl, SystemPricingPlansFile);
file_url!(SystemRegionsFileUrl, SystemRegionsFile);
file_url!(VehicleStatusFileUrl, VehicleStatusFile);
file_url!(VehicleTypesFileUrl, VehicleTypesFile);
