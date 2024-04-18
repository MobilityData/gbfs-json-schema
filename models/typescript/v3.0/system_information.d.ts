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
 * Details including system operator, system location, year implemented, URL, contact info,
 * time zone.
 */
export interface SystemInformation {
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
     * GBFS version number to which the feed conforms, according to the versioning framework
     * (added in v1.1).
     */
    version: Version;
    [property: string]: any;
}

/**
 * Response data in the form of name:value pairs.
 */
export interface Data {
    /**
     * If the feed license requires attribution, name of the organization to which attribution
     * should be provided. An array with one object per supported language with the following
     * keys:
     */
    attribution_organization_name?: AttributionOrganizationName[];
    /**
     * URL of the organization to which attribution should be provided.
     */
    attribution_url?: string;
    /**
     * An object where each key defines one of the items listed below (added in v2.3-RC).
     */
    brand_assets?: BrandAssets;
    /**
     * Email address actively monitored by the operator's customer service department.
     */
    email?: string;
    /**
     * A single contact email address for consumers of this feed to report technical issues
     * (added in v1.1).
     */
    feed_contact_email: string;
    /**
     * List of languages used in translated strings. Each element in the list must be of type
     * Language.
     */
    languages: string[];
    /**
     * REQUIRED if the dataset is provided under a standard license. An identifier for a
     * standard license from the SPDX License List. Provide license_id rather than license_url
     * if the license is included in the SPDX License List. See the GBFS wiki for a comparison
     * of a subset of standard licenses. If the license_id and license_url fields are blank or
     * omitted, this indicates that the feed is provided under the Creative Commons Universal
     * Public Domain Dedication.
     */
    license_id?: LicenseID;
    /**
     * A fully qualified URL of a page that defines the license terms for the GBFS data for this
     * system.
     */
    license_url?: string;
    /**
     * REQUIRED if the producer publishes datasets for more than one system geography, for
     * example Berlin and Paris. A fully qualified URL pointing to the manifest.json file for
     * the publisher.
     */
    manifest_url?: string;
    /**
     * Name of the system to be displayed to customers. An array with one object per supported
     * language with the following keys:
     */
    name: Name[];
    /**
     * Hours and dates of operation for the system in OSM opening_hours format. (added in v3.0)
     */
    opening_hours: string;
    /**
     * Name of the system operator. An array with one object per supported language with the
     * following keys:
     */
    operator?: Operator[];
    /**
     * This OPTIONAL field SHOULD contain a single voice telephone number for the specified
     * system’s customer service department. MUST be in E.164 format as defined in Field Types.
     */
    phone_number?: string;
    /**
     * The date that the privacy policy provided at privacy_url was last updated (added in
     * v2.3-RC).
     */
    privacy_last_updated?: Date;
    /**
     * A fully qualified URL pointing to the privacy policy for the service. An array with one
     * object per supported language with the following keys:
     */
    privacy_url?: PrivacyURL[];
    /**
     * URL where a customer can purchase a membership.
     */
    purchase_url?: string;
    /**
     * Contains rental app information in the android and ios JSON objects (added in v1.1).
     */
    rental_apps?: RentalApps;
    /**
     * Abbreviation for a system. An array with one object per supported language with the
     * following keys:
     */
    short_name?: ShortName[];
    /**
     * Date that the system began operations.
     */
    start_date?: Date;
    /**
     * Identifier for this vehicle share system. This should be globally unique (even between
     * different systems).
     */
    system_id: string;
    /**
     * Date after which this data source will no longer be available to consuming applications.
     */
    termination_date?: Date;
    /**
     * The date that the terms of service provided at terms_url were last updated (added in
     * v2.3-RC)
     */
    terms_last_updated?: Date;
    /**
     * A fully qualified URL pointing to the terms of service (also often called "terms of use"
     * or "terms and conditions") for the service. An array with one object per supported
     * language with the following keys:
     */
    terms_url?: TermsURL[];
    /**
     * The time zone where the system is located.
     */
    timezone: Timezone;
    /**
     * The URL of the vehicle share system.
     */
    url?: string;
}

