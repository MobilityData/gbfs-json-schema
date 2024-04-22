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
 * An index of gbfs.json URLs for each GBFS data set produced by a publisher. A single
 * instance of this file should be published at a single stable URL, for example:
 * https://example.com/gbfs/manifest.json.
 */
export interface Manifest {
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
    version: ManifestVersion;
    [property: string]: any;
}

export interface Data {
    /**
     * An array of objects, each containing the keys below.
     */
    datasets: Dataset[];
}

export interface Dataset {
    /**
     * The system_id from system_information.json for the corresponding data set(s).
     */
    system_id: string;
    /**
     * Contains one object, as defined below, for each of the available versions of a feed. The
     * array MUST be sorted by increasing MAJOR and MINOR version number.
     */
    versions: VersionElement[];
    [property: string]: any;
}

export interface VersionElement {
    /**
     * URL of the corresponding gbfs.json endpoint
     */
    url: string;
    /**
     * The semantic version of the feed in the form X.Y
     */
    version: VersionVersion;
    [property: string]: any;
}

/**
 * The semantic version of the feed in the form X.Y
 */
export type VersionVersion = "1.0" | "1.1" | "2.0" | "2.1" | "2.2" | "2.3" | "3.0";

export type ManifestVersion = "3.0";
