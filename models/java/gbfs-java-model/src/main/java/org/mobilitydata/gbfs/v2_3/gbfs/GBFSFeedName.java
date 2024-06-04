package org.mobilitydata.gbfs.v2_3.gbfs;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import org.mobilitydata.gbfs.v2_3.free_bike_status.GBFSFreeBikeStatus;
import org.mobilitydata.gbfs.v2_3.system_calendar.GBFSSystemCalendar;
import org.mobilitydata.gbfs.v2_3.system_hours.GBFSSystemHours;
import org.mobilitydata.gbfs.v2_3.system_information.GBFSSystemInformation;
import org.mobilitydata.gbfs.v2_3.gbfs_versions.GBFSGbfsVersions;
import org.mobilitydata.gbfs.v2_3.geofencing_zones.GBFSGeofencingZones;
import org.mobilitydata.gbfs.v2_3.station_information.GBFSStationInformation;
import org.mobilitydata.gbfs.v2_3.station_status.GBFSStationStatus;
import org.mobilitydata.gbfs.v2_3.system_alerts.GBFSSystemAlerts;
import org.mobilitydata.gbfs.v2_3.system_pricing_plans.GBFSSystemPricingPlans;
import org.mobilitydata.gbfs.v2_3.system_regions.GBFSSystemRegions;
import org.mobilitydata.gbfs.v2_3.vehicle_types.GBFSVehicleTypes;

import java.util.HashMap;
import java.util.Map;

public enum GBFSFeedName {
    GBFS("gbfs", org.mobilitydata.gbfs.v2_3.gbfs.GBFS.class),
    GBFSVersions("gbfs_versions", GBFSGbfsVersions.class),
    SystemInformation("system_information", GBFSSystemInformation.class),
    VehicleTypes("vehicle_types", GBFSVehicleTypes.class),
    StationInformation("station_information", GBFSStationInformation.class),
    StationStatus("station_status", GBFSStationStatus.class),
    FreeBikeStatus("free_bike_status", GBFSFreeBikeStatus.class),
    SystemHours("system_hours", GBFSSystemHours.class),
    SystemAlerts("system_alerts", GBFSSystemAlerts.class),
    SystemCalendar("system_calendar", GBFSSystemCalendar.class),
    SystemRegions("system_regions", GBFSSystemRegions.class),
    SystemPricingPlans("system_pricing_plans", GBFSSystemPricingPlans.class),
    GeofencingZones("geofencing_zones", GBFSGeofencingZones.class);

    private final String value;
    private final Class<?> implementingClass;
    private final static Map<String, GBFSFeedName> CONSTANTS = new HashMap<String, GBFSFeedName>();
    private final static Map<Class<?>, GBFSFeedName> CLASSES = new HashMap<Class<?>, GBFSFeedName>();

    static {
        for (GBFSFeedName c: values()) {
            CONSTANTS.put(c.value, c);
            CLASSES.put(c.implementingClass, c);
        }
    }

    private GBFSFeedName(String value, Class<?> implementingClass) {
        this.value = value;
        this.implementingClass = implementingClass;
    }

    @Override
    public String toString() {
        return this.value;
    }

    @JsonValue
    public String value() {
        return this.value;
    }

    public Class<?> implementingClass() {
        return implementingClass;
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

    public static GBFSFeedName fromClass(Class<?> clazz) {
        GBFSFeedName constant = CLASSES.get(clazz);
        if (constant == null) {
            throw new IllegalArgumentException(clazz.getName());
        } else {
            return constant;
        }
    }
}

