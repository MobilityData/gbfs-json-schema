package org.entur.gbfs.v2_2.station_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class StationStatusTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/station_status.json", GBFSStationStatus.class);
    }
}
