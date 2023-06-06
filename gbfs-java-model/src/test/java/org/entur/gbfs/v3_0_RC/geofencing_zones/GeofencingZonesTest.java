package org.entur.gbfs.v3_0_RC.geofencing_zones;

import org.entur.gbfs.TestBase;
import org.geojson.LngLatAlt;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GeofencingZonesTest extends TestBase {
    @Test
    void testUnmarshal() {
        GBFSGeofencingZones zones = assertUnmarshalDoesNotThrow("v3_0_RC/geofencing_zones.json", GBFSGeofencingZones.class);
        GBFSFeature feature = zones.getData().getGeofencingZones().getFeatures().get(0);
        LngLatAlt coord = feature.getGeometry().getCoordinates().get(0).get(0).get(0);
        assertEquals(45.562982, coord.getLatitude(), 0.01);
        assertEquals(-122.578067, coord.getLongitude());
    }
}
