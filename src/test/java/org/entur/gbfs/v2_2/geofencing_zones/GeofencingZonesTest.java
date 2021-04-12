package org.entur.gbfs.v2_2.geofencing_zones;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class GeofencingZonesTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test @Disabled
    public void testUnmarshal() throws IOException {
        //URL resource = getClass().getClassLoader().getResource("geofencing_zones.json");
        //objectMapper.readValue(resource, GBFSGeofencingZones.class);
    }
}
