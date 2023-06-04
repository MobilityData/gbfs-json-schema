package org.entur.gbfs.v3_0_RC.station_status;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class StationStatusTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v3_0_RC/station_status.json");
        objectMapper.readValue(resource, GBFSStationStatus.class);
    }
}
