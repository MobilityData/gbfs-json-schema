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
 * Describes the types of vehicles that System operator has available for rent (added in
 * v2.1-RC).
 */
export interface VehicleTypes {
    /**
     * Response data in the form of name:value pairs.
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
     * GBFS version number to which the feed conforms, according to the versioning framework.
     */
    version: Version;
    [property: string]: any;
}

/**
 * Response data in the form of name:value pairs.
 */
export interface Data {
    /**
     * Array that contains one object per vehicle type in the system as defined below.
     */
    vehicle_types: VehicleType[];
    [property: string]: any;
}

export interface VehicleType {
    /**
     * The capacity of the vehicle cargo space (excluding passengers), expressed in kilograms.
     */
    cargo_load_capacity?: number;
    /**
     * Cargo volume available in the vehicle, expressed in liters.
     */
    cargo_volume_capacity?: number;
    /**
     * The color of the vehicle. Added in v2.3
     */
    color?: string;
    /**
     * A plan_id as defined in system_pricing_plans.json added in v2.3-RC.
     */
    default_pricing_plan_id?: string;
    /**
     * Maximum time in minutes that a vehicle can be reserved before a rental begins added in
     * v2.3-RC.
     */
    default_reserve_time?: number;
    /**
     * Customer-readable description of the vehicle type outlining special features or how-tos.
     * An array with one object per supported language with the following keys:
     */
    description?: Description[];
    /**
     * Vehicle air quality certificate. added in v2.3.
     */
    eco_labels?: EcoLabel[];
    /**
     * The vehicle's general form factor.
     */
    form_factor: FormFactor;
    /**
     * Maximum quantity of CO2, in grams, emitted per kilometer, according to the WLTP. Added in
     * v2.3
     */
    g_CO2_km?: number;
    /**
     * The name of the vehicle manufacturer. An array with one object per supported language
     * with the following keys:
     */
    make?: Make[];
    /**
     * The maximum speed in kilometers per hour this vehicle is permitted to reach in accordance
     * with local permit and regulations. Added in v2.3
     */
    max_permitted_speed?: number;
    /**
     * The furthest distance in meters that the vehicle can travel without recharging or
     * refueling when it has the maximum amount of energy potential.
     */
    max_range_meters?: number;
    /**
     * The name of the vehicle model. An array with one object per supported language with the
     * following keys:
     */
    model?: Model[];
    /**
     * The public name of this vehicle type. An array with one object per supported language
     * with the following keys:
     */
    name?: Name[];
    /**
     * Array of all pricing plan IDs as defined in system_pricing_plans.json added in v2.3-RC.
     */
    pricing_plan_ids?: string[];
    /**
     * The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing
     */
    propulsion_type: PropulsionType;
    /**
     * The rated power of the motor for this vehicle type in watts. Added in v2.3
     */
    rated_power?: number;
    /**
     * The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as
     * return_type, and updated to return_constraint in v2.3.
     */
    return_constraint?: ReturnConstraint;
    /**
     * The number of riders (driver included) the vehicle can legally accommodate
     */
    rider_capacity?: number;
    /**
     * Description of accessories available in the vehicle.
     */
    vehicle_accessories?: VehicleAccessory[];
    /**
     * An object where each key defines one of the items listed below added in v2.3-RC.
     */
    vehicle_assets?: VehicleAssets;
    /**
     * URL to an image that would assist the user in identifying the vehicle. JPEG or PNG. Added
     * in v2.3
     */
    vehicle_image?: string;
    /**
     * Unique identifier of a vehicle type.
     */
    vehicle_type_id: string;
    /**
     * Number of wheels this vehicle type has. Added in v2.3
     */
    wheel_count?: number;
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

export interface EcoLabel {
    /**
     * Country code following the ISO 3166-1 alpha-2 notation. Added in v2.3.
     */
    country_code?: string;
    /**
     * Name of the eco label. Added in v2.3.
     */
    eco_sticker?: string;
    [property: string]: any;
}

/**
 * The vehicle's general form factor.
 */
export type FormFactor = "bicycle" | "cargo_bicycle" | "car" | "moped" | "scooter_standing" | "scooter_seated" | "other" | "scooter";

export interface Make {
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

export interface Model {
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

/**
 * The primary propulsion type of the vehicle. Updated in v2.3 to represent car-sharing
 */
export type PropulsionType = "human" | "electric_assist" | "electric" | "combustion" | "combustion_diesel" | "hybrid" | "plug_in_hybrid" | "hydrogen_fuel_cell";

/**
 * The conditions for returning the vehicle at the end of the trip. Added in v2.3-RC as
 * return_type, and updated to return_constraint in v2.3.
 */
export type ReturnConstraint = "free_floating" | "roundtrip_station" | "any_station" | "hybrid";

export type VehicleAccessory = "air_conditioning" | "automatic" | "manual" | "convertible" | "cruise_control" | "doors_2" | "doors_3" | "doors_4" | "doors_5" | "navigation";

/**
 * An object where each key defines one of the items listed below added in v2.3-RC.
 */
export interface VehicleAssets {
    /**
     * Date that indicates the last time any included vehicle icon images were modified or
     * updated added in v2.3-RC.
     */
    icon_last_modified: Date;
    /**
     * A fully qualified URL pointing to the location of a graphic icon file that MAY be used to
     * represent this vehicle type on maps and in other applications added in v2.3-RC.
     */
    icon_url: string;
    /**
     * A fully qualified URL pointing to the location of a graphic icon file to be used to
     * represent this vehicle type when in dark mode added in v2.3-RC.
     */
    icon_url_dark?: string;
    [property: string]: any;
}

export type Version = "3.0";
