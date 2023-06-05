package org.entur.gbfs.v3_0_RC.gbfs;

import org.entur.gbfs.v3_0_RC.gbfs_versions.GBFSGbfsVersions;
import org.entur.gbfs.v3_0_RC.geofencing_zones.GBFSGeofencingZones;
import org.entur.gbfs.v3_0_RC.manifest.GBFSManifest;
import org.entur.gbfs.v3_0_RC.station_information.GBFSStationInformation;
import org.entur.gbfs.v3_0_RC.station_status.GBFSStationStatus;
import org.entur.gbfs.v3_0_RC.system_alerts.GBFSSystemAlerts;
import org.entur.gbfs.v3_0_RC.system_information.GBFSSystemInformation;
import org.entur.gbfs.v3_0_RC.system_pricing_plans.GBFSSystemPricingPlans;
import org.entur.gbfs.v3_0_RC.system_regions.GBFSSystemRegions;
import org.entur.gbfs.v3_0_RC.vehicle_status.GBFSVehicleStatus;
import org.entur.gbfs.v3_0_RC.vehicle_types.GBFSVehicleTypes;

import java.util.Map;
import java.util.stream.Collectors;

import static java.util.Map.entry;

public class GBFSFeedName {
    static private Map<GBFSFeed.Name, Class<?>> feedNameMap = Map.ofEntries(
            entry(GBFSFeed.Name.GBFS, GBFSFeed.class),
            entry(GBFSFeed.Name.GBFS_VERSIONS, GBFSGbfsVersions.class),
            entry(GBFSFeed.Name.GEOFENCING_ZONES, GBFSGeofencingZones.class),
            entry(GBFSFeed.Name.MANIFEST, GBFSManifest.class),
            entry(GBFSFeed.Name.STATION_INFORMATION, GBFSStationInformation.class),
            entry(GBFSFeed.Name.STATION_STATUS, GBFSStationStatus.class),
            entry(GBFSFeed.Name.SYSTEM_ALERTS, GBFSSystemAlerts.class),
            entry(GBFSFeed.Name.SYSTEM_INFORMATION, GBFSSystemInformation.class),
            entry(GBFSFeed.Name.SYSTEM_PRICING_PLANS, GBFSSystemPricingPlans.class),
            entry(GBFSFeed.Name.SYSTEM_REGIONS, GBFSSystemRegions.class),
            entry(GBFSFeed.Name.VEHICLE_STATUS, GBFSVehicleStatus.class),
            entry(GBFSFeed.Name.VEHICLE_TYPES, GBFSVehicleTypes.class)
    );

    static private Map<Class<?>, GBFSFeed.Name> classMap = feedNameMap.entrySet().stream().collect(Collectors.toMap(Map.Entry::getValue, Map.Entry::getKey));

    public static Class<?> implementingClass(GBFSFeed.Name feedName) {
        return feedNameMap.get(feedName);
    }

    public static GBFSFeed.Name fromClass(Class<?> clazz) {
        return classMap.get(clazz);
    }
}
