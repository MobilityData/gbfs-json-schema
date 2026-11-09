package org.mobilitydata.gbfs;

import tools.jackson.databind.ObjectMapper;

import java.io.InputStream;
import java.net.URL;
import java.util.concurrent.atomic.AtomicReference;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public abstract class TestBase {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    public <T> T assertUnmarshalDoesNotThrow(String path, Class<T> clazz) {
        URL resource = getClass().getClassLoader().getResource(path);
        AtomicReference<T> returnValue = new AtomicReference<>();
        assertDoesNotThrow(() -> {
            try (InputStream in = resource.openStream()) {
                returnValue.set(objectMapper.readValue(in, clazz));
            }
        });
        return returnValue.get();
    }
}
