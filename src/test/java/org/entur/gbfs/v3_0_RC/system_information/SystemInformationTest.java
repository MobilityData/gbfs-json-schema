package org.entur.gbfs.v3_0_RC.system_information;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

class SystemInformationTest extends TestBase {
    @Test
    void testUnmarshal() {
        assertUnmarshalDoesNotThrow("v3_0_RC/system_information.json", GBFSSystemInformation.class);
    }
}
