#![allow(non_snake_case)]

#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(VehicleTypesFile, VehicleTypesData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per vehicle that is currently deployed in the field.
pub struct VehicleTypesData {
    pub vehicle_types: Vec<VehicleType>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Vehicle that is currently deployed in the field
pub struct VehicleType {
    /// Unique identifier of a vehicle type.
    pub vehicle_type_id: VehicleTypeID,
    /// The vehicle's general form factor.
    pub form_factor: Option<String>,
    /// The number of riders (driver included) the vehicle can legally accommodate.
    pub rider_capacity: Option<u8>,
    /// Cargo volume available in the vehicle, expressed in liters. For cars, it corresponds to the space between the boot floor, including the storage under the hatch, to the rear shelf in the trunk.
    pub cargo_volume_capacity: Option<i32>,
    /// The capacity of the vehicle cargo space (excluding passengers), expressed in kilograms.
    pub cargo_load_capacity: Option<f64>,
    /// The primary propulsion type of the vehicle.
    pub propulsion_type: Option<String>,
    /// Vehicle air quality certificate. Official anti-pollution certificate, based on the information on the vehicle's registration certificate, attesting to its level of pollutant emissions based on a defined standard. In Europe, for example, it is the European emission standard. The aim of this measure is to encourage the use of the least polluting vehicles by allowing them to drive during pollution peaks or in low emission zones.
    pub eco_labels: Option<Vec<EcoLabel>>,
    /// This represents the furthest distance in meters that the vehicle can travel without recharging or refueling when it has the maximum amount of energy potential (for example, a full battery or full tank of gas).
    pub max_range_meters: Option<f64>,
    /// The public name of this vehicle type.
    pub name: Option<Vec<LocalizedString>>,
    /// Description of accessories available in the vehicle. These accessories are part of the vehicle and are not supposed to change frequently.
    pub vehicle_accessories: Option<Vec<String>>,
    /// Maximum quantity of CO2, in grams, emitted per kilometer, according to the [WLTP](https://en.wikipedia.org/wiki/Worldwide_Harmonised_Light_Vehicles_Test_Procedure).
    pub g_CO2_km: Option<i64>,
    /// URL to an image that would assist the user in identifying the vehicle (for example, an image of the vehicle or a logo).
    pub vehicle_image: Option<String>,
    /// The name of the vehicle manufacturer.
    pub make: Option<Vec<LocalizedString>>,
    /// The name of the vehicle model.
    pub model: Option<Vec<LocalizedString>>,
    /// The color of the vehicle.
    pub color: Option<String>,
    /// Customer-readable description of the vehicle type outlining special features or how-tos.
    pub description: Option<Vec<LocalizedString>>,
    /// Number of wheels this vehicle type has.
    pub wheel_count: Option<u8>,
    /// The maximum speed in kilometers per hour this vehicle is permitted to reach in accordance with local permit and regulations.
    pub max_permitted_speed: Option<u8>,
    /// The rated power of the motor for this vehicle type in watts.
    pub rated_power: Option<u32>,
    /// Maximum time in minutes that a vehicle can be reserved before a rental begins.
    /// If default_reserve_time is set to 0, the vehicle type cannot be reserved.
    pub default_reserve_time: Option<u32>,
    /// The conditions for returning the vehicle at the end of the rental.
    pub return_constraint: Option<String>,
    pub vehicle_assets: Option<VehicleAsset>,
    /// A plan_id, as defined in system_pricing_plans.json, that identifies a default pricing plan for this vehicle to be used by trip planning applications for purposes of calculating the cost of a single trip using this vehicle type.
    /// This default pricing plan is superseded by `pricing_plan_id` when `pricing_plan_id` is defined in `vehicle_status.json`.
    pub default_pricing_plan_id: Option<PricingPlanID>,
    /// All pricing plan IDs that are applied to this vehicle type.
    pub pricing_plan_ids: Option<Vec<PricingPlanID>>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct EcoLabel {
    /// Country where the `eco_sticker` applies.
    pub country_code: CountryCode,
    /// Name of the eco label.
    pub eco_sticker: String,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VehicleAsset {
    /// A fully qualified URL pointing to the location of a graphic icon file that MAY be used to represent this vehicle type on maps and in other applications.
    pub icon_url: String,
    /// A fully qualified URL pointing to the location of a graphic icon file to be used to represent this vehicle type when in dark mode on maps and in other applications.
    pub icon_url_dark: Option<String>,
    /// Date that indicates the last time any included vehicle icon images were modified or updated.
    pub icon_last_modified: Option<Date>,
}
