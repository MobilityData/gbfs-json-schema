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
 * Describes ad-hoc changes to the system.
 */
export interface SystemAlerts {
    /**
     * Array that contains ad-hoc alerts for the system.
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
 * Array that contains ad-hoc alerts for the system.
 */
export interface Data {
    alerts: Alert[];
    [property: string]: any;
}

export interface Alert {
    /**
     * Identifier for this alert.
     */
    alert_id: string;
    /**
     * Detailed description of the alert.
     */
    description?: Description[];
    /**
     * Indicates the last time the info for the alert was updated.
     */
    last_updated?: Date;
    /**
     * Array of identifiers of the regions for which this alert applies.
     */
    region_ids?: string[];
    /**
     * Array of identifiers of the stations for which this alert applies.
     */
    station_ids?: string[];
    /**
     * A short summary of this alert to be displayed to the customer.
     */
    summary: Summary[];
    /**
     * Array of objects indicating when the alert is in effect.
     */
    times?: Time[];
    /**
     * Type of alert.
     */
    type: Type;
    /**
     * URL where the customer can learn more information about this alert.
     */
    url?: URL[];
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

export interface Summary {
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

export interface Time {
    /**
     * End time of the alert.
     */
    end?: Date;
    /**
     * Start time of the alert.
     */
    start?: Date;
    [property: string]: any;
}

/**
 * Type of alert.
 */
export type Type = "system_closure" | "station_closure" | "station_move" | "other";

export interface URL {
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

export type Version = "3.0";
