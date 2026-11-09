package org.mobilitydata.gbfs.v3_0;

import org.junit.jupiter.api.Test;
import org.mobilitydata.gbfs.TestBase;
import org.mobilitydata.gbfs.v3_0.gbfs_versions.GBFSGbfsVersions;
import tools.jackson.core.JacksonException;
import tools.jackson.databind.ObjectMapper;
import tools.jackson.databind.cfg.DateTimeFeature;
import tools.jackson.databind.json.JsonMapper;

import java.time.Instant;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

/**
 * GBFS v3 defines timestamps as RFC 3339 date-time strings.
 */
class DateSerializationTest extends TestBase {
    private static final ObjectMapper DEFAULT_MAPPER = new ObjectMapper();

    private static GBFSGbfsVersions read(ObjectMapper mapper, String lastUpdated) {
        return mapper.readValue("{\"last_updated\": " + lastUpdated + "}", GBFSGbfsVersions.class);
    }

    @Test
    void serializesAsRfc3339String() {
        GBFSGbfsVersions versions = assertUnmarshalDoesNotThrow("v3_0/gbfs_versions.json", GBFSGbfsVersions.class);

        assertEquals("2021-12-30T17:59:23Z", DEFAULT_MAPPER.valueToTree(versions).get("last_updated").stringValue());
    }

    @Test
    void serializationDoesNotDependOnMapperConfiguration() {
        ObjectMapper timestampMapper = JsonMapper.builder().enable(DateTimeFeature.WRITE_DATES_AS_TIMESTAMPS).build();
        GBFSGbfsVersions versions = read(DEFAULT_MAPPER, "\"2021-12-30T17:59:23.5Z\"");

        assertEquals("2021-12-30T17:59:23.500Z", timestampMapper.valueToTree(versions).get("last_updated").stringValue());
    }

    @Test
    void deserializesRfc3339DateTimes() {
        Instant expected = Instant.parse("2023-07-17T11:34:13Z");
        List<String> values = List.of(
                "2023-07-17T11:34:13Z",
                "2023-07-17T13:34:13+02:00",
                "2023-07-17T04:34:13-07:00",
                "2023-07-17T11:34:13.000Z",
                "2023-07-17T11:34:13.000000Z",
                "2023-07-17t11:34:13z"
        );
        for (String value : values) {
            assertEquals(expected, read(DEFAULT_MAPPER, "\"" + value + "\"").getLastUpdated().toInstant(), value);
        }
    }

    @Test
    void deserializesNullAsNull() {
        assertNull(read(DEFAULT_MAPPER, "null").getLastUpdated());
    }

    @Test
    void rejectsValuesThatAreNotRfc3339DateTimes() {
        List<String> values = List.of(
                "1640887163",
                "\"1640887163\"",
                "\"2023-07-17\"",
                "\"2023-07-17T11:34:13\"",
                "\"2023-07-17T13:34:13+0200\"",
                "\" 2023-07-17T11:34:13Z\"",
                "\"not a date\""
        );
        for (String value : values) {
            assertThrows(JacksonException.class, () -> read(DEFAULT_MAPPER, value), value);
        }
    }
}
