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
 * Describes the vehicles that are available for rent (as of v3.0, formerly
 * free_bike_status).
 */
export interface VehicleStatus {
    /**
     * Array that contains one object per vehicle as defined below.
     */
    data: Data;
    /**
     * Last time the data in the feed was updated in RFC3339 format.
     */
    last_updated: string;
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
 * Array that contains one object per vehicle as defined below.
 */
export interface Data {
    vehicles: Vehicle[];
    [property: string]: any;
}

export interface Vehicle {
    /**
     * The date and time when any rental of the vehicle must be completed. Added in v2.3.
     */
    available_until?: string;
    /**
     * This value represents the current percentage, expressed from 0 to 1, of fuel or battery
     * power remaining in the vehicle. Added in v2.3-RC.
     */
    current_fuel_percent?: number;
    /**
     * The furthest distance in meters that the vehicle can travel without recharging or
     * refueling with the vehicle's current charge or fuel (added in v2.1-RC).
     */
    current_range_meters?: number;
    /**
     * The station_id of the station this vehicle must be returned to (added in v2.3-RC).
     */
    home_station_id?: string;
    /**
     * Is the vehicle currently disabled (broken)?
     */
    is_disabled: boolean;
    /**
     * Is the vehicle currently reserved?
     */
    is_reserved: boolean;
    /**
     * The last time this vehicle reported its status to the operator's backend in RFC3339
     * format (added in v2.1-RC).
     */
    last_reported?: string;
    /**
     * The latitude of the vehicle.
     */
    lat?: number;
    /**
     * The longitude of the vehicle.
     */
    lon?: number;
    /**
     * The plan_id of the pricing plan this vehicle is eligible for (added in v2.2).
     */
    pricing_plan_id?: string;
    /**
     * Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
     * in v1.1).
     */
    rental_uris?: RentalUris;
    /**
     * Identifier referencing the station_id if the vehicle is currently at a station (added in
     * v2.1-RC2).
     */
    station_id?: string;
    /**
     * List of vehicle equipment provided by the operator in addition to the accessories already
     * provided in the vehicle. Added in v2.3.
     */
    vehicle_equipment?: VehicleEquipment[];
    /**
     * Rotating (as of v2.0) identifier of a vehicle.
     */
    vehicle_id: string;
    /**
     * The vehicle_type_id of this vehicle (added in v2.1-RC).
     */
    vehicle_type_id?: string;
    [property: string]: any;
}

/**
 * Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
 * in v1.1).
 */
export interface RentalUris {
    /**
     * URI that can be passed to an Android app with an intent (added in v1.1).
     */
    android?: string;
    /**
     * URI that can be used on iOS to launch the rental app for this vehicle (added in v1.1).
     */
    ios?: string;
    /**
     * URL that can be used by a web browser to show more information about renting this vehicle
     * (added in v1.1).
     */
    web?: string;
    [property: string]: any;
}

export type VehicleEquipment = "child_seat_a" | "child_seat_b" | "child_seat_c" | "winter_tires" | "snow_chains";

export type Version = "3.0";
