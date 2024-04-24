package org.mobilitydata.gbfs.v3_0.gbfs;

import org.mobilitydata.gbfs.v3_0.GBFSGbfsVersions;
import org.mobilitydata.gbfs.v3_0.GBFSGeofencingZones;
import org.mobilitydata.gbfs.v3_0.GBFSStationInformation;
import org.mobilitydata.gbfs.v3_0.GBFSStationStatus;
import org.mobilitydata.gbfs.v3_0.GBFSSystemAlerts;
import org.mobilitydata.gbfs.v3_0.GBFSSystemInformation;
import org.mobilitydata.gbfs.v3_0.GBFSSystemPricingPlans;
import org.mobilitydata.gbfs.v3_0.GBFSSystemRegions;
import org.mobilitydata.gbfs.v3_0.GBFSVehicleStatus;
import org.mobilitydata.gbfs.v3_0.GBFSVehicleTypes;
import org.mobilitydata.gbfs.v3_0.GBFSFeed;
import org.mobilitydata.gbfs.v3_0.GBFSGbfs;

import java.util.EnumMap;
import java.util.Map;
import java.util.stream.Collectors;

public class GBFSFeedName {
    private GBFSFeedName() {}


    private static final Map<GBFSFeed.Name, Class<?>> feedNameMap = new EnumMap<>(GBFSFeed.Name.class);

    private static final Map<Class<?>, GBFSFeed.Name> classMap;

    static {
        feedNameMap.put(GBFSFeed.Name.GBFS, GBFSGbfs.class);
        feedNameMap.put(GBFSFeed.Name.GBFS_VERSIONS, GBFSGbfsVersions.class);
        feedNameMap.put(GBFSFeed.Name.GEOFENCING_ZONES, GBFSGeofencingZones.class);
        feedNameMap.put(GBFSFeed.Name.STATION_INFORMATION, GBFSStationInformation.class);
        feedNameMap.put(GBFSFeed.Name.STATION_STATUS, GBFSStationStatus.class);
        feedNameMap.put(GBFSFeed.Name.SYSTEM_ALERTS, GBFSSystemAlerts.class);
        feedNameMap.put(GBFSFeed.Name.SYSTEM_INFORMATION, GBFSSystemInformation.class);
        feedNameMap.put(GBFSFeed.Name.SYSTEM_PRICING_PLANS, GBFSSystemPricingPlans.class);
        feedNameMap.put(GBFSFeed.Name.SYSTEM_REGIONS, GBFSSystemRegions.class);
        feedNameMap.put(GBFSFeed.Name.VEHICLE_STATUS, GBFSVehicleStatus.class);
        feedNameMap.put(GBFSFeed.Name.VEHICLE_TYPES, GBFSVehicleTypes.class);
        classMap = feedNameMap.entrySet().stream().collect(Collectors.toMap(Map.Entry::getValue, Map.Entry::getKey));
    }

    public static Class<?> implementingClass(GBFSFeed.Name feedName) {
        return feedNameMap.get(feedName);
    }

    public static GBFSFeed.Name fromClass(Class<?> clazz) {
        return classMap.get(clazz);
    }
}
