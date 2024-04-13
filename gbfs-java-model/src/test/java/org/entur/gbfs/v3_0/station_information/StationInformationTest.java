package org.entur.gbfs.v3_0.station_information;

import org.entur.gbfs.TestBase;
import org.entur.gbfs.v3_0.station_information.GBFSStationInformation;
import org.junit.jupiter.api.Test;

class StationInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/station_information.json", GBFSStationInformation.class);
    }
}
