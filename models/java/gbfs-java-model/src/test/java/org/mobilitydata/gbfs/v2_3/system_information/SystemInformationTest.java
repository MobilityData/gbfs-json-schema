package org.mobilitydata.gbfs.v2_3.system_information;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v2_X/system_information.json", GBFSSystemInformation.class);
    }
}
