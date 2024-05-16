package org.mobilitydata.gbfs.v2_3.station_status;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class StationStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/station_status.json", GBFSStationStatus.class);
    }
}
