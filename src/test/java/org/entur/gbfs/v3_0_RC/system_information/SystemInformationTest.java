package org.entur.gbfs.v3_0_RC.system_information;

import org.entur.gbfs.TestBase;
import org.junit.jupiter.api.Test;

import java.io.IOException;

class SystemInformationTest extends TestBase {
    @Test
    void testUnmarshal() throws IOException {
        assertUnmarshalDoesNotThrow("v3_0_RC/system_information.json", GBFSSystemInformation.class);
    }
}
