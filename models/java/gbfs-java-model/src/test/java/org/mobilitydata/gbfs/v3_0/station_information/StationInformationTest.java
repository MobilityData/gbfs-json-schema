package org.mobilitydata.gbfs.v3_0.station_information;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class StationInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/station_information.json", GBFSStationInformation.class);
    }
}
