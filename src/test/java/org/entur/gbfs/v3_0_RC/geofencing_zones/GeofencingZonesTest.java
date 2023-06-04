package org.entur.gbfs.v3_0_RC.geofencing_zones;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.geojson.LngLatAlt;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class GeofencingZonesTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v3_0_RC/geofencing_zones.json");
        GBFSGeofencingZones zones = objectMapper.readValue(resource, GBFSGeofencingZones.class);

        GBFSFeature feature = zones.getData().getGeofencingZones().getFeatures().get(0);

        LngLatAlt coord = feature.getGeometry().getCoordinates().get(0).get(0).get(0);

        assertEquals(45.562982, coord.getLatitude(), 0.01);
        assertEquals(-122.578067, coord.getLongitude());
    }
}
