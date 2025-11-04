macro_rules! file_struct {
    ( $ty: ident, $data: ty ) => {
        #[cfg_attr(feature = "napi", napi(object))]
        #[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
        #[serde_with::skip_serializing_none]
        #[derive(Serialize, Deserialize, Debug, Clone)]
        pub struct $ty {
            /// Indicates the last time data in the feed was updated. This timestamp represents the publisher's knowledge of the current state of the system at this point in time.
            pub last_updated: Timestamp,
            /// Number of seconds before the data in the feed will be updated again (0 if the data should always be refreshed).
            pub ttl: u32,
            /// GBFS version number to which the feed conforms, according to the versioning framework.
            pub version: String,
            /// Response data.
            pub data: $data,
        }
    };
}

pub mod gbfs;
pub mod gbfs_versions;
pub mod geofencing_zones;
pub mod manifest;
pub mod station_information;
pub mod station_status;
pub mod system_alerts;
pub mod system_information;
pub mod system_pricing_plans;
pub mod system_regions;
pub mod vehicle_status;
pub mod vehicle_types;

pub use self::gbfs::{GbfsData, GbfsFile};
pub use self::gbfs_versions::{GbfsVersionsData, GbfsVersionsFile};
pub use self::geofencing_zones::{GeofencingZonesData, GeofencingZonesFile};
pub use self::manifest::{ManifestData, ManifestFile};
pub use self::station_information::{StationInformationData, StationInformationFile};
pub use self::station_status::{StationStatusData, StationStatusFile};
pub use self::system_alerts::{SystemAlertsData, SystemAlertsFile};
pub use self::system_information::{SystemInformationData, SystemInformationFile};
pub use self::system_pricing_plans::{SystemPricingPlansData, SystemPricingPlansFile};
pub use self::system_regions::{SystemRegionsData, SystemRegionsFile};
pub use self::vehicle_status::{VehicleStatusData, VehicleStatusFile};
pub use self::vehicle_types::{VehicleTypesData, VehicleTypesFile};

#[cfg(test)]
mod tests {
    use pretty_assertions::assert_eq;

    fn test_file<T: serde::de::DeserializeOwned + serde::Serialize>(example: &str) {
        let json_value: serde_json::Value = serde_json::from_str(example).unwrap();

        let file: T = serde_json::from_str(example).unwrap();

        let json_value2: serde_json::Value =
            serde_json::from_str(&serde_json::to_string(&file).unwrap()).unwrap();

        assert_eq!(json_value, json_value2);
    }

    #[test]
    fn gbfs() {
        let gbfs = include_str!("./examples/specification/gbfs.json");

        test_file::<super::gbfs::GbfsFile>(gbfs);
    }

    #[test]
    fn manifest() {
        let manifest = include_str!("./examples/specification/manifest.json");

        test_file::<super::manifest::ManifestFile>(manifest);
    }

    #[test]
    fn gbfs_versions() {
        let gbfs_versions = include_str!("./examples/specification/gbfs_versions.json");

        test_file::<super::gbfs_versions::GbfsVersionsFile>(gbfs_versions);
    }

    #[test]
    fn system_information() {
        let system_information = include_str!("./examples/specification/system_information-1.json");

        test_file::<super::system_information::SystemInformationFile>(system_information);

        let system_information = include_str!("./examples/specification/system_information-2.json");

        test_file::<super::system_information::SystemInformationFile>(system_information);
    }

    #[test]
    fn vehicle_types() {
        let vehicle_types = include_str!("./examples/specification/vehicle_types.json");

        test_file::<super::vehicle_types::VehicleTypesFile>(vehicle_types);
    }

    #[test]
    fn station_information() {
        let station_information =
            include_str!("./examples/specification/station_information-1.json");

        test_file::<super::station_information::StationInformationFile>(station_information);

        let station_information =
            include_str!("./examples/specification/station_information-2.json");

        test_file::<super::station_information::StationInformationFile>(station_information);

        let station_information =
            include_str!("./examples/specification/station_information-3.json");

        test_file::<super::station_information::StationInformationFile>(station_information);
    }

    #[test]
    fn station_status() {
        let station_status = include_str!("./examples/specification/station_status.json");

        test_file::<super::station_status::StationStatusFile>(station_status);
    }

    #[test]
    fn vehicle_status() {
        let vehicle_status = include_str!("./examples/specification/vehicle_status-1.json");

        test_file::<super::vehicle_status::VehicleStatusFile>(vehicle_status);

        let vehicle_status = include_str!("./examples/specification/vehicle_status-2.json");

        test_file::<super::vehicle_status::VehicleStatusFile>(vehicle_status);
    }

    #[test]
    fn system_regions() {
        let system_regions = include_str!("./examples/specification/system_regions.json");

        test_file::<super::system_regions::SystemRegionsFile>(system_regions);
    }

    #[test]
    fn system_pricing_plans() {
        let system_pricing_plans =
            include_str!("./examples/specification/system_pricing_plans-1.json");

        test_file::<super::system_pricing_plans::SystemPricingPlansFile>(system_pricing_plans);

        let system_pricing_plans =
            include_str!("./examples/specification/system_pricing_plans-2.json");

        test_file::<super::system_pricing_plans::SystemPricingPlansFile>(system_pricing_plans);
    }

    #[test]
    fn system_alerts() {
        let system_alerts = include_str!("./examples/specification/system_alerts.json");

        test_file::<super::system_alerts::SystemAlertsFile>(system_alerts);
    }

    #[test]
    fn geofencing_zones() {
        let geofencing_zones = include_str!("./examples/specification/geofencing_zones.json");

        test_file::<super::geofencing_zones::GeofencingZonesFile>(geofencing_zones);
    }
}
