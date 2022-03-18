package org.entur.gbfs.v2_3.system_calendar;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemCalendarTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("system_calendar.json");
        objectMapper.readValue(resource, GBFSSystemCalendar.class);
    }
}
