package org.entur.gbfs.v2_2.free_bike_status;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class FreeBikeStatusTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v2_X/free_bike_status.json");
        objectMapper.readValue(resource, GBFSFreeBikeStatus.class);
    }
}
