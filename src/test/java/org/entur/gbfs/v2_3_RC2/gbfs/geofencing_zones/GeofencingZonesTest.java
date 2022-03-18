package org.entur.gbfs.v2_3_RC2.gbfs.geofencing_zones;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.entur.gbfs.v2_3_RC2.geofencing_zones.GBFSGeofencingZones;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class GeofencingZonesTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("geofencing_zones.json");
        objectMapper.readValue(resource, GBFSGeofencingZones.class);
    }
}
