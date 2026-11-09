package org.mobilitydata.gbfs.v3_1_RC3;

import org.junit.jupiter.api.Test;
import org.mobilitydata.gbfs.JsonEquivalence;
import org.mobilitydata.gbfs.v3_1_RC3.gbfs.GBFSGbfs;
import org.mobilitydata.gbfs.v3_1_RC3.gbfs_versions.GBFSGbfsVersions;
import org.mobilitydata.gbfs.v3_1_RC3.geofencing_zones.GBFSGeofencingZones;
import org.mobilitydata.gbfs.v3_1_RC3.manifest.GBFSManifest;
import org.mobilitydata.gbfs.v3_1_RC3.station_information.GBFSStationInformation;
import org.mobilitydata.gbfs.v3_1_RC3.station_status.GBFSStationStatus;
import org.mobilitydata.gbfs.v3_1_RC3.system_alerts.GBFSSystemAlerts;
import org.mobilitydata.gbfs.v3_1_RC3.system_information.GBFSSystemInformation;
import org.mobilitydata.gbfs.v3_1_RC3.system_pricing_plans.GBFSSystemPricingPlans;
import org.mobilitydata.gbfs.v3_1_RC3.system_regions.GBFSSystemRegions;
import org.mobilitydata.gbfs.v3_1_RC3.vehicle_availability.GBFSVehicleAvailability;
import org.mobilitydata.gbfs.v3_1_RC3.vehicle_status.GBFSVehicleStatus;
import org.mobilitydata.gbfs.v3_1_RC3.vehicle_types.GBFSVehicleTypes;
import tools.jackson.databind.JsonNode;
import tools.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.io.InputStream;
import java.util.LinkedHashMap;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertTrue;

/**
 * Serializing a deserialized feed must give back the same JSON, see {@link JsonEquivalence} for what counts as the same.
 */
class RoundTripTest {
    private static final ObjectMapper MAPPER = new ObjectMapper();

    private static final Map<String, Class<?>> FEEDS = new LinkedHashMap<>();

    static {
        FEEDS.put("gbfs", GBFSGbfs.class);
        FEEDS.put("gbfs_versions", GBFSGbfsVersions.class);
        FEEDS.put("geofencing_zones", GBFSGeofencingZones.class);
        FEEDS.put("manifest", GBFSManifest.class);
        FEEDS.put("station_information", GBFSStationInformation.class);
        FEEDS.put("station_status", GBFSStationStatus.class);
        FEEDS.put("system_alerts", GBFSSystemAlerts.class);
        FEEDS.put("system_information", GBFSSystemInformation.class);
        FEEDS.put("system_pricing_plans", GBFSSystemPricingPlans.class);
        FEEDS.put("system_regions", GBFSSystemRegions.class);
        FEEDS.put("vehicle_availability", GBFSVehicleAvailability.class);
        FEEDS.put("vehicle_status", GBFSVehicleStatus.class);
        FEEDS.put("vehicle_types", GBFSVehicleTypes.class);
    }

    @Test
    void v31Rc3FeedsRoundTrip() throws IOException {
        for (Map.Entry<String, Class<?>> feed : FEEDS.entrySet()) {
            String path = "v3_1_RC3/" + feed.getKey() + ".json";
            JsonNode original;
            Object model;
            try (InputStream in = getClass().getClassLoader().getResourceAsStream(path)) {
                byte[] bytes = in.readAllBytes();
                original = MAPPER.readTree(bytes);
                model = MAPPER.readValue(bytes, feed.getValue());
            }
            JsonNode written = MAPPER.valueToTree(model);
            assertTrue(JsonEquivalence.equivalent(original, written), path + "\nexpected: " + original + "\nwas: " + written);
        }
    }
}
