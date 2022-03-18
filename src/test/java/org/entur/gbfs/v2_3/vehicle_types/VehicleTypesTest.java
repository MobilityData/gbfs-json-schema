package org.entur.gbfs.v2_3.vehicle_types;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class VehicleTypesTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("vehicle_types.json");
        objectMapper.readValue(resource, GBFSVehicleTypes.class);
    }
}
