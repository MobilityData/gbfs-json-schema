package org.entur.gbfs.v2_3.system_information;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemInformationTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("v2_X/system_information.json");
        objectMapper.readValue(resource, GBFSSystemInformation.class);
    }
}
