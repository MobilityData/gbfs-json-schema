package org.entur.gbfs.v2_3_RC2.gbfs.station_status;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.entur.gbfs.v2_3_RC2.station_status.GBFSStationStatus;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class StationStatusTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("station_status.json");
        objectMapper.readValue(resource, GBFSStationStatus.class);
    }
}
