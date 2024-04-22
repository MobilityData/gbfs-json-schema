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
 * Auto-discovery file that links to all of the other files published by the system.
 */
export interface Gbfs {
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
}

export interface Data {
    /**
     * An array of all of the feeds that are published by the auto-discovery file. Each element
     * in the array is an object with the keys below.
     */
    feeds: Feed[];
    [property: string]: any;
}

export interface Feed {
    /**
     * Key identifying the type of feed this is. The key must be the base file name defined in
     * the spec for the corresponding feed type.
     */
    name: Name;
    /**
     * URL for the feed.
     */
    url: string;
    [property: string]: any;
}

/**
 * Key identifying the type of feed this is. The key must be the base file name defined in
 * the spec for the corresponding feed type.
 */
export type Name = "gbfs" | "gbfs_versions" | "system_information" | "vehicle_types" | "station_information" | "station_status" | "vehicle_status" | "system_alerts" | "system_regions" | "system_pricing_plans" | "geofencing_zones";

export type Version = "3.0";
