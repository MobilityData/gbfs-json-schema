#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(VehicleStatusFile, VehicleStatusData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per vehicle that is currently deployed in the field.
pub struct VehicleStatusData {
    /// Array that contains one object per vehicle that is currently deployed in the field and not part of an active rental.
    pub vehicles: Vec<Vehicle>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Vehicle that is currently deployed in the field
pub struct Vehicle {
    /// Identifier of a vehicle.
    pub vehicle_id: VehicleID,
    /// Latitude of the vehicle in decimal degrees.
    pub lat: Option<Latitude>,
    /// Longitude of the vehicle in decimal degrees.
    pub lon: Option<Longitude>,
    /// Is the vehicle currently reserved?
    pub is_reserved: bool,
    /// Is the vehicle currently disabled?
    pub is_disabled: bool,
    /// Rental URIs for Android, iOS, and web
    pub rental_uris: Option<RentalUris>,
    /// The [VehicleType](crate::v3_0::files::vehicle_types::VehicleType) of this vehicle.
    pub vehicle_type_id: Option<VehicleTypeID>,
    /// The last time this vehicle reported its status to the operator's backend.
    pub last_reported: Option<Timestamp>,
    /// This value represents the furthest distance in meters that the vehicle can travel with the vehicle's current charge or fuel (without recharging or refueling).
    /// Note that in the case of carsharing, the given range is indicative and can be different from the one displayed on the vehicle's dashboard.
    pub current_range_meters: Option<NonNegativeFloat>,
    /// This value represents the current percentage, expressed from 0 to 1, of fuel or battery power remaining in the vehicle.
    pub current_fuel_percent: Option<NonNegativeFloat>,
    /// If the vehicle is currently at a [Station](crate::v3_0::files::station_information::Station).
    pub station_id: Option<StationID>,
    /// The [Station](crate::v3_0::files::station_information::Station) where this vehicle must be returned to.
    pub home_station_id: Option<StationID>,
    /// The [PricingPlan](crate::v3_0::files::system_pricing_plans::SystemPricingPlan) this vehicle is eligible for.
    ///
    /// If this field is defined it supersedes [default_pricing_plan_id](crate::v3_0::files::system_information::SystemInformation::default_pricing_plan_id).
    pub pricing_plan_id: Option<PricingPlanID>,
    /// List of vehicle equipment provided by the operator in addition to the accessories already provided in the vehicle (field vehicle_accessories of vehicle_types.json) but subject to more frequent updates.
    pub vehicle_equipment: Option<Vec<String>>,
    /// The date and time when any rental of the vehicle must be completed.
    /// The vehicle must be returned and made available for the next user by this time.
    /// If this field is empty, it indicates that the vehicle is available indefinitely.
    pub available_until: Option<Datetime>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RentalUris {
    /// URI that can be passed to an Android app with an `android.intent.action.VIEW` Android intent to support Android Deep Links (https://developer.android.com/training/app-links/deep-linking).
    pub android: Option<URI>,
    /// URI that can be used on iOS to launch the rental app for this vehicle.
    pub ios: Option<URI>,
    /// URL that can be used by a web browser to show more information about renting a vehicle at this vehicle.
    pub web: Option<URL>,
}
