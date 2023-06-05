package org.entur.gbfs.v2_3.geofencing_zones;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class GeofencingZonesTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/geofencing_zones.json", GBFSGeofencingZones.class);
    }
}
