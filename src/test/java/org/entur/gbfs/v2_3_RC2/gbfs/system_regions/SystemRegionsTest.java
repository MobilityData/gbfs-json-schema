package org.entur.gbfs.v2_3_RC2.gbfs.system_regions;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.entur.gbfs.v2_3_RC2.system_regions.GBFSSystemRegions;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemRegionsTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("system_regions.json");
        objectMapper.readValue(resource, GBFSSystemRegions.class);
    }
}
