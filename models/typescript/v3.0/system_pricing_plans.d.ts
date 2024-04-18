// Copyright 2024 MobilityData
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Describes the pricing schemes of the system.
 */
export interface SystemPricingPlans {
    /**
     * Array that contains one object per plan as defined below.
     */
    data: Data;
    /**
     * Last time the data in the feed was updated in RFC3339 format.
     */
    last_updated: Date;
    /**
     * Number of seconds before the data in the feed will be updated again (0 if the data should
     * always be refreshed).
     */
    ttl: number;
    /**
     * GBFS version number to which the feed conforms, according to the versioning framework
     * (added in v1.1).
     */
    version: Version;
    [property: string]: any;
}

/**
 * Array that contains one object per plan as defined below.
 */
export interface Data {
    plans: Plan[];
    [property: string]: any;
}

export interface Plan {
    /**
     * Currency used to pay the fare in ISO 4217 code.
     */
    currency: string;
    /**
     * Customer-readable description of the pricing plan.
     */
    description: Description[];
    /**
     * Will additional tax be added to the base price?
     */
    is_taxable: boolean;
    /**
     * Name of this pricing plan.
     */
    name: Name[];
    /**
     * Array of segments when the price is a function of distance travelled, displayed in
     * kilometers (added in v2.1-RC2).
     */
    per_km_pricing?: PerKMPricing[];
    /**
     * Array of segments when the price is a function of time travelled, displayed in minutes
     * (added in v2.1-RC2).
     */
    per_min_pricing?: PerMinPricing[];
    /**
     * Identifier of a pricing plan in the system.
     */
    plan_id: string;
    /**
     * Fare price.
     */
    price: number;
    /**
     * Is there currently an increase in price in response to increased demand in this pricing
     * plan? (added in v2.1-RC2)
     */
    surge_pricing?: boolean;
    /**
     * URL where the customer can learn more about this pricing plan.
     */
    url?: string;
    [property: string]: any;
}

export interface Description {
    /**
     * IETF BCP 47 language code.
     */
    language: string;
    /**
     * The translated text.
     */
    text: string;
    [property: string]: any;
}

export interface Name {
    /**
     * IETF BCP 47 language code.
     */
    language: string;
    /**
     * The translated text.
     */
    text: string;
    [property: string]: any;
}

export interface PerKMPricing {
    /**
     * The kilometer at which the rate will no longer apply (added in v2.1-RC2).
     */
    end?: number;
    /**
     * Interval in kilometers at which the rate of this segment is either reapplied
     * indefinitely, or if defined, up until (but not including) end kilometer (added in
     * v2.1-RC2).
     */
    interval: number;
    /**
     * Rate that is charged for each kilometer interval after the start (added in v2.1-RC2).
     */
    rate: number;
    /**
     * Number of kilometers that have to elapse before this segment starts applying (added in
     * v2.1-RC2).
     */
    start: number;
    [property: string]: any;
}

export interface PerMinPricing {
    /**
     * The minute at which the rate will no longer apply (added in v2.1-RC2).
     */
    end?: number;
    /**
     * Interval in minutes at which the rate of this segment is either reapplied (added in
     * v2.1-RC2).
     */
    interval: number;
    /**
     * Rate that is charged for each minute interval after the start (added in v2.1-RC2).
     */
    rate: number;
    /**
     * Number of minutes that have to elapse before this segment starts applying (added in
     * v2.1-RC2).
     */
    start: number;
    [property: string]: any;
}

export type Version = "3.0";
