package org.entur.gbfs.v3_0_RC.gbfs_versions;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class GBFSVersionsTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v3_0_RC/gbfs_versions.json");
        objectMapper.readValue(resource, GBFSGbfsVersions.class);
    }
}
