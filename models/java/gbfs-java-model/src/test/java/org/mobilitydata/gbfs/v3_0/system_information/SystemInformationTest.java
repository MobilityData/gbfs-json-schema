package org.mobilitydata.gbfs.v3_0.system_information;

import org.mobilitydata.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0/system_information.json", GBFSSystemInformation.class);
    }
}