export interface AttributionOrganizationName {
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
 * An object where each key defines one of the items listed below (added in v2.3-RC).
 */
export interface BrandAssets {
    /**
     * A fully qualified URL pointing to the location of a graphic file representing the brand
     * for the service (added in v2.3-RC).
     */
    brand_image_url: string;
    /**
     * A fully qualified URL pointing to the location of a graphic file representing the brand
     * for the service for use in dark mode (added in v2.3-RC).
     */
    brand_image_url_dark?: string;
    /**
     * Date that indicates the last time any included brand assets were updated (added in
     * v2.3-RC).
     */
    brand_last_modified: Date;
    /**
     * A fully qualified URL pointing to the location of a page that defines the license terms
     * of brand icons, colors or other trademark information (added in v2.3-RC).
     */
    brand_terms_url?: string;
    /**
     * Color used to represent the brand for the service (added in v2.3-RC)
     */
    color?: string;
    [property: string]: any;
}

/**
 * REQUIRED if the dataset is provided under a standard license. An identifier for a
 * standard license from the SPDX License List. Provide license_id rather than license_url
 * if the license is included in the SPDX License List. See the GBFS wiki for a comparison
 * of a subset of standard licenses. If the license_id and license_url fields are blank or
 * omitted, this indicates that the feed is provided under the Creative Commons Universal
 * Public Domain Dedication.
 */
export type LicenseID = "0BSD" | "AAL" | "Abstyles" | "AdaCore-doc" | "Adobe-2006" | "Adobe-Glyph" | "ADSL" | "AFL-1.1" | "AFL-1.2" | "AFL-2.0" | "AFL-2.1" | "AFL-3.0" | "Afmparse" | "AGPL-1.0-only" | "AGPL-1.0-or-later" | "AGPL-3.0-only" | "AGPL-3.0-or-later" | "Aladdin" | "AMDPLPA" | "AML" | "AMPAS" | "ANTLR-PD" | "ANTLR-PD-fallback" | "Apache-1.0" | "Apache-1.1" | "Apache-2.0" | "APAFML" | "APL-1.0" | "App-s2p" | "APSL-1.0" | "APSL-1.1" | "APSL-1.2" | "APSL-2.0" | "Arphic-1999" | "Artistic-1.0" | "Artistic-1.0-cl8" | "Artistic-1.0-Perl" | "Artistic-2.0" | "Baekmuk" | "Bahyph" | "Barr" | "Beerware" | "Bitstream-Charter" | "Bitstream-Vera" | "BitTorrent-1.0" | "BitTorrent-1.1" | "blessing" | "BlueOak-1.0.0" | "Borceux" | "Brian-Gladman-3-Clause" | "BSD-1-Clause" | "BSD-2-Clause" | "BSD-2-Clause-Patent" | "BSD-2-Clause-Views" | "BSD-3-Clause" | "BSD-3-Clause-Attribution" | "BSD-3-Clause-Clear" | "BSD-3-Clause-LBNL" | "BSD-3-Clause-Modification" | "BSD-3-Clause-No-Military-License" | "BSD-3-Clause-No-Nuclear-License" | "BSD-3-Clause-No-Nuclear-License-2014" | "BSD-3-Clause-No-Nuclear-Warranty" | "BSD-3-Clause-Open-MPI" | "BSD-4-Clause" | "BSD-4-Clause-Shortened" | "BSD-4-Clause-UC" | "BSD-4.3RENO" | "BSD-4.3TAHOE" | "BSD-Advertising-Acknowledgement" | "BSD-Attribution-HPND-disclaimer" | "BSD-Protection" | "BSD-Source-Code" | "BSL-1.0" | "BUSL-1.1" | "bzip2-1.0.6" | "C-UDA-1.0" | "CAL-1.0" | "CAL-1.0-Combined-Work-Exception" | "Caldera" | "CATOSL-1.1" | "CC-BY-1.0" | "CC-BY-2.0" | "CC-BY-2.5" | "CC-BY-2.5-AU" | "CC-BY-3.0" | "CC-BY-3.0-AT" | "CC-BY-3.0-DE" | "CC-BY-3.0-IGO" | "CC-BY-3.0-NL" | "CC-BY-3.0-US" | "CC-BY-4.0" | "CC-BY-NC-1.0" | "CC-BY-NC-2.0" | "CC-BY-NC-2.5" | "CC-BY-NC-3.0" | "CC-BY-NC-3.0-DE" | "CC-BY-NC-4.0" | "CC-BY-NC-ND-1.0" | "CC-BY-NC-ND-2.0" | "CC-BY-NC-ND-2.5" | "CC-BY-NC-ND-3.0" | "CC-BY-NC-ND-3.0-DE" | "CC-BY-NC-ND-3.0-IGO" | "CC-BY-NC-ND-4.0" | "CC-BY-NC-SA-1.0" | "CC-BY-NC-SA-2.0" | "CC-BY-NC-SA-2.0-DE" | "CC-BY-NC-SA-2.0-FR" | "CC-BY-NC-SA-2.0-UK" | "CC-BY-NC-SA-2.5" | "CC-BY-NC-SA-3.0" | "CC-BY-NC-SA-3.0-DE" | "CC-BY-NC-SA-3.0-IGO" | "CC-BY-NC-SA-4.0" | "CC-BY-ND-1.0" | "CC-BY-ND-2.0" | "CC-BY-ND-2.5" | "CC-BY-ND-3.0" | "CC-BY-ND-3.0-DE" | "CC-BY-ND-4.0" | "CC-BY-SA-1.0" | "CC-BY-SA-2.0" | "CC-BY-SA-2.0-UK" | "CC-BY-SA-2.1-JP" | "CC-BY-SA-2.5" | "CC-BY-SA-3.0" | "CC-BY-SA-3.0-AT" | "CC-BY-SA-3.0-DE" | "CC-BY-SA-4.0" | "CC-PDDC" | "CC0-1.0" | "CDDL-1.0" | "CDDL-1.1" | "CDL-1.0" | "CDLA-Permissive-1.0" | "CDLA-Permissive-2.0" | "CDLA-Sharing-1.0" | "CECILL-1.0" | "CECILL-1.1" | "CECILL-2.0" | "CECILL-2.1" | "CECILL-B" | "CECILL-C" | "CERN-OHL-1.1" | "CERN-OHL-1.2" | "CERN-OHL-P-2.0" | "CERN-OHL-S-2.0" | "CERN-OHL-W-2.0" | "CFITSIO" | "checkmk" | "ClArtistic" | "Clips" | "CMU-Mach" | "CNRI-Jython" | "CNRI-Python" | "CNRI-Python-GPL-Compatible" | "COIL-1.0" | "Community-Spec-1.0" | "Condor-1.1" | "copyleft-next-0.3.0" | "copyleft-next-0.3.1" | "Cornell-Lossless-JPEG" | "CPAL-1.0" | "CPL-1.0" | "CPOL-1.02" | "Crossword" | "CrystalStacker" | "CUA-OPL-1.0" | "Cube" | "curl" | "D-FSL-1.0" | "diffmark" | "DL-DE-BY-2.0" | "DOC" | "Dotseqn" | "DRL-1.0" | "DSDP" | "dvipdfm" | "ECL-1.0" | "ECL-2.0" | "EFL-1.0" | "EFL-2.0" | "eGenix" | "Elastic-2.0" | "Entessa" | "EPICS" | "EPL-1.0" | "EPL-2.0" | "ErlPL-1.1" | "etalab-2.0" | "EUDatagrid" | "EUPL-1.0" | "EUPL-1.1" | "EUPL-1.2" | "Eurosym" | "Fair" | "FDK-AAC" | "Frameworx-1.0" | "FreeBSD-DOC" | "FreeImage" | "FSFAP" | "FSFUL" | "FSFULLR" | "FSFULLRWD" | "FTL" | "GD" | "GFDL-1.1-invariants-only" | "GFDL-1.1-invariants-or-later" | "GFDL-1.1-no-invariants-only" | "GFDL-1.1-no-invariants-or-later" | "GFDL-1.1-only" | "GFDL-1.1-or-later" | "GFDL-1.2-invariants-only" | "GFDL-1.2-invariants-or-later" | "GFDL-1.2-no-invariants-only" | "GFDL-1.2-no-invariants-or-later" | "GFDL-1.2-only" | "GFDL-1.2-or-later" | "GFDL-1.3-invariants-only" | "GFDL-1.3-invariants-or-later" | "GFDL-1.3-no-invariants-only" | "GFDL-1.3-no-invariants-or-later" | "GFDL-1.3-only" | "GFDL-1.3-or-later" | "Giftware" | "GL2PS" | "Glide" | "Glulxe" | "GLWTPL" | "gnuplot" | "GPL-1.0-only" | "GPL-1.0-or-later" | "GPL-2.0-only" | "GPL-2.0-or-later" | "GPL-3.0-only" | "GPL-3.0-or-later" | "Graphics-Gems" | "gSOAP-1.3b" | "HaskellReport" | "Hippocratic-2.1" | "HP-1986" | "HPND" | "HPND-export-US" | "HPND-Markus-Kuhn" | "HPND-sell-variant" | "HPND-sell-variant-MIT-disclaimer" | "HTMLTIDY" | "IBM-pibs" | "ICU" | "IEC-Code-Components-EULA" | "IJG" | "IJG-short" | "ImageMagick" | "iMatix" | "Imlib2" | "Info-ZIP" | "Intel" | "Intel-ACPI" | "Interbase-1.0" | "IPA" | "IPL-1.0" | "ISC" | "Jam" | "JasPer-2.0" | "JPL-image" | "JPNIC" | "JSON" | "Kazlib" | "Knuth-CTAN" | "LAL-1.2" | "LAL-1.3" | "Latex2e" | "Leptonica" | "LGPL-2.0-only" | "LGPL-2.0-or-later" | "LGPL-2.1-only" | "LGPL-2.1-or-later" | "LGPL-3.0-only" | "LGPL-3.0-or-later" | "LGPLLR" | "Libpng" | "libpng-2.0" | "libselinux-1.0" | "libtiff" | "libutil-David-Nugent" | "LiLiQ-P-1.1" | "LiLiQ-R-1.1" | "LiLiQ-Rplus-1.1" | "Linux-man-pages-copyleft" | "Linux-OpenIB" | "LOOP" | "LPL-1.0" | "LPL-1.02" | "LPPL-1.0" | "LPPL-1.1" | "LPPL-1.2" | "LPPL-1.3a" | "LPPL-1.3c" | "LZMA-SDK-9.11-to-9.20" | "LZMA-SDK-9.22" | "MakeIndex" | "Martin-Birgmeier" | "Minpack" | "MirOS" | "MIT" | "MIT-0" | "MIT-advertising" | "MIT-CMU" | "MIT-enna" | "MIT-feh" | "MIT-Modern-Variant" | "MIT-open-group" | "MIT-Wu" | "MITNFA" | "Motosoto" | "mpi-permissive" | "mpich2" | "MPL-1.0" | "MPL-1.1" | "MPL-2.0" | "MPL-2.0-no-copyleft-exception" | "mplus" | "MS-LPL" | "MS-PL" | "MS-RL" | "MTLL" | "MulanPSL-1.0" | "MulanPSL-2.0" | "Multics" | "Mup" | "NAIST-2003" | "NASA-1.3" | "Naumen" | "NBPL-1.0" | "NCGL-UK-2.0" | "NCSA" | "Net-SNMP" | "NetCDF" | "Newsletr" | "NGPL" | "NICTA-1.0" | "NIST-PD" | "NIST-PD-fallback" | "NLOD-1.0" | "NLOD-2.0" | "NLPL" | "Nokia" | "NOSL" | "Noweb" | "NPL-1.0" | "NPL-1.1" | "NPOSL-3.0" | "NRL" | "NTP" | "NTP-0" | "O-UDA-1.0" | "OCCT-PL" | "OCLC-2.0" | "ODbL-1.0" | "ODC-By-1.0" | "OFFIS" | "OFL-1.0" | "OFL-1.0-no-RFN" | "OFL-1.0-RFN" | "OFL-1.1" | "OFL-1.1-no-RFN" | "OFL-1.1-RFN" | "OGC-1.0" | "OGDL-Taiwan-1.0" | "OGL-Canada-2.0" | "OGL-UK-1.0" | "OGL-UK-2.0" | "OGL-UK-3.0" | "OGTSL" | "OLDAP-1.1" | "OLDAP-1.2" | "OLDAP-1.3" | "OLDAP-1.4" | "OLDAP-2.0" | "OLDAP-2.0.1" | "OLDAP-2.1" | "OLDAP-2.2" | "OLDAP-2.2.1" | "OLDAP-2.2.2" | "OLDAP-2.3" | "OLDAP-2.4" | "OLDAP-2.5" | "OLDAP-2.6" | "OLDAP-2.7" | "OLDAP-2.8" | "OML" | "OpenPBS-2.3" | "OpenSSL" | "OPL-1.0" | "OPUBL-1.0" | "OSET-PL-2.1" | "OSL-1.0" | "OSL-1.1" | "OSL-2.0" | "OSL-2.1" | "OSL-3.0" | "Parity-6.0.0" | "Parity-7.0.0" | "PDDL-1.0" | "PHP-3.0" | "PHP-3.01" | "Plexus" | "PolyForm-Noncommercial-1.0.0" | "PolyForm-Small-Business-1.0.0" | "PostgreSQL" | "PSF-2.0" | "psfrag" | "psutils" | "Python-2.0" | "Python-2.0.1" | "Qhull" | "QPL-1.0" | "QPL-1.0-INRIA-2004" | "Rdisc" | "RHeCos-1.1" | "RPL-1.1" | "RPL-1.5" | "RPSL-1.0" | "RSA-MD" | "RSCPL" | "Ruby" | "SAX-PD" | "Saxpath" | "SCEA" | "SchemeReport" | "Sendmail" | "Sendmail-8.23" | "SGI-B-1.0" | "SGI-B-1.1" | "SGI-B-2.0" | "SHL-0.5" | "SHL-0.51" | "SimPL-2.0" | "SISSL" | "SISSL-1.2" | "Sleepycat" | "SMLNJ" | "SMPPL" | "SNIA" | "snprintf" | "Spencer-86" | "Spencer-94" | "Spencer-99" | "SPL-1.0" | "SSH-OpenSSH" | "SSH-short" | "SSPL-1.0" | "SugarCRM-1.1.3" | "SunPro" | "SWL" | "Symlinks" | "TAPR-OHL-1.0" | "TCL" | "TCP-wrappers" | "TMate" | "TORQUE-1.1" | "TOSL" | "TPDL" | "TPL-1.0" | "TTWL" | "TU-Berlin-1.0" | "TU-Berlin-2.0" | "UCAR" | "UCL-1.0" | "Unicode-DFS-2015" | "Unicode-DFS-2016" | "Unicode-TOU" | "Unlicense" | "UPL-1.0" | "Vim" | "VOSTROM" | "VSL-1.0" | "W3C" | "W3C-19980720" | "W3C-20150513" | "w3m" | "Watcom-1.0" | "Wsuipa" | "WTFPL" | "X11" | "X11-distribute-modifications-variant" | "Xerox" | "XFree86-1.1" | "xinetd" | "xlock" | "Xnet" | "xpp" | "XSkat" | "YPL-1.0" | "YPL-1.1" | "Zed" | "Zend-2.0" | "Zimbra-1.3" | "Zimbra-1.4" | "Zlib" | "zlib-acknowledgement" | "ZPL-1.1" | "ZPL-2.0" | "ZPL-2.1";

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

export interface Operator {
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

export interface PrivacyURL {
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
 * Contains rental app information in the android and ios JSON objects (added in v1.1).
 */
export interface RentalApps {
    /**
     * Contains rental app download and app discovery information for the Android platform.
     * (added in v1.1)
     */
    android?: Android;
    /**
     * Contains rental information for the iOS platform (added in v1.1).
     */
    ios?: Ios;
    [property: string]: any;
}

/**
 * Contains rental app download and app discovery information for the Android platform.
 * (added in v1.1)
 */
export interface Android {
    /**
     * URI that can be used to discover if the rental Android app is installed on the device
     * (added in v1.1).
     */
    discovery_uri: string;
    /**
     * URI where the rental Android app can be downloaded from (added in v1.1).
     */
    store_uri: string;
    [property: string]: any;
}

/**
 * Contains rental information for the iOS platform (added in v1.1).
 */
export interface Ios {
    /**
     * URI that can be used to discover if the rental iOS app is installed on the device (added
     * in v1.1).
     */
    discovery_uri: string;
    /**
     * URI where the rental iOS app can be downloaded from (added in v1.1).
     */
    store_uri: string;
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

export interface TermsURL {
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
 * The time zone where the system is located.
 */
export type Timezone = "Africa/Abidjan" | "Africa/Algiers" | "Africa/Bissau" | "Africa/Cairo" | "Africa/Casablanca" | "Africa/Ceuta" | "Africa/El_Aaiun" | "Africa/Johannesburg" | "Africa/Juba" | "Africa/Khartoum" | "Africa/Lagos" | "Africa/Maputo" | "Africa/Monrovia" | "Africa/Nairobi" | "Africa/Ndjamena" | "Africa/Sao_Tome" | "Africa/Tripoli" | "Africa/Tunis" | "Africa/Windhoek" | "America/Adak" | "America/Anchorage" | "America/Araguaina" | "America/Argentina/Buenos_Aires" | "America/Argentina/Catamarca" | "America/Argentina/Cordoba" | "America/Argentina/Jujuy" | "America/Argentina/La_Rioja" | "America/Argentina/Mendoza" | "America/Argentina/Rio_Gallegos" | "America/Argentina/Salta" | "America/Argentina/San_Juan" | "America/Argentina/San_Luis" | "America/Argentina/Tucuman" | "America/Argentina/Ushuaia" | "America/Asuncion" | "America/Bahia" | "America/Bahia_Banderas" | "America/Barbados" | "America/Belem" | "America/Belize" | "America/Boa_Vista" | "America/Bogota" | "America/Boise" | "America/Cambridge_Bay" | "America/Campo_Grande" | "America/Cancun" | "America/Caracas" | "America/Cayenne" | "America/Chicago" | "America/Chihuahua" | "America/Costa_Rica" | "America/Cuiaba" | "America/Danmarkshavn" | "America/Dawson" | "America/Dawson_Creek" | "America/Denver" | "America/Detroit" | "America/Edmonton" | "America/Eirunepe" | "America/El_Salvador" | "America/Fort_Nelson" | "America/Fortaleza" | "America/Glace_Bay" | "America/Goose_Bay" | "America/Grand_Turk" | "America/Guatemala" | "America/Guayaquil" | "America/Guyana" | "America/Halifax" | "America/Havana" | "America/Hermosillo" | "America/Indiana/Indianapolis" | "America/Indiana/Knox" | "America/Indiana/Marengo" | "America/Indiana/Petersburg" | "America/Indiana/Tell_City" | "America/Indiana/Vevay" | "America/Indiana/Vincennes" | "America/Indiana/Winamac" | "America/Inuvik" | "America/Iqaluit" | "America/Jamaica" | "America/Juneau" | "America/Kentucky/Louisville" | "America/Kentucky/Monticello" | "America/La_Paz" | "America/Lima" | "America/Los_Angeles" | "America/Maceio" | "America/Managua" | "America/Manaus" | "America/Martinique" | "America/Matamoros" | "America/Mazatlan" | "America/Menominee" | "America/Merida" | "America/Metlakatla" | "America/Mexico_City" | "America/Miquelon" | "America/Moncton" | "America/Monterrey" | "America/Montevideo" | "America/New_York" | "America/Nipigon" | "America/Nome" | "America/Noronha" | "America/North_Dakota/Beulah" | "America/North_Dakota/Center" | "America/North_Dakota/New_Salem" | "America/Nuuk" | "America/Ojinaga" | "America/Panama" | "America/Pangnirtung" | "America/Paramaribo" | "America/Phoenix" | "America/Port-au-Prince" | "America/Porto_Velho" | "America/Puerto_Rico" | "America/Punta_Arenas" | "America/Rainy_River" | "America/Rankin_Inlet" | "America/Recife" | "America/Regina" | "America/Resolute" | "America/Rio_Branco" | "America/Santarem" | "America/Santiago" | "America/Santo_Domingo" | "America/Sao_Paulo" | "America/Scoresbysund" | "America/Sitka" | "America/St_Johns" | "America/Swift_Current" | "America/Tegucigalpa" | "America/Thule" | "America/Thunder_Bay" | "America/Tijuana" | "America/Toronto" | "America/Vancouver" | "America/Whitehorse" | "America/Winnipeg" | "America/Yakutat" | "America/Yellowknife" | "Antarctica/Casey" | "Antarctica/Davis" | "Antarctica/Macquarie" | "Antarctica/Mawson" | "Antarctica/Palmer" | "Antarctica/Rothera" | "Antarctica/Troll" | "Antarctica/Vostok" | "Asia/Almaty" | "Asia/Amman" | "Asia/Anadyr" | "Asia/Aqtau" | "Asia/Aqtobe" | "Asia/Ashgabat" | "Asia/Atyrau" | "Asia/Baghdad" | "Asia/Baku" | "Asia/Bangkok" | "Asia/Barnaul" | "Asia/Beirut" | "Asia/Bishkek" | "Asia/Brunei" | "Asia/Chita" | "Asia/Choibalsan" | "Asia/Colombo" | "Asia/Damascus" | "Asia/Dhaka" | "Asia/Dili" | "Asia/Dubai" | "Asia/Dushanbe" | "Asia/Famagusta" | "Asia/Gaza" | "Asia/Hebron" | "Asia/Ho_Chi_Minh" | "Asia/Hong_Kong" | "Asia/Hovd" | "Asia/Irkutsk" | "Asia/Jakarta" | "Asia/Jayapura" | "Asia/Jerusalem" | "Asia/Kabul" | "Asia/Kamchatka" | "Asia/Karachi" | "Asia/Kathmandu" | "Asia/Khandyga" | "Asia/Kolkata" | "Asia/Krasnoyarsk" | "Asia/Kuala_Lumpur" | "Asia/Kuching" | "Asia/Macau" | "Asia/Magadan" | "Asia/Makassar" | "Asia/Manila" | "Asia/Nicosia" | "Asia/Novokuznetsk" | "Asia/Novosibirsk" | "Asia/Omsk" | "Asia/Oral" | "Asia/Pontianak" | "Asia/Pyongyang" | "Asia/Qatar" | "Asia/Qostanay" | "Asia/Qyzylorda" | "Asia/Riyadh" | "Asia/Sakhalin" | "Asia/Samarkand" | "Asia/Seoul" | "Asia/Shanghai" | "Asia/Singapore" | "Asia/Srednekolymsk" | "Asia/Taipei" | "Asia/Tashkent" | "Asia/Tbilisi" | "Asia/Tehran" | "Asia/Thimphu" | "Asia/Tokyo" | "Asia/Tomsk" | "Asia/Ulaanbaatar" | "Asia/Urumqi" | "Asia/Ust-Nera" | "Asia/Vladivostok" | "Asia/Yakutsk" | "Asia/Yangon" | "Asia/Yekaterinburg" | "Asia/Yerevan" | "Atlantic/Azores" | "Atlantic/Bermuda" | "Atlantic/Canary" | "Atlantic/Cape_Verde" | "Atlantic/Faroe" | "Atlantic/Madeira" | "Atlantic/Reykjavik" | "Atlantic/South_Georgia" | "Atlantic/Stanley" | "Australia/Adelaide" | "Australia/Brisbane" | "Australia/Broken_Hill" | "Australia/Darwin" | "Australia/Eucla" | "Australia/Hobart" | "Australia/Lindeman" | "Australia/Lord_Howe" | "Australia/Melbourne" | "Australia/Perth" | "Australia/Sydney" | "CET" | "CST6CDT" | "EET" | "EST" | "EST5EDT" | "Etc/GMT" | "Etc/GMT-1" | "Etc/GMT-10" | "Etc/GMT-11" | "Etc/GMT-12" | "Etc/GMT-13" | "Etc/GMT-14" | "Etc/GMT-2" | "Etc/GMT-3" | "Etc/GMT-4" | "Etc/GMT-5" | "Etc/GMT-6" | "Etc/GMT-7" | "Etc/GMT-8" | "Etc/GMT-9" | "Etc/GMT+1" | "Etc/GMT+10" | "Etc/GMT+11" | "Etc/GMT+12" | "Etc/GMT+2" | "Etc/GMT+3" | "Etc/GMT+4" | "Etc/GMT+5" | "Etc/GMT+6" | "Etc/GMT+7" | "Etc/GMT+8" | "Etc/GMT+9" | "Etc/UTC" | "Europe/Amsterdam" | "Europe/Andorra" | "Europe/Astrakhan" | "Europe/Athens" | "Europe/Belgrade" | "Europe/Berlin" | "Europe/Brussels" | "Europe/Bucharest" | "Europe/Budapest" | "Europe/Chisinau" | "Europe/Copenhagen" | "Europe/Dublin" | "Europe/Gibraltar" | "Europe/Helsinki" | "Europe/Istanbul" | "Europe/Kaliningrad" | "Europe/Kiev" | "Europe/Kirov" | "Europe/Lisbon" | "Europe/London" | "Europe/Luxembourg" | "Europe/Madrid" | "Europe/Malta" | "Europe/Minsk" | "Europe/Monaco" | "Europe/Moscow" | "Europe/Oslo" | "Europe/Paris" | "Europe/Prague" | "Europe/Riga" | "Europe/Rome" | "Europe/Samara" | "Europe/Saratov" | "Europe/Simferopol" | "Europe/Sofia" | "Europe/Stockholm" | "Europe/Tallinn" | "Europe/Tirane" | "Europe/Ulyanovsk" | "Europe/Uzhgorod" | "Europe/Vienna" | "Europe/Vilnius" | "Europe/Volgograd" | "Europe/Warsaw" | "Europe/Zaporozhye" | "Europe/Zurich" | "HST" | "Indian/Chagos" | "Indian/Christmas" | "Indian/Cocos" | "Indian/Kerguelen" | "Indian/Mahe" | "Indian/Maldives" | "Indian/Mauritius" | "Indian/Reunion" | "MET" | "MST" | "MST7MDT" | "Pacific/Apia" | "Pacific/Auckland" | "Pacific/Bougainville" | "Pacific/Chatham" | "Pacific/Chuuk" | "Pacific/Easter" | "Pacific/Efate" | "Pacific/Fakaofo" | "Pacific/Fiji" | "Pacific/Funafuti" | "Pacific/Galapagos" | "Pacific/Gambier" | "Pacific/Guadalcanal" | "Pacific/Guam" | "Pacific/Honolulu" | "Pacific/Kanton" | "Pacific/Kiritimati" | "Pacific/Kosrae" | "Pacific/Kwajalein" | "Pacific/Majuro" | "Pacific/Marquesas" | "Pacific/Nauru" | "Pacific/Niue" | "Pacific/Norfolk" | "Pacific/Noumea" | "Pacific/Pago_Pago" | "Pacific/Palau" | "Pacific/Pitcairn" | "Pacific/Pohnpei" | "Pacific/Port_Moresby" | "Pacific/Rarotonga" | "Pacific/Tahiti" | "Pacific/Tarawa" | "Pacific/Tongatapu" | "Pacific/Wake" | "Pacific/Wallis" | "PST8PDT" | "WET";

export type Version = "3.0";
