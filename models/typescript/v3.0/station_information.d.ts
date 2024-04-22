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
 * List of all stations, their capacities and locations. REQUIRED of systems utilizing docks.
 */
export interface StationInformation {
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
     * Address where station is located.
     */
    address?: string;
    /**
     * Number of total docking points installed at this station, both available and unavailable.
     */
    capacity?: number;
    /**
     * Contact phone of the station. Added in v2.3
     */
    contact_phone?: string;
    /**
     * Cross street or landmark where the station is located.
     */
    cross_street?: string;
    /**
     * Does the station support charging of electric vehicles? (added in v2.3-RC)
     */
    is_charging_station?: boolean;
    /**
     * Are valet services provided at this station? (added in v2.1-RC)
     */
    is_valet_station?: boolean;
    /**
     * Is this station a location with or without physical infrastructure? (added in v2.1-RC)
     */
    is_virtual_station?: boolean;
    /**
     * The latitude of the station.
     */
    lat: number;
    /**
     * The longitude fo the station.
     */
    lon: number;
    /**
     * Public name of the station.
     */
    name: Name[];
    /**
     * Are parking hoops present at this station? Added in v2.3
     */
    parking_hoop?: boolean;
    /**
     * Type of parking station. Added in v2.3
     */
    parking_type?: ParkingType;
    /**
     * Postal code where station is located.
     */
    post_code?: string;
    /**
     * Identifier of the region where the station is located.
     */
    region_id?: string;
    /**
     * Payment methods accepted at this station.
     */
    rental_methods?: RentalMethod[];
    /**
     * Contains rental uris for Android, iOS, and web in the android, ios, and web fields (added
     * in v1.1).
     */
    rental_uris?: RentalUris;
    /**
     * Short name or other type of identifier.
     */
    short_name?: ShortName[];
    /**
     * A multipolygon that describes the area of a virtual station (added in v2.1-RC).
     */
    station_area?: StationArea;
    /**
     * Identifier of a station.
     */
    station_id: string;
    /**
     * Hours of operation for the station in OSM opening_hours format.
     */
    station_opening_hours?: string;
    /**
     * This field's value is an array of objects containing the keys vehicle_type_ids and count
     * defined below. These objects are used to model the total docking capacity of a station,
     * both available and unavailable, for each type of vehicle that may dock at this station.
     */
    vehicle_docks_capacity?: VehicleDocksCapacity[];
    /**
     * This field's value is an array of objects containing the keys vehicle_type_ids and count
     * defined below. These objects are used to model the parking capacity of virtual stations
     * (defined using the is_virtual_station field) for each vehicle type that can be returned
     * to this station.
     */
    vehicle_types_capacity?: VehicleTypesCapacity[];
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
 * Type of parking station. Added in v2.3
 */
export type ParkingType = "parking_lot" | "street_parking" | "underground_parking" | "sidewalk_parking" | "other";

export type RentalMethod = "key" | "creditcard" | "paypass" | "applepay" | "androidpay" | "transitcard" | "accountnumber" | "phone";

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
     * URI that can be used on iOS to launch the rental app for this station (added in v1.1).
     */
    ios?: string;
    /**
     * URL that can be used by a web browser to show more information about renting a vehicle at
     * this station (added in v1.1).
     */
    web?: string;
    [property: string]: any;
}

export interface ShortName {
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
 * A multipolygon that describes the area of a virtual station (added in v2.1-RC).
 */
export interface StationArea {
    coordinates: Array<Array<Array<number[]>>>;
    type:        Type;
    [property: string]: any;
}

export type Type = "MultiPolygon";

export interface VehicleDocksCapacity {
    /**
     * A number representing the total number of docks at the station, both available and
     * unavailable, that may accept the vehicle types specified by vehicle_type_ids.
     */
    count: number;
    /**
     * An array of strings where each string represents a vehicle_type_id that is able to use a
     * particular type of dock at the station.
     */
    vehicle_type_ids: string[];
    [property: string]: any;
}

export interface VehicleTypesCapacity {
    /**
     * A number representing the total number of vehicles of the specified vehicle_type_ids that
     * can park within the virtual station.
     */
    count: number;
    /**
     * The vehicle_type_ids, as defined in vehicle_types.json, that may park at the virtual
     * station.
     */
    vehicle_type_ids: string[];
    [property: string]: any;
}

export type Version = "3.0";
