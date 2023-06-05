package org.entur.gbfs.v3_0_RC.station_information;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class StationInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/station_information.json", GBFSStationInformation.class);
    }
}
