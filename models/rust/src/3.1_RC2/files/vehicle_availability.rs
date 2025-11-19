#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_1_rc2::types::*;

file_struct!(VehicleAvailabilityFile, VehicleAvailabilityData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per vehicle.
pub struct VehicleAvailabilityData {
    /// Describes the future availability of each vehicle. Useful for systems that allow vehicles to be reserved in advance (e.g. carsharing, cargo bike share, etc).
    pub vehicles: Vec<Vehicle>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Vehicle {
    /// Identifier of a vehicle.
    pub vehicle_id: VehicleID,
    /// Unique identifier of a vehicle type as defined in vehicle_types.json.
    pub vehicle_type_id: VehicleTypeID,
    /// The station_id of the station where this vehicle is located when available as defined in station_information.json.
    pub station_id: StationID,
    /// The plan_id of the pricing plan this vehicle is eligible for as described in system_pricing_plans.json. If this field is defined it supersedes default_pricing_plan_id in vehicle_types.json.
    pub pricing_plan_id: Option<PricingPlanID>,
    /// List of vehicle equipment provided by the operator in addition to the accessories already provided in the vehicle (field vehicle_accessories of vehicle_types.json) but subject to more frequent updates.
    pub vehicle_equipment: Option<Vec<String>>,
    /// Array of time slots during which the specified vehicle is available.
    pub availabilities: Vec<Availability>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Availability {
    /// Start date and time of available time slot.
    pub from: Datetime,
    /// End date and time of available time slot. If this field is empty, it means that the vehicle is available all the time from the date in the `from` field.
    pub until: Option<Datetime>,
}
