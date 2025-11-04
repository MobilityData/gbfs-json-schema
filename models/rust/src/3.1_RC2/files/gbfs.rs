#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use futures::try_join;

use super::super::GbfsObjects;

use super::*;
use crate::v3_0::types::*;
use crate::v3_0::urls::*;

file_struct!(GbfsFile, GbfsData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsData {
    /// All of the feeds that are published by this auto-discovery file.
    pub feeds: Vec<GbfsDataFeeds>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsDataFeeds {
    /// Key identifying the type of feed this is. The key MUST be the base file name defined in the spec for the corresponding feed type
    pub name: FeedType,
    /// URL for the feed. Note that the actual feed endpoints (urls) may not be defined in the `file_name.json` format.
    /// For example, a valid feed endpoint could end with `station_info` instead of `station_information.json`.
    pub url: URL,
}

impl GbfsData {
    fn get_feed_url(&self, name: &str) -> Option<String> {
        self.feeds
            .iter()
            .find(|feed| feed.name == name)
            .map(|feed| feed.url.clone())
    }

    pub async fn get_objects(&self) -> Result<GbfsObjects, reqwest::Error> {
        let vehicle_status = self.get_vehicle_status();
        let station_information = self.get_station_information();
        let station_status = self.get_station_status();
        let system_pricing_plans = self.get_system_pricing_plans();
        let vehicle_types = self.get_vehicle_types();

        let (
            vehicle_status,
            station_information,
            station_status,
            system_pricing_plans,
            vehicle_types,
        ) = try_join!(
            vehicle_status,
            station_information,
            station_status,
            system_pricing_plans,
            vehicle_types
        )?;

        let vehicle_status = vehicle_status.map(|f| f.data);
        let station_information = station_information.map(|f| f.data);
        let station_status = station_status.map(|f| f.data);
        let system_pricing_plans = system_pricing_plans.map(|f| f.data);
        let vehicle_types = vehicle_types.map(|f| f.data);

        Ok(GbfsObjects {
            vehicle_status,
            station_information,
            station_status,
            system_pricing_plans,
            vehicle_types,
        })
    }
}

macro_rules! add_file_type {
    ($fn_url_name: ident, $url_ty: ident, $fn_name: ident, $ty: ident, $name: expr) => {
        impl GbfsData {
            pub fn $fn_url_name(&self) -> Option<$url_ty> {
                let url = self.get_feed_url($name)?;

                Some($url_ty::new(&url))
            }

            pub async fn $fn_name(&self) -> Result<Option<$ty>, reqwest::Error> {
                let url = self.$fn_url_name();

                if let Some(url) = url {
                    Ok(Some(url.fetch_async().await?))
                } else {
                    Ok(None)
                }
            }
        }
    };
}

add_file_type!(
    get_gbfs_versions_url,
    GbfsVersionsFileUrl,
    get_gbfs_versions,
    GbfsVersionsFile,
    "gbfs_versions"
);

add_file_type!(
    get_geofencing_zones_url,
    GeofencingZonesFileUrl,
    get_geofencing_zones,
    GeofencingZonesFile,
    "geofencing_zones"
);

add_file_type!(
    get_station_information_url,
    StationInformationFileUrl,
    get_station_information,
    StationInformationFile,
    "station_information"
);

add_file_type!(
    get_station_status_url,
    StationStatusFileUrl,
    get_station_status,
    StationStatusFile,
    "station_status"
);

add_file_type!(
    get_system_alerts_url,
    SystemAlertsFileUrl,
    get_system_alerts,
    SystemAlertsFile,
    "system_alerts"
);

add_file_type!(
    get_system_information_url,
    SystemInformationFileUrl,
    get_system_information,
    SystemInformationFile,
    "system_information"
);

add_file_type!(
    get_system_pricing_plans_url,
    SystemPricingPlansFileUrl,
    get_system_pricing_plans,
    SystemPricingPlansFile,
    "system_pricing_plans"
);

add_file_type!(
    get_system_regions_url,
    SystemRegionsFileUrl,
    get_system_regions,
    SystemRegionsFile,
    "system_regions"
);

add_file_type!(
    get_vehicle_status_url,
    VehicleStatusFileUrl,
    get_vehicle_status,
    VehicleStatusFile,
    "vehicle_status"
);

add_file_type!(
    get_vehicle_types_url,
    VehicleTypesFileUrl,
    get_vehicle_types,
    VehicleTypesFile,
    "vehicle_types"
);
