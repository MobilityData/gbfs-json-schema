package org.entur.gbfs.v2_3.gbfs;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.net.URL;

public class GBFSTest {

    private static final ObjectMapper objectMapper = new ObjectMapper();

    @Test
    public void testUnmarshal() throws IOException {
        URL resource = getClass().getClassLoader().getResource("gbfs.json");
        objectMapper.readValue(resource, GBFS.class);
    }
}
