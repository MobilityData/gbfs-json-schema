package org.entur.gbfs.v3_0_RC.station_status;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class StationStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/station_status.json", GBFSStationStatus.class);
    }
}
