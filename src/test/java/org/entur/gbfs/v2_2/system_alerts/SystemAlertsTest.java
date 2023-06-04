package org.entur.gbfs.v2_2.system_alerts;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemAlertsTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v2_X/system_alerts.json");
        objectMapper.readValue(resource, GBFSSystemAlerts.class);
    }
}
