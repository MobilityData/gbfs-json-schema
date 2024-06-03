#[cfg(feature = "pyo3")]
use pyo3::prelude::*;

#[cfg(feature = "napi")]
use napi_derive::napi;

use serde::{Deserialize, Serialize};

use crate::v3_0::types::*;

file_struct!(SystemPricingPlansFile, SystemPricingPlansData);

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
/// Contains one object per vehicle that is currently deployed in the field.
pub struct SystemPricingPlansData {
    pub plans: Vec<SystemPricingPlan>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SystemPricingPlan {
    /// Identifier for a pricing plan in the system.
    pub plan_id: PricingPlanID,
    /// URL where the customer can learn more about this pricing plan.
    pub url: Option<URL>,
    /// Name of this pricing plan.
    pub name: Option<Vec<LocalizedString>>,
    /// Currency used to pay the fare.
    pub currency: Currency,
    /// Fare price, in the unit specified by currency.
    ///
    /// In case of non-rate price, this field is the total price.
    /// In case of rate price, this field is the base price that is charged only once per trip (typically the price for unlocking) in addition to per_km_pricing and/or per_min_pricing.
    pub price: Option<NonNegativeFloat>,
    /// Will additional tax be added to the base price?
    pub is_taxable: Option<bool>,
    /// Customer-readable description of the pricing plan.
    /// This SHOULD include the duration, price, conditions, etc. that the publisher would like users to see.
    pub description: Option<Vec<LocalizedString>>,
    /// When the price is a function of distance traveled, displayed in kilometers.
    pub per_km_pricing: Option<Vec<PricingRate>>,
    /// When the price is a function of time traveled, displayed in minutes.
    pub per_min_pricing: Option<Vec<PricingRate>>,
    /// Is there currently an increase in price in response to increased demand in this pricing plan? If this field is empty, it means there is no surge pricing in effect.
    pub surge_pricing: Option<bool>,
}

#[cfg_attr(feature = "napi", napi(object))]
#[cfg_attr(feature = "pyo3", pyclass(get_all, set_all))]
#[serde_with::skip_serializing_none]
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PricingRate {
    /// The unit at which this segment rate starts being charged (inclusive).
    pub start: i32,
    /// Rate that is charged for each unit interval after the start.
    /// Can be a negative number, which indicates that the traveler will receive a discount.
    pub rate: f64,
    /// Interval at which the rate of this segment is either reapplied indefinitely, or if defined, up until (but not including) `end` unit.
    ///
    /// An interval of 0 indicates the rate is only charged once.
    pub interval: NonNegativeFloat,
    /// The unit at which the rate will no longer apply (exclusive) for example, if end is `20` the rate no longer applies at `20.00`.
    ///
    ///  If this field is empty, the price issued for this segment is charged until the trip ends, in addition to the cost of any subsequent segments.
    pub end: Option<i32>,
}
