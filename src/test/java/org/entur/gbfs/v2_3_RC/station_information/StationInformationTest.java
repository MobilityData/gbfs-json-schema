package org.entur.gbfs.v2_3_RC.station_information;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class StationInformationTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("station_information.json");
        objectMapper.readValue(resource, GBFSStationInformation.class);
    }
}
