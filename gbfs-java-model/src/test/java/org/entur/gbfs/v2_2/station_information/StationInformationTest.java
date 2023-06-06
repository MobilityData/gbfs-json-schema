package org.entur.gbfs.v2_2.station_information;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class StationInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/station_information.json", GBFSStationInformation.class);
    }
}
