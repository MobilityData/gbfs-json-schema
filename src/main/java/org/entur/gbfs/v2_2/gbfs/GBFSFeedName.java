package org.entur.gbfs.v2_2.gbfs;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

import java.util.HashMap;
import java.util.Map;

public enum GBFSFeedName {
    GBFS("gbfs"),
    GBFSVersions("gbfs_versions"),
    SystemInformation("system_information"),
    VehicleTypes("vehicle_types"),
    StationInformation("station_information"),
    StationStatus("station_status"),
    FreeBikeStatus("free_bike_status"),
    SystemHours("system_hours"),
    SystemAlerts("system_alerts"),
    SystemCalendar("system_calendar"),
    SystemRegions("system_regions"),
    SystemPricingPlans("system_pricing_plans"),
    GeofencingZones("geofencing_zones");

    private final String value;
    private final static Map<String, GBFSFeedName> CONSTANTS = new HashMap<String, GBFSFeedName>();

    static {
        for (GBFSFeedName c: values()) {
            CONSTANTS.put(c.value, c);
        }
    }

    private GBFSFeedName(String value) {
        this.value = value;
    }

    @Override
    public String toString() {
        return this.value;
    }

    @JsonValue
    public String value() {
        return this.value;
    }

    @JsonCreator
    public static GBFSFeedName fromValue(String value) {
        GBFSFeedName constant = CONSTANTS.get(value);
        if (constant == null) {
            throw new IllegalArgumentException(value);
        } else {
            return constant;
        }
    }

}

