package org.mobilitydata.gbfs.v2_2.geofencing_zones;

import org.mobilitydata.gbfs.TestBase;
import org.geojson.LngLatAlt;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class GeofencingZonesTest extends TestBase {
    @Test
    void testUnmarshal() {
        GBFSGeofencingZones zones = assertUnmarshalDoesNotThrow("v2_X/geofencing_zones.json", GBFSGeofencingZones.class);
        GBFSFeature feature = zones.getData().getGeofencingZones().getFeatures().get(0);

        LngLatAlt coord = feature.getGeometry().getCoordinates().get(0).get(0).get(0);

        assertEquals(60.137400634341, coord.getLatitude(), 0.01);
        assertEquals(11.53037789067, coord.getLongitude());
    }
}
