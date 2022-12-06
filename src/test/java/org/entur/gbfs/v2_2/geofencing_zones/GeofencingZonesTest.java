package org.entur.gbfs.v2_2.geofencing_zones;

import static org.junit.jupiter.api.Assertions.assertEquals;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.geojson.LngLatAlt;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class GeofencingZonesTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("geofencing_zones.json");
        GBFSGeofencingZones zones = objectMapper.readValue(resource, GBFSGeofencingZones.class);

        GBFSFeature feature = zones.getData().getGeofencingZones().getFeatures().get(0);

        LngLatAlt coord = feature.getGeometry().getCoordinates().get(0).get(0).get(0);

        assertEquals(60.137400634341, coord.getLatitude(), 0.01);
        assertEquals(11.53037789067, coord.getLongitude());
    }
}
