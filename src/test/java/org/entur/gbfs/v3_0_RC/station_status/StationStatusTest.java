package org.entur.gbfs.v3_0_RC.station_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class StationStatusTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/station_status.json", GBFSStationStatus.class);
    }
}
