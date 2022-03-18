package org.entur.gbfs.v2_3_RC2.gbfs.system_information;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.entur.gbfs.v2_3_RC2.system_information.GBFSSystemInformation;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class SystemInformationTest {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("system_information.json");
        objectMapper.readValue(resource, GBFSSystemInformation.class);
    }
}
