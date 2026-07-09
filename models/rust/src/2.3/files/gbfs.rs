#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use futures::try_join;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;

use super::super::GbfsObjects;

use super::*;
use crate::v2_3::types::*;
use crate::v2_3::urls::*;

file_struct!(GbfsFile, GbfsData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsData {
    #[serde(flatten)]
    pub feeds: HashMap<Language, GbfsLanguageFeeds>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsLanguageFeeds {
    pub feeds: Vec<GbfsDataFeed>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct GbfsDataFeed {
    /// Key identifying the type of feed this is. The key MUST be the base file name defined in the spec for the corresponding feed type
    pub name: FeedType,
    /// URL for the feed. Note that the actual feed endpoints (urls) may not be defined in the `file_name.json` format.
    /// For example, a valid feed endpoint could end with `station_info` instead of `station_information.json`.
    pub url: URL,
}

impl GbfsData {
    fn get_feed_url(&self, name: &str, language: &str) -> Option<String> {
        self.feeds
            .get(language)?
            .feeds
            .iter()
            .find(|feed| feed.name == name)
            .map(|feed| feed.url.clone())
    }

    pub async fn get_objects(&self, language: &str) -> Result<GbfsObjects, reqwest::Error> {
        let free_bike_status = self.get_free_bike_status(language);
        let station_information = self.get_station_information(language);
        let station_status = self.get_station_status(language);
        let system_pricing_plans = self.get_system_pricing_plans(language);
        let vehicle_types = self.get_vehicle_types(language);

        let (
            free_bike_status,
            station_information,
            station_status,
            system_pricing_plans,
            vehicle_types,
        ) = try_join!(
            free_bike_status,
            station_information,
            station_status,
            system_pricing_plans,
            vehicle_types
        )
        .unwrap();

        let free_bike_status = free_bike_status.map(|f| f.data);
        let station_information = station_information.map(|f| f.data);
        let station_status = station_status.map(|f| f.data);
        let system_pricing_plans = system_pricing_plans.map(|f| f.data);
        let vehicle_types = vehicle_types.map(|f| f.data);

        Ok(GbfsObjects {
            free_bike_status,
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
            pub fn $fn_url_name(&self, language: &str) -> Option<$url_ty> {
                let url = self.get_feed_url($name, language)?;

                Some($url_ty::new(&url))
            }

            pub async fn $fn_name(&self, language: &str) -> Result<Option<$ty>, reqwest::Error> {
                let url = self.$fn_url_name(language);

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
    get_system_alerts_url,
    SystemAlertsFileUrl,
    get_system_alerts,
    SystemAlertsFile,
    "system_alerts"
);

add_file_type!(
    get_station_status_url,
    StationStatusFileUrl,
    get_station_status,
    StationStatusFile,
    "station_status"
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
    get_system_hours_url,
    SystemHoursFileUrl,
    get_system_hours,
    SystemHoursFile,
    "system_hours"
);

add_file_type!(
    get_system_calendar_url,
    SystemCalendarFileUrl,
    get_system_calendar,
    SystemCalendarFile,
    "system_calendar"
);

add_file_type!(
    get_system_regions_url,
    SystemRegionsFileUrl,
    get_system_regions,
    SystemRegionsFile,
    "system_regions"
);

add_file_type!(
    get_free_bike_status_url,
    FreeBikeStatusFileUrl,
    get_free_bike_status,
    FreeBikeStatusFile,
    "free_bike_status"
);

add_file_type!(
    get_vehicle_types_url,
    VehicleTypesFileUrl,
    get_vehicle_types,
    VehicleTypesFile,
    "vehicle_types"
);
