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
 * Describes the capacity and rental availability of the station
 */
export interface StationStatus {
    /**
     * Array that contains one object per station as defined below.
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
 * Array that contains one object per station as defined below.
 */
export interface Data {
    stations: Station[];
    [property: string]: any;
}

export interface Station {
    /**
     * Is the station currently on the street?
     */
    is_installed: boolean;
    /**
     * Is the station currently renting vehicles?
     */
    is_renting: boolean;
    /**
     * Is the station accepting vehicle returns?
     */
    is_returning: boolean;
    /**
     * The last time this station reported its status to the operator's backend in RFC3339
     * format.
     */
    last_reported: string;
    /**
     * Number of functional docks physically at the station.
     */
    num_docks_available?: number;
    /**
     * Number of empty but disabled docks at the station.
     */
    num_docks_disabled?: number;
    /**
     * Number of vehicles of any type physically available for rental at the station.
     */
    num_vehicles_available: number;
    /**
     * Number of disabled vehicles of any type at the station.
     */
    num_vehicles_disabled?: number;
    /**
     * Identifier of a station.
     */
    station_id: string;
    /**
     * Object displaying available docks by vehicle type (added in v2.1-RC).
     */
    vehicle_docks_available?: VehicleDocksAvailable[];
    /**
     * Array of objects displaying the total number of each vehicle type at the station (added
     * in v2.1-RC).
     */
    vehicle_types_available?: VehicleTypesAvailable[];
    [property: string]: any;
}

export interface VehicleDocksAvailable {
    /**
     * A number representing the total number of available docks for the defined vehicle type
     * (added in v2.1-RC).
     */
    count: number;
    /**
     * An array of strings where each string represents a vehicle_type_id that is able to use a
     * particular type of dock at the station (added in v2.1-RC).
     */
    vehicle_type_ids: string[];
    [property: string]: any;
}

export interface VehicleTypesAvailable {
    /**
     * A number representing the total amount of this vehicle type at the station (added in
     * v2.1-RC).
     */
    count: number;
    /**
     * The vehicle_type_id of vehicle at the station (added in v2.1-RC).
     */
    vehicle_type_id: string;
    [property: string]: any;
}

export type Version = "3.0";
