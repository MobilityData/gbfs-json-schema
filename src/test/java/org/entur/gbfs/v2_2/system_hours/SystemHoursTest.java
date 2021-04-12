package org.entur.gbfs.v2_2.system_hours;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemHoursTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("system_hours.json");
        objectMapper.readValue(resource, GBFSSystemHours.class);
    }
}
