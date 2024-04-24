package org.mobilitydata.gbfs.v3_0.station_status;

import org.mobilitydata.gbfs.TestBase;
import org.mobilitydata.v3_0.GBFSStationStatus;
import org.junit.jupiter.api.Test;

class StationStatusTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/station_status.json", GBFSStationStatus.class);
    }
}
