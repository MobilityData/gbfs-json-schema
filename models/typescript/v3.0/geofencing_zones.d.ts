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
 * Describes geofencing zones and their associated rules and attributes (added in v2.1-RC).
 */
export interface GeofencingZones {
    /**
     * Array that contains geofencing information for the system.
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
 * Array that contains geofencing information for the system.
 */
export interface Data {
    /**
     * Each geofenced zone and its associated rules and attributes is described as an object
     * within the array of features.
     */
    geofencing_zones: GeofencingZonesObject;
    /**
     * Array of Rule objects defining restrictions that apply globally in all areas as the
     * default restrictions, except where overridden with an explicit geofencing zone.
     */
    global_rules: GlobalRule[];
    [property: string]: any;
}

/**
 * Each geofenced zone and its associated rules and attributes is described as an object
 * within the array of features.
 */
export interface GeofencingZonesObject {
    /**
     * Array of objects.
     */
    features: GeoJSONFeature[];
    /**
     * FeatureCollection as per IETF RFC 7946.
     */
    type: GeofencingZonesType;
    [property: string]: any;
}

export interface GeoJSONFeature {
    /**
     * A polygon that describes where rides might not be able to start, end, go through, or have
     * other limitations. Must follow the right-hand rule.
     */
    geometry: GeoJSONMultiPolygon;
    /**
     * Describing travel allowances and limitations.
     */
    properties: Properties;
    type:       FeatureType;
    [property: string]: any;
}

/**
 * A polygon that describes where rides might not be able to start, end, go through, or have
 * other limitations. Must follow the right-hand rule.
 */
export interface GeoJSONMultiPolygon {
    coordinates: Array<Array<Array<number[]>>>;
    type:        GeometryType;
    [property: string]: any;
}

export type GeometryType = "MultiPolygon";

/**
 * Describing travel allowances and limitations.
 */
export interface Properties {
    /**
     * End time of the geofencing zone in RFC3339 format.
     */
    end?: Date;
    /**
     * Public name of the geofencing zone.
     */
    name?: Name[];
    /**
     * Array of Rule objects defining restrictions that apply within the area of the polygon.
     */
    rules?: Rule[];
    /**
     * Start time of the geofencing zone in RFC3339 format.
     */
    start?: Date;
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

export interface Rule {
    /**
     * What is the maximum speed allowed, in kilometers per hour?
     */
    maximum_speed_kph?: number;
    /**
     * Is the ride allowed to end in this zone?
     */
    ride_end_allowed: boolean;
    /**
     * Is the ride allowed to start in this zone?
     */
    ride_start_allowed: boolean;
    /**
     * Is the ride allowed to travel through this zone?
     */
    ride_through_allowed: boolean;
    /**
     * Vehicle MUST be parked at stations defined in station_information.json within this
     * geofence zone
     */
    station_parking?: boolean;
    /**
     * Array of vehicle type IDs for which these restrictions apply.
     */
    vehicle_type_ids?: string[];
    [property: string]: any;
}

export type FeatureType = "Feature";

/**
 * FeatureCollection as per IETF RFC 7946.
 */
export type GeofencingZonesType = "FeatureCollection";

export interface GlobalRule {
    /**
     * What is the maximum speed allowed, in kilometers per hour?
     */
    maximum_speed_kph?: number;
    /**
     * Is the ride allowed to end in this zone?
     */
    ride_end_allowed: boolean;
    /**
     * Is the ride allowed to start in this zone?
     */
    ride_start_allowed: boolean;
    /**
     * Is the ride allowed to travel through this zone?
     */
    ride_through_allowed: boolean;
    /**
     * Vehicle MUST be parked at stations defined in station_information.json within this
     * geofence zone
     */
    station_parking?: boolean;
    /**
     * Array of vehicle type IDs for which these restrictions apply.
     */
    vehicle_type_ids?: string[];
    [property: string]: any;
}

export type Version = "3.0";
