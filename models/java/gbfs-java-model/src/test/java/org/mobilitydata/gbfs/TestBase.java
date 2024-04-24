package org.mobilitydata.gbfs;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.net.URL;
import java.util.concurrent.atomic.AtomicReference;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;

public abstract class TestBase {
    private static final ObjectMapper objectMapper = new ObjectMapper();

    public <T> T assertUnmarshalDoesNotThrow(String path, Class<T> clazz) {
        URL resource = getClass().getClassLoader().getResource(path);
        AtomicReference<T> returnValue = new AtomicReference<>();
        assertDoesNotThrow(() -> returnValue.set(objectMapper.readValue(resource, clazz)));
        return returnValue.get();
    }
}
