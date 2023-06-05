package org.entur.gbfs.v2_3.geofencing_zones;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class GeofencingZonesTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/geofencing_zones.json", GBFSGeofencingZones.class);
    }
}
