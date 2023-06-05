package org.entur.gbfs.v2_2.station_information;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

public class StationInformationTest extends TestBase {
    @Test
    public void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v2_X/station_information.json", GBFSStationInformation.class);
    }
}
