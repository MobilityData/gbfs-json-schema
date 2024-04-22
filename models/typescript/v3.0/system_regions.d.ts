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
 * Describes regions for a system that is broken up by geographic or political region.
 */
export interface SystemRegions {
    /**
     * Describe regions for a system that is broken up by geographic or political region.
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
 * Describe regions for a system that is broken up by geographic or political region.
 */
export interface Data {
    /**
     * Array of regions.
     */
    regions: Region[];
    [property: string]: any;
}

export interface Region {
    /**
     * Public name for this region.
     */
    name: Name[];
    /**
     * identifier of the region.
     */
    region_id: string;
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

export type Version = "3.0";
